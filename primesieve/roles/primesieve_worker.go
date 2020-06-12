package roles

import "ScribbleBenchmark/primesieve/channels/primesieve"
import "ScribbleBenchmark/primesieve/invitations"
import "ScribbleBenchmark/primesieve/callbacks"
import primesieve_2 "ScribbleBenchmark/primesieve/results/primesieve"
import "sync"

func PrimeSieve_Worker(wg *sync.WaitGroup, roleChannels primesieve.Worker_Chan, inviteChannels invitations.PrimeSieve_Worker_InviteChan, env callbacks.PrimeSieve_Worker_Env) primesieve_2.Worker_Result {
	firstprime_msg := <-roleChannels.Master_FirstPrime
	env.FirstPrime_From_Master(firstprime_msg)

	ubound_msg := <-roleChannels.Master_UBound
	env.UBound_From_Master(ubound_msg)

	worker_choice := env.Worker_Choice()
	switch worker_choice {
	case callbacks.PrimeSieve_Worker_Prime:
		prime_msg := env.Prime_To_Master()
		roleChannels.Master_Prime <- prime_msg

		env.Sieve_Setup()
		sieve_rolechan := invitations.Sieve_RoleSetupChan{
			M_Chan:  inviteChannels.Invite_Master_To_Sieve_M,
			W1_Chan: inviteChannels.Invite_Worker_To_Sieve_W1,
		}
		sieve_invitechan := invitations.Sieve_InviteSetupChan{
			M_InviteChan:  inviteChannels.Invite_Master_To_Sieve_M_InviteChan,
			W1_InviteChan: inviteChannels.Invite_Worker_To_Sieve_W1_InviteChan,
		}
		Sieve_SendCommChannels(wg, sieve_rolechan, sieve_invitechan)

		sieve_w1_chan := <-inviteChannels.Invite_Worker_To_Sieve_W1
		sieve_w1_inviteChan := <-inviteChannels.Invite_Worker_To_Sieve_W1_InviteChan
		sieve_w1_env := env.To_Sieve_W1_Env()
		sieve_w1_result := Sieve_W1(wg, sieve_w1_chan, sieve_w1_inviteChan, sieve_w1_env)
		env.ResultFrom_Sieve_W1(sieve_w1_result)

		return env.Done()
	case callbacks.PrimeSieve_Worker_Finish:
		finish_msg := env.Finish_To_Master()
		roleChannels.Master_Finish <- finish_msg

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
}
