package protocol

import "NestedScribbleBenchmark/primesieve/messages/primesieve"
import "NestedScribbleBenchmark/primesieve/channels/sieve"
import primesieve_2 "NestedScribbleBenchmark/primesieve/channels/primesieve"
import "NestedScribbleBenchmark/primesieve/invitations"
import "NestedScribbleBenchmark/primesieve/callbacks"
import primesieve_3 "NestedScribbleBenchmark/primesieve/results/primesieve"
import "NestedScribbleBenchmark/primesieve/roles"
import "sync"

type PrimeSieve_Env interface {
	New_Master_Env() callbacks.PrimeSieve_Master_Env
	New_Worker_Env() callbacks.PrimeSieve_Worker_Env
	Master_Result(result primesieve_3.Master_Result)
	Worker_Result(result primesieve_3.Worker_Result)
}

func Start_PrimeSieve_Master(protocolEnv PrimeSieve_Env, wg *sync.WaitGroup, roleChannels primesieve_2.Master_Chan, inviteChannels invitations.PrimeSieve_Master_InviteChan, env callbacks.PrimeSieve_Master_Env) {
	defer wg.Done()
	result := roles.PrimeSieve_Master(wg, roleChannels, inviteChannels, env)
	protocolEnv.Master_Result(result)
}

func Start_PrimeSieve_Worker(protocolEnv PrimeSieve_Env, wg *sync.WaitGroup, roleChannels primesieve_2.Worker_Chan, inviteChannels invitations.PrimeSieve_Worker_InviteChan, env callbacks.PrimeSieve_Worker_Env) {
	defer wg.Done()
	result := roles.PrimeSieve_Worker(wg, roleChannels, inviteChannels, env)
	protocolEnv.Worker_Result(result)
}

func PrimeSieve(protocolEnv PrimeSieve_Env) {
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

	var wg sync.WaitGroup

	wg.Add(2)

	master_env := protocolEnv.New_Master_Env()
	worker_env := protocolEnv.New_Worker_Env()

	go Start_PrimeSieve_Master(protocolEnv, &wg, master_chan, master_inviteChan, master_env)
	go Start_PrimeSieve_Worker(protocolEnv, &wg, worker_chan, worker_inviteChan, worker_env)

	wg.Wait()
}
