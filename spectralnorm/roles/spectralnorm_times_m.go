package roles

import "NestedScribbleBenchmark/spectralnorm/messages"
import "NestedScribbleBenchmark/spectralnorm/channels/spectralnorm_times"
import "NestedScribbleBenchmark/spectralnorm/invitations"
import "NestedScribbleBenchmark/spectralnorm/callbacks"
import spectralnorm_times_2 "NestedScribbleBenchmark/spectralnorm/results/spectralnorm_times"
import "sync"

func SpectralNorm_Times_M(wg *sync.WaitGroup, roleChannels spectralnorm_times.M_Chan, inviteChannels invitations.SpectralNorm_Times_M_InviteChan, env callbacks.SpectralNorm_Times_M_Env) spectralnorm_times_2.M_Result {
	m_choice := env.M_Choice()
	switch m_choice {
	case callbacks.SpectralNorm_Times_M_TimesTask:
		ii, n, u, v := env.TimesTask_To_W()
		roleChannels.Label_To_W <- messages.TimesTask
		roleChannels.Int_To_W <- ii
		roleChannels.Int_To_W <- n
		roleChannels.Vec_To_W <- u
		roleChannels.Vec_To_W <- v

		env.SpectralNorm_Times_Setup()

		spectralnorm_times_rolechan := invitations.SpectralNorm_Times_RoleSetupChan{
			M_Chan: inviteChannels.Invite_M_To_SpectralNorm_Times_M,
		}
		spectralnorm_times_invitechan := invitations.SpectralNorm_Times_InviteSetupChan{
			M_InviteChan: inviteChannels.Invite_M_To_SpectralNorm_Times_M_InviteChan,
		}
		SpectralNorm_Times_SendCommChannels(wg, spectralnorm_times_rolechan, spectralnorm_times_invitechan)

		spectralnorm_times_m_chan := <-inviteChannels.Invite_M_To_SpectralNorm_Times_M
		spectralnorm_times_m_inviteChan := <-inviteChannels.Invite_M_To_SpectralNorm_Times_M_InviteChan
		spectralnorm_times_m_env := env.To_SpectralNorm_Times_M_Env()
		spectralnorm_times_m_result := SpectralNorm_Times_M(wg, spectralnorm_times_m_chan, spectralnorm_times_m_inviteChan, spectralnorm_times_m_env)
		env.ResultFrom_SpectralNorm_Times_M(spectralnorm_times_m_result)

		<-roleChannels.Label_From_W
		res := <-roleChannels.Vec_From_W
		env.TimesResult_From_W(res)

		return env.Done()
	case callbacks.SpectralNorm_Times_M_Finish:
		env.Finish_To_W()
		roleChannels.Label_To_W <- messages.Finish

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
}
