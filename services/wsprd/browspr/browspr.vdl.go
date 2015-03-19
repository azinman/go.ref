// This file was auto-generated by the vanadium vdl tool.
// Source: browspr.vdl

package browspr

import (
	// VDL system imports
	"v.io/v23/vdl"

	// VDL user imports
	"v.io/x/ref/services/identity"
	"v.io/x/ref/services/wsprd/account"
)

type StartMessage struct {
	Identityd             string
	IdentitydBlessingRoot identity.BlessingRootResponse
	Proxy                 string
	NamespaceRoot         string
	LogLevel              int32
	LogModule             string
}

func (StartMessage) __VDLReflect(struct {
	Name string "v.io/x/ref/services/wsprd/browspr.StartMessage"
}) {
}

type AssociateAccountMessage struct {
	Account string
	Origin  string
	Caveats []account.Caveat
}

func (AssociateAccountMessage) __VDLReflect(struct {
	Name string "v.io/x/ref/services/wsprd/browspr.AssociateAccountMessage"
}) {
}

type CreateAccountMessage struct {
	Token string
}

func (CreateAccountMessage) __VDLReflect(struct {
	Name string "v.io/x/ref/services/wsprd/browspr.CreateAccountMessage"
}) {
}

type CleanupMessage struct {
	InstanceId int32
}

func (CleanupMessage) __VDLReflect(struct {
	Name string "v.io/x/ref/services/wsprd/browspr.CleanupMessage"
}) {
}

type OriginHasAccountMessage struct {
	Origin string
}

func (OriginHasAccountMessage) __VDLReflect(struct {
	Name string "v.io/x/ref/services/wsprd/browspr.OriginHasAccountMessage"
}) {
}

type GetAccountsMessage struct {
}

func (GetAccountsMessage) __VDLReflect(struct {
	Name string "v.io/x/ref/services/wsprd/browspr.GetAccountsMessage"
}) {
}

type CreateInstanceMessage struct {
	InstanceId     int32
	Origin         string
	NamespaceRoots []string
	Proxy          string
}

func (CreateInstanceMessage) __VDLReflect(struct {
	Name string "v.io/x/ref/services/wsprd/browspr.CreateInstanceMessage"
}) {
}

func init() {
	vdl.Register((*StartMessage)(nil))
	vdl.Register((*AssociateAccountMessage)(nil))
	vdl.Register((*CreateAccountMessage)(nil))
	vdl.Register((*CleanupMessage)(nil))
	vdl.Register((*OriginHasAccountMessage)(nil))
	vdl.Register((*GetAccountsMessage)(nil))
	vdl.Register((*CreateInstanceMessage)(nil))
}
