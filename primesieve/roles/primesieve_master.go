package roles

import "ScribbleBenchmark/primesieve/channels/primesieve"
import "ScribbleBenchmark/primesieve/invitations"
import "ScribbleBenchmark/primesieve/callbacks"
import primesieve_2 "ScribbleBenchmark/primesieve/results/primesieve"
import "sync"

func PrimeSieve_Master(wg *sync.WaitGroup, roleChannels primesieve.Master_Chan, inviteChannels invitations.PrimeSieve_Master_InviteChan, env callbacks.PrimeSieve_Master_Env) primesieve_2.Master_Result {
	firstprime_msg := env.FirstPrime_To_Worker()
	roleChannels.Worker_FirstPrime <- firstprime_msg

	ubound_msg := env.UBound_To_Worker()
	roleChannels.Worker_UBound <- ubound_msg

	select {
	case prime_msg := <-roleChannels.Worker_Prime:
		env.Prime_From_Worker(prime_msg)

		sieve_m_chan := <-inviteChannels.Worker_Invite_To_Sieve_M
		sieve_m_inviteChan := <-inviteChannels.Worker_Invite_To_Sieve_M_InviteChan
		sieve_m_env := env.To_Sieve_M_Env()
		sieve_m_result := Sieve_M(wg, sieve_m_chan, sieve_m_inviteChan, sieve_m_env)
		env.ResultFrom_Sieve_M(sieve_m_result)

		return env.Done()
	case finish_msg := <-roleChannels.Worker_Finish:
		env.Finish_From_Worker(finish_msg)

		return env.Done()
	}
}
