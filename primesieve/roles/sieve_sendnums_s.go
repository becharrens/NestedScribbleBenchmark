package roles

import "NestedScribbleBenchmark/primesieve/messages"
import "NestedScribbleBenchmark/primesieve/channels/sieve_sendnums"
import "NestedScribbleBenchmark/primesieve/invitations"
import "NestedScribbleBenchmark/primesieve/callbacks"
import sieve_sendnums_2 "NestedScribbleBenchmark/primesieve/results/sieve_sendnums"
import "sync"

func Sieve_SendNums_S(wg *sync.WaitGroup, roleChannels sieve_sendnums.S_Chan, inviteChannels invitations.Sieve_SendNums_S_InviteChan, env callbacks.Sieve_SendNums_S_Env) sieve_sendnums_2.S_Result {
	s_choice_2 := env.S_Choice()
	switch s_choice_2 {
	case callbacks.Sieve_SendNums_S_Num:
		n := env.Num_To_R()
		roleChannels.Label_To_R <- messages.Num
		roleChannels.Int_To_R <- n

	SEND:
		for {
			s_choice := env.S_Choice_2()
			switch s_choice {
			case callbacks.Sieve_SendNums_S_Num_2:
				n_2 := env.Num_To_R_2()
				roleChannels.Label_To_R <- messages.Num
				roleChannels.Int_To_R <- n_2

				continue SEND
			case callbacks.Sieve_SendNums_S_End_2:
				env.End_To_R()
				roleChannels.Label_To_R <- messages.End

				return env.Done()
			default:
				panic("Invalid choice was made")
			}
		}
	case callbacks.Sieve_SendNums_S_End:
		env.End_To_R_2()
		roleChannels.Label_To_R <- messages.End

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
}
