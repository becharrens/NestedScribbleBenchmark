package invitations

import "NestedScribbleBenchmark/clientserver/channels/dyntaskgen"

type DynTaskGen_RoleSetupChan struct {
	S_Chan chan dyntaskgen.S_Chan
}

type DynTaskGen_InviteSetupChan struct {
	S_InviteChan chan DynTaskGen_S_InviteChan
}

type DynTaskGen_S_InviteChan struct {
	Invite_S_To_DynTaskGen_S chan dyntaskgen.S_Chan
	Invite_S_To_DynTaskGen_S_InviteChan chan DynTaskGen_S_InviteChan
}

type DynTaskGen_W_InviteChan struct {

}