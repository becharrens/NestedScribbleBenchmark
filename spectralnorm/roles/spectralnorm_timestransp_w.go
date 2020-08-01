package roles

import "NestedScribbleBenchmark/spectralnorm/channels/spectralnorm_timestransp"
import "NestedScribbleBenchmark/spectralnorm/invitations"
import "NestedScribbleBenchmark/spectralnorm/callbacks"
import "sync"

func SpectralNorm_TimesTransp_W(wg *sync.WaitGroup, roleChannels spectralnorm_timestransp.W_Chan, inviteChannels invitations.SpectralNorm_TimesTransp_W_InviteChan, env callbacks.SpectralNorm_TimesTransp_W_Env) {
	defer wg.Done()
	select {
	case timestransptask_msg := <-roleChannels.M_TimesTranspTask:
		env.TimesTranspTask_From_M(timestransptask_msg)

		timestranspresult_msg := env.TimesTranspResult_To_M()
		roleChannels.M_TimesTranspResult <- timestranspresult_msg

		env.Done()
		return
	case finish_msg := <-roleChannels.M_Finish:
		env.Finish_From_M(finish_msg)

		env.Done()
		return
	}
}
