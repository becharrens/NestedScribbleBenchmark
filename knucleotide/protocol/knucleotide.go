package protocol

import "NestedScribbleBenchmark/knucleotide/messages/knucleotide"
import "NestedScribbleBenchmark/knucleotide/channels/schedulejobs"
import knucleotide_2 "NestedScribbleBenchmark/knucleotide/channels/knucleotide"
import "NestedScribbleBenchmark/knucleotide/invitations"
import "NestedScribbleBenchmark/knucleotide/callbacks"
import knucleotide_3 "NestedScribbleBenchmark/knucleotide/results/knucleotide"
import "NestedScribbleBenchmark/knucleotide/roles"
import "sync"

type KNucleotide_Env interface {
	New_Master_Env() callbacks.KNucleotide_Master_Env
	New_Worker_Env() callbacks.KNucleotide_Worker_Env
	Master_Result(result knucleotide_3.Master_Result)
	Worker_Result(result knucleotide_3.Worker_Result)
}

func Start_KNucleotide_Master(protocolEnv KNucleotide_Env, wg *sync.WaitGroup, roleChannels knucleotide_2.Master_Chan, inviteChannels invitations.KNucleotide_Master_InviteChan, env callbacks.KNucleotide_Master_Env) {
	defer wg.Done()
	result := roles.KNucleotide_Master(wg, roleChannels, inviteChannels, env)
	protocolEnv.Master_Result(result)
}

func Start_KNucleotide_Worker(protocolEnv KNucleotide_Env, wg *sync.WaitGroup, roleChannels knucleotide_2.Worker_Chan, inviteChannels invitations.KNucleotide_Worker_InviteChan, env callbacks.KNucleotide_Worker_Env) {
	defer wg.Done()
	result := roles.KNucleotide_Worker(wg, roleChannels, inviteChannels, env)
	protocolEnv.Worker_Result(result)
}

func KNucleotide(protocolEnv KNucleotide_Env) {
	worker_master_sequenceresult := make(chan knucleotide.SequenceResult, 1)
	master_invite_master := make(chan schedulejobs.M_Chan, 1)
	master_invite_master_invitechan := make(chan invitations.ScheduleJobs_M_InviteChan, 1)
	master_worker_sequencejob := make(chan knucleotide.SequenceJob, 1)

	worker_chan := knucleotide_2.Worker_Chan{
		Master_SequenceResult: worker_master_sequenceresult,
		Master_SequenceJob:    master_worker_sequencejob,
	}
	master_chan := knucleotide_2.Master_Chan{
		Worker_SequenceResult: worker_master_sequenceresult,
		Worker_SequenceJob:    master_worker_sequencejob,
	}

	worker_inviteChan := invitations.KNucleotide_Worker_InviteChan{}
	master_inviteChan := invitations.KNucleotide_Master_InviteChan{
		Invite_Master_To_ScheduleJobs_M_InviteChan: master_invite_master_invitechan,
		Invite_Master_To_ScheduleJobs_M:            master_invite_master,
	}

	var wg sync.WaitGroup

	wg.Add(2)

	master_env := protocolEnv.New_Master_Env()
	worker_env := protocolEnv.New_Worker_Env()

	go Start_KNucleotide_Master(protocolEnv, &wg, master_chan, master_inviteChan, master_env)
	go Start_KNucleotide_Worker(protocolEnv, &wg, worker_chan, worker_inviteChan, worker_env)

	wg.Wait()
}
