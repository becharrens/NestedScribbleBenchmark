package roles

import "NestedScribbleBenchmark/fibonacci/messages"
import "NestedScribbleBenchmark/fibonacci/channels/fib"
import "NestedScribbleBenchmark/fibonacci/invitations"
import "NestedScribbleBenchmark/fibonacci/callbacks"
import "sync"

func Fib_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.Fib_RoleSetupChan, inviteChannels invitations.Fib_InviteSetupChan)  {
	f3_invite_f3 := make(chan fib.F2_Chan, 1)
	f3_invite_f3_invitechan := make(chan invitations.Fib_F2_InviteChan, 1)
	f3_f2_label := make(chan messages.Fibonacci_Label, 1)
	f3_invite_f2 := make(chan fib.F1_Chan, 1)
	f3_invite_f2_invitechan := make(chan invitations.Fib_F1_InviteChan, 1)
	f3_invite_res := make(chan fib.Res_Chan, 1)
	f3_invite_res_invitechan := make(chan invitations.Fib_Res_InviteChan, 1)
	f3_res_int := make(chan int, 1)
	f3_res_label := make(chan messages.Fibonacci_Label, 1)
	f2_f3_int := make(chan int, 1)
	f2_f3_label := make(chan messages.Fibonacci_Label, 1)
	f1_f3_int := make(chan int, 1)
	f1_f3_label := make(chan messages.Fibonacci_Label, 1)

	res_chan := fib.Res_Chan{
		Label_From_F3: f3_res_label,
		Int_From_F3: f3_res_int,
	}
	f3_chan := fib.F3_Chan{
		Label_To_Res: f3_res_label,
		Label_To_F2: f3_f2_label,
		Label_From_F2: f2_f3_label,
		Label_From_F1: f1_f3_label,
		Int_To_Res: f3_res_int,
		Int_From_F2: f2_f3_int,
		Int_From_F1: f1_f3_int,
	}
	f2_chan := fib.F2_Chan{
		Label_To_F3: f2_f3_label,
		Label_From_F3: f3_f2_label,
		Int_To_F3: f2_f3_int,
	}
	f1_chan := fib.F1_Chan{
		Label_To_F3: f1_f3_label,
		Int_To_F3: f1_f3_int,
	}

	res_inviteChan := invitations.Fib_Res_InviteChan{
		F3_Invite_To_Fib_Res_InviteChan: f3_invite_res_invitechan,
		F3_Invite_To_Fib_Res: f3_invite_res,
	}
	f3_inviteChan := invitations.Fib_F3_InviteChan{
		Invite_Res_To_Fib_Res_InviteChan: f3_invite_res_invitechan,
		Invite_Res_To_Fib_Res: f3_invite_res,
		Invite_F3_To_Fib_F2_InviteChan: f3_invite_f3_invitechan,
		Invite_F3_To_Fib_F2: f3_invite_f3,
		Invite_F2_To_Fib_F1_InviteChan: f3_invite_f2_invitechan,
		Invite_F2_To_Fib_F1: f3_invite_f2,
	}
	f2_inviteChan := invitations.Fib_F2_InviteChan{
		F3_Invite_To_Fib_F1_InviteChan: f3_invite_f2_invitechan,
		F3_Invite_To_Fib_F1: f3_invite_f2,
	}
	f1_inviteChan := invitations.Fib_F1_InviteChan{

	}

	roleChannels.Res_Chan <- res_chan
	roleChannels.F1_Chan <- f1_chan
	roleChannels.F2_Chan <- f2_chan

	inviteChannels.Res_InviteChan <- res_inviteChan
	inviteChannels.F1_InviteChan <- f1_inviteChan
	inviteChannels.F2_InviteChan <- f2_inviteChan

	wg.Add(1)

	f3_env := callbacks.New_Fib_F3_State()
	go Fib_F3(wg, f3_chan, f3_inviteChan, f3_env)
} 