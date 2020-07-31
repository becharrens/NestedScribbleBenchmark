package protocol

import "ScribbleBenchmark/regexredux/messages/regexredux"
import "ScribbleBenchmark/regexredux/channels/regexredux2"
import regexredux_2 "ScribbleBenchmark/regexredux/channels/regexredux"
import "ScribbleBenchmark/regexredux/invitations"
import "ScribbleBenchmark/regexredux/callbacks"
import regexredux_3 "ScribbleBenchmark/regexredux/results/regexredux"
import "ScribbleBenchmark/regexredux/roles"
import "sync"

type RegexRedux_Env interface {
	New_Master_Env() callbacks.RegexRedux_Master_Env
	New_Worker_Env() callbacks.RegexRedux_Worker_Env
	Master_Result(result regexredux_3.Master_Result) 
	Worker_Result(result regexredux_3.Worker_Result) 
}

func Start_RegexRedux_Master(protocolEnv RegexRedux_Env, wg *sync.WaitGroup, roleChannels regexredux_2.Master_Chan, inviteChannels invitations.RegexRedux_Master_InviteChan, env callbacks.RegexRedux_Master_Env)  {
	defer wg.Done()
	result := roles.RegexRedux_Master(wg, roleChannels, inviteChannels, env)
	protocolEnv.Master_Result(result)
} 

func Start_RegexRedux_Worker(protocolEnv RegexRedux_Env, wg *sync.WaitGroup, roleChannels regexredux_2.Worker_Chan, inviteChannels invitations.RegexRedux_Worker_InviteChan, env callbacks.RegexRedux_Worker_Env)  {
	defer wg.Done()
	result := roles.RegexRedux_Worker(wg, roleChannels, inviteChannels, env)
	protocolEnv.Worker_Result(result)
} 

func RegexRedux(protocolEnv RegexRedux_Env)  {
	worker_master_nummatches := make(chan regexredux.NumMatches, 1)
	master_invite_master := make(chan regexredux2.M_Chan, 1)
	master_invite_master_invitechan := make(chan invitations.RegexRedux2_M_InviteChan, 1)
	master_worker_task := make(chan regexredux.Task, 1)

	worker_chan := regexredux_2.Worker_Chan{
		Master_Task: master_worker_task,
		Master_NumMatches: worker_master_nummatches,
	}
	master_chan := regexredux_2.Master_Chan{
		Worker_Task: master_worker_task,
		Worker_NumMatches: worker_master_nummatches,
	}

	worker_inviteChan := invitations.RegexRedux_Worker_InviteChan{

	}
	master_inviteChan := invitations.RegexRedux_Master_InviteChan{
		Invite_Master_To_RegexRedux2_M_InviteChan: master_invite_master_invitechan,
		Invite_Master_To_RegexRedux2_M: master_invite_master,
	}

	var wg sync.WaitGroup

	wg.Add(2)

	master_env := protocolEnv.New_Master_Env()
	worker_env := protocolEnv.New_Worker_Env()

	go Start_RegexRedux_Master(protocolEnv, &wg, master_chan, master_inviteChan, master_env)
	go Start_RegexRedux_Worker(protocolEnv, &wg, worker_chan, worker_inviteChan, worker_env)

	wg.Wait()
} 