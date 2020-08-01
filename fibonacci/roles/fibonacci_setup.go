package roles

import "NestedScribbleBenchmark/fibonacci/messages/fibonacci"
import fibonacci_2 "NestedScribbleBenchmark/fibonacci/channels/fibonacci"
import "NestedScribbleBenchmark/fibonacci/channels/fib"
import "NestedScribbleBenchmark/fibonacci/invitations"
import "sync"

func Fibonacci_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.Fibonacci_RoleSetupChan, inviteChannels invitations.Fibonacci_InviteSetupChan) {
	start_invite_f2 := make(chan fib.F2_Chan, 1)
	start_invite_f2_invitechan := make(chan invitations.Fib_F2_InviteChan, 1)
	start_invite_f1 := make(chan fib.F1_Chan, 1)
	start_invite_f1_invitechan := make(chan invitations.Fib_F1_InviteChan, 1)
	start_invite_start := make(chan fib.Res_Chan, 1)
	start_invite_start_invitechan := make(chan invitations.Fib_Res_InviteChan, 1)
	start_f2_startfib2 := make(chan fibonacci.StartFib2, 1)
	start_f1_startfib1 := make(chan fibonacci.StartFib1, 1)

	start_chan := fibonacci_2.Start_Chan{
		F2_StartFib2: start_f2_startfib2,
		F1_StartFib1: start_f1_startfib1,
	}
	f2_chan := fibonacci_2.F2_Chan{
		Start_StartFib2: start_f2_startfib2,
	}
	f1_chan := fibonacci_2.F1_Chan{
		Start_StartFib1: start_f1_startfib1,
	}

	start_inviteChan := invitations.Fibonacci_Start_InviteChan{
		Invite_Start_To_Fib_Res_InviteChan: start_invite_start_invitechan,
		Invite_Start_To_Fib_Res:            start_invite_start,
		Invite_F2_To_Fib_F2_InviteChan:     start_invite_f2_invitechan,
		Invite_F2_To_Fib_F2:                start_invite_f2,
		Invite_F1_To_Fib_F1_InviteChan:     start_invite_f1_invitechan,
		Invite_F1_To_Fib_F1:                start_invite_f1,
	}
	f2_inviteChan := invitations.Fibonacci_F2_InviteChan{
		Start_Invite_To_Fib_F2_InviteChan: start_invite_f2_invitechan,
		Start_Invite_To_Fib_F2:            start_invite_f2,
	}
	f1_inviteChan := invitations.Fibonacci_F1_InviteChan{
		Start_Invite_To_Fib_F1_InviteChan: start_invite_f1_invitechan,
		Start_Invite_To_Fib_F1:            start_invite_f1,
	}

	roleChannels.Start_Chan <- start_chan
	roleChannels.F1_Chan <- f1_chan
	roleChannels.F2_Chan <- f2_chan

	inviteChannels.Start_InviteChan <- start_inviteChan
	inviteChannels.F1_InviteChan <- f1_inviteChan
	inviteChannels.F2_InviteChan <- f2_inviteChan
}
