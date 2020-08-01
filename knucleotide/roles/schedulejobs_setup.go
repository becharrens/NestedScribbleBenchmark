package roles

import "NestedScribbleBenchmark/knucleotide/messages/schedulejobs"
import schedulejobs_2 "NestedScribbleBenchmark/knucleotide/channels/schedulejobs"
import "NestedScribbleBenchmark/knucleotide/invitations"
import "NestedScribbleBenchmark/knucleotide/callbacks"
import "sync"

func ScheduleJobs_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.ScheduleJobs_RoleSetupChan, inviteChannels invitations.ScheduleJobs_InviteSetupChan) {
	m_w_finish := make(chan schedulejobs.Finish, 1)
	w_m_frequencyresult := make(chan schedulejobs.FrequencyResult, 1)
	m_invite_m_2 := make(chan schedulejobs_2.M_Chan, 1)
	m_invite_m_invitechan_2 := make(chan invitations.ScheduleJobs_M_InviteChan, 1)
	m_w_frequencyjob := make(chan schedulejobs.FrequencyJob, 1)
	w_m_sequenceresult := make(chan schedulejobs.SequenceResult, 1)
	m_invite_m := make(chan schedulejobs_2.M_Chan, 1)
	m_invite_m_invitechan := make(chan invitations.ScheduleJobs_M_InviteChan, 1)
	m_w_sequencejob := make(chan schedulejobs.SequenceJob, 1)

	w_chan := schedulejobs_2.W_Chan{
		M_SequenceResult:  w_m_sequenceresult,
		M_SequenceJob:     m_w_sequencejob,
		M_FrequencyResult: w_m_frequencyresult,
		M_FrequencyJob:    m_w_frequencyjob,
		M_Finish:          m_w_finish,
	}
	m_chan := schedulejobs_2.M_Chan{
		W_SequenceResult:  w_m_sequenceresult,
		W_SequenceJob:     m_w_sequencejob,
		W_FrequencyResult: w_m_frequencyresult,
		W_FrequencyJob:    m_w_frequencyjob,
		W_Finish:          m_w_finish,
	}

	w_inviteChan := invitations.ScheduleJobs_W_InviteChan{}
	m_inviteChan := invitations.ScheduleJobs_M_InviteChan{
		Invite_M_To_ScheduleJobs_M_InviteChan_2: m_invite_m_invitechan_2,
		Invite_M_To_ScheduleJobs_M_InviteChan:   m_invite_m_invitechan,
		Invite_M_To_ScheduleJobs_M_2:            m_invite_m_2,
		Invite_M_To_ScheduleJobs_M:              m_invite_m,
	}

	roleChannels.M_Chan <- m_chan

	inviteChannels.M_InviteChan <- m_inviteChan

	wg.Add(1)

	w_env := callbacks.New_ScheduleJobs_W_State()
	go ScheduleJobs_W(wg, w_chan, w_inviteChan, w_env)
}
