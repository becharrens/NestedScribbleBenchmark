package roles

import "NestedScribbleBenchmark/old_fibonacci/channels/fib"
import "NestedScribbleBenchmark/old_fibonacci/invitations"
import "NestedScribbleBenchmark/old_fibonacci/callbacks"
import fib_2 "NestedScribbleBenchmark/old_fibonacci/results/fib"
import "sync"

func Fib_Res(wg *sync.WaitGroup, roleChannels fib.Res_Chan, inviteChannels invitations.Fib_Res_InviteChan, env callbacks.Fib_Res_Env) fib_2.Res_Result {
	select {
	case fib_res_chan := <-inviteChannels.F3_Invite_To_Fib_Res:
		fib_res_inviteChan := <-inviteChannels.F3_Invite_To_Fib_Res_InviteChan
		fib_res_env := env.To_Fib_Res_Env()
		fib_res_result := Fib_Res(wg, fib_res_chan, fib_res_inviteChan, fib_res_env)
		env.ResultFrom_Fib_Res(fib_res_result)

		return env.Done()
	case result_msg := <-roleChannels.F3_Result:
		env.Result_From_F3(result_msg)

		return env.Done()
	}
}
