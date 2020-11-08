package roles

import "NestedScribbleBenchmark/boundedfibonacci/channels/boundedfibonacci"
import "NestedScribbleBenchmark/boundedfibonacci/invitations"
import "NestedScribbleBenchmark/boundedfibonacci/callbacks"
import boundedfibonacci_2 "NestedScribbleBenchmark/boundedfibonacci/results/boundedfibonacci"
import "sync"

func BoundedFibonacci_F1(wg *sync.WaitGroup, roleChannels boundedfibonacci.F1_Chan, inviteChannels invitations.BoundedFibonacci_F1_InviteChan, env callbacks.BoundedFibonacci_F1_Env) boundedfibonacci_2.F1_Result {
	startfib1_msg := <-roleChannels.Start_StartFib1
	env.StartFib1_From_Start(startfib1_msg)

	boundedfib_f1_chan := <-inviteChannels.Start_Invite_To_BoundedFib_F1
	boundedfib_f1_inviteChan := <-inviteChannels.Start_Invite_To_BoundedFib_F1_InviteChan
	boundedfib_f1_env := env.To_BoundedFib_F1_Env()
	boundedfib_f1_result := BoundedFib_F1(wg, boundedfib_f1_chan, boundedfib_f1_inviteChan, boundedfib_f1_env)
	env.ResultFrom_BoundedFib_F1(boundedfib_f1_result)

	return env.Done()
} 