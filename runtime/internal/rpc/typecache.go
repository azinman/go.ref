// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rpc

import (
	"sync"

	"v.io/v23/flow"
	"v.io/v23/vom"
)

type typeCacheEntry struct {
	enc    *vom.TypeEncoder
	dec    *vom.TypeDecoder
	writer bool
	ready  chan struct{}
}

const typeCacheCollectInterval = 1000

type typeCache struct {
	mu     sync.Mutex
	flows  map[flow.ManagedConn]*typeCacheEntry
	writes int
}

func newTypeCache() *typeCache {
	return &typeCache{flows: make(map[flow.ManagedConn]*typeCacheEntry)}
}

func (tc *typeCache) writer(c flow.ManagedConn) (write func(flow.Flow)) {
	tc.mu.Lock()
	tce := tc.flows[c]
	if tce == nil {
		tce = &typeCacheEntry{ready: make(chan struct{})}
		tc.flows[c] = tce
	}
	if !tce.writer {
		// TODO(mattr): This is a very course garbage collection policy.
		// Every 1000 connections we clean out expired entries.
		if tc.writes++; tc.writes%typeCacheCollectInterval == 0 {
			go tc.collect()
		}
		tce.writer = true
		write = func(f flow.Flow) {
			tce.enc = vom.NewTypeEncoder(f)
			tce.dec = vom.NewTypeDecoder(f)
			close(tce.ready)
		}
	}
	tc.mu.Unlock()
	return
}

func (tc *typeCache) get(c flow.ManagedConn) (*vom.TypeEncoder, *vom.TypeDecoder) {
	tc.mu.Lock()
	tce := tc.flows[c]
	if tce == nil {
		tce = &typeCacheEntry{ready: make(chan struct{})}
		tc.flows[c] = tce
	}
	tc.mu.Unlock()
	<-tce.ready
	return tce.enc, tce.dec
}

func (tc *typeCache) collect() {
	tc.mu.Lock()
	conns := make([]flow.ManagedConn, len(tc.flows))
	i := 0
	for c, _ := range tc.flows {
		conns[i] = c
		i++
	}
	tc.mu.Unlock()

	last := 0
	for _, c := range conns {
		select {
		case <-c.Closed():
			conns[last] = c
			last++
		default:
		}
	}

	if last > 0 {
		tc.mu.Lock()
		for idx := 0; idx < last; idx++ {
			delete(tc.flows, conns[idx])
		}
		tc.mu.Unlock()
	}
}