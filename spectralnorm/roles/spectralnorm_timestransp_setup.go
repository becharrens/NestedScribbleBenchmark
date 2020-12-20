package roles

import "NestedScribbleBenchmark/spectralnorm/messages"
import "NestedScribbleBenchmark/spectralnorm/channels/spectralnorm_timestransp"
import "NestedScribbleBenchmark/spectralnorm/invitations"
import "NestedScribbleBenchmark/spectralnorm/callbacks"
import "sync"

func SpectralNorm_TimesTransp_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.SpectralNorm_TimesTransp_RoleSetupChan, inviteChannels invitations.SpectralNorm_TimesTransp_InviteSetupChan) {
	w_m_vec := make(chan []float64, 1)
	w_m_label := make(chan messages.SpectralNorm_Label, 1)
	m_invite_m := make(chan spectralnorm_timestransp.M_Chan, 1)
	m_invite_m_invitechan := make(chan invitations.SpectralNorm_TimesTransp_M_InviteChan, 1)
	m_w_vec := make(chan []float64, 1)
	m_w_int := make(chan int, 1)
	m_w_label := make(chan messages.SpectralNorm_Label, 1)

	w_chan := spectralnorm_timestransp.W_Chan{
		Vec_To_M:     w_m_vec,
		Vec_From_M:   m_w_vec,
		Label_To_M:   w_m_label,
		Label_From_M: m_w_label,
		Int_From_M:   m_w_int,
	}
	m_chan := spectralnorm_timestransp.M_Chan{
		Vec_To_W:     m_w_vec,
		Vec_From_W:   w_m_vec,
		Label_To_W:   m_w_label,
		Label_From_W: w_m_label,
		Int_To_W:     m_w_int,
	}

	w_inviteChan := invitations.SpectralNorm_TimesTransp_W_InviteChan{}
	m_inviteChan := invitations.SpectralNorm_TimesTransp_M_InviteChan{
		Invite_M_To_SpectralNorm_TimesTransp_M_InviteChan: m_invite_m_invitechan,
		Invite_M_To_SpectralNorm_TimesTransp_M:            m_invite_m,
	}

	roleChannels.M_Chan <- m_chan

	inviteChannels.M_InviteChan <- m_inviteChan

	wg.Add(1)

	w_env := callbacks.New_SpectralNorm_TimesTransp_W_State()
	go SpectralNorm_TimesTransp_W(wg, w_chan, w_inviteChan, w_env)
}
