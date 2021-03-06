package roles

import (
	"NestedScribbleBenchmark/primesieve_base/primesieve2/callbacks"
	"NestedScribbleBenchmark/primesieve_base/primesieve2/channels/primesieve"
	"NestedScribbleBenchmark/primesieve_base/primesieve2/invitations"
	"NestedScribbleBenchmark/primesieve_base/primesieve2/messages"
	"sync"
)

func PrimeSieve_Worker(wg *sync.WaitGroup, roleChannels primesieve.Worker_Chan, inviteChannels invitations.PrimeSieve_Worker_InviteChan, env callbacks.PrimeSieve_Worker_Env) {
	<-roleChannels.Label_From_Master
	prime := <-roleChannels.Int_From_Master
	env.FirstPrime_From_Master(prime)

	<-roleChannels.Label_From_Master
	n := <-roleChannels.Int_From_Master
	env.UBound_From_Master(n)

	worker_choice := env.Worker_Choice()
	switch worker_choice {
	case callbacks.PrimeSieve_Worker_Prime:
		n_2 := env.Prime_To_Master()
		roleChannels.Label_To_Master <- messages.Prime
		roleChannels.Int_To_Master <- n_2

		env.Sieve_Setup()
		roleChannels.Label_To_Master <- messages.Sieve_Master_Worker

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
		Sieve_W1(wg, sieve_w1_chan, sieve_w1_inviteChan, sieve_w1_env)

		env.Done()
		return
	case callbacks.PrimeSieve_Worker_Finish:
		env.Finish_To_Master()
		roleChannels.Label_To_Master <- messages.Finish

		env.Done()
		return
	default:
		panic("Invalid choice was made")
	}
}
