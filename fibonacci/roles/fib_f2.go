package roles

import "NestedScribbleBenchmark/fibonacci/messages"
import "NestedScribbleBenchmark/fibonacci/channels/fib"
import "NestedScribbleBenchmark/fibonacci/invitations"
import "NestedScribbleBenchmark/fibonacci/callbacks"
import fib_2 "NestedScribbleBenchmark/fibonacci/results/fib"
import "sync"

func Fib_F2(wg *sync.WaitGroup, roleChannels fib.F2_Chan, inviteChannels invitations.Fib_F2_InviteChan, env callbacks.Fib_F2_Env) fib_2.F2_Result {
	val := env.Fib2_To_F3()
	roleChannels.Label_To_F3 <- messages.Fib2
	roleChannels.Int_To_F3 <- val

	<-roleChannels.Label_From_F3
	fib_f1_chan := <-inviteChannels.F3_Invite_To_Fib_F1
	fib_f1_inviteChan := <-inviteChannels.F3_Invite_To_Fib_F1_InviteChan
	fib_f1_env := env.To_Fib_F1_Env()
	fib_f1_result := Fib_F1(wg, fib_f1_chan, fib_f1_inviteChan, fib_f1_env)
	env.ResultFrom_Fib_F1(fib_f1_result)

	return env.Done()
}
