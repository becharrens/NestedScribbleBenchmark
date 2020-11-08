package roles

import "NestedScribbleBenchmark/boundedfibonacci/channels/boundedfib"
import "NestedScribbleBenchmark/boundedfibonacci/invitations"
import "NestedScribbleBenchmark/boundedfibonacci/callbacks"
import boundedfib_2 "NestedScribbleBenchmark/boundedfibonacci/results/boundedfib"
import "sync"

func BoundedFib_F2(wg *sync.WaitGroup, roleChannels boundedfib.F2_Chan, inviteChannels invitations.BoundedFib_F2_InviteChan, env callbacks.BoundedFib_F2_Env) boundedfib_2.F2_Result {
	fib2_msg := env.Fib2_To_F3()
	roleChannels.F3_Fib2 <- fib2_msg

	select {
		case end_msg := <-roleChannels.F3_End:
			env.End_From_F3(end_msg)

			return env.Done()
		case boundedfib_f1_chan := <-inviteChannels.F3_Invite_To_BoundedFib_F1:
			boundedfib_f1_inviteChan := <-inviteChannels.F3_Invite_To_BoundedFib_F1_InviteChan
			boundedfib_f1_env := env.To_BoundedFib_F1_Env()
			boundedfib_f1_result := BoundedFib_F1(wg, boundedfib_f1_chan, boundedfib_f1_inviteChan, boundedfib_f1_env)
			env.ResultFrom_BoundedFib_F1(boundedfib_f1_result)

			return env.Done()
	}
} 