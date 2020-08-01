package invitations

import "NestedScribbleBenchmark/ring/channels/forward"

type Forward_RoleSetupChan struct {
	E_Chan chan forward.E_Chan
	S_Chan chan forward.S_Chan
}

type Forward_InviteSetupChan struct {
	E_InviteChan chan Forward_E_InviteChan
	S_InviteChan chan Forward_S_InviteChan
}

type Forward_S_InviteChan struct {
}

type Forward_E_InviteChan struct {
	RingNode_Invite_To_Forward_E            chan forward.E_Chan
	RingNode_Invite_To_Forward_E_InviteChan chan Forward_E_InviteChan
}

type Forward_RingNode_InviteChan struct {
	Invite_E_To_Forward_E                   chan forward.E_Chan
	Invite_E_To_Forward_E_InviteChan        chan Forward_E_InviteChan
	Invite_RingNode_To_Forward_S            chan forward.S_Chan
	Invite_RingNode_To_Forward_S_InviteChan chan Forward_S_InviteChan
}
