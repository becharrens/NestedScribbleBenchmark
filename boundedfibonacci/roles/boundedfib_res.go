package roles

import "NestedScribbleBenchmark/boundedfibonacci/channels/boundedfib"
import "NestedScribbleBenchmark/boundedfibonacci/invitations"
import "NestedScribbleBenchmark/boundedfibonacci/callbacks"
import boundedfib_2 "NestedScribbleBenchmark/boundedfibonacci/results/boundedfib"
import "sync"

func BoundedFib_Res(wg *sync.WaitGroup, roleChannels boundedfib.Res_Chan, inviteChannels invitations.BoundedFib_Res_InviteChan, env callbacks.BoundedFib_Res_Env) boundedfib_2.Res_Result {
	select {
		case boundedfib_res_chan := <-inviteChannels.F3_Invite_To_BoundedFib_Res:
			boundedfib_res_inviteChan := <-inviteChannels.F3_Invite_To_BoundedFib_Res_InviteChan
			boundedfib_res_env := env.To_BoundedFib_Res_Env()
			boundedfib_res_result := BoundedFib_Res(wg, boundedfib_res_chan, boundedfib_res_inviteChan, boundedfib_res_env)
			env.ResultFrom_BoundedFib_Res(boundedfib_res_result)

			return env.Done()
		case result_msg := <-roleChannels.F3_Result:
			env.Result_From_F3(result_msg)

			return env.Done()
	}
} 