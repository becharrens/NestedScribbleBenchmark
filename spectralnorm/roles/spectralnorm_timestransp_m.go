package roles

import "ScribbleBenchmark/spectralnorm/channels/spectralnorm_timestransp"
import "ScribbleBenchmark/spectralnorm/invitations"
import "ScribbleBenchmark/spectralnorm/callbacks"
import spectralnorm_timestransp_2 "ScribbleBenchmark/spectralnorm/results/spectralnorm_timestransp"
import "sync"

func SpectralNorm_TimesTransp_M(wg *sync.WaitGroup, roleChannels spectralnorm_timestransp.M_Chan, inviteChannels invitations.SpectralNorm_TimesTransp_M_InviteChan, env callbacks.SpectralNorm_TimesTransp_M_Env) spectralnorm_timestransp_2.M_Result {
	m_choice := env.M_Choice()
	switch m_choice {
	case callbacks.SpectralNorm_TimesTransp_M_TimesTranspTask:
		timestransptask_msg := env.TimesTranspTask_To_W()
		roleChannels.W_TimesTranspTask <- timestransptask_msg

		env.SpectralNorm_TimesTransp_Setup()
		spectralnorm_timestransp_rolechan := invitations.SpectralNorm_TimesTransp_RoleSetupChan{
			M_Chan: inviteChannels.Invite_M_To_SpectralNorm_TimesTransp_M,
		}
		spectralnorm_timestransp_invitechan := invitations.SpectralNorm_TimesTransp_InviteSetupChan{
			M_InviteChan: inviteChannels.Invite_M_To_SpectralNorm_TimesTransp_M_InviteChan,
		}
		SpectralNorm_TimesTransp_SendCommChannels(wg, spectralnorm_timestransp_rolechan, spectralnorm_timestransp_invitechan)

		spectralnorm_timestransp_m_chan := <-inviteChannels.Invite_M_To_SpectralNorm_TimesTransp_M
		spectralnorm_timestransp_m_inviteChan := <-inviteChannels.Invite_M_To_SpectralNorm_TimesTransp_M_InviteChan
		spectralnorm_timestransp_m_env := env.To_SpectralNorm_TimesTransp_M_Env()
		spectralnorm_timestransp_m_result := SpectralNorm_TimesTransp_M(wg, spectralnorm_timestransp_m_chan, spectralnorm_timestransp_m_inviteChan, spectralnorm_timestransp_m_env)
		env.ResultFrom_SpectralNorm_TimesTransp_M(spectralnorm_timestransp_m_result)

		timestranspresult_msg := <-roleChannels.W_TimesTranspResult
		env.TimesTranspResult_From_W(timestranspresult_msg)

		return env.Done()
	case callbacks.SpectralNorm_TimesTransp_M_Finish:
		finish_msg := env.Finish_To_W()
		roleChannels.W_Finish <- finish_msg

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
} 