package protocol

import "NestedScribbleBenchmark/regexredux/channels/regexredux2"
import "NestedScribbleBenchmark/regexredux/channels/regexredux"
import "NestedScribbleBenchmark/regexredux/invitations"
import "NestedScribbleBenchmark/regexredux/callbacks"
import regexredux_2 "NestedScribbleBenchmark/regexredux/results/regexredux"
import "NestedScribbleBenchmark/regexredux/roles"
import "sync"

type RegexRedux_Env interface {
	New_Master_Env() callbacks.RegexRedux_Master_Env
	Master_Result(result regexredux_2.Master_Result)
}

func Start_RegexRedux_Master(protocolEnv RegexRedux_Env, wg *sync.WaitGroup, roleChannels regexredux.Master_Chan, inviteChannels invitations.RegexRedux_Master_InviteChan, env callbacks.RegexRedux_Master_Env) {
	defer wg.Done()
	result := roles.RegexRedux_Master(wg, roleChannels, inviteChannels, env)
	protocolEnv.Master_Result(result)
}

func RegexRedux(protocolEnv RegexRedux_Env) {
	master_invite_master := make(chan regexredux2.M_Chan, 1)
	master_invite_master_invitechan := make(chan invitations.RegexRedux2_M_InviteChan, 1)

	master_chan := regexredux.Master_Chan{}

	master_inviteChan := invitations.RegexRedux_Master_InviteChan{
		Invite_Master_To_RegexRedux2_M_InviteChan: master_invite_master_invitechan,
		Invite_Master_To_RegexRedux2_M:            master_invite_master,
	}

	var wg sync.WaitGroup

	wg.Add(1)

	master_env := protocolEnv.New_Master_Env()

	go Start_RegexRedux_Master(protocolEnv, &wg, master_chan, master_inviteChan, master_env)

	wg.Wait()
}
