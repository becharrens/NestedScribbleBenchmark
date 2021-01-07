package roles

import (
	"NestedScribbleBenchmark/primesieve_base/primesieve2/messages"
)
import "NestedScribbleBenchmark/primesieve_base/primesieve2/channels/sieve_sendnums"
import "NestedScribbleBenchmark/primesieve_base/primesieve2/invitations"
import "NestedScribbleBenchmark/primesieve_base/primesieve2/callbacks"
import "sync"

func Sieve_SendNums_R(wg *sync.WaitGroup, roleChannels sieve_sendnums.R_Chan, inviteChannels invitations.Sieve_SendNums_R_InviteChan, env callbacks.Sieve_SendNums_R_Env) []int {
	s_choice_2 := <-roleChannels.Label_From_S
	switch s_choice_2 {
	case messages.Num:
		n := <-roleChannels.Int_From_S
		env.Num_From_S(n)

	SEND:
		for {
			s_choice := <-roleChannels.Label_From_S
			switch s_choice {
			case messages.Num:
				n_2 := <-roleChannels.Int_From_S
				env.Num_From_S_2(n_2)

				continue SEND
			case messages.End:
				env.End_From_S()

				return env.Done()
			default:
				panic("Invalid choice was made")
			}
		}
	case messages.End:
		env.End_From_S_2()

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
}
