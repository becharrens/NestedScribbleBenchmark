package roles

import (
	"NestedScribbleBenchmark/boundedfibonacci/messages/boundedfib"
)
import boundedfib_2 "NestedScribbleBenchmark/boundedfibonacci/channels/boundedfib"
import "NestedScribbleBenchmark/boundedfibonacci/invitations"
import "NestedScribbleBenchmark/boundedfibonacci/callbacks"
import "sync"

func BoundedFib_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.BoundedFib_RoleSetupChan, inviteChannels invitations.BoundedFib_InviteSetupChan) {
	f3_f2_end := make(chan boundedfib.End, 1)
	f3_res_result := make(chan boundedfib.Result, 1)
	f3_invite_f3 := make(chan boundedfib_2.F2_Chan, 1)
	f3_invite_f3_invitechan := make(chan invitations.BoundedFib_F2_InviteChan, 1)
	f3_invite_f2 := make(chan boundedfib_2.F1_Chan, 1)
	f3_invite_f2_invitechan := make(chan invitations.BoundedFib_F1_InviteChan, 1)
	f3_invite_res := make(chan boundedfib_2.Res_Chan, 1)
	f3_invite_res_invitechan := make(chan invitations.BoundedFib_Res_InviteChan, 1)
	f2_f3_fib2 := make(chan boundedfib.Fib2, 1)
	f1_f3_fib1 := make(chan boundedfib.Fib1, 1)

	res_chan := boundedfib_2.Res_Chan{
		F3_Result: f3_res_result,
	}
	f3_chan := boundedfib_2.F3_Chan{
		Res_Result: f3_res_result,
		F2_Fib2:    f2_f3_fib2,
		F2_End:     f3_f2_end,
		F1_Fib1:    f1_f3_fib1,
	}
	f2_chan := boundedfib_2.F2_Chan{
		F3_Fib2: f2_f3_fib2,
		F3_End:  f3_f2_end,
	}
	f1_chan := boundedfib_2.F1_Chan{
		F3_Fib1: f1_f3_fib1,
	}

	res_inviteChan := invitations.BoundedFib_Res_InviteChan{
		F3_Invite_To_BoundedFib_Res_InviteChan: f3_invite_res_invitechan,
		F3_Invite_To_BoundedFib_Res:            f3_invite_res,
	}
	f3_inviteChan := invitations.BoundedFib_F3_InviteChan{
		Invite_Res_To_BoundedFib_Res_InviteChan: f3_invite_res_invitechan,
		Invite_Res_To_BoundedFib_Res:            f3_invite_res,
		Invite_F3_To_BoundedFib_F2_InviteChan:   f3_invite_f3_invitechan,
		Invite_F3_To_BoundedFib_F2:              f3_invite_f3,
		Invite_F2_To_BoundedFib_F1_InviteChan:   f3_invite_f2_invitechan,
		Invite_F2_To_BoundedFib_F1:              f3_invite_f2,
	}
	f2_inviteChan := invitations.BoundedFib_F2_InviteChan{
		F3_Invite_To_BoundedFib_F1_InviteChan: f3_invite_f2_invitechan,
		F3_Invite_To_BoundedFib_F1:            f3_invite_f2,
	}
	f1_inviteChan := invitations.BoundedFib_F1_InviteChan{}

	roleChannels.Res_Chan <- res_chan
	roleChannels.F1_Chan <- f1_chan
	roleChannels.F2_Chan <- f2_chan

	inviteChannels.Res_InviteChan <- res_inviteChan
	inviteChannels.F1_InviteChan <- f1_inviteChan
	inviteChannels.F2_InviteChan <- f2_inviteChan

	wg.Add(1)

	f3_env := callbacks.New_BoundedFib_F3_State()
	go BoundedFib_F3(wg, f3_chan, f3_inviteChan, f3_env)
}
