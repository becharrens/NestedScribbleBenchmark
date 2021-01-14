package roles

import "NestedScribbleBenchmark/fibonacci/channels/fib"
import "NestedScribbleBenchmark/fibonacci/invitations"
import "NestedScribbleBenchmark/fibonacci/callbacks"
import fib_2 "NestedScribbleBenchmark/fibonacci/results/fib"
import "sync"

func Fib_Res(wg *sync.WaitGroup, roleChannels fib.Res_Chan, inviteChannels invitations.Fib_Res_InviteChan, env callbacks.Fib_Res_Env) fib_2.Res_Result {
	<-roleChannels.Label_From_F3
	val := <-roleChannels.Int_From_F3
	env.NextFib_From_F3(val)

	<-roleChannels.Label_From_F3
	fib_res_chan := <-inviteChannels.F3_Invite_To_Fib_Res
	fib_res_inviteChan := <-inviteChannels.F3_Invite_To_Fib_Res_InviteChan
	fib_res_env := env.To_Fib_Res_Env()
	fib_res_result := Fib_Res(wg, fib_res_chan, fib_res_inviteChan, fib_res_env)
	env.ResultFrom_Fib_Res(fib_res_result)

	return env.Done()
} 