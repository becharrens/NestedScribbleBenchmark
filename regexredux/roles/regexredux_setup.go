package roles

import "NestedScribbleBenchmark/regexredux/messages/regexredux"
import "NestedScribbleBenchmark/regexredux/channels/regexredux2"
import regexredux_2 "NestedScribbleBenchmark/regexredux/channels/regexredux"
import "NestedScribbleBenchmark/regexredux/invitations"
import "sync"

func RegexRedux_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.RegexRedux_RoleSetupChan, inviteChannels invitations.RegexRedux_InviteSetupChan) {
	worker_master_nummatches := make(chan regexredux.NumMatches, 1)
	master_invite_master := make(chan regexredux2.M_Chan, 1)
	master_invite_master_invitechan := make(chan invitations.RegexRedux2_M_InviteChan, 1)
	master_worker_task := make(chan regexredux.Task, 1)

	worker_chan := regexredux_2.Worker_Chan{
		Master_Task:       master_worker_task,
		Master_NumMatches: worker_master_nummatches,
	}
	master_chan := regexredux_2.Master_Chan{
		Worker_Task:       master_worker_task,
		Worker_NumMatches: worker_master_nummatches,
	}

	worker_inviteChan := invitations.RegexRedux_Worker_InviteChan{}
	master_inviteChan := invitations.RegexRedux_Master_InviteChan{
		Invite_Master_To_RegexRedux2_M_InviteChan: master_invite_master_invitechan,
		Invite_Master_To_RegexRedux2_M:            master_invite_master,
	}

	roleChannels.Master_Chan <- master_chan
	roleChannels.Worker_Chan <- worker_chan

	inviteChannels.Master_InviteChan <- master_inviteChan
	inviteChannels.Worker_InviteChan <- worker_inviteChan
}
