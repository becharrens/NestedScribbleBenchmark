package roles

import "NestedScribbleBenchmark/fibonacci_prev/channels/fibonacci"
import "NestedScribbleBenchmark/fibonacci_prev/invitations"
import "NestedScribbleBenchmark/fibonacci_prev/callbacks"
import fibonacci_2 "NestedScribbleBenchmark/fibonacci_prev/results/fibonacci"
import "sync"

func Fibonacci_F1(wg *sync.WaitGroup, roleChannels fibonacci.F1_Chan, inviteChannels invitations.Fibonacci_F1_InviteChan, env callbacks.Fibonacci_F1_Env) fibonacci_2.F1_Result {
	startfib1_msg := <-roleChannels.Start_StartFib1
	env.StartFib1_From_Start(startfib1_msg)

	fib_f1_chan := <-inviteChannels.Start_Invite_To_Fib_F1
	fib_f1_inviteChan := <-inviteChannels.Start_Invite_To_Fib_F1_InviteChan
	fib_f1_env := env.To_Fib_F1_Env()
	fib_f1_result := Fib_F1(wg, fib_f1_chan, fib_f1_inviteChan, fib_f1_env)
	env.ResultFrom_Fib_F1(fib_f1_result)

	return env.Done()
}
