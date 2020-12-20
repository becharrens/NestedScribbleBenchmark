package roles

import "NestedScribbleBenchmark/boundedfibonacci/channels/boundedfibonacci"
import "NestedScribbleBenchmark/boundedfibonacci/invitations"
import "NestedScribbleBenchmark/boundedfibonacci/callbacks"
import boundedfibonacci_2 "NestedScribbleBenchmark/boundedfibonacci/results/boundedfibonacci"
import "sync"

func BoundedFibonacci_F1(wg *sync.WaitGroup, roleChannels boundedfibonacci.F1_Chan, inviteChannels invitations.BoundedFibonacci_F1_InviteChan, env callbacks.BoundedFibonacci_F1_Env) boundedfibonacci_2.F1_Result {
	<-roleChannels.Label_From_Start
	n := <-roleChannels.Int_From_Start
	val := <-roleChannels.Int_From_Start
	env.StartFib1_From_Start(n, val)

	<-roleChannels.Label_From_Start
	boundedfib_f1_chan := <-inviteChannels.Start_Invite_To_BoundedFib_F1
	boundedfib_f1_inviteChan := <-inviteChannels.Start_Invite_To_BoundedFib_F1_InviteChan
	boundedfib_f1_env := env.To_BoundedFib_F1_Env()
	boundedfib_f1_result := BoundedFib_F1(wg, boundedfib_f1_chan, boundedfib_f1_inviteChan, boundedfib_f1_env)
	env.ResultFrom_BoundedFib_F1(boundedfib_f1_result)

	return env.Done()
}
