package roles

import "ScribbleBenchmark/spectralnorm/channels/spectralnorm"
import "ScribbleBenchmark/spectralnorm/invitations"
import "ScribbleBenchmark/spectralnorm/callbacks"
import spectralnorm_2 "ScribbleBenchmark/spectralnorm/results/spectralnorm"
import "sync"

func SpectralNorm_Master(wg *sync.WaitGroup, roleChannels spectralnorm.Master_Chan, inviteChannels invitations.SpectralNorm_Master_InviteChan, env callbacks.SpectralNorm_Master_Env) spectralnorm_2.Master_Result {
	master_choice := env.Master_Choice()
	switch master_choice {
	case callbacks.SpectralNorm_Master_TimesTask:
		timestask_msg := env.TimesTask_To_Worker()
		roleChannels.Worker_TimesTask <- timestask_msg

		env.SpectralNorm_Times_Setup()
		spectralnorm_times_rolechan := invitations.SpectralNorm_Times_RoleSetupChan{
			M_Chan: inviteChannels.Invite_Master_To_SpectralNorm_Times_M,
		}
		spectralnorm_times_invitechan := invitations.SpectralNorm_Times_InviteSetupChan{
			M_InviteChan: inviteChannels.Invite_Master_To_SpectralNorm_Times_M_InviteChan,
		}
		SpectralNorm_Times_SendCommChannels(wg, spectralnorm_times_rolechan, spectralnorm_times_invitechan)

		spectralnorm_times_m_chan := <-inviteChannels.Invite_Master_To_SpectralNorm_Times_M
		spectralnorm_times_m_inviteChan := <-inviteChannels.Invite_Master_To_SpectralNorm_Times_M_InviteChan
		spectralnorm_times_m_env := env.To_SpectralNorm_Times_M_Env()
		spectralnorm_times_m_result := SpectralNorm_Times_M(wg, spectralnorm_times_m_chan, spectralnorm_times_m_inviteChan, spectralnorm_times_m_env)
		env.ResultFrom_SpectralNorm_Times_M(spectralnorm_times_m_result)

		timesresult_msg := <-roleChannels.Worker_TimesResult
		env.TimesResult_From_Worker(timesresult_msg)

		env.SpectralNorm_TimesTransp_Setup()
		spectralnorm_timestransp_rolechan := invitations.SpectralNorm_TimesTransp_RoleSetupChan{
			M_Chan: inviteChannels.Invite_Master_To_SpectralNorm_TimesTransp_M,
		}
		spectralnorm_timestransp_invitechan := invitations.SpectralNorm_TimesTransp_InviteSetupChan{
			M_InviteChan: inviteChannels.Invite_Master_To_SpectralNorm_TimesTransp_M_InviteChan,
		}
		SpectralNorm_TimesTransp_SendCommChannels(wg, spectralnorm_timestransp_rolechan, spectralnorm_timestransp_invitechan)

		spectralnorm_timestransp_m_chan := <-inviteChannels.Invite_Master_To_SpectralNorm_TimesTransp_M
		spectralnorm_timestransp_m_inviteChan := <-inviteChannels.Invite_Master_To_SpectralNorm_TimesTransp_M_InviteChan
		spectralnorm_timestransp_m_env := env.To_SpectralNorm_TimesTransp_M_Env()
		spectralnorm_timestransp_m_result := SpectralNorm_TimesTransp_M(wg, spectralnorm_timestransp_m_chan, spectralnorm_timestransp_m_inviteChan, spectralnorm_timestransp_m_env)
		env.ResultFrom_SpectralNorm_TimesTransp_M(spectralnorm_timestransp_m_result)

		env.SpectralNorm_Times_Setup_2()
		spectralnorm_times_rolechan_2 := invitations.SpectralNorm_Times_RoleSetupChan{
			M_Chan: inviteChannels.Invite_Master_To_SpectralNorm_Times_M_2,
		}
		spectralnorm_times_invitechan_2 := invitations.SpectralNorm_Times_InviteSetupChan{
			M_InviteChan: inviteChannels.Invite_Master_To_SpectralNorm_Times_M_InviteChan_2,
		}
		SpectralNorm_Times_SendCommChannels(wg, spectralnorm_times_rolechan_2, spectralnorm_times_invitechan_2)

		spectralnorm_times_m_chan_2 := <-inviteChannels.Invite_Master_To_SpectralNorm_Times_M_2
		spectralnorm_times_m_inviteChan_2 := <-inviteChannels.Invite_Master_To_SpectralNorm_Times_M_InviteChan_2
		spectralnorm_times_m_env_2 := env.To_SpectralNorm_Times_M_Env_2()
		spectralnorm_times_m_result_2 := SpectralNorm_Times_M(wg, spectralnorm_times_m_chan_2, spectralnorm_times_m_inviteChan_2, spectralnorm_times_m_env_2)
		env.ResultFrom_SpectralNorm_Times_M_2(spectralnorm_times_m_result_2)

		env.SpectralNorm_TimesTransp_Setup_2()
		spectralnorm_timestransp_rolechan_2 := invitations.SpectralNorm_TimesTransp_RoleSetupChan{
			M_Chan: inviteChannels.Invite_Master_To_SpectralNorm_TimesTransp_M_2,
		}
		spectralnorm_timestransp_invitechan_2 := invitations.SpectralNorm_TimesTransp_InviteSetupChan{
			M_InviteChan: inviteChannels.Invite_Master_To_SpectralNorm_TimesTransp_M_InviteChan_2,
		}
		SpectralNorm_TimesTransp_SendCommChannels(wg, spectralnorm_timestransp_rolechan_2, spectralnorm_timestransp_invitechan_2)

		spectralnorm_timestransp_m_chan_2 := <-inviteChannels.Invite_Master_To_SpectralNorm_TimesTransp_M_2
		spectralnorm_timestransp_m_inviteChan_2 := <-inviteChannels.Invite_Master_To_SpectralNorm_TimesTransp_M_InviteChan_2
		spectralnorm_timestransp_m_env_2 := env.To_SpectralNorm_TimesTransp_M_Env_2()
		spectralnorm_timestransp_m_result_2 := SpectralNorm_TimesTransp_M(wg, spectralnorm_timestransp_m_chan_2, spectralnorm_timestransp_m_inviteChan_2, spectralnorm_timestransp_m_env_2)
		env.ResultFrom_SpectralNorm_TimesTransp_M_2(spectralnorm_timestransp_m_result_2)

		env.SpectralNorm_Setup()
		spectralnorm_rolechan := invitations.SpectralNorm_RoleSetupChan{
			Master_Chan: inviteChannels.Invite_Master_To_SpectralNorm_Master,
			Worker_Chan: inviteChannels.Invite_Worker_To_SpectralNorm_Worker,
		}
		spectralnorm_invitechan := invitations.SpectralNorm_InviteSetupChan{
			Master_InviteChan: inviteChannels.Invite_Master_To_SpectralNorm_Master_InviteChan,
			Worker_InviteChan: inviteChannels.Invite_Worker_To_SpectralNorm_Worker_InviteChan,
		}
		SpectralNorm_SendCommChannels(wg, spectralnorm_rolechan, spectralnorm_invitechan)

		spectralnorm_master_chan := <-inviteChannels.Invite_Master_To_SpectralNorm_Master
		spectralnorm_master_inviteChan := <-inviteChannels.Invite_Master_To_SpectralNorm_Master_InviteChan
		spectralnorm_master_env := env.To_SpectralNorm_Master_Env()
		spectralnorm_master_result := SpectralNorm_Master(wg, spectralnorm_master_chan, spectralnorm_master_inviteChan, spectralnorm_master_env)
		env.ResultFrom_SpectralNorm_Master(spectralnorm_master_result)

		return env.Done()
	case callbacks.SpectralNorm_Master_Finish:
		finish_msg := env.Finish_To_Worker()
		roleChannels.Worker_Finish <- finish_msg

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
} 