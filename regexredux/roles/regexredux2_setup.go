package roles

import "NestedScribbleBenchmark/regexredux/messages"
import "NestedScribbleBenchmark/regexredux/channels/regexredux2"
import "NestedScribbleBenchmark/regexredux/invitations"
import "NestedScribbleBenchmark/regexredux/callbacks"
import "sync"

func RegexRedux2_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.RegexRedux2_RoleSetupChan, inviteChannels invitations.RegexRedux2_InviteSetupChan) {
	w_m_int := make(chan int, 1)
	w_m_label := make(chan messages.RegexRedux_Label, 1)
	m_invite_m := make(chan regexredux2.M_Chan, 1)
	m_invite_m_invitechan := make(chan invitations.RegexRedux2_M_InviteChan, 1)
	m_w_bytearr := make(chan []byte, 1)
	m_w_string := make(chan string, 1)
	m_w_label := make(chan messages.RegexRedux_Label, 1)

	w_chan := regexredux2.W_Chan{
		String_From_M:  m_w_string,
		Label_To_M:     w_m_label,
		Label_From_M:   m_w_label,
		Int_To_M:       w_m_int,
		ByteArr_From_M: m_w_bytearr,
	}
	m_chan := regexredux2.M_Chan{
		String_To_W:  m_w_string,
		Label_To_W:   m_w_label,
		Label_From_W: w_m_label,
		Int_From_W:   w_m_int,
		ByteArr_To_W: m_w_bytearr,
	}

	w_inviteChan := invitations.RegexRedux2_W_InviteChan{}
	m_inviteChan := invitations.RegexRedux2_M_InviteChan{
		Invite_M_To_RegexRedux2_M_InviteChan: m_invite_m_invitechan,
		Invite_M_To_RegexRedux2_M:            m_invite_m,
	}

	roleChannels.M_Chan <- m_chan

	inviteChannels.M_InviteChan <- m_inviteChan

	wg.Add(1)

	w_env := callbacks.New_RegexRedux2_W_State()
	go RegexRedux2_W(wg, w_chan, w_inviteChan, w_env)
}
