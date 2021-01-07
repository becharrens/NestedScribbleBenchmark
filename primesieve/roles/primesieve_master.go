package roles

import (
	"NestedScribbleBenchmark/primesieve/callbacks"
	"NestedScribbleBenchmark/primesieve/channels/primesieve"
	"NestedScribbleBenchmark/primesieve/invitations"
	"NestedScribbleBenchmark/primesieve/messages"
	primesieve_2 "NestedScribbleBenchmark/primesieve/results/primesieve"
	"sync"
)

func PrimeSieve_Master(wg *sync.WaitGroup, roleChannels primesieve.Master_Chan, inviteChannels invitations.PrimeSieve_Master_InviteChan, env callbacks.PrimeSieve_Master_Env) primesieve_2.Master_Result {
	prime := env.FirstPrime_To_Worker()
	roleChannels.Label_To_Worker <- messages.FirstPrime
	roleChannels.Int_To_Worker <- prime

	n := env.UBound_To_Worker()
	roleChannels.Label_To_Worker <- messages.UBound
	roleChannels.Int_To_Worker <- n

	worker_choice := <-roleChannels.Label_From_Worker
	switch worker_choice {
	case messages.Prime:
		n_2 := <-roleChannels.Int_From_Worker
		env.Prime_From_Worker(n_2)

		<-roleChannels.Label_From_Worker
		sieve_m_chan := <-inviteChannels.Worker_Invite_To_Sieve_M
		sieve_m_inviteChan := <-inviteChannels.Worker_Invite_To_Sieve_M_InviteChan
		sieve_m_env := env.To_Sieve_M_Env()
		sieve_m_result := Sieve_M(wg, sieve_m_chan, sieve_m_inviteChan, sieve_m_env)
		env.ResultFrom_Sieve_M(sieve_m_result)

		return env.Done()
	case messages.Finish:
		env.Finish_From_Worker()

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
}
