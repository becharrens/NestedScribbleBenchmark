package roles

import "NestedScribbleBenchmark/fannkuch/messages"
import "NestedScribbleBenchmark/fannkuch/channels/fannkuchrecursive"
import "NestedScribbleBenchmark/fannkuch/channels/fannkuch"
import "NestedScribbleBenchmark/fannkuch/invitations"
import "sync"

func Fannkuch_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.Fannkuch_RoleSetupChan, inviteChannels invitations.Fannkuch_InviteSetupChan) {
	worker_main_int := make(chan int, 1)
	worker_invite_worker := make(chan fannkuchrecursive.Worker_Chan, 1)
	worker_invite_worker_invitechan := make(chan invitations.FannkuchRecursive_Worker_InviteChan, 1)
	worker_main_label := make(chan messages.Fannkuch_Label, 1)
	worker_invite_main := make(chan fannkuchrecursive.Source_Chan, 1)
	worker_invite_main_invitechan := make(chan invitations.FannkuchRecursive_Source_InviteChan, 1)
	main_worker_int := make(chan int, 1)
	main_worker_label := make(chan messages.Fannkuch_Label, 1)

	worker_chan := fannkuch.Worker_Chan{
		Label_To_Main:   worker_main_label,
		Label_From_Main: main_worker_label,
		Int_To_Main:     worker_main_int,
		Int_From_Main:   main_worker_int,
	}
	main_chan := fannkuch.Main_Chan{
		Label_To_Worker:   main_worker_label,
		Label_From_Worker: worker_main_label,
		Int_To_Worker:     main_worker_int,
		Int_From_Worker:   worker_main_int,
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
