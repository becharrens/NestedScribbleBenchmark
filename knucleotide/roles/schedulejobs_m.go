package roles

import "ScribbleBenchmark/knucleotide/channels/schedulejobs"
import "ScribbleBenchmark/knucleotide/invitations"
import "ScribbleBenchmark/knucleotide/callbacks"
import schedulejobs_2 "ScribbleBenchmark/knucleotide/results/schedulejobs"
import "sync"

func ScheduleJobs_M(wg *sync.WaitGroup, roleChannels schedulejobs.M_Chan, inviteChannels invitations.ScheduleJobs_M_InviteChan, env callbacks.ScheduleJobs_M_Env) schedulejobs_2.M_Result {
	m_choice := env.M_Choice()
	switch m_choice {
	case callbacks.ScheduleJobs_M_SequenceJob:
		sequencejob_msg := env.SequenceJob_To_W()
		roleChannels.W_SequenceJob <- sequencejob_msg

		env.ScheduleJobs_Setup()
		schedulejobs_rolechan := invitations.ScheduleJobs_RoleSetupChan{
			M_Chan: inviteChannels.Invite_M_To_ScheduleJobs_M,
		}
		schedulejobs_invitechan := invitations.ScheduleJobs_InviteSetupChan{
			M_InviteChan: inviteChannels.Invite_M_To_ScheduleJobs_M_InviteChan,
		}
		ScheduleJobs_SendCommChannels(wg, schedulejobs_rolechan, schedulejobs_invitechan)

		schedulejobs_m_chan := <-inviteChannels.Invite_M_To_ScheduleJobs_M
		schedulejobs_m_inviteChan := <-inviteChannels.Invite_M_To_ScheduleJobs_M_InviteChan
		schedulejobs_m_env := env.To_ScheduleJobs_M_Env()
		schedulejobs_m_result := ScheduleJobs_M(wg, schedulejobs_m_chan, schedulejobs_m_inviteChan, schedulejobs_m_env)
		env.ResultFrom_ScheduleJobs_M(schedulejobs_m_result)

		sequenceresult_msg := <-roleChannels.W_SequenceResult
		env.SequenceResult_From_W(sequenceresult_msg)

		return env.Done()
	case callbacks.ScheduleJobs_M_FrequencyJob:
		frequencyjob_msg := env.FrequencyJob_To_W()
		roleChannels.W_FrequencyJob <- frequencyjob_msg

		env.ScheduleJobs_Setup_2()
		schedulejobs_rolechan_2 := invitations.ScheduleJobs_RoleSetupChan{
			M_Chan: inviteChannels.Invite_M_To_ScheduleJobs_M_2,
		}
		schedulejobs_invitechan_2 := invitations.ScheduleJobs_InviteSetupChan{
			M_InviteChan: inviteChannels.Invite_M_To_ScheduleJobs_M_InviteChan_2,
		}
		ScheduleJobs_SendCommChannels(wg, schedulejobs_rolechan_2, schedulejobs_invitechan_2)

		schedulejobs_m_chan_2 := <-inviteChannels.Invite_M_To_ScheduleJobs_M_2
		schedulejobs_m_inviteChan_2 := <-inviteChannels.Invite_M_To_ScheduleJobs_M_InviteChan_2
		schedulejobs_m_env_2 := env.To_ScheduleJobs_M_Env_2()
		schedulejobs_m_result_2 := ScheduleJobs_M(wg, schedulejobs_m_chan_2, schedulejobs_m_inviteChan_2, schedulejobs_m_env_2)
		env.ResultFrom_ScheduleJobs_M_2(schedulejobs_m_result_2)

		frequencyresult_msg := <-roleChannels.W_FrequencyResult
		env.FrequencyResult_From_W(frequencyresult_msg)

		return env.Done()
	case callbacks.ScheduleJobs_M_Finish:
		finish_msg := env.Finish_To_W()
		roleChannels.W_Finish <- finish_msg

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
} 