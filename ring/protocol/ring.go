package protocol

import "NestedScribbleBenchmark/ring/messages/ring"
import ring_2 "NestedScribbleBenchmark/ring/channels/ring"
import "NestedScribbleBenchmark/ring/channels/forward"
import "NestedScribbleBenchmark/ring/invitations"
import "NestedScribbleBenchmark/ring/callbacks"
import ring_3 "NestedScribbleBenchmark/ring/results/ring"
import "NestedScribbleBenchmark/ring/roles"
import "sync"

type Ring_Env interface {
	New_Start_Env() callbacks.Ring_Start_Env
	New_End_Env() callbacks.Ring_End_Env
	Start_Result(result ring_3.Start_Result)
	End_Result(result ring_3.End_Result)
}

func Start_Ring_Start(protocolEnv Ring_Env, wg *sync.WaitGroup, roleChannels ring_2.Start_Chan, inviteChannels invitations.Ring_Start_InviteChan, env callbacks.Ring_Start_Env) {
	defer wg.Done()
	result := roles.Ring_Start(wg, roleChannels, inviteChannels, env)
	protocolEnv.Start_Result(result)
}

func Start_Ring_End(protocolEnv Ring_Env, wg *sync.WaitGroup, roleChannels ring_2.End_Chan, inviteChannels invitations.Ring_End_InviteChan, env callbacks.Ring_End_Env) {
	defer wg.Done()
	result := roles.Ring_End(wg, roleChannels, inviteChannels, env)
	protocolEnv.End_Result(result)
}

func Ring(protocolEnv Ring_Env) {
	end_start_msg_2 := make(chan ring.Msg, 1)
	start_end_msg := make(chan ring.Msg, 1)
	end_start_msg := make(chan ring.Msg, 1)
	start_invite_end := make(chan forward.E_Chan, 1)
	start_invite_end_invitechan := make(chan invitations.Forward_E_InviteChan, 1)
	start_invite_start := make(chan forward.S_Chan, 1)
	start_invite_start_invitechan := make(chan invitations.Forward_S_InviteChan, 1)

	start_chan := ring_2.Start_Chan{
		End_Msg_3: end_start_msg_2,
		End_Msg_2: start_end_msg,
		End_Msg:   end_start_msg,
	}
	end_chan := ring_2.End_Chan{
		Start_Msg_3: end_start_msg_2,
		Start_Msg_2: start_end_msg,
		Start_Msg:   end_start_msg,
	}

	start_inviteChan := invitations.Ring_Start_InviteChan{
		Invite_Start_To_Forward_S_InviteChan: start_invite_start_invitechan,
		Invite_Start_To_Forward_S:            start_invite_start,
		Invite_End_To_Forward_E_InviteChan:   start_invite_end_invitechan,
		Invite_End_To_Forward_E:              start_invite_end,
	}
	end_inviteChan := invitations.Ring_End_InviteChan{
		Start_Invite_To_Forward_E_InviteChan: start_invite_end_invitechan,
		Start_Invite_To_Forward_E:            start_invite_end,
	}

	var wg sync.WaitGroup

	wg.Add(2)

	start_env := protocolEnv.New_Start_Env()
	end_env := protocolEnv.New_End_Env()

	go Start_Ring_Start(protocolEnv, &wg, start_chan, start_inviteChan, start_env)
	go Start_Ring_End(protocolEnv, &wg, end_chan, end_inviteChan, end_env)

	wg.Wait()
}
