package roles

import "NestedScribbleBenchmark/primesieve_base/primesieve3/messages"
import "NestedScribbleBenchmark/primesieve_base/primesieve3/channels/sieve"
import "NestedScribbleBenchmark/primesieve_base/primesieve3/channels/primesieve"
import "NestedScribbleBenchmark/primesieve_base/primesieve3/invitations"
import "sync"

func PrimeSieve_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.PrimeSieve_RoleSetupChan, inviteChannels invitations.PrimeSieve_InviteSetupChan) {
	worker_invite_worker := make(chan sieve.W1_Chan, 1)
	worker_invite_worker_invitechan := make(chan invitations.Sieve_W1_InviteChan, 1)
	worker_invite_master := make(chan sieve.M_Chan, 1)
	worker_invite_master_invitechan := make(chan invitations.Sieve_M_InviteChan, 1)
	worker_master_int := make(chan int, 1)
	worker_master_label := make(chan messages.PrimeSieve_Label, 1)
	master_worker_int := make(chan int, 1)
	master_worker_label := make(chan messages.PrimeSieve_Label, 1)

	worker_chan := primesieve.Worker_Chan{
		Label_To_Master:   worker_master_label,
		Label_From_Master: master_worker_label,
		Int_To_Master:     worker_master_int,
		Int_From_Master:   master_worker_int,
	}
	master_chan := primesieve.Master_Chan{
		Label_To_Worker:   master_worker_label,
		Label_From_Worker: worker_master_label,
		Int_To_Worker:     master_worker_int,
		Int_From_Worker:   worker_master_int,
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
