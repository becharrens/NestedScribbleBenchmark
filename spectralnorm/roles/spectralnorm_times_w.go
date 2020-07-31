package roles

import "ScribbleBenchmark/spectralnorm/channels/spectralnorm_times"
import "ScribbleBenchmark/spectralnorm/invitations"
import "ScribbleBenchmark/spectralnorm/callbacks"
import "sync"

func SpectralNorm_Times_W(wg *sync.WaitGroup, roleChannels spectralnorm_times.W_Chan, inviteChannels invitations.SpectralNorm_Times_W_InviteChan, env callbacks.SpectralNorm_Times_W_Env)  {
	defer wg.Done()
	select {
		case timestask_msg := <-roleChannels.M_TimesTask:
			env.TimesTask_From_M(timestask_msg)

			timesresult_msg := env.TimesResult_To_M()
			roleChannels.M_TimesResult <- timesresult_msg

			env.Done()
			return 
		case finish_msg := <-roleChannels.M_Finish:
			env.Finish_From_M(finish_msg)

			env.Done()
			return 
	}
} 