package roles

import "ScribbleBenchmark/primesieve/channels/sieve"
import "ScribbleBenchmark/primesieve/invitations"
import "ScribbleBenchmark/primesieve/callbacks"
import "sync"

func Sieve_W2(wg *sync.WaitGroup, roleChannels sieve.W2_Chan, inviteChannels invitations.Sieve_W2_InviteChan, env callbacks.Sieve_W2_Env) {
	defer wg.Done()
	filterprime_msg := <-roleChannels.W1_FilterPrime
	env.FilterPrime_From_W1(filterprime_msg)

	sieve_sendnums_r_chan := <-inviteChannels.W1_Invite_To_Sieve_SendNums_R
	sieve_sendnums_r_inviteChan := <-inviteChannels.W1_Invite_To_Sieve_SendNums_R_InviteChan
	sieve_sendnums_r_env := env.To_Sieve_SendNums_R_Env()
	sieve_sendnums_r_result := Sieve_SendNums_R(wg, sieve_sendnums_r_chan, sieve_sendnums_r_inviteChan, sieve_sendnums_r_env)
	env.ResultFrom_Sieve_SendNums_R(sieve_sendnums_r_result)

	w2_choice := env.W2_Choice()
	switch w2_choice {
	case callbacks.Sieve_W2_Prime:
		prime_msg := env.Prime_To_M()
		roleChannels.M_Prime <- prime_msg

		env.Sieve_Setup()
		sieve_rolechan := invitations.Sieve_RoleSetupChan{
			M_Chan:  inviteChannels.Invite_M_To_Sieve_M,
			W1_Chan: inviteChannels.Invite_W2_To_Sieve_W1,
		}
		sieve_invitechan := invitations.Sieve_InviteSetupChan{
			M_InviteChan:  inviteChannels.Invite_M_To_Sieve_M_InviteChan,
			W1_InviteChan: inviteChannels.Invite_W2_To_Sieve_W1_InviteChan,
		}
		Sieve_SendCommChannels(wg, sieve_rolechan, sieve_invitechan)

		sieve_w1_chan := <-inviteChannels.Invite_W2_To_Sieve_W1
		sieve_w1_inviteChan := <-inviteChannels.Invite_W2_To_Sieve_W1_InviteChan
		sieve_w1_env := env.To_Sieve_W1_Env()
		sieve_w1_result := Sieve_W1(wg, sieve_w1_chan, sieve_w1_inviteChan, sieve_w1_env)
		env.ResultFrom_Sieve_W1(sieve_w1_result)

		env.Done()
		return
	case callbacks.Sieve_W2_Finish:
		finish_msg := env.Finish_To_M()
		roleChannels.M_Finish <- finish_msg

		env.Done()
		return
	default:
		panic("Invalid choice was made")
	}
}
