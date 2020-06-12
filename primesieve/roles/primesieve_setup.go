package roles

import "ScribbleBenchmark/primesieve/messages/primesieve"
import "ScribbleBenchmark/primesieve/channels/sieve"
import primesieve_2 "ScribbleBenchmark/primesieve/channels/primesieve"
import "ScribbleBenchmark/primesieve/invitations"
import "sync"

func PrimeSieve_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.PrimeSieve_RoleSetupChan, inviteChannels invitations.PrimeSieve_InviteSetupChan) {
	worker_master_finish := make(chan primesieve.Finish, 1)
	worker_invite_worker := make(chan sieve.W1_Chan, 1)
	worker_invite_worker_invitechan := make(chan invitations.Sieve_W1_InviteChan, 1)
	worker_invite_master := make(chan sieve.M_Chan, 1)
	worker_invite_master_invitechan := make(chan invitations.Sieve_M_InviteChan, 1)
	worker_master_prime := make(chan primesieve.Prime, 1)
	master_worker_ubound := make(chan primesieve.UBound, 1)
	master_worker_firstprime := make(chan primesieve.FirstPrime, 1)

	worker_chan := primesieve_2.Worker_Chan{
		Master_UBound:     master_worker_ubound,
		Master_Prime:      worker_master_prime,
		Master_FirstPrime: master_worker_firstprime,
		Master_Finish:     worker_master_finish,
	}
	master_chan := primesieve_2.Master_Chan{
		Worker_UBound:     master_worker_ubound,
		Worker_Prime:      worker_master_prime,
		Worker_FirstPrime: master_worker_firstprime,
		Worker_Finish:     worker_master_finish,
	}

	worker_inviteChan := invitations.PrimeSieve_Worker_InviteChan{
		Invite_Worker_To_Sieve_W1_InviteChan: worker_invite_worker_invitechan,
		Invite_Worker_To_Sieve_W1:            worker_invite_worker,
		Invite_Master_To_Sieve_M_InviteChan:  worker_invite_master_invitechan,
		Invite_Master_To_Sieve_M:             worker_invite_master,
	}
	master_inviteChan := invitations.PrimeSieve_Master_InviteChan{
		Worker_Invite_To_Sieve_M_InviteChan: worker_invite_master_invitechan,
		Worker_Invite_To_Sieve_M:            worker_invite_master,
	}

	roleChannels.Master_Chan <- master_chan
	roleChannels.Worker_Chan <- worker_chan

	inviteChannels.Master_InviteChan <- master_inviteChan
	inviteChannels.Worker_InviteChan <- worker_inviteChan
}
