package roles

import "ScribbleBenchmark/fannkuch/messages/fannkuchrecursive"
import fannkuchrecursive_2 "ScribbleBenchmark/fannkuch/channels/fannkuchrecursive"
import "ScribbleBenchmark/fannkuch/invitations"
import "ScribbleBenchmark/fannkuch/callbacks"
import "sync"

func FannkuchRecursive_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.FannkuchRecursive_RoleSetupChan, inviteChannels invitations.FannkuchRecursive_InviteSetupChan)  {
	newworker_source_result_2 := make(chan fannkuchrecursive.Result, 1)
	newworker_source_result := make(chan fannkuchrecursive.Result, 1)
	newworker_invite_newworker := make(chan fannkuchrecursive_2.Worker_Chan, 1)
	newworker_invite_newworker_invitechan := make(chan invitations.FannkuchRecursive_Worker_InviteChan, 1)
	newworker_invite_source := make(chan fannkuchrecursive_2.Source_Chan, 1)
	newworker_invite_source_invitechan := make(chan invitations.FannkuchRecursive_Source_InviteChan, 1)
	worker_newworker_task := make(chan fannkuchrecursive.Task, 1)

	worker_chan := fannkuchrecursive_2.Worker_Chan{
		NewWorker_Task: worker_newworker_task,
	}
	source_chan := fannkuchrecursive_2.Source_Chan{
		NewWorker_Result_2: newworker_source_result_2,
		NewWorker_Result: newworker_source_result,
	}
	newworker_chan := fannkuchrecursive_2.NewWorker_Chan{
		Worker_Task: worker_newworker_task,
		Source_Result_2: newworker_source_result_2,
		Source_Result: newworker_source_result,
	}

	worker_inviteChan := invitations.FannkuchRecursive_Worker_InviteChan{

	}
	source_inviteChan := invitations.FannkuchRecursive_Source_InviteChan{
		NewWorker_Invite_To_FannkuchRecursive_Source_InviteChan: newworker_invite_source_invitechan,
		NewWorker_Invite_To_FannkuchRecursive_Source: newworker_invite_source,
	}
	newworker_inviteChan := invitations.FannkuchRecursive_NewWorker_InviteChan{
		Invite_Source_To_FannkuchRecursive_Source_InviteChan: newworker_invite_source_invitechan,
		Invite_Source_To_FannkuchRecursive_Source: newworker_invite_source,
		Invite_NewWorker_To_FannkuchRecursive_Worker_InviteChan: newworker_invite_newworker_invitechan,
		Invite_NewWorker_To_FannkuchRecursive_Worker: newworker_invite_newworker,
	}

	roleChannels.Source_Chan <- source_chan
	roleChannels.Worker_Chan <- worker_chan

	inviteChannels.Source_InviteChan <- source_inviteChan
	inviteChannels.Worker_InviteChan <- worker_inviteChan

	wg.Add(1)

	newworker_env := callbacks.New_FannkuchRecursive_NewWorker_State()
	go FannkuchRecursive_NewWorker(wg, newworker_chan, newworker_inviteChan, newworker_env)
} 