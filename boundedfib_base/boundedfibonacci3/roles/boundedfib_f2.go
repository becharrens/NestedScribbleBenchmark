package roles

import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/messages"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/channels/boundedfib"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/invitations"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/callbacks"
import boundedfib_2 "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/results/boundedfib"
import "sync"

func BoundedFib_F2(wg *sync.WaitGroup, roleChannels boundedfib.F2_Chan, inviteChannels invitations.BoundedFib_F2_InviteChan, env callbacks.BoundedFib_F2_Env) boundedfib_2.F2_Result {
	// idx, val := env.Fib2_To_F3()
	val := env.Fib2_To_F3()
	roleChannels.Label_To_F3 <- messages.Fib2
	// roleChannels.Int_To_F3 <- idx
	roleChannels.Int_To_F3 <- val

	f3_choice := <-roleChannels.Label_From_F3
	switch f3_choice {
	case messages.End:
		env.End_From_F3()

		return env.Done()
	case messages.BoundedFib_Res_F2_F3:
		boundedfib_f1_chan := <-inviteChannels.F3_Invite_To_BoundedFib_F1
		boundedfib_f1_env := env.To_BoundedFib_F1_Env()
		boundedfib_f1_result := BoundedFib_F1(wg, boundedfib_f1_chan, boundedfib_f1_env)
		env.ResultFrom_BoundedFib_F1(boundedfib_f1_result)

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
}
