package roles

import "NestedScribbleBenchmark/primesieve_base/primesieve3/messages"
import "NestedScribbleBenchmark/primesieve_base/primesieve3/channels/sieve"
import "NestedScribbleBenchmark/primesieve_base/primesieve3/invitations"

import "sync"

func Sieve_M(wg *sync.WaitGroup, primes []int, roleChannels sieve.M_Chan, inviteChannels invitations.Sieve_M_InviteChan) []int {
	w2_choice := <-roleChannels.Label_From_W2
	switch w2_choice {
	case messages.Prime:
		n := <-roleChannels.Int_From_W2
		primes = append(primes, n)

		<-roleChannels.Label_From_W2
		sieve_m_chan := <-inviteChannels.W2_Invite_To_Sieve_M
		sieve_m_inviteChan := <-inviteChannels.W2_Invite_To_Sieve_M_InviteChan
		primes := Sieve_M(wg, primes, sieve_m_chan, sieve_m_inviteChan)

		return primes
	case messages.Finish:
		return primes
	default:
		panic("Invalid choice was made")
	}
}
