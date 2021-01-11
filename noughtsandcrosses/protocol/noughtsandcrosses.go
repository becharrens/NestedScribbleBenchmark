package protocol

import "NestedScribbleBenchmark/noughtsandcrosses/messages"
import "NestedScribbleBenchmark/noughtsandcrosses/channels/noughtsandcrosses"
import "NestedScribbleBenchmark/noughtsandcrosses/channels/calcmove"
import "NestedScribbleBenchmark/noughtsandcrosses/invitations"
import "NestedScribbleBenchmark/noughtsandcrosses/callbacks"
import noughtsandcrosses_2 "NestedScribbleBenchmark/noughtsandcrosses/results/noughtsandcrosses"
import "NestedScribbleBenchmark/noughtsandcrosses/roles"
import "sync"

type NoughtsAndCrosses_Env interface {
	New_P1_Env() callbacks.NoughtsAndCrosses_P1_Env
	New_P2_Env() callbacks.NoughtsAndCrosses_P2_Env
	P1_Result(result noughtsandcrosses_2.P1_Result) 
	P2_Result(result noughtsandcrosses_2.P2_Result) 
}

func Start_NoughtsAndCrosses_P1(protocolEnv NoughtsAndCrosses_Env, wg *sync.WaitGroup, roleChannels noughtsandcrosses.P1_Chan, inviteChannels invitations.NoughtsAndCrosses_P1_InviteChan, env callbacks.NoughtsAndCrosses_P1_Env)  {
	defer wg.Done()
	result := roles.NoughtsAndCrosses_P1(wg, roleChannels, inviteChannels, env)
	protocolEnv.P1_Result(result)
} 

func Start_NoughtsAndCrosses_P2(protocolEnv NoughtsAndCrosses_Env, wg *sync.WaitGroup, roleChannels noughtsandcrosses.P2_Chan, inviteChannels invitations.NoughtsAndCrosses_P2_InviteChan, env callbacks.NoughtsAndCrosses_P2_Env)  {
	defer wg.Done()
	result := roles.NoughtsAndCrosses_P2(wg, roleChannels, inviteChannels, env)
	protocolEnv.P2_Result(result)
} 

func NoughtsAndCrosses(protocolEnv NoughtsAndCrosses_Env)  {
	p2_p1_int := make(chan int, 1)
	p2_p1_label := make(chan messages.NoughtsAndCrosses_Label, 1)
	p2_invite_p2 := make(chan calcmove.P_Chan, 1)
	p2_invite_p2_invitechan := make(chan invitations.CalcMove_P_InviteChan, 1)
	p1_p2_int := make(chan int, 1)
	p1_p2_label := make(chan messages.NoughtsAndCrosses_Label, 1)
	p1_invite_p1 := make(chan calcmove.P_Chan, 1)
	p1_invite_p1_invitechan := make(chan invitations.CalcMove_P_InviteChan, 1)

	p2_chan := noughtsandcrosses.P2_Chan{
		Label_To_P1: p2_p1_label,
		Label_From_P1: p1_p2_label,
		Int_To_P1: p2_p1_int,
		Int_From_P1: p1_p2_int,
	}
	p1_chan := noughtsandcrosses.P1_Chan{
		Label_To_P2: p1_p2_label,
		Label_From_P2: p2_p1_label,
		Int_To_P2: p1_p2_int,
		Int_From_P2: p2_p1_int,
	}

	p2_inviteChan := invitations.NoughtsAndCrosses_P2_InviteChan{
		Invite_P2_To_CalcMove_P_InviteChan: p2_invite_p2_invitechan,
		Invite_P2_To_CalcMove_P: p2_invite_p2,
	}
	p1_inviteChan := invitations.NoughtsAndCrosses_P1_InviteChan{
		Invite_P1_To_CalcMove_P_InviteChan: p1_invite_p1_invitechan,
		Invite_P1_To_CalcMove_P: p1_invite_p1,
	}

	var wg sync.WaitGroup

	wg.Add(2)

	p1_env := protocolEnv.New_P1_Env()
	p2_env := protocolEnv.New_P2_Env()

	go Start_NoughtsAndCrosses_P1(protocolEnv, &wg, p1_chan, p1_inviteChan, p1_env)
	go Start_NoughtsAndCrosses_P2(protocolEnv, &wg, p2_chan, p2_inviteChan, p2_env)

	wg.Wait()
} 