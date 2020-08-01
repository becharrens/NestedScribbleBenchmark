package invitations

import "NestedScribbleBenchmark/regexredux/channels/regexredux2"

type RegexRedux2_RoleSetupChan struct {
	M_Chan chan regexredux2.M_Chan
}

type RegexRedux2_InviteSetupChan struct {
	M_InviteChan chan RegexRedux2_M_InviteChan
}

type RegexRedux2_M_InviteChan struct {
	Invite_M_To_RegexRedux2_M            chan regexredux2.M_Chan
	Invite_M_To_RegexRedux2_M_InviteChan chan RegexRedux2_M_InviteChan
}

type RegexRedux2_W_InviteChan struct {
}
