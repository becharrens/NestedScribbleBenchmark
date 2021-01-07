package protocol

import "NestedScribbleBenchmark/primesieve_base/primesieve4/messages"
import "NestedScribbleBenchmark/primesieve_base/primesieve4/channels/sieve"
import "NestedScribbleBenchmark/primesieve_base/primesieve4/channels/primesieve"
import "NestedScribbleBenchmark/primesieve_base/primesieve4/invitations"
import "NestedScribbleBenchmark/primesieve_base/primesieve4/callbacks"
import primesieve_2 "NestedScribbleBenchmark/primesieve_base/primesieve4/results/primesieve"
import "NestedScribbleBenchmark/primesieve_base/primesieve4/roles"
import "sync"

type PrimeSieve_Env interface {
	New_Master_Env() callbacks.PrimeSieve_Master_Env
	New_Worker_Env() callbacks.PrimeSieve_Worker_Env
	Master_Result(result primesieve_2.Master_Result)
	Worker_Result(result primesieve_2.Worker_Result)
}

func Start_PrimeSieve_Master(protocolEnv PrimeSieve_Env, wg *sync.WaitGroup, roleChannels primesieve.Master_Chan, inviteChannels invitations.PrimeSieve_Master_InviteChan, env callbacks.PrimeSieve_Master_Env) {
	defer wg.Done()
	result := roles.PrimeSieve_Master(wg, roleChannels, inviteChannels, env)
	protocolEnv.Master_Result(result)
}

func Start_PrimeSieve_Worker(protocolEnv PrimeSieve_Env, wg *sync.WaitGroup, roleChannels primesieve.Worker_Chan, inviteChannels invitations.PrimeSieve_Worker_InviteChan, env callbacks.PrimeSieve_Worker_Env) {
	defer wg.Done()
	result := roles.PrimeSieve_Worker(wg, roleChannels, inviteChannels, env)
	protocolEnv.Worker_Result(result)
}

func PrimeSieve(protocolEnv PrimeSieve_Env) {
	// worker_invite_worker := make(chan sieve.W1_Chan, 1)
	// worker_invite_worker_invitechan := make(chan invitations.Sieve_W1_InviteChan, 1)
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
		// Invite_Worker_To_Sieve_W1_InviteChan: worker_invite_worker_invitechan,
		// Invite_Worker_To_Sieve_W1:           worker_invite_worker,
		Invite_Master_To_Sieve_M_InviteChan: worker_invite_master_invitechan,
		Invite_Master_To_Sieve_M:            worker_invite_master,
	}
	master_inviteChan := invitations.PrimeSieve_Master_InviteChan{
		Worker_Invite_To_Sieve_M_InviteChan: worker_invite_master_invitechan,
		Worker_Invite_To_Sieve_M:            worker_invite_master,
	}

	var wg sync.WaitGroup

	wg.Add(2)

	master_env := protocolEnv.New_Master_Env()
	worker_env := protocolEnv.New_Worker_Env()

	go Start_PrimeSieve_Master(protocolEnv, &wg, master_chan, master_inviteChan, master_env)
	go Start_PrimeSieve_Worker(protocolEnv, &wg, worker_chan, worker_inviteChan, worker_env)

	wg.Wait()
}
