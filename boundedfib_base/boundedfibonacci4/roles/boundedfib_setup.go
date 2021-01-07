package roles

import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci4/messages"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci4/channels/boundedfib"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci4/invitations"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci4/callbacks"
import "sync"

func BoundedFib_SendCommChannels_Start(wg *sync.WaitGroup, inviteChannels invitations.BoundedFibonacci_Start_InviteChan) (boundedfib.Res_Chan, invitations.BoundedFib_Res_InviteChan) {
	f3_res_int := make(chan int, 1)
	// f3_invite_f3 := make(chan boundedfib.F2_Chan, 1)
	// f3_invite_f3_invitechan := make(chan invitations.BoundedFib_F2_InviteChan, 1)
	f3_f2_label := make(chan messages.BoundedFibonacci_Label, 1)
	f3_invite_f2 := make(chan boundedfib.F1_Chan, 1)
	// f3_invite_f2_invitechan := make(chan invitations.BoundedFib_F1_InviteChan, 1)
	f3_res_label := make(chan messages.BoundedFibonacci_Label, 1)
	f3_invite_res := make(chan boundedfib.Res_Chan, 1)
	f3_invite_res_invitechan := make(chan invitations.BoundedFib_Res_InviteChan, 1)
	f2_f3_int := make(chan int, 1)
	f2_f3_label := make(chan messages.BoundedFibonacci_Label, 1)
	f1_f3_int := make(chan int, 1)
	f1_f3_label := make(chan messages.BoundedFibonacci_Label, 1)

	res_chan := boundedfib.Res_Chan{
		Label_From_F3: f3_res_label,
		Int_From_F3:   f3_res_int,
	}
	f3_chan := boundedfib.F3_Chan{
		Label_To_Res:  f3_res_label,
		Label_To_F2:   f3_f2_label,
		Label_From_F2: f2_f3_label,
		Label_From_F1: f1_f3_label,
		Int_To_Res:    f3_res_int,
		Int_From_F2:   f2_f3_int,
		Int_From_F1:   f1_f3_int,
	}
	f2_chan := boundedfib.F2_Chan{
		Label_To_F3:   f2_f3_label,
		Label_From_F3: f3_f2_label,
		Int_To_F3:     f2_f3_int,
	}
	f1_chan := boundedfib.F1_Chan{
		Label_To_F3: f1_f3_label,
		Int_To_F3:   f1_f3_int,
	}

	res_inviteChan := invitations.BoundedFib_Res_InviteChan{
		F3_Invite_To_BoundedFib_Res_InviteChan: f3_invite_res_invitechan,
		F3_Invite_To_BoundedFib_Res:            f3_invite_res,
	}
	f3_inviteChan := invitations.BoundedFib_F3_InviteChan{
		Invite_Res_To_BoundedFib_Res_InviteChan: f3_invite_res_invitechan,
		Invite_Res_To_BoundedFib_Res:            f3_invite_res,
		// Invite_F3_To_BoundedFib_F2_InviteChan:   f3_invite_f3_invitechan,
		// Invite_F3_To_BoundedFib_F2:              f3_invite_f3,
		// Invite_F2_To_BoundedFib_F1_InviteChan:   f3_invite_f2_invitechan,
		Invite_F2_To_BoundedFib_F1: f3_invite_f2,
	}
	f2_inviteChan := invitations.BoundedFib_F2_InviteChan{
		// F3_Invite_To_BoundedFib_F1_InviteChan: f3_invite_f2_invitechan,
		F3_Invite_To_BoundedFib_F1: f3_invite_f2,
	}
	// f1_inviteChan := invitations.BoundedFib_F1_InviteChan{}

	// Res_Chan <- res_chan
	inviteChannels.Invite_F1_To_BoundedFib_F1 <- f1_chan
	inviteChannels.Invite_F2_To_BoundedFib_F2 <- f2_chan

	// inviteChannels.Invite_F1_To_BoundedFib_F1_InviteChan <- f1_inviteChan
	inviteChannels.Invite_F2_To_BoundedFib_F2_InviteChan <- f2_inviteChan
	// inviteChannels.Res_InviteChan <- res_inviteChan
	// inviteChannels.F1_InviteChan <- f1_inviteChan
	// inviteChannels.F2_InviteChan <- f2_inviteChan

	wg.Add(1)

	f3_env := callbacks.New_BoundedFib_F3_State()
	go BoundedFib_F3(wg, f3_chan, f3_inviteChan, f3_env)
	return res_chan, res_inviteChan
}

