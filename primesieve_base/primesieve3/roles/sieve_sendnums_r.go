package roles

import (
	"NestedScribbleBenchmark/primesieve_base/primesieve3/messages"
)
import "NestedScribbleBenchmark/primesieve_base/primesieve3/channels/sieve_sendnums"
import "NestedScribbleBenchmark/primesieve_base/primesieve3/invitations"

import "sync"

func Sieve_SendNums_R(wg *sync.WaitGroup, filterPrime int, roleChannels sieve_sendnums.R_Chan, inviteChannels invitations.Sieve_SendNums_R_InviteChan) []int {
	var primes []int
	s_choice_2 := <-roleChannels.Label_From_S
	switch s_choice_2 {
	case messages.Num:
		num := <-roleChannels.Int_From_S
		primes = append(primes, num)
	SEND:
		for {
			s_choice := <-roleChannels.Label_From_S
			switch s_choice {
			case messages.Num:
				num2 := <-roleChannels.Int_From_S
				if num2%filterPrime != 0 {
					primes = append(primes, num2)
				}

				continue SEND
			case messages.End:
				return primes
			default:
				panic("Invalid choice was made")
			}
		}
	case messages.End:
		return primes
	default:
		panic("Invalid choice was made")
	}
}
