package roles

import "NestedScribbleBenchmark/fannkuch/messages/fannkuch"
import "NestedScribbleBenchmark/fannkuch/channels/fannkuchrecursive"
import fannkuch_2 "NestedScribbleBenchmark/fannkuch/channels/fannkuch"
import "NestedScribbleBenchmark/fannkuch/invitations"
import "sync"

func Fannkuch_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.Fannkuch_RoleSetupChan, inviteChannels invitations.Fannkuch_InviteSetupChan) {
	worker_main_result_2 := make(chan fannkuch.Result, 1)
	worker_main_result := make(chan fannkuch.Result, 1)
	worker_invite_worker := make(chan fannkuchrecursive.Worker_Chan, 1)
	worker_invite_worker_invitechan := make(chan invitations.FannkuchRecursive_Worker_InviteChan, 1)
	worker_invite_main := make(chan fannkuchrecursive.Source_Chan, 1)
	worker_invite_main_invitechan := make(chan invitations.FannkuchRecursive_Source_InviteChan, 1)
	main_worker_task := make(chan fannkuch.Task, 1)

	worker_chan := fannkuch_2.Worker_Chan{
		Main_Task:     main_worker_task,
		Main_Result_2: worker_main_result_2,
		Main_Result:   worker_main_result,
	}
	main_chan := fannkuch_2.Main_Chan{
		Worker_Task:     main_worker_task,
		Worker_Result_2: worker_main_result_2,
		Worker_Result:   worker_main_result,
	}

	worker_inviteChan := invitations.Fannkuch_Worker_InviteChan{
		Invite_Worker_To_FannkuchRecursive_Worker_InviteChan: worker_invite_worker_invitechan,
		Invite_Worker_To_FannkuchRecursive_Worker:            worker_invite_worker,
		Invite_Main_To_FannkuchRecursive_Source_InviteChan:   worker_invite_main_invitechan,
		Invite_Main_To_FannkuchRecursive_Source:              worker_invite_main,
	}
	main_inviteChan := invitations.Fannkuch_Main_InviteChan{
		Worker_Invite_To_FannkuchRecursive_Source_InviteChan: worker_invite_main_invitechan,
		Worker_Invite_To_FannkuchRecursive_Source:            worker_invite_main,
	}

	roleChannels.Main_Chan <- main_chan
	roleChannels.Worker_Chan <- worker_chan

	inviteChannels.Main_InviteChan <- main_inviteChan
	inviteChannels.Worker_InviteChan <- worker_inviteChan
}
