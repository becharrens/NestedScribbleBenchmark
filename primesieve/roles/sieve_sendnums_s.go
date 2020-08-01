package roles

import "NestedScribbleBenchmark/primesieve/channels/sieve_sendnums"
import "NestedScribbleBenchmark/primesieve/invitations"
import "NestedScribbleBenchmark/primesieve/callbacks"
import sieve_sendnums_2 "NestedScribbleBenchmark/primesieve/results/sieve_sendnums"
import "sync"

func Sieve_SendNums_S(wg *sync.WaitGroup, roleChannels sieve_sendnums.S_Chan, inviteChannels invitations.Sieve_SendNums_S_InviteChan, env callbacks.Sieve_SendNums_S_Env) sieve_sendnums_2.S_Result {
	s_choice := env.S_Choice()
	switch s_choice {
	case callbacks.Sieve_SendNums_S_Num:
		num_msg := env.Num_To_R()
		roleChannels.R_Num <- num_msg

		sieve_sendnums_send_s_chan := <-inviteChannels.R_Invite_To_Sieve_SendNums_SEND_S
		sieve_sendnums_send_s_inviteChan := <-inviteChannels.R_Invite_To_Sieve_SendNums_SEND_S_InviteChan
		sieve_sendnums_send_s_env := env.To_Sieve_SendNums_SEND_S_Env()
		sieve_sendnums_send_s_result := Sieve_SendNums_SEND_S(wg, sieve_sendnums_send_s_chan, sieve_sendnums_send_s_inviteChan, sieve_sendnums_send_s_env)
		env.ResultFrom_Sieve_SendNums_SEND_S(sieve_sendnums_send_s_result)

		return env.Done()
	case callbacks.Sieve_SendNums_S_End:
		end_msg := env.End_To_R()
		roleChannels.R_End <- end_msg

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
}
