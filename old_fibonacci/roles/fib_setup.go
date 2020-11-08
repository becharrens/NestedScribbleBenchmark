package roles

import "NestedScribbleBenchmark/old_fibonacci/messages/fib"
import fib_2 "NestedScribbleBenchmark/old_fibonacci/channels/fib"
import "NestedScribbleBenchmark/old_fibonacci/invitations"
import "NestedScribbleBenchmark/old_fibonacci/callbacks"
import "sync"

func Fib_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.Fib_RoleSetupChan, inviteChannels invitations.Fib_InviteSetupChan) {
	f3_f2_end := make(chan fib.End, 1)
	f3_res_result := make(chan fib.Result, 1)
	f3_invite_f3 := make(chan fib_2.F2_Chan, 1)
	f3_invite_f3_invitechan := make(chan invitations.Fib_F2_InviteChan, 1)
	f3_invite_f2 := make(chan fib_2.F1_Chan, 1)
	f3_invite_f2_invitechan := make(chan invitations.Fib_F1_InviteChan, 1)
	f3_invite_res := make(chan fib_2.Res_Chan, 1)
	f3_invite_res_invitechan := make(chan invitations.Fib_Res_InviteChan, 1)
	f2_f3_fib2 := make(chan fib.Fib2, 1)
	f1_f3_fib1 := make(chan fib.Fib1, 1)

	res_chan := fib_2.Res_Chan{
		F3_Result: f3_res_result,
	}
	f3_chan := fib_2.F3_Chan{
		Res_Result: f3_res_result,
		F2_Fib2:    f2_f3_fib2,
		F2_End:     f3_f2_end,
		F1_Fib1:    f1_f3_fib1,
	}
	f2_chan := fib_2.F2_Chan{
		F3_Fib2: f2_f3_fib2,
		F3_End:  f3_f2_end,
	}
	f1_chan := fib_2.F1_Chan{
		F3_Fib1: f1_f3_fib1,
	}

	res_inviteChan := invitations.Fib_Res_InviteChan{
		F3_Invite_To_Fib_Res_InviteChan: f3_invite_res_invitechan,
		F3_Invite_To_Fib_Res:            f3_invite_res,
	}
	f3_inviteChan := invitations.Fib_F3_InviteChan{
		Invite_Res_To_Fib_Res_InviteChan: f3_invite_res_invitechan,
		Invite_Res_To_Fib_Res:            f3_invite_res,
		Invite_F3_To_Fib_F2_InviteChan:   f3_invite_f3_invitechan,
		Invite_F3_To_Fib_F2:              f3_invite_f3,
		Invite_F2_To_Fib_F1_InviteChan:   f3_invite_f2_invitechan,
		Invite_F2_To_Fib_F1:              f3_invite_f2,
	}
	f2_inviteChan := invitations.Fib_F2_InviteChan{
		F3_Invite_To_Fib_F1_InviteChan: f3_invite_f2_invitechan,
		F3_Invite_To_Fib_F1:            f3_invite_f2,
	}
	f1_inviteChan := invitations.Fib_F1_InviteChan{}

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
