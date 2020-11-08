package protocol

import "NestedScribbleBenchmark/boundedfibonacci/messages/boundedfibonacci"
import boundedfibonacci_2 "NestedScribbleBenchmark/boundedfibonacci/channels/boundedfibonacci"
import "NestedScribbleBenchmark/boundedfibonacci/channels/boundedfib"
import "NestedScribbleBenchmark/boundedfibonacci/invitations"
import "NestedScribbleBenchmark/boundedfibonacci/callbacks"
import boundedfibonacci_3 "NestedScribbleBenchmark/boundedfibonacci/results/boundedfibonacci"
import "NestedScribbleBenchmark/boundedfibonacci/roles"
import "sync"

type BoundedFibonacci_Env interface {
	New_Start_Env() callbacks.BoundedFibonacci_Start_Env
	New_F1_Env() callbacks.BoundedFibonacci_F1_Env
	New_F2_Env() callbacks.BoundedFibonacci_F2_Env
	Start_Result(result boundedfibonacci_3.Start_Result) 
	F1_Result(result boundedfibonacci_3.F1_Result) 
	F2_Result(result boundedfibonacci_3.F2_Result) 
}

func Start_BoundedFibonacci_Start(protocolEnv BoundedFibonacci_Env, wg *sync.WaitGroup, roleChannels boundedfibonacci_2.Start_Chan, inviteChannels invitations.BoundedFibonacci_Start_InviteChan, env callbacks.BoundedFibonacci_Start_Env)  {
	defer wg.Done()
	result := roles.BoundedFibonacci_Start(wg, roleChannels, inviteChannels, env)
	protocolEnv.Start_Result(result)
} 

func Start_BoundedFibonacci_F1(protocolEnv BoundedFibonacci_Env, wg *sync.WaitGroup, roleChannels boundedfibonacci_2.F1_Chan, inviteChannels invitations.BoundedFibonacci_F1_InviteChan, env callbacks.BoundedFibonacci_F1_Env)  {
	defer wg.Done()
	result := roles.BoundedFibonacci_F1(wg, roleChannels, inviteChannels, env)
	protocolEnv.F1_Result(result)
} 

func Start_BoundedFibonacci_F2(protocolEnv BoundedFibonacci_Env, wg *sync.WaitGroup, roleChannels boundedfibonacci_2.F2_Chan, inviteChannels invitations.BoundedFibonacci_F2_InviteChan, env callbacks.BoundedFibonacci_F2_Env)  {
	defer wg.Done()
	result := roles.BoundedFibonacci_F2(wg, roleChannels, inviteChannels, env)
	protocolEnv.F2_Result(result)
} 

func BoundedFibonacci(protocolEnv BoundedFibonacci_Env)  {
	start_invite_f2 := make(chan boundedfib.F2_Chan, 1)
	start_invite_f2_invitechan := make(chan invitations.BoundedFib_F2_InviteChan, 1)
	start_invite_f1 := make(chan boundedfib.F1_Chan, 1)
	start_invite_f1_invitechan := make(chan invitations.BoundedFib_F1_InviteChan, 1)
	start_invite_start := make(chan boundedfib.Res_Chan, 1)
	start_invite_start_invitechan := make(chan invitations.BoundedFib_Res_InviteChan, 1)
	start_f2_startfib2 := make(chan boundedfibonacci.StartFib2, 1)
	start_f1_startfib1 := make(chan boundedfibonacci.StartFib1, 1)

	start_chan := boundedfibonacci_2.Start_Chan{
		F2_StartFib2: start_f2_startfib2,
		F1_StartFib1: start_f1_startfib1,
	}
	f2_chan := boundedfibonacci_2.F2_Chan{
		Start_StartFib2: start_f2_startfib2,
	}
	f1_chan := boundedfibonacci_2.F1_Chan{
		Start_StartFib1: start_f1_startfib1,
	}

	start_inviteChan := invitations.BoundedFibonacci_Start_InviteChan{
		Invite_Start_To_BoundedFib_Res_InviteChan: start_invite_start_invitechan,
		Invite_Start_To_BoundedFib_Res: start_invite_start,
		Invite_F2_To_BoundedFib_F2_InviteChan: start_invite_f2_invitechan,
		Invite_F2_To_BoundedFib_F2: start_invite_f2,
		Invite_F1_To_BoundedFib_F1_InviteChan: start_invite_f1_invitechan,
		Invite_F1_To_BoundedFib_F1: start_invite_f1,
	}
	f2_inviteChan := invitations.BoundedFibonacci_F2_InviteChan{
		Start_Invite_To_BoundedFib_F2_InviteChan: start_invite_f2_invitechan,
		Start_Invite_To_BoundedFib_F2: start_invite_f2,
	}
	f1_inviteChan := invitations.BoundedFibonacci_F1_InviteChan{
		Start_Invite_To_BoundedFib_F1_InviteChan: start_invite_f1_invitechan,
		Start_Invite_To_BoundedFib_F1: start_invite_f1,
	}

	var wg sync.WaitGroup

	wg.Add(3)

	start_env := protocolEnv.New_Start_Env()
	f1_env := protocolEnv.New_F1_Env()
	f2_env := protocolEnv.New_F2_Env()

	go Start_BoundedFibonacci_Start(protocolEnv, &wg, start_chan, start_inviteChan, start_env)
	go Start_BoundedFibonacci_F1(protocolEnv, &wg, f1_chan, f1_inviteChan, f1_env)
	go Start_BoundedFibonacci_F2(protocolEnv, &wg, f2_chan, f2_inviteChan, f2_env)

	wg.Wait()
} 