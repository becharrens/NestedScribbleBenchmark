package roles

import "NestedScribbleBenchmark/primesieve/messages"
import "NestedScribbleBenchmark/primesieve/channels/sieve_sendnums"
import "NestedScribbleBenchmark/primesieve/invitations"
import "NestedScribbleBenchmark/primesieve/callbacks"
import sieve_sendnums_2 "NestedScribbleBenchmark/primesieve/results/sieve_sendnums"
import "sync"

func Sieve_SendNums_R(wg *sync.WaitGroup, roleChannels sieve_sendnums.R_Chan, inviteChannels invitations.Sieve_SendNums_R_InviteChan, env callbacks.Sieve_SendNums_R_Env) sieve_sendnums_2.R_Result {
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
