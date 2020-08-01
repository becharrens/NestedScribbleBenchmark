package roles

import "NestedScribbleBenchmark/primesieve/channels/sieve"
import "NestedScribbleBenchmark/primesieve/invitations"
import "NestedScribbleBenchmark/primesieve/callbacks"
import sieve_2 "NestedScribbleBenchmark/primesieve/results/sieve"
import "sync"

func Sieve_W1(wg *sync.WaitGroup, roleChannels sieve.W1_Chan, inviteChannels invitations.Sieve_W1_InviteChan, env callbacks.Sieve_W1_Env) sieve_2.W1_Result {
	filterprime_msg := env.FilterPrime_To_W2()
	roleChannels.W2_FilterPrime <- filterprime_msg

	env.Sieve_SendNums_Setup()
	sieve_sendnums_rolechan := invitations.Sieve_SendNums_RoleSetupChan{
		S_Chan: inviteChannels.Invite_W1_To_Sieve_SendNums_S,
		R_Chan: inviteChannels.Invite_W2_To_Sieve_SendNums_R,
	}
	sieve_sendnums_invitechan := invitations.Sieve_SendNums_InviteSetupChan{
		S_InviteChan: inviteChannels.Invite_W1_To_Sieve_SendNums_S_InviteChan,
		R_InviteChan: inviteChannels.Invite_W2_To_Sieve_SendNums_R_InviteChan,
	}
	Sieve_SendNums_SendCommChannels(wg, sieve_sendnums_rolechan, sieve_sendnums_invitechan)

	sieve_sendnums_s_chan := <-inviteChannels.Invite_W1_To_Sieve_SendNums_S
	sieve_sendnums_s_inviteChan := <-inviteChannels.Invite_W1_To_Sieve_SendNums_S_InviteChan
	sieve_sendnums_s_env := env.To_Sieve_SendNums_S_Env()
	sieve_sendnums_s_result := Sieve_SendNums_S(wg, sieve_sendnums_s_chan, sieve_sendnums_s_inviteChan, sieve_sendnums_s_env)
	env.ResultFrom_Sieve_SendNums_S(sieve_sendnums_s_result)

	return env.Done()
}
