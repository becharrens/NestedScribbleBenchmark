package roles

import "NestedScribbleBenchmark/primesieve_base/primesieve4/messages"
import "NestedScribbleBenchmark/primesieve_base/primesieve4/channels/sieve"
import "NestedScribbleBenchmark/primesieve_base/primesieve4/invitations"
import "NestedScribbleBenchmark/primesieve_base/primesieve4/callbacks"
import sieve_2 "NestedScribbleBenchmark/primesieve_base/primesieve4/results/sieve"
import "sync"

func Sieve_M(wg *sync.WaitGroup, roleChannels sieve.M_Chan, inviteChannels invitations.Sieve_M_InviteChan, env callbacks.Sieve_M_Env) sieve_2.M_Result {
	w2_choice := <-roleChannels.Label_From_W2
	switch w2_choice {
	case messages.Prime:
		n := <-roleChannels.Int_From_W2
		env.Prime_From_W2(n)

		<-roleChannels.Label_From_W2
		sieve_m_chan := <-inviteChannels.W2_Invite_To_Sieve_M
		sieve_m_inviteChan := <-inviteChannels.W2_Invite_To_Sieve_M_InviteChan
		sieve_m_env := env.To_Sieve_M_Env()
		sieve_m_result := Sieve_M(wg, sieve_m_chan, sieve_m_inviteChan, sieve_m_env)
		env.ResultFrom_Sieve_M(sieve_m_result)

		return env.Done()
	case messages.Finish:
		env.Finish_From_W2()

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
}
