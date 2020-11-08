package roles

import "NestedScribbleBenchmark/boundedfibonacci/channels/boundedfibonacci"
import "NestedScribbleBenchmark/boundedfibonacci/invitations"
import "NestedScribbleBenchmark/boundedfibonacci/callbacks"
import boundedfibonacci_2 "NestedScribbleBenchmark/boundedfibonacci/results/boundedfibonacci"
import "sync"

func BoundedFibonacci_Start(wg *sync.WaitGroup, roleChannels boundedfibonacci.Start_Chan, inviteChannels invitations.BoundedFibonacci_Start_InviteChan, env callbacks.BoundedFibonacci_Start_Env) boundedfibonacci_2.Start_Result {
	startfib1_msg := env.StartFib1_To_F1()
	roleChannels.F1_StartFib1 <- startfib1_msg

	startfib2_msg := env.StartFib2_To_F2()
	roleChannels.F2_StartFib2 <- startfib2_msg

	env.BoundedFib_Setup()
	boundedfib_rolechan := invitations.BoundedFib_RoleSetupChan{
		Res_Chan: inviteChannels.Invite_Start_To_BoundedFib_Res,
		F1_Chan: inviteChannels.Invite_F1_To_BoundedFib_F1,
		F2_Chan: inviteChannels.Invite_F2_To_BoundedFib_F2,
	}
	boundedfib_invitechan := invitations.BoundedFib_InviteSetupChan{
		Res_InviteChan: inviteChannels.Invite_Start_To_BoundedFib_Res_InviteChan,
		F1_InviteChan: inviteChannels.Invite_F1_To_BoundedFib_F1_InviteChan,
		F2_InviteChan: inviteChannels.Invite_F2_To_BoundedFib_F2_InviteChan,
	}
	BoundedFib_SendCommChannels(wg, boundedfib_rolechan, boundedfib_invitechan)

	boundedfib_res_chan := <-inviteChannels.Invite_Start_To_BoundedFib_Res
	boundedfib_res_inviteChan := <-inviteChannels.Invite_Start_To_BoundedFib_Res_InviteChan
	boundedfib_res_env := env.To_BoundedFib_Res_Env()
	boundedfib_res_result := BoundedFib_Res(wg, boundedfib_res_chan, boundedfib_res_inviteChan, boundedfib_res_env)
	env.ResultFrom_BoundedFib_Res(boundedfib_res_result)

	return env.Done()
} 