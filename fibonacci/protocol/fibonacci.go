package protocol

import "ScribbleBenchmark/fibonacci/messages/fibonacci"
import fibonacci_2 "ScribbleBenchmark/fibonacci/channels/fibonacci"
import "ScribbleBenchmark/fibonacci/channels/fib"
import "ScribbleBenchmark/fibonacci/invitations"
import "ScribbleBenchmark/fibonacci/callbacks"
import fibonacci_3 "ScribbleBenchmark/fibonacci/results/fibonacci"
import "ScribbleBenchmark/fibonacci/roles"
import "sync"

type Fibonacci_Env interface {
	New_Start_Env() callbacks.Fibonacci_Start_Env
	New_F1_Env() callbacks.Fibonacci_F1_Env
	New_F2_Env() callbacks.Fibonacci_F2_Env
	Start_Result(result fibonacci_3.Start_Result) 
	F1_Result(result fibonacci_3.F1_Result) 
	F2_Result(result fibonacci_3.F2_Result) 
}

func Start_Fibonacci_Start(protocolEnv Fibonacci_Env, wg *sync.WaitGroup, roleChannels fibonacci_2.Start_Chan, inviteChannels invitations.Fibonacci_Start_InviteChan, env callbacks.Fibonacci_Start_Env)  {
	defer wg.Done()
	result := roles.Fibonacci_Start(wg, roleChannels, inviteChannels, env)
	protocolEnv.Start_Result(result)
} 

func Start_Fibonacci_F1(protocolEnv Fibonacci_Env, wg *sync.WaitGroup, roleChannels fibonacci_2.F1_Chan, inviteChannels invitations.Fibonacci_F1_InviteChan, env callbacks.Fibonacci_F1_Env)  {
	defer wg.Done()
	result := roles.Fibonacci_F1(wg, roleChannels, inviteChannels, env)
	protocolEnv.F1_Result(result)
} 

func Start_Fibonacci_F2(protocolEnv Fibonacci_Env, wg *sync.WaitGroup, roleChannels fibonacci_2.F2_Chan, inviteChannels invitations.Fibonacci_F2_InviteChan, env callbacks.Fibonacci_F2_Env)  {
	defer wg.Done()
	result := roles.Fibonacci_F2(wg, roleChannels, inviteChannels, env)
	protocolEnv.F2_Result(result)
} 

func Fibonacci(protocolEnv Fibonacci_Env)  {
	start_invite_f2 := make(chan fib.F2_Chan, 1)
	start_invite_f2_invitechan := make(chan invitations.Fib_F2_InviteChan, 1)
	start_invite_f1 := make(chan fib.F1_Chan, 1)
	start_invite_f1_invitechan := make(chan invitations.Fib_F1_InviteChan, 1)
	start_invite_start := make(chan fib.Res_Chan, 1)
	start_invite_start_invitechan := make(chan invitations.Fib_Res_InviteChan, 1)
	start_f2_startfib2 := make(chan fibonacci.StartFib2, 1)
	start_f1_startfib1 := make(chan fibonacci.StartFib1, 1)

	start_chan := fibonacci_2.Start_Chan{
		F2_StartFib2: start_f2_startfib2,
		F1_StartFib1: start_f1_startfib1,
	}
	f2_chan := fibonacci_2.F2_Chan{
		Start_StartFib2: start_f2_startfib2,
	}
	f1_chan := fibonacci_2.F1_Chan{
		Start_StartFib1: start_f1_startfib1,
	}

	start_inviteChan := invitations.Fibonacci_Start_InviteChan{
		Invite_Start_To_Fib_Res_InviteChan: start_invite_start_invitechan,
		Invite_Start_To_Fib_Res: start_invite_start,
		Invite_F2_To_Fib_F2_InviteChan: start_invite_f2_invitechan,
		Invite_F2_To_Fib_F2: start_invite_f2,
		Invite_F1_To_Fib_F1_InviteChan: start_invite_f1_invitechan,
		Invite_F1_To_Fib_F1: start_invite_f1,
	}
	f2_inviteChan := invitations.Fibonacci_F2_InviteChan{
		Start_Invite_To_Fib_F2_InviteChan: start_invite_f2_invitechan,
		Start_Invite_To_Fib_F2: start_invite_f2,
	}
	f1_inviteChan := invitations.Fibonacci_F1_InviteChan{
		Start_Invite_To_Fib_F1_InviteChan: start_invite_f1_invitechan,
		Start_Invite_To_Fib_F1: start_invite_f1,
	}

	var wg sync.WaitGroup

	wg.Add(3)

	start_env := protocolEnv.New_Start_Env()
	f1_env := protocolEnv.New_F1_Env()
	f2_env := protocolEnv.New_F2_Env()

	go Start_Fibonacci_Start(protocolEnv, &wg, start_chan, start_inviteChan, start_env)
	go Start_Fibonacci_F1(protocolEnv, &wg, f1_chan, f1_inviteChan, f1_env)
	go Start_Fibonacci_F2(protocolEnv, &wg, f2_chan, f2_inviteChan, f2_env)

	wg.Wait()
} 