package roles

import (
	"NestedScribbleBenchmark/primesieve_base/primesieve3/channels/primesieve"
	"NestedScribbleBenchmark/primesieve_base/primesieve3/invitations"
	"NestedScribbleBenchmark/primesieve_base/primesieve3/messages"
	"sync"
)

func PrimeSieve_Worker(wg *sync.WaitGroup, roleChannels primesieve.Worker_Chan, inviteChannels invitations.PrimeSieve_Worker_InviteChan) {
	<-roleChannels.Label_From_Master
	prime := <-roleChannels.Int_From_Master

	<-roleChannels.Label_From_Master
	n := <-roleChannels.Int_From_Master
	primes := initPossiblePrimes(prime, n)

	if len(primes) > 0 {
		newPrime := primes[0]
		roleChannels.Label_To_Master <- messages.Prime
		roleChannels.Int_To_Master <- newPrime

		roleChannels.Label_To_Master <- messages.Sieve_Master_Worker

		sieve_rolechan := invitations.Sieve_RoleSetupChan{
			M_Chan:  inviteChannels.Invite_Master_To_Sieve_M,
			W1_Chan: inviteChannels.Invite_Worker_To_Sieve_W1,
		}
		sieve_invitechan := invitations.Sieve_InviteSetupChan{
			M_InviteChan:  inviteChannels.Invite_Master_To_Sieve_M_InviteChan,
			W1_InviteChan: inviteChannels.Invite_Worker_To_Sieve_W1_InviteChan,
		}
		Sieve_SendCommChannels(wg, sieve_rolechan, sieve_invitechan)

		sieve_w1_chan := <-inviteChannels.Invite_Worker_To_Sieve_W1
		sieve_w1_inviteChan := <-inviteChannels.Invite_Worker_To_Sieve_W1_InviteChan
		Sieve_W1(wg, newPrime, primes[1:], sieve_w1_chan, sieve_w1_inviteChan)
	} else {
		roleChannels.Label_To_Master <- messages.Finish
	}
}

func initPossiblePrimes(firstPrime, ubound int) []int {
	var result []int
	for i := 3; i <= ubound; i++ {
		if i%firstPrime > 0 {
			result = append(result, i)
		}
	}
	// fmt.Println("worker:", len(result), firstPrime, ubound)
	return result
}
