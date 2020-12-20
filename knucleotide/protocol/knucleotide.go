package protocol

import "NestedScribbleBenchmark/knucleotide/channels/schedulejobs"
import "NestedScribbleBenchmark/knucleotide/channels/knucleotide"
import "NestedScribbleBenchmark/knucleotide/invitations"
import "NestedScribbleBenchmark/knucleotide/callbacks"
import knucleotide_2 "NestedScribbleBenchmark/knucleotide/results/knucleotide"
import "NestedScribbleBenchmark/knucleotide/roles"
import "sync"

type KNucleotide_Env interface {
	New_Master_Env() callbacks.KNucleotide_Master_Env
	Master_Result(result knucleotide_2.Master_Result)
}

func Start_KNucleotide_Master(protocolEnv KNucleotide_Env, wg *sync.WaitGroup, roleChannels knucleotide.Master_Chan, inviteChannels invitations.KNucleotide_Master_InviteChan, env callbacks.KNucleotide_Master_Env) {
	defer wg.Done()
	result := roles.KNucleotide_Master(wg, roleChannels, inviteChannels, env)
	protocolEnv.Master_Result(result)
}

func KNucleotide(protocolEnv KNucleotide_Env) {
	master_invite_master := make(chan schedulejobs.M_Chan, 1)
	master_invite_master_invitechan := make(chan invitations.ScheduleJobs_M_InviteChan, 1)

	master_chan := knucleotide.Master_Chan{}

	master_inviteChan := invitations.KNucleotide_Master_InviteChan{
		Invite_Master_To_ScheduleJobs_M_InviteChan: master_invite_master_invitechan,
		Invite_Master_To_ScheduleJobs_M:            master_invite_master,
	}

	var wg sync.WaitGroup

	wg.Add(1)

	master_env := protocolEnv.New_Master_Env()

	go Start_KNucleotide_Master(protocolEnv, &wg, master_chan, master_inviteChan, master_env)

	wg.Wait()
}
