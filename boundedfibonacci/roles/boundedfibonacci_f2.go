package roles

import "NestedScribbleBenchmark/boundedfibonacci/channels/boundedfibonacci"
import "NestedScribbleBenchmark/boundedfibonacci/invitations"
import "NestedScribbleBenchmark/boundedfibonacci/callbacks"
import boundedfibonacci_2 "NestedScribbleBenchmark/boundedfibonacci/results/boundedfibonacci"
import "sync"

func BoundedFibonacci_F2(wg *sync.WaitGroup, roleChannels boundedfibonacci.F2_Chan, inviteChannels invitations.BoundedFibonacci_F2_InviteChan, env callbacks.BoundedFibonacci_F2_Env) boundedfibonacci_2.F2_Result {
	<-roleChannels.Label_From_Start
	n := <-roleChannels.Int_From_Start
	val := <-roleChannels.Int_From_Start
	env.StartFib2_From_Start(n, val)

	<-roleChannels.Label_From_Start
	boundedfib_f2_chan := <-inviteChannels.Start_Invite_To_BoundedFib_F2
	boundedfib_f2_inviteChan := <-inviteChannels.Start_Invite_To_BoundedFib_F2_InviteChan
	boundedfib_f2_env := env.To_BoundedFib_F2_Env()
	boundedfib_f2_result := BoundedFib_F2(wg, boundedfib_f2_chan, boundedfib_f2_inviteChan, boundedfib_f2_env)
	env.ResultFrom_BoundedFib_F2(boundedfib_f2_result)

	return env.Done()
}
