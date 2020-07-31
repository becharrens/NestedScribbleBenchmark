package roles

import "ScribbleBenchmark/spectralnorm/messages/spectralnorm_times"
import spectralnorm_times_2 "ScribbleBenchmark/spectralnorm/channels/spectralnorm_times"
import "ScribbleBenchmark/spectralnorm/invitations"
import "ScribbleBenchmark/spectralnorm/callbacks"
import "sync"

func SpectralNorm_Times_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.SpectralNorm_Times_RoleSetupChan, inviteChannels invitations.SpectralNorm_Times_InviteSetupChan)  {
	m_w_finish := make(chan spectralnorm_times.Finish, 1)
	w_m_timesresult := make(chan spectralnorm_times.TimesResult, 1)
	m_invite_m := make(chan spectralnorm_times_2.M_Chan, 1)
	m_invite_m_invitechan := make(chan invitations.SpectralNorm_Times_M_InviteChan, 1)
	m_w_timestask := make(chan spectralnorm_times.TimesTask, 1)

	w_chan := spectralnorm_times_2.W_Chan{
		M_TimesTask: m_w_timestask,
		M_TimesResult: w_m_timesresult,
		M_Finish: m_w_finish,
	}
	m_chan := spectralnorm_times_2.M_Chan{
		W_TimesTask: m_w_timestask,
		W_TimesResult: w_m_timesresult,
		W_Finish: m_w_finish,
	}

	w_inviteChan := invitations.SpectralNorm_Times_W_InviteChan{

	}
	m_inviteChan := invitations.SpectralNorm_Times_M_InviteChan{
		Invite_M_To_SpectralNorm_Times_M_InviteChan: m_invite_m_invitechan,
		Invite_M_To_SpectralNorm_Times_M: m_invite_m,
	}

	roleChannels.M_Chan <- m_chan

	inviteChannels.M_InviteChan <- m_inviteChan

	wg.Add(1)

	w_env := callbacks.New_SpectralNorm_Times_W_State()
	go SpectralNorm_Times_W(wg, w_chan, w_inviteChan, w_env)
} 