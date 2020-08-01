package roles

import "NestedScribbleBenchmark/primesieve/channels/sieve"
import "NestedScribbleBenchmark/primesieve/invitations"
import "NestedScribbleBenchmark/primesieve/callbacks"
import sieve_2 "NestedScribbleBenchmark/primesieve/results/sieve"
import "sync"

func Sieve_M(wg *sync.WaitGroup, roleChannels sieve.M_Chan, inviteChannels invitations.Sieve_M_InviteChan, env callbacks.Sieve_M_Env) sieve_2.M_Result {
	select {
	case prime_msg := <-roleChannels.W2_Prime:
		env.Prime_From_W2(prime_msg)

		sieve_m_chan := <-inviteChannels.W2_Invite_To_Sieve_M
		sieve_m_inviteChan := <-inviteChannels.W2_Invite_To_Sieve_M_InviteChan
		sieve_m_env := env.To_Sieve_M_Env()
		sieve_m_result := Sieve_M(wg, sieve_m_chan, sieve_m_inviteChan, sieve_m_env)
		env.ResultFrom_Sieve_M(sieve_m_result)

		return env.Done()
	case finish_msg := <-roleChannels.W2_Finish:
		env.Finish_From_W2(finish_msg)

		return env.Done()
	}
}
