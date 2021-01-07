package roles

import (
	"NestedScribbleBenchmark/primesieve_base/primesieve3/channels/sieve"
	"NestedScribbleBenchmark/primesieve_base/primesieve3/invitations"
	"NestedScribbleBenchmark/primesieve_base/primesieve3/messages"
	"sync"
)

func Sieve_W2(wg *sync.WaitGroup, roleChannels sieve.W2_Chan, inviteChannels invitations.Sieve_W2_InviteChan) {
	defer wg.Done()
	<-roleChannels.Label_From_W1
	filterPrime := <-roleChannels.Int_From_W1

	<-roleChannels.Label_From_W1
	sieve_sendnums_r_chan := <-inviteChannels.W1_Invite_To_Sieve_SendNums_R
	sieve_sendnums_r_inviteChan := <-inviteChannels.W1_Invite_To_Sieve_SendNums_R_InviteChan
	primeCandidates := Sieve_SendNums_R(wg, filterPrime, sieve_sendnums_r_chan, sieve_sendnums_r_inviteChan)

	if len(primeCandidates) > 0 {
		newPrime := primeCandidates[0]
		roleChannels.Label_To_M <- messages.Prime
		roleChannels.Int_To_M <- newPrime

		roleChannels.Label_To_M <- messages.Sieve_M_W2

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
		Sieve_W1(wg, newPrime, primeCandidates[1:], sieve_w1_chan, sieve_w1_inviteChan)
	} else {
		roleChannels.Label_To_M <- messages.Finish
	}
}
