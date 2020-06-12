package roles

import "ScribbleBenchmark/fibonacci/channels/fib"
import "ScribbleBenchmark/fibonacci/invitations"
import "ScribbleBenchmark/fibonacci/callbacks"
import fib_2 "ScribbleBenchmark/fibonacci/results/fib"
import "sync"

func Fib_F1(wg *sync.WaitGroup, roleChannels fib.F1_Chan, inviteChannels invitations.Fib_F1_InviteChan, env callbacks.Fib_F1_Env) fib_2.F1_Result {
	fib1_msg := env.Fib1_To_F3()
	roleChannels.F3_Fib1 <- fib1_msg

	return env.Done()
} 