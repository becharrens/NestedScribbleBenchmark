package roles

import "NestedScribbleBenchmark/fibonacci/messages"
import "NestedScribbleBenchmark/fibonacci/channels/fib"
import "NestedScribbleBenchmark/fibonacci/invitations"
import "NestedScribbleBenchmark/fibonacci/callbacks"
import "sync"

func Fib_F3(wg *sync.WaitGroup, roleChannels fib.F3_Chan, inviteChannels invitations.Fib_F3_InviteChan, env callbacks.Fib_F3_Env)  {
	defer wg.Done()
	<-roleChannels.Label_From_F1
	val := <-roleChannels.Int_From_F1
	env.Fib1_From_F1(val)

	<-roleChannels.Label_From_F2
	val_2 := <-roleChannels.Int_From_F2
	env.Fib2_From_F2(val_2)

	val_3 := env.NextFib_To_Res()
	roleChannels.Label_To_Res <- messages.NextFib
	roleChannels.Int_To_Res <- val_3

	env.Fib_Setup()
	roleChannels.Label_To_Res <- messages.Fib_Res_F2_F3
	roleChannels.Label_To_F2 <- messages.Fib_Res_F2_F3
	
	fib_rolechan := invitations.Fib_RoleSetupChan{
		Res_Chan: inviteChannels.Invite_Res_To_Fib_Res,
		F1_Chan: inviteChannels.Invite_F2_To_Fib_F1,
		F2_Chan: inviteChannels.Invite_F3_To_Fib_F2,
	}
	fib_invitechan := invitations.Fib_InviteSetupChan{
		Res_InviteChan: inviteChannels.Invite_Res_To_Fib_Res_InviteChan,
		F1_InviteChan: inviteChannels.Invite_F2_To_Fib_F1_InviteChan,
		F2_InviteChan: inviteChannels.Invite_F3_To_Fib_F2_InviteChan,
	}
	Fib_SendCommChannels(wg, fib_rolechan, fib_invitechan)

	fib_f2_chan := <-inviteChannels.Invite_F3_To_Fib_F2
	fib_f2_inviteChan := <-inviteChannels.Invite_F3_To_Fib_F2_InviteChan
	fib_f2_env := env.To_Fib_F2_Env()
	fib_f2_result := Fib_F2(wg, fib_f2_chan, fib_f2_inviteChan, fib_f2_env)
	env.ResultFrom_Fib_F2(fib_f2_result)

	env.Done()
	return 
} 