package roles

import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/messages"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/channels/boundedfib"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/invitations"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/callbacks"
import "sync"

func BoundedFib_F3(wg *sync.WaitGroup, roleChannels boundedfib.F3_Chan, inviteChannels invitations.BoundedFib_F3_InviteChan, env callbacks.BoundedFib_F3_Env) {
	defer wg.Done()
	<-roleChannels.Label_From_F1
	ubound := <-roleChannels.Int_From_F1
	idx := <-roleChannels.Int_From_F1
	val := <-roleChannels.Int_From_F1
	env.Fib1_From_F1(ubound, idx, val)

	<-roleChannels.Label_From_F2
	// idx_2 := <-roleChannels.Int_From_F2
	val_2 := <-roleChannels.Int_From_F2
	env.Fib2_From_F2(val_2)
	// env.Fib2_From_F2(idx_2, val_2)

	f3_choice := env.F3_Choice()
	switch f3_choice {
	case callbacks.BoundedFib_F3_BoundedFib:
		env.BoundedFib_Setup()
		roleChannels.Label_To_Res <- messages.BoundedFib_Res_F2_F3
		roleChannels.Label_To_F2 <- messages.BoundedFib_Res_F2_F3

		boundedfib_rolechan := invitations.BoundedFib_RoleSetupChan{
			Res_Chan: inviteChannels.Invite_Res_To_BoundedFib_Res,
			F1_Chan:  inviteChannels.Invite_F2_To_BoundedFib_F1,
			F2_Chan:  inviteChannels.Invite_F3_To_BoundedFib_F2,
		}
		boundedfib_invitechan := invitations.BoundedFib_InviteSetupChan{
			Res_InviteChan: inviteChannels.Invite_Res_To_BoundedFib_Res_InviteChan,
			F2_InviteChan:  inviteChannels.Invite_F3_To_BoundedFib_F2_InviteChan,
		}
		BoundedFib_SendCommChannels(wg, boundedfib_rolechan, boundedfib_invitechan)

		boundedfib_f2_chan := <-inviteChannels.Invite_F3_To_BoundedFib_F2
		boundedfib_f2_inviteChan := <-inviteChannels.Invite_F3_To_BoundedFib_F2_InviteChan
		boundedfib_f2_env := env.To_BoundedFib_F2_Env()
		boundedfib_f2_result := BoundedFib_F2(wg, boundedfib_f2_chan, boundedfib_f2_inviteChan, boundedfib_f2_env)
		env.ResultFrom_BoundedFib_F2(boundedfib_f2_result)

		env.Done()
		return
	case callbacks.BoundedFib_F3_Result:
		fib := env.Result_To_Res()
		roleChannels.Label_To_Res <- messages.Result
		roleChannels.Int_To_Res <- fib

		env.End_To_F2()
		roleChannels.Label_To_F2 <- messages.End

		env.Done()
		return
	default:
		panic("Invalid choice was made")
	}
}
