package roles

import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/messages"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/channels/boundedfib"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/invitations"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/callbacks"
import boundedfib_2 "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/results/boundedfib"
import "sync"

func BoundedFib_Res(wg *sync.WaitGroup, roleChannels boundedfib.Res_Chan, inviteChannels invitations.BoundedFib_Res_InviteChan, env callbacks.BoundedFib_Res_Env) boundedfib_2.Res_Result {
	f3_choice := <-roleChannels.Label_From_F3
	switch f3_choice {
	case messages.BoundedFib_Res_F2_F3:
		boundedfib_res_chan := <-inviteChannels.F3_Invite_To_BoundedFib_Res
		boundedfib_res_inviteChan := <-inviteChannels.F3_Invite_To_BoundedFib_Res_InviteChan
		boundedfib_res_env := env.To_BoundedFib_Res_Env()
		boundedfib_res_result := BoundedFib_Res(wg, boundedfib_res_chan, boundedfib_res_inviteChan, boundedfib_res_env)
		env.ResultFrom_BoundedFib_Res(boundedfib_res_result)

		return env.Done()
	case messages.Result:
		fib := <-roleChannels.Int_From_F3
		env.Result_From_F3(fib)

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
}