func BoundedFib_SendCommChannels_F3(wg *sync.WaitGroup, inviteChannels invitations.BoundedFib_F3_InviteChan) (boundedfib.F2_Chan, invitations.BoundedFib_F2_InviteChan) {
	f3_res_int := make(chan int, 1)
	// f3_invite_f3 := make(chan boundedfib.F2_Chan, 1)
	// f3_invite_f3_invitechan := make(chan invitations.BoundedFib_F2_InviteChan, 1)
	f3_f2_label := make(chan messages.BoundedFibonacci_Label, 1)
	f3_invite_f2 := make(chan boundedfib.F1_Chan, 1)
	// f3_invite_f2_invitechan := make(chan invitations.BoundedFib_F1_InviteChan, 1)
	f3_res_label := make(chan messages.BoundedFibonacci_Label, 1)
	f3_invite_res := make(chan boundedfib.Res_Chan, 1)
	f3_invite_res_invitechan := make(chan invitations.BoundedFib_Res_InviteChan, 1)
	f2_f3_int := make(chan int, 1)
	f2_f3_label := make(chan messages.BoundedFibonacci_Label, 1)
	f1_f3_int := make(chan int, 1)
	f1_f3_label := make(chan messages.BoundedFibonacci_Label, 1)

	res_chan := boundedfib.Res_Chan{
		Label_From_F3: f3_res_label,
		Int_From_F3:   f3_res_int,
	}
	f3_chan := boundedfib.F3_Chan{
		Label_To_Res:  f3_res_label,
		Label_To_F2:   f3_f2_label,
		Label_From_F2: f2_f3_label,
		Label_From_F1: f1_f3_label,
		Int_To_Res:    f3_res_int,
		Int_From_F2:   f2_f3_int,
		Int_From_F1:   f1_f3_int,
	}
	f2_chan := boundedfib.F2_Chan{
		Label_To_F3:   f2_f3_label,
		Label_From_F3: f3_f2_label,
		Int_To_F3:     f2_f3_int,
	}
	f1_chan := boundedfib.F1_Chan{
		Label_To_F3: f1_f3_label,
		Int_To_F3:   f1_f3_int,
	}

	res_inviteChan := invitations.BoundedFib_Res_InviteChan{
		F3_Invite_To_BoundedFib_Res_InviteChan: f3_invite_res_invitechan,
		F3_Invite_To_BoundedFib_Res:            f3_invite_res,
	}
	f3_inviteChan := invitations.BoundedFib_F3_InviteChan{
		Invite_Res_To_BoundedFib_Res_InviteChan: f3_invite_res_invitechan,
		Invite_Res_To_BoundedFib_Res:            f3_invite_res,
		// Invite_F3_To_BoundedFib_F2_InviteChan:   f3_invite_f3_invitechan,
		// Invite_F3_To_BoundedFib_F2: f3_invite_f3,
		// Invite_F2_To_BoundedFib_F1_InviteChan:   f3_invite_f2_invitechan,
		Invite_F2_To_BoundedFib_F1: f3_invite_f2,
	}
	f2_inviteChan := invitations.BoundedFib_F2_InviteChan{
		// F3_Invite_To_BoundedFib_F1_InviteChan: f3_invite_f2_invitechan,
		F3_Invite_To_BoundedFib_F1: f3_invite_f2,
	}
	// f1_inviteChan := invitations.BoundedFib_F1_InviteChan{}

	inviteChannels.Invite_Res_To_BoundedFib_Res <- res_chan
	inviteChannels.Invite_F2_To_BoundedFib_F1 <- f1_chan
	// inviteChannels.F2_Chan <- f2_chan

	inviteChannels.Invite_Res_To_BoundedFib_Res_InviteChan <- res_inviteChan
	// inviteChannels.Invite_F2_To_BoundedFib_F1_InviteChan <- f1_inviteChan
	// inviteChannels.F2_InviteChan <- f2_inviteChan

	wg.Add(1)

	f3_env := callbacks.New_BoundedFib_F3_State()
	go BoundedFib_F3(wg, f3_chan, f3_inviteChan, f3_env)
	return f2_chan, f2_inviteChan
}
