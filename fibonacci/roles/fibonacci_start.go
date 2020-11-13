package roles

import "NestedScribbleBenchmark/fibonacci/channels/fibonacci"
import "NestedScribbleBenchmark/fibonacci/invitations"
import "NestedScribbleBenchmark/fibonacci/callbacks"
import fibonacci_2 "NestedScribbleBenchmark/fibonacci/results/fibonacci"
import "sync"

func Fibonacci_Start(wg *sync.WaitGroup, roleChannels fibonacci.Start_Chan, inviteChannels invitations.Fibonacci_Start_InviteChan, env callbacks.Fibonacci_Start_Env) fibonacci_2.Start_Result {
	startfib1_msg := env.StartFib1_To_F1()
	roleChannels.F1_StartFib1 <- startfib1_msg

	startfib2_msg := env.StartFib2_To_F2()
	roleChannels.F2_StartFib2 <- startfib2_msg

	env.Fib_Setup()
	fib_rolechan := invitations.Fib_RoleSetupChan{
		Res_Chan: inviteChannels.Invite_Start_To_Fib_Res,
		F1_Chan: inviteChannels.Invite_F1_To_Fib_F1,
		F2_Chan: inviteChannels.Invite_F2_To_Fib_F2,
	}
	fib_invitechan := invitations.Fib_InviteSetupChan{
		Res_InviteChan: inviteChannels.Invite_Start_To_Fib_Res_InviteChan,
		F1_InviteChan: inviteChannels.Invite_F1_To_Fib_F1_InviteChan,
		F2_InviteChan: inviteChannels.Invite_F2_To_Fib_F2_InviteChan,
	}
	Fib_SendCommChannels(wg, fib_rolechan, fib_invitechan)

	fib_res_chan := <-inviteChannels.Invite_Start_To_Fib_Res
	fib_res_inviteChan := <-inviteChannels.Invite_Start_To_Fib_Res_InviteChan
	fib_res_env := env.To_Fib_Res_Env()
	fib_res_result := Fib_Res(wg, fib_res_chan, fib_res_inviteChan, fib_res_env)
	env.ResultFrom_Fib_Res(fib_res_result)

	return env.Done()
} 