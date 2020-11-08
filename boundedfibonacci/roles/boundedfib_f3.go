package roles

import "NestedScribbleBenchmark/boundedfibonacci/channels/boundedfib"
import "NestedScribbleBenchmark/boundedfibonacci/invitations"
import "NestedScribbleBenchmark/boundedfibonacci/callbacks"
import "sync"

func BoundedFib_F3(wg *sync.WaitGroup, roleChannels boundedfib.F3_Chan, inviteChannels invitations.BoundedFib_F3_InviteChan, env callbacks.BoundedFib_F3_Env)  {
	defer wg.Done()
	fib1_msg := <-roleChannels.F1_Fib1
	env.Fib1_From_F1(fib1_msg)

	fib2_msg := <-roleChannels.F2_Fib2
	env.Fib2_From_F2(fib2_msg)

	f3_choice := env.F3_Choice()
	switch f3_choice {
	case callbacks.BoundedFib_F3_BoundedFib:
		env.BoundedFib_Setup()
		boundedfib_rolechan := invitations.BoundedFib_RoleSetupChan{
			Res_Chan: inviteChannels.Invite_Res_To_BoundedFib_Res,
			F1_Chan: inviteChannels.Invite_F2_To_BoundedFib_F1,
			F2_Chan: inviteChannels.Invite_F3_To_BoundedFib_F2,
		}
		boundedfib_invitechan := invitations.BoundedFib_InviteSetupChan{
			Res_InviteChan: inviteChannels.Invite_Res_To_BoundedFib_Res_InviteChan,
			F1_InviteChan: inviteChannels.Invite_F2_To_BoundedFib_F1_InviteChan,
			F2_InviteChan: inviteChannels.Invite_F3_To_BoundedFib_F2_InviteChan,
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
		result_msg := env.Result_To_Res()
		roleChannels.Res_Result <- result_msg

		end_msg := env.End_To_F2()
		roleChannels.F2_End <- end_msg

		env.Done()
		return 
	default:
		panic("Invalid choice was made")
	}
} 