package roles

import "ScribbleBenchmark/spectralnorm/channels/spectralnorm"
import "ScribbleBenchmark/spectralnorm/invitations"
import "ScribbleBenchmark/spectralnorm/callbacks"
import spectralnorm_2 "ScribbleBenchmark/spectralnorm/results/spectralnorm"
import "sync"

func SpectralNorm_Worker(wg *sync.WaitGroup, roleChannels spectralnorm.Worker_Chan, inviteChannels invitations.SpectralNorm_Worker_InviteChan, env callbacks.SpectralNorm_Worker_Env) spectralnorm_2.Worker_Result {
	select {
		case timestask_msg := <-roleChannels.Master_TimesTask:
			env.TimesTask_From_Master(timestask_msg)

			timesresult_msg := env.TimesResult_To_Master()
			roleChannels.Master_TimesResult <- timesresult_msg

			spectralnorm_worker_chan := <-inviteChannels.Master_Invite_To_SpectralNorm_Worker
			spectralnorm_worker_inviteChan := <-inviteChannels.Master_Invite_To_SpectralNorm_Worker_InviteChan
			spectralnorm_worker_env := env.To_SpectralNorm_Worker_Env()
			spectralnorm_worker_result := SpectralNorm_Worker(wg, spectralnorm_worker_chan, spectralnorm_worker_inviteChan, spectralnorm_worker_env)
			env.ResultFrom_SpectralNorm_Worker(spectralnorm_worker_result)

			return env.Done()
		case finish_msg := <-roleChannels.Master_Finish:
			env.Finish_From_Master(finish_msg)

			return env.Done()
	}
} 