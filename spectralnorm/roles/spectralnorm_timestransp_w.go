package roles

import "NestedScribbleBenchmark/spectralnorm/messages"
import "NestedScribbleBenchmark/spectralnorm/channels/spectralnorm_timestransp"
import "NestedScribbleBenchmark/spectralnorm/invitations"
import "NestedScribbleBenchmark/spectralnorm/callbacks"
import "sync"

func SpectralNorm_TimesTransp_W(wg *sync.WaitGroup, roleChannels spectralnorm_timestransp.W_Chan, inviteChannels invitations.SpectralNorm_TimesTransp_W_InviteChan, env callbacks.SpectralNorm_TimesTransp_W_Env) {
	defer wg.Done()
	m_choice := <-roleChannels.Label_From_M
	switch m_choice {
	case messages.TimesTranspTask:
		ii := <-roleChannels.Int_From_M
		n := <-roleChannels.Int_From_M
		u := <-roleChannels.Vec_From_M
		v := <-roleChannels.Vec_From_M
		env.TimesTranspTask_From_M(ii, n, u, v)

		res := env.TimesTranspResult_To_M()
		roleChannels.Label_To_M <- messages.TimesTranspResult
		roleChannels.Vec_To_M <- res

		env.Done()
		return
	case messages.Finish:
		env.Finish_From_M()

		env.Done()
		return
	default:
		panic("Invalid choice was made")
	}
}
