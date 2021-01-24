package protocol

import "NestedScribbleBenchmark/ring/messages"
import "NestedScribbleBenchmark/ring/channels/ring"
import "NestedScribbleBenchmark/ring/channels/forward"
import "NestedScribbleBenchmark/ring/invitations"
import "NestedScribbleBenchmark/ring/callbacks"
import ring_2 "NestedScribbleBenchmark/ring/results/ring"
import "NestedScribbleBenchmark/ring/roles"
import "sync"

type Ring_Env interface {
	New_Start_Env() callbacks.Ring_Start_Env
	New_End_Env() callbacks.Ring_End_Env
	Start_Result(result ring_2.Start_Result)
	End_Result(result ring_2.End_Result)
}

func Start_Ring_Start(protocolEnv Ring_Env, wg *sync.WaitGroup, roleChannels ring.Start_Chan, inviteChannels invitations.Ring_Start_InviteChan, env callbacks.Ring_Start_Env) {
	defer wg.Done()
	result := roles.Ring_Start(wg, roleChannels, inviteChannels, env)
	protocolEnv.Start_Result(result)
}

func Start_Ring_End(protocolEnv Ring_Env, wg *sync.WaitGroup, roleChannels ring.End_Chan, inviteChannels invitations.Ring_End_InviteChan, env callbacks.Ring_End_Env) {
	defer wg.Done()
	result := roles.Ring_End(wg, roleChannels, inviteChannels, env)
	protocolEnv.End_Result(result)
}

func Ring(protocolEnv Ring_Env) {
	start_end_int := make(chan int, 1)
	start_end_string := make(chan string, 1)
	end_start_int := make(chan int, 1)
	end_start_string := make(chan string, 1)
	end_start_label := make(chan messages.Ring_Label, 1)
	start_end_label := make(chan messages.Ring_Label, 1)
	start_invite_end := make(chan forward.E_Chan, 1)
	start_invite_end_invitechan := make(chan invitations.Forward_E_InviteChan, 1)
	start_invite_start := make(chan forward.S_Chan, 1)
	start_invite_start_invitechan := make(chan invitations.Forward_S_InviteChan, 1)

	start_chan := ring.Start_Chan{
		String_To_End:   start_end_string,
		String_From_End: end_start_string,
		Label_To_End:    start_end_label,
		Label_From_End:  end_start_label,
		Int_To_End:      start_end_int,
		Int_From_End:    end_start_int,
	}
	end_chan := ring.End_Chan{
		String_To_Start:   end_start_string,
		String_From_Start: start_end_string,
		Label_To_Start:    end_start_label,
		Label_From_Start:  start_end_label,
		Int_To_Start:      end_start_int,
		Int_From_Start:    start_end_int,
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
