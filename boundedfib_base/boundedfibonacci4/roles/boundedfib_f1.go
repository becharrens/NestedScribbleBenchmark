package roles

import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci4/messages"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci4/channels/boundedfib"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci4/callbacks"
import boundedfib_2 "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci4/results/boundedfib"
import "sync"

func BoundedFib_F1(wg *sync.WaitGroup, roleChannels boundedfib.F1_Chan, env callbacks.BoundedFib_F1_Env) boundedfib_2.F1_Result {
	ubound, idx, val := env.Fib1_To_F3()
	roleChannels.Label_To_F3 <- messages.Fib1
	roleChannels.Int_To_F3 <- ubound
	roleChannels.Int_To_F3 <- idx
	roleChannels.Int_To_F3 <- val

	return env.Done()
}
