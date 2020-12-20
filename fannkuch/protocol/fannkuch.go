package protocol

import "NestedScribbleBenchmark/fannkuch/messages"
import "NestedScribbleBenchmark/fannkuch/channels/fannkuchrecursive"
import "NestedScribbleBenchmark/fannkuch/channels/fannkuch"
import "NestedScribbleBenchmark/fannkuch/invitations"
import "NestedScribbleBenchmark/fannkuch/callbacks"
import fannkuch_2 "NestedScribbleBenchmark/fannkuch/results/fannkuch"
import "NestedScribbleBenchmark/fannkuch/roles"
import "sync"

type Fannkuch_Env interface {
	New_Main_Env() callbacks.Fannkuch_Main_Env
	New_Worker_Env() callbacks.Fannkuch_Worker_Env
	Main_Result(result fannkuch_2.Main_Result)
	Worker_Result(result fannkuch_2.Worker_Result)
}

func Start_Fannkuch_Main(protocolEnv Fannkuch_Env, wg *sync.WaitGroup, roleChannels fannkuch.Main_Chan, inviteChannels invitations.Fannkuch_Main_InviteChan, env callbacks.Fannkuch_Main_Env) {
	defer wg.Done()
	result := roles.Fannkuch_Main(wg, roleChannels, inviteChannels, env)
	protocolEnv.Main_Result(result)
}

func Start_Fannkuch_Worker(protocolEnv Fannkuch_Env, wg *sync.WaitGroup, roleChannels fannkuch.Worker_Chan, inviteChannels invitations.Fannkuch_Worker_InviteChan, env callbacks.Fannkuch_Worker_Env) {
	defer wg.Done()
	result := roles.Fannkuch_Worker(wg, roleChannels, inviteChannels, env)
	protocolEnv.Worker_Result(result)
}

func Fannkuch(protocolEnv Fannkuch_Env) {
	worker_main_int := make(chan int, 1)
	worker_invite_worker := make(chan fannkuchrecursive.Worker_Chan, 1)
	worker_invite_worker_invitechan := make(chan invitations.FannkuchRecursive_Worker_InviteChan, 1)
	worker_main_label := make(chan messages.Fannkuch_Label, 1)
	worker_invite_main := make(chan fannkuchrecursive.Source_Chan, 1)
	worker_invite_main_invitechan := make(chan invitations.FannkuchRecursive_Source_InviteChan, 1)
	main_worker_int := make(chan int, 1)
	main_worker_label := make(chan messages.Fannkuch_Label, 1)

	worker_chan := fannkuch.Worker_Chan{
		Label_To_Main:   worker_main_label,
		Label_From_Main: main_worker_label,
		Int_To_Main:     worker_main_int,
		Int_From_Main:   main_worker_int,
	}
	main_chan := fannkuch.Main_Chan{
		Label_To_Worker:   main_worker_label,
		Label_From_Worker: worker_main_label,
		Int_To_Worker:     main_worker_int,
		Int_From_Worker:   worker_main_int,
	}

	worker_inviteChan := invitations.Fannkuch_Worker_InviteChan{
		Invite_Worker_To_FannkuchRecursive_Worker_InviteChan: worker_invite_worker_invitechan,
		Invite_Worker_To_FannkuchRecursive_Worker:            worker_invite_worker,
		Invite_Main_To_FannkuchRecursive_Source_InviteChan:   worker_invite_main_invitechan,
		Invite_Main_To_FannkuchRecursive_Source:              worker_invite_main,
	}
	main_inviteChan := invitations.Fannkuch_Main_InviteChan{
		Worker_Invite_To_FannkuchRecursive_Source_InviteChan: worker_invite_main_invitechan,
		Worker_Invite_To_FannkuchRecursive_Source:            worker_invite_main,
	}

	var wg sync.WaitGroup

	wg.Add(2)

	main_env := protocolEnv.New_Main_Env()
	worker_env := protocolEnv.New_Worker_Env()

	go Start_Fannkuch_Main(protocolEnv, &wg, main_chan, main_inviteChan, main_env)
	go Start_Fannkuch_Worker(protocolEnv, &wg, worker_chan, worker_inviteChan, worker_env)

	wg.Wait()
}
