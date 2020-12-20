package roles

import "NestedScribbleBenchmark/boundedfibonacci/messages"
import "NestedScribbleBenchmark/boundedfibonacci/channels/boundedfib"
import "NestedScribbleBenchmark/boundedfibonacci/invitations"
import "NestedScribbleBenchmark/boundedfibonacci/callbacks"
import boundedfib_2 "NestedScribbleBenchmark/boundedfibonacci/results/boundedfib"
import "sync"

func BoundedFib_F1(wg *sync.WaitGroup, roleChannels boundedfib.F1_Chan, inviteChannels invitations.BoundedFib_F1_InviteChan, env callbacks.BoundedFib_F1_Env) boundedfib_2.F1_Result {
	ubound, idx, val := env.Fib1_To_F3()
	roleChannels.Label_To_F3 <- messages.Fib1
	roleChannels.Int_To_F3 <- ubound
	roleChannels.Int_To_F3 <- idx
	roleChannels.Int_To_F3 <- val

	return env.Done()
}
