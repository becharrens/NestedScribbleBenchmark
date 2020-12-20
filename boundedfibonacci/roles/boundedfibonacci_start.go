package roles

import "NestedScribbleBenchmark/boundedfibonacci/messages"
import "NestedScribbleBenchmark/boundedfibonacci/channels/boundedfibonacci"
import "NestedScribbleBenchmark/boundedfibonacci/invitations"
import "NestedScribbleBenchmark/boundedfibonacci/callbacks"
import boundedfibonacci_2 "NestedScribbleBenchmark/boundedfibonacci/results/boundedfibonacci"
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
	boundedfib_rolechan := invitations.BoundedFib_RoleSetupChan{
		Res_Chan: inviteChannels.Invite_Start_To_BoundedFib_Res,
		F1_Chan:  inviteChannels.Invite_F1_To_BoundedFib_F1,
		F2_Chan:  inviteChannels.Invite_F2_To_BoundedFib_F2,
	}
	boundedfib_invitechan := invitations.BoundedFib_InviteSetupChan{
		Res_InviteChan: inviteChannels.Invite_Start_To_BoundedFib_Res_InviteChan,
		F1_InviteChan:  inviteChannels.Invite_F1_To_BoundedFib_F1_InviteChan,
		F2_InviteChan:  inviteChannels.Invite_F2_To_BoundedFib_F2_InviteChan,
	}
	BoundedFib_SendCommChannels(wg, boundedfib_rolechan, boundedfib_invitechan)

	boundedfib_res_chan := <-inviteChannels.Invite_Start_To_BoundedFib_Res
	boundedfib_res_inviteChan := <-inviteChannels.Invite_Start_To_BoundedFib_Res_InviteChan
	boundedfib_res_env := env.To_BoundedFib_Res_Env()
	boundedfib_res_result := BoundedFib_Res(wg, boundedfib_res_chan, boundedfib_res_inviteChan, boundedfib_res_env)
	env.ResultFrom_BoundedFib_Res(boundedfib_res_result)

	return env.Done()
}
