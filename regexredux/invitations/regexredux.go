package invitations

import "NestedScribbleBenchmark/regexredux/channels/regexredux2"
import "NestedScribbleBenchmark/regexredux/channels/regexredux"

type RegexRedux_RoleSetupChan struct {
	Master_Chan chan regexredux.Master_Chan
}

type RegexRedux_InviteSetupChan struct {
	Master_InviteChan chan RegexRedux_Master_InviteChan
}

type RegexRedux_Master_InviteChan struct {
	Invite_Master_To_RegexRedux2_M            chan regexredux2.M_Chan
	Invite_Master_To_RegexRedux2_M_InviteChan chan RegexRedux2_M_InviteChan
}
