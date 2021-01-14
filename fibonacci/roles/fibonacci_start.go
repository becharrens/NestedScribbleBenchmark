package roles

import "NestedScribbleBenchmark/fibonacci/messages"
import "NestedScribbleBenchmark/fibonacci/channels/fibonacci"
import "NestedScribbleBenchmark/fibonacci/invitations"
import "NestedScribbleBenchmark/fibonacci/callbacks"
import fibonacci_2 "NestedScribbleBenchmark/fibonacci/results/fibonacci"
import "sync"

func Fibonacci_Start(wg *sync.WaitGroup, roleChannels fibonacci.Start_Chan, inviteChannels invitations.Fibonacci_Start_InviteChan, env callbacks.Fibonacci_Start_Env) fibonacci_2.Start_Result {
	val := env.StartFib1_To_F1()
	roleChannels.Label_To_F1 <- messages.StartFib1
	roleChannels.Int_To_F1 <- val

	val_2 := env.StartFib2_To_F2()
	roleChannels.Label_To_F2 <- messages.StartFib2
	roleChannels.Int_To_F2 <- val_2

	env.Fib_Setup()
	
	roleChannels.Label_To_F1 <- messages.Fib_Start_F1_F2
	roleChannels.Label_To_F2 <- messages.Fib_Start_F1_F2
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