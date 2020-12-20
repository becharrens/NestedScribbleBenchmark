package roles

import "NestedScribbleBenchmark/fannkuch/messages"
import "NestedScribbleBenchmark/fannkuch/channels/fannkuchrecursive"
import "NestedScribbleBenchmark/fannkuch/invitations"
import "NestedScribbleBenchmark/fannkuch/callbacks"
import "sync"

func FannkuchRecursive_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.FannkuchRecursive_RoleSetupChan, inviteChannels invitations.FannkuchRecursive_InviteSetupChan) {
	newworker_source_int := make(chan int, 1)
	newworker_invite_newworker := make(chan fannkuchrecursive.Worker_Chan, 1)
	newworker_invite_newworker_invitechan := make(chan invitations.FannkuchRecursive_Worker_InviteChan, 1)
	newworker_source_label := make(chan messages.Fannkuch_Label, 1)
	newworker_invite_source := make(chan fannkuchrecursive.Source_Chan, 1)
	newworker_invite_source_invitechan := make(chan invitations.FannkuchRecursive_Source_InviteChan, 1)
	worker_newworker_int := make(chan int, 1)
	worker_newworker_label := make(chan messages.Fannkuch_Label, 1)

	worker_chan := fannkuchrecursive.Worker_Chan{
		Label_To_NewWorker: worker_newworker_label,
		Int_To_NewWorker:   worker_newworker_int,
	}
	source_chan := fannkuchrecursive.Source_Chan{
		Label_From_NewWorker: newworker_source_label,
		Int_From_NewWorker:   newworker_source_int,
	}
	newworker_chan := fannkuchrecursive.NewWorker_Chan{
		Label_To_Source:   newworker_source_label,
		Label_From_Worker: worker_newworker_label,
		Int_To_Source:     newworker_source_int,
		Int_From_Worker:   worker_newworker_int,
	}

	worker_inviteChan := invitations.FannkuchRecursive_Worker_InviteChan{}
	source_inviteChan := invitations.FannkuchRecursive_Source_InviteChan{
		NewWorker_Invite_To_FannkuchRecursive_Source_InviteChan: newworker_invite_source_invitechan,
		NewWorker_Invite_To_FannkuchRecursive_Source:            newworker_invite_source,
	}
	newworker_inviteChan := invitations.FannkuchRecursive_NewWorker_InviteChan{
		Invite_Source_To_FannkuchRecursive_Source_InviteChan:    newworker_invite_source_invitechan,
		Invite_Source_To_FannkuchRecursive_Source:               newworker_invite_source,
		Invite_NewWorker_To_FannkuchRecursive_Worker_InviteChan: newworker_invite_newworker_invitechan,
		Invite_NewWorker_To_FannkuchRecursive_Worker:            newworker_invite_newworker,
	}

	roleChannels.Source_Chan <- source_chan
	roleChannels.Worker_Chan <- worker_chan

	inviteChannels.Source_InviteChan <- source_inviteChan
	inviteChannels.Worker_InviteChan <- worker_inviteChan

	wg.Add(1)

	newworker_env := callbacks.New_FannkuchRecursive_NewWorker_State()
	go FannkuchRecursive_NewWorker(wg, newworker_chan, newworker_inviteChan, newworker_env)
}
