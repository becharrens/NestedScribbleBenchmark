package roles

import "NestedScribbleBenchmark/knucleotide/messages"
import "NestedScribbleBenchmark/knucleotide/channels/schedulejobs"
import "NestedScribbleBenchmark/knucleotide/invitations"
import "NestedScribbleBenchmark/knucleotide/callbacks"
import "sync"

func ScheduleJobs_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.ScheduleJobs_RoleSetupChan, inviteChannels invitations.ScheduleJobs_InviteSetupChan) {
	m_w_int := make(chan int, 1)
	w_m_string := make(chan string, 1)
	w_m_label := make(chan messages.KNucleotide_Label, 1)
	m_invite_m := make(chan schedulejobs.M_Chan, 1)
	m_invite_m_invitechan := make(chan invitations.ScheduleJobs_M_InviteChan, 1)
	m_w_bytearr := make(chan []byte, 1)
	m_w_string := make(chan string, 1)
	m_w_label := make(chan messages.KNucleotide_Label, 1)

	w_chan := schedulejobs.W_Chan{
		String_To_M:    w_m_string,
		String_From_M:  m_w_string,
		Label_To_M:     w_m_label,
		Label_From_M:   m_w_label,
		Int_From_M:     m_w_int,
		ByteArr_From_M: m_w_bytearr,
	}
	m_chan := schedulejobs.M_Chan{
		String_To_W:   m_w_string,
		String_From_W: w_m_string,
		Label_To_W:    m_w_label,
		Label_From_W:  w_m_label,
		Int_To_W:      m_w_int,
		ByteArr_To_W:  m_w_bytearr,
	}

	w_inviteChan := invitations.ScheduleJobs_W_InviteChan{}
	m_inviteChan := invitations.ScheduleJobs_M_InviteChan{
		Invite_M_To_ScheduleJobs_M_InviteChan: m_invite_m_invitechan,
		Invite_M_To_ScheduleJobs_M:            m_invite_m,
	}

	roleChannels.M_Chan <- m_chan

	inviteChannels.M_InviteChan <- m_inviteChan

	wg.Add(1)

	w_env := callbacks.New_ScheduleJobs_W_State()
	go ScheduleJobs_W(wg, w_chan, w_inviteChan, w_env)
}
