package roles

import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci4/messages"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci4/channels/boundedfibonacci"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci4/invitations"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci4/callbacks"
import boundedfibonacci_2 "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci4/results/boundedfibonacci"
import "sync"

func BoundedFibonacci_Start(wg *sync.WaitGroup, roleChannels boundedfibonacci.Start_Chan, inviteChannels invitations.BoundedFibonacci_Start_InviteChan, env callbacks.BoundedFibonacci_Start_Env) boundedfibonacci_2.Start_Result {
	n, val := env.StartFib1_To_F1()
	roleChannels.Label_To_F1 <- messages.StartFib1
	roleChannels.Int_To_F1 <- n
	roleChannels.Int_To_F1 <- val

	n_2, val_2 := env.StartFib2_To_F2()
	roleChannels.Label_To_F2 <- messages.StartFib2
	roleChannels.Int_To_F2 <- n_2
	roleChannels.Int_To_F2 <- val_2

	env.BoundedFib_Setup()

	roleChannels.Label_To_F1 <- messages.BoundedFib_Start_F1_F2
	roleChannels.Label_To_F2 <- messages.BoundedFib_Start_F1_F2

	// boundedfib_res_chan := <-inviteChannels.Invite_Start_To_BoundedFib_Res
	// boundedfib_res_inviteChan := <-inviteChannels.Invite_Start_To_BoundedFib_Res_InviteChan
	boundedfib_res_chan, boundedfib_res_inviteChan := BoundedFib_SendCommChannels_Start(wg, inviteChannels)
	boundedfib_res_env := env.To_BoundedFib_Res_Env()
	boundedfib_res_result := BoundedFib_Res(wg, boundedfib_res_chan, boundedfib_res_inviteChan, boundedfib_res_env)
	env.ResultFrom_BoundedFib_Res(boundedfib_res_result)

	return env.Done()
}
