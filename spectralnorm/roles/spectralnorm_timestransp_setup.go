package roles

import "NestedScribbleBenchmark/spectralnorm/messages/spectralnorm_timestransp"
import spectralnorm_timestransp_2 "NestedScribbleBenchmark/spectralnorm/channels/spectralnorm_timestransp"
import "NestedScribbleBenchmark/spectralnorm/invitations"
import "NestedScribbleBenchmark/spectralnorm/callbacks"
import "sync"

func SpectralNorm_TimesTransp_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.SpectralNorm_TimesTransp_RoleSetupChan, inviteChannels invitations.SpectralNorm_TimesTransp_InviteSetupChan) {
	m_w_finish := make(chan spectralnorm_timestransp.Finish, 1)
	w_m_timestranspresult := make(chan spectralnorm_timestransp.TimesTranspResult, 1)
	m_invite_m := make(chan spectralnorm_timestransp_2.M_Chan, 1)
	m_invite_m_invitechan := make(chan invitations.SpectralNorm_TimesTransp_M_InviteChan, 1)
	m_w_timestransptask := make(chan spectralnorm_timestransp.TimesTranspTask, 1)

	w_chan := spectralnorm_timestransp_2.W_Chan{
		M_TimesTranspTask:   m_w_timestransptask,
		M_TimesTranspResult: w_m_timestranspresult,
		M_Finish:            m_w_finish,
	}
	m_chan := spectralnorm_timestransp_2.M_Chan{
		W_TimesTranspTask:   m_w_timestransptask,
		W_TimesTranspResult: w_m_timestranspresult,
		W_Finish:            m_w_finish,
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
