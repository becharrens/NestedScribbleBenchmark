package roles

import "NestedScribbleBenchmark/old_fibonacci/channels/fib"
import "NestedScribbleBenchmark/old_fibonacci/invitations"
import "NestedScribbleBenchmark/old_fibonacci/callbacks"
import "sync"

func Fib_F3(wg *sync.WaitGroup, roleChannels fib.F3_Chan, inviteChannels invitations.Fib_F3_InviteChan, env callbacks.Fib_F3_Env) {
	defer wg.Done()
	fib1_msg := <-roleChannels.F1_Fib1
	env.Fib1_From_F1(fib1_msg)

	fib2_msg := <-roleChannels.F2_Fib2
	env.Fib2_From_F2(fib2_msg)

	f3_choice := env.F3_Choice()
	switch f3_choice {
	case callbacks.Fib_F3_Fib:
		env.Fib_Setup()
		fib_rolechan := invitations.Fib_RoleSetupChan{
			Res_Chan: inviteChannels.Invite_Res_To_Fib_Res,
			F1_Chan:  inviteChannels.Invite_F2_To_Fib_F1,
			F2_Chan:  inviteChannels.Invite_F3_To_Fib_F2,
		}
		fib_invitechan := invitations.Fib_InviteSetupChan{
			Res_InviteChan: inviteChannels.Invite_Res_To_Fib_Res_InviteChan,
			F1_InviteChan:  inviteChannels.Invite_F2_To_Fib_F1_InviteChan,
			F2_InviteChan:  inviteChannels.Invite_F3_To_Fib_F2_InviteChan,
		}
		Fib_SendCommChannels(wg, fib_rolechan, fib_invitechan)

		fib_f2_chan := <-inviteChannels.Invite_F3_To_Fib_F2
		fib_f2_inviteChan := <-inviteChannels.Invite_F3_To_Fib_F2_InviteChan
		fib_f2_env := env.To_Fib_F2_Env()
		fib_f2_result := Fib_F2(wg, fib_f2_chan, fib_f2_inviteChan, fib_f2_env)
		env.ResultFrom_Fib_F2(fib_f2_result)

		env.Done()
		return
	case callbacks.Fib_F3_Result:
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
