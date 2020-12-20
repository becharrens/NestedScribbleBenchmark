package roles

import "NestedScribbleBenchmark/regexredux/channels/regexredux2"
import "NestedScribbleBenchmark/regexredux/channels/regexredux"
import "NestedScribbleBenchmark/regexredux/invitations"
import "sync"

func RegexRedux_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.RegexRedux_RoleSetupChan, inviteChannels invitations.RegexRedux_InviteSetupChan) {
	master_invite_master := make(chan regexredux2.M_Chan, 1)
	master_invite_master_invitechan := make(chan invitations.RegexRedux2_M_InviteChan, 1)

	master_chan := regexredux.Master_Chan{}

	master_inviteChan := invitations.RegexRedux_Master_InviteChan{
		Invite_Master_To_RegexRedux2_M_InviteChan: master_invite_master_invitechan,
		Invite_Master_To_RegexRedux2_M:            master_invite_master,
	}

	roleChannels.Master_Chan <- master_chan

	inviteChannels.Master_InviteChan <- master_inviteChan
}
