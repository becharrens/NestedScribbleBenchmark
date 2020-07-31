package roles

import "ScribbleBenchmark/regexredux/messages/regexredux2"
import regexredux2_2 "ScribbleBenchmark/regexredux/channels/regexredux2"
import "ScribbleBenchmark/regexredux/invitations"
import "ScribbleBenchmark/regexredux/callbacks"
import "sync"

func RegexRedux2_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.RegexRedux2_RoleSetupChan, inviteChannels invitations.RegexRedux2_InviteSetupChan)  {
	w_m_length := make(chan regexredux2.Length, 1)
	m_w_calclength := make(chan regexredux2.CalcLength, 1)
	w_m_nummatches := make(chan regexredux2.NumMatches, 1)
	m_invite_m := make(chan regexredux2_2.M_Chan, 1)
	m_invite_m_invitechan := make(chan invitations.RegexRedux2_M_InviteChan, 1)
	m_w_task := make(chan regexredux2.Task, 1)

	w_chan := regexredux2_2.W_Chan{
		M_Task: m_w_task,
		M_NumMatches: w_m_nummatches,
		M_Length: w_m_length,
		M_CalcLength: m_w_calclength,
	}
	m_chan := regexredux2_2.M_Chan{
		W_Task: m_w_task,
		W_NumMatches: w_m_nummatches,
		W_Length: w_m_length,
		W_CalcLength: m_w_calclength,
	}

	w_inviteChan := invitations.RegexRedux2_W_InviteChan{

	}
	m_inviteChan := invitations.RegexRedux2_M_InviteChan{
		Invite_M_To_RegexRedux2_M_InviteChan: m_invite_m_invitechan,
		Invite_M_To_RegexRedux2_M: m_invite_m,
	}

	roleChannels.M_Chan <- m_chan

	inviteChannels.M_InviteChan <- m_inviteChan

	wg.Add(1)

	w_env := callbacks.New_RegexRedux2_W_State()
	go RegexRedux2_W(wg, w_chan, w_inviteChan, w_env)
} 