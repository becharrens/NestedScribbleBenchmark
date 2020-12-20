package roles

import "NestedScribbleBenchmark/knucleotide/channels/knucleotide"
import "NestedScribbleBenchmark/knucleotide/invitations"
import "NestedScribbleBenchmark/knucleotide/callbacks"
import knucleotide_2 "NestedScribbleBenchmark/knucleotide/results/knucleotide"
import "sync"

func KNucleotide_Master(wg *sync.WaitGroup, roleChannels knucleotide.Master_Chan, inviteChannels invitations.KNucleotide_Master_InviteChan, env callbacks.KNucleotide_Master_Env) knucleotide_2.Master_Result {
	env.ScheduleJobs_Setup()

	schedulejobs_rolechan := invitations.ScheduleJobs_RoleSetupChan{
		M_Chan: inviteChannels.Invite_Master_To_ScheduleJobs_M,
	}
	schedulejobs_invitechan := invitations.ScheduleJobs_InviteSetupChan{
		M_InviteChan: inviteChannels.Invite_Master_To_ScheduleJobs_M_InviteChan,
	}
	ScheduleJobs_SendCommChannels(wg, schedulejobs_rolechan, schedulejobs_invitechan)

	schedulejobs_m_chan := <-inviteChannels.Invite_Master_To_ScheduleJobs_M
	schedulejobs_m_inviteChan := <-inviteChannels.Invite_Master_To_ScheduleJobs_M_InviteChan
	schedulejobs_m_env := env.To_ScheduleJobs_M_Env()
	schedulejobs_m_result := ScheduleJobs_M(wg, schedulejobs_m_chan, schedulejobs_m_inviteChan, schedulejobs_m_env)
	env.ResultFrom_ScheduleJobs_M(schedulejobs_m_result)

	return env.Done()
}
