package roles

import "NestedScribbleBenchmark/spectralnorm/messages"
import "NestedScribbleBenchmark/spectralnorm/channels/spectralnorm_times"
import "NestedScribbleBenchmark/spectralnorm/invitations"
import "NestedScribbleBenchmark/spectralnorm/callbacks"
import "sync"

func SpectralNorm_Times_W(wg *sync.WaitGroup, roleChannels spectralnorm_times.W_Chan, inviteChannels invitations.SpectralNorm_Times_W_InviteChan, env callbacks.SpectralNorm_Times_W_Env) {
	defer wg.Done()
	m_choice := <-roleChannels.Label_From_M
	switch m_choice {
	case messages.TimesTask:
		ii := <-roleChannels.Int_From_M
		n := <-roleChannels.Int_From_M
		u := <-roleChannels.Vec_From_M
		v := <-roleChannels.Vec_From_M
		env.TimesTask_From_M(ii, n, u, v)

		res := env.TimesResult_To_M()
		roleChannels.Label_To_M <- messages.TimesResult
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
