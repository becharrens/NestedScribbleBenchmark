package roles

import "NestedScribbleBenchmark/fibonacci/channels/fibonacci"
import "NestedScribbleBenchmark/fibonacci/invitations"
import "NestedScribbleBenchmark/fibonacci/callbacks"
import fibonacci_2 "NestedScribbleBenchmark/fibonacci/results/fibonacci"
import "sync"

func Fibonacci_F2(wg *sync.WaitGroup, roleChannels fibonacci.F2_Chan, inviteChannels invitations.Fibonacci_F2_InviteChan, env callbacks.Fibonacci_F2_Env) fibonacci_2.F2_Result {
	startfib2_msg := <-roleChannels.Start_StartFib2
	env.StartFib2_From_Start(startfib2_msg)

	fib_f2_chan := <-inviteChannels.Start_Invite_To_Fib_F2
	fib_f2_inviteChan := <-inviteChannels.Start_Invite_To_Fib_F2_InviteChan
	fib_f2_env := env.To_Fib_F2_Env()
	fib_f2_result := Fib_F2(wg, fib_f2_chan, fib_f2_inviteChan, fib_f2_env)
	env.ResultFrom_Fib_F2(fib_f2_result)

	return env.Done()
}
