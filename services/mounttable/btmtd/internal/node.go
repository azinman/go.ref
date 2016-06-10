// Copyright 2016 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package internal

import (
	"encoding/hex"
	"encoding/json"
	"hash/fnv"
	"io/ioutil"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"

	"google.golang.org/cloud/bigtable"

	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/naming"
	"v.io/v23/security/access"
	vdltime "v.io/v23/vdlroot/time"
	"v.io/v23/verror"

	"v.io/x/ref/lib/timekeeper"
)

type mtNode struct {
	bt           *BigTable
	name         string
	sticky       bool
	creationTime bigtable.Timestamp
	permissions  access.Permissions
	version      string
	creator      string
	mountFlags   mtFlags
	servers      []naming.MountedServer
	children     []string
}

type mtFlags struct {
	MT   bool `json:"mt,omitempty"`
	Leaf bool `json:"leaf,omitempty"`
}

func longTimeout(ctx *context.T) (*context.T, func()) {
	return context.WithTimeout(v23.GetBackgroundContext(ctx), time.Hour)
}

func getNode(ctx *context.T, bt *BigTable, name string) (*mtNode, error) {
	row, err := bt.readRow(ctx, rowKey(name), bigtable.RowFilter(bigtable.LatestNFilter(1)))
	if err != nil {
		return nil, err
	}
	return nodeFromRow(ctx, bt, row, clock), nil
}

func rowKey(name string) string {
	// The row key is a hash of the node name followed by the node name
	// itself.
	// This spreads the rows evenly across the tablet servers to avoid
	// traffic imbalance.
	name = naming.Clean("/" + name)
	h := fnv.New32()
	h.Write([]byte(name))
	return hex.EncodeToString(h.Sum(nil)) + name
}

func nodeFromRow(ctx *context.T, bt *BigTable, row bigtable.Row, clock timekeeper.TimeKeeper) *mtNode {
	const offset = 9 // 32-bit value in hex + '/'
	name := row.Key()
	if len(name) < offset {
		return nil
	}
	n := &mtNode{
		bt:   bt,
		name: name[offset:],
	}
	for _, i := range row[metadataFamily] {
		col := strings.TrimPrefix(i.Column, metadataFamily+":")
		switch col {
		case stickyColumn:
			n.sticky = true
		case versionColumn:
			n.version = string(i.Value)
		case permissionsColumn:
			if err := json.Unmarshal(i.Value, &n.permissions); err != nil {
				ctx.Errorf("Failed to decode permissions for %s", name)
				return nil
			}
		case creatorColumn:
			n.creationTime = i.Timestamp
			n.creator = string(i.Value)
		}
	}
	n.servers = make([]naming.MountedServer, 0, len(row[serversFamily]))
	for _, i := range row[serversFamily] {
		deadline := i.Timestamp.Time()
		if deadline.Before(clock.Now()) {
			continue
		}
		if err := json.Unmarshal(i.Value, &n.mountFlags); err != nil {
			ctx.Errorf("Failed to decode mount flags for %s", name)
			return nil
		}
		n.servers = append(n.servers, naming.MountedServer{
			Server:   i.Column[2:],
			Deadline: vdltime.Deadline{deadline},
		})
	}
	n.children = make([]string, 0, len(row[childrenFamily]))
	for _, i := range row[childrenFamily] {
		child := strings.TrimPrefix(i.Column, childrenFamily+":")
		n.children = append(n.children, child)
	}
	return n
}

func (n *mtNode) createChild(ctx *context.T, child string, perms access.Permissions, creator string) (*mtNode, error) {
	ts := n.bt.now()
	mut := bigtable.NewMutation()
	mut.Set(childrenFamily, child, ts, []byte{1})
	if err := n.mutate(ctx, mut, false); err != nil {
		return nil, err
	}

	// If the current process dies right here, it will leave the parent with
	// a reference to a child row that doesn't exist. This means that the
	// parent will never be seen as "empty" and will not be garbage
	// collected. This will be corrected when:
	//  - the child is created again, or
	//  - the parent is forcibly deleted with Delete().

	childName := naming.Join(n.name, child)
	longCtx, cancel := longTimeout(ctx)
	defer cancel()
	if err := n.bt.createRow(longCtx, childName, perms, creator, ts); err != nil {
		return nil, err
	}
	n, err := getNode(ctx, n.bt, childName)
	if err != nil {
		return nil, err
	}
	if n == nil {
		return nil, verror.New(errConcurrentAccess, ctx, childName)
	}
	return n, nil
}

func (n *mtNode) mount(ctx *context.T, server string, deadline time.Time, flags naming.MountFlag) error {
	mut := bigtable.NewMutation()
	if flags&naming.Replace != 0 {
		mut.DeleteCellsInFamily(serversFamily)
	}
	f := mtFlags{
		MT:   flags&naming.MT != 0,
		Leaf: flags&naming.Leaf != 0,
	}
	jsonValue, err := json.Marshal(f)
	if err != nil {
		return err
	}
	mut.Set(serversFamily, server, n.bt.time(deadline), jsonValue)
	if err := n.mutate(ctx, mut, false); err != nil {
		return err
	}
	if flags&naming.Replace != 0 {
		n.servers = nil
	}
	return nil
}

func (n *mtNode) unmount(ctx *context.T, server string) error {
	mut := bigtable.NewMutation()
	if server == "" {
		// HACK ALERT
		// The bttest server doesn't support DeleteCellsInFamily
		if !n.bt.testMode {
			mut.DeleteCellsInFamily(serversFamily)
		} else {
			for _, s := range n.servers {
				mut.DeleteCellsInColumn(serversFamily, s.Server)
			}
		}
	} else {
		mut.DeleteCellsInColumn(serversFamily, server)
	}
	if err := n.mutate(ctx, mut, false); err != nil {
		return err
	}
	if n, err := getNode(ctx, n.bt, n.name); err == nil {
		n.gc(ctx)
	}
	return nil
}

func (n *mtNode) gc(ctx *context.T) (deletedAtLeastOne bool, err error) {
	for n != nil && n.name != "" && !n.sticky && len(n.children) == 0 && len(n.servers) == 0 {
		if err = n.delete(ctx, false); err != nil {
			break
		}
		ctx.Infof("Deleted empty node %q", n.name)
		deletedAtLeastOne = true
		parent := path.Dir(n.name)
		if parent == "." {
			break
		}
		if n, err = getNode(ctx, n.bt, parent); err != nil {
			break
		}
	}
	return
}

func (n *mtNode) deleteAndGC(ctx *context.T, deleteSubtree bool) error {
	if err := n.delete(ctx, deleteSubtree); err != nil {
		return err
	}
	parentName, _ := path.Split(n.name)
	if parent, err := getNode(ctx, n.bt, parentName); err == nil {
		parent.gc(ctx)
	}
	return nil
}

func (n *mtNode) delete(ctx *context.T, deleteSubtree bool) error {
	if !deleteSubtree && len(n.children) > 0 {
		return verror.New(errNotEmpty, ctx, n.name)
	}

	// TODO(rthellend): This naive approach could be very expensive in
	// terms of memory. A smarter, but slower, approach would be to walk
	// the tree without holding on to all the node data.
	for _, c := range n.children {
		cn, err := getNode(ctx, n.bt, naming.Join(n.name, c))
		if err != nil {
			return err
		}
		if cn == nil {
			// Node 'n' has a reference to a child that doesn't
			// exist. It could be that it is being created or
			// deleted concurrently. To be sure, we have to create
			// it before deleting it.
			if cn, err = n.createChild(ctx, c, n.permissions, ""); err != nil {
				return err
			}
		}
		if err := cn.delete(ctx, true); err != nil {
			return err
		}
	}

	mut := bigtable.NewMutation()
	mut.DeleteRow()
	if err := n.mutate(ctx, mut, true); err != nil {
		return err
	}

	// If the current process dies right here, it will leave the parent with
	// a reference to a child row that no longer exists. This means that the
	// parent will never be seen as "empty" and will not be garbage
	// collected. This will be corrected when:
	//  - the child is re-created, or
	//  - the parent is forcibly deleted with Delete().

	// Delete from parent node.
	parent, child := path.Split(n.name)
	mut = bigtable.NewMutation()
	// DeleteTimestampRange deletes the cells whose timestamp is in the
	// half open range [start,end). We need to delete the cell with
	// timestamp n.creationTime (and any older ones).
	mut.DeleteTimestampRange(childrenFamily, child, 0, n.bt.timeNext(n.creationTime))

	longCtx, cancel := longTimeout(ctx)
	defer cancel()
	if err := n.bt.apply(longCtx, rowKey(parent), mut); err != nil {
		return err
	}
	return n.bt.incrementCreatorNodeCount(ctx, n.creator, -1)
}

func (n *mtNode) setPermissions(ctx *context.T, perms access.Permissions) error {
	jsonPerms, err := json.Marshal(perms)
	if err != nil {
		return err
	}
	mut := bigtable.NewMutation()
	mut.Set(metadataFamily, permissionsColumn, bigtable.ServerTime, jsonPerms)
	mut.Set(metadataFamily, stickyColumn, bigtable.ServerTime, []byte{1})
	if err := n.mutate(ctx, mut, false); err != nil {
		return err
	}
	return nil
}

func (n *mtNode) mutate(ctx *context.T, mut *bigtable.Mutation, delete bool) error {
	if !delete {
		v, err := strconv.ParseUint(n.version, 10, 64)
		if err != nil {
			return err
		}
		newVersion := strconv.FormatUint(v+1, 10)
		mut.Set(metadataFamily, versionColumn, bigtable.ServerTime, []byte(newVersion))
	}

	// The mutation will succeed iff the row already exists with the
	// expected version.
	filter := bigtable.ChainFilters(
		bigtable.FamilyFilter(metadataFamily),
		bigtable.ColumnFilter(versionColumn),
		bigtable.LatestNFilter(1),
		bigtable.ValueFilter(n.version),
	)
	condMut := bigtable.NewCondMutation(filter, mut, nil)
	var success bool
	if err := n.bt.apply(ctx, rowKey(n.name), condMut, bigtable.GetCondMutationResult(&success)); err != nil {
		return err
	}
	if !success {
		return verror.New(errConcurrentAccess, ctx, n.name)
	}
	return nil
}

func createNodesFromFile(ctx *context.T, bt *BigTable, fileName string) error {
	var nodes map[string]access.Permissions
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &nodes); err != nil {
		return err
	}
	// This loop adds backward compatibility with the older template format,
	// e.g. "a/b/%%" { "Admin": { "In": [ "%%" ] } }
	// With the new format, this is equivalent to:
	// "a/b" { "%%/Admin": { "In": [ "%%" ] } }
	for node, perms := range nodes {
		if strings.HasSuffix(node, "/%%") {
			delete(nodes, node)
			node = strings.TrimSuffix(node, "/%%")
			p := nodes[node]
			if p == nil {
				p = make(access.Permissions)
			}
			for tag := range perms {
				p["%%/"+tag] = perms[tag]
			}
			nodes[node] = p
		}
	}

	// Create the nodes in alphanumeric order so that children are
	// created after their parents.
	sortedNodes := []string{}
	for node := range nodes {
		sortedNodes = append(sortedNodes, node)
	}
	sort.Strings(sortedNodes)

	ts := bt.now()
	for _, node := range sortedNodes {
		perms := nodes[node]
		if node == "" {
			if err := bt.createRow(ctx, "", perms, "", ts); err != nil {
				return err
			}
			continue
		}
		parentName := ""
		for _, e := range strings.Split(node, "/") {
			n, err := getNode(ctx, bt, naming.Join(parentName, e))
			if err != nil {
				return err
			}
			if n == nil {
				parent, err := getNode(ctx, bt, parentName)
				if err != nil {
					return err
				}
				if n, err = parent.createChild(ctx, e, parent.permissions, ""); err != nil {
					return err
				}
			}
			if n.name == node {
				// setPermissions also makes the node sticky.
				if err := n.setPermissions(ctx, perms); err != nil {
					return err
				}
			}
			parentName = n.name
		}
	}
	return nil
}
