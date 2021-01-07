package roles

import (
	"NestedScribbleBenchmark/primesieve_base/primesieve3/channels/primesieve"
	"NestedScribbleBenchmark/primesieve_base/primesieve3/invitations"
	"NestedScribbleBenchmark/primesieve_base/primesieve3/messages"
	"sync"
)

func PrimeSieve_Master(wg *sync.WaitGroup, n int, primes []int, roleChannels primesieve.Master_Chan, inviteChannels invitations.PrimeSieve_Master_InviteChan) []int {
	prime := primes[0]
	roleChannels.Label_To_Worker <- messages.FirstPrime
	roleChannels.Int_To_Worker <- prime

	roleChannels.Label_To_Worker <- messages.UBound
	roleChannels.Int_To_Worker <- n

	worker_choice := <-roleChannels.Label_From_Worker
	switch worker_choice {
	case messages.Prime:
		prime := <-roleChannels.Int_From_Worker
		primes = append(primes, prime)

		<-roleChannels.Label_From_Worker
		sieve_m_chan := <-inviteChannels.Worker_Invite_To_Sieve_M
		sieve_m_inviteChan := <-inviteChannels.Worker_Invite_To_Sieve_M_InviteChan
		primes := Sieve_M(wg, primes, sieve_m_chan, sieve_m_inviteChan)
		return primes
	case messages.Finish:
		return primes
	default:
		panic("Invalid choice was made")
	}
}
