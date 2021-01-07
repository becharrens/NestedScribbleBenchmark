package roles

import (
	"NestedScribbleBenchmark/primesieve_base/primesieve2/messages"
)
import "NestedScribbleBenchmark/primesieve_base/primesieve2/channels/sieve"
import "NestedScribbleBenchmark/primesieve_base/primesieve2/invitations"
import "NestedScribbleBenchmark/primesieve_base/primesieve2/callbacks"

import "sync"

func Sieve_W1(wg *sync.WaitGroup, roleChannels sieve.W1_Chan, inviteChannels invitations.Sieve_W1_InviteChan, env callbacks.Sieve_W1_Env) {
	int := env.FilterPrime_To_W2()
	roleChannels.Label_To_W2 <- messages.FilterPrime
	roleChannels.Int_To_W2 <- int

	env.Sieve_SendNums_Setup()

	roleChannels.Label_To_W2 <- messages.Sieve_SendNums_W1_W2
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
	Sieve_SendNums_S(wg, sieve_sendnums_s_chan, sieve_sendnums_s_inviteChan, sieve_sendnums_s_env)

	env.Done()
	return
}
