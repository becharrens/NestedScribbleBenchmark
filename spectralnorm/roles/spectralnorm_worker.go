package roles

import "NestedScribbleBenchmark/spectralnorm/messages"
import "NestedScribbleBenchmark/spectralnorm/channels/spectralnorm"
import "NestedScribbleBenchmark/spectralnorm/invitations"
import "NestedScribbleBenchmark/spectralnorm/callbacks"
import spectralnorm_2 "NestedScribbleBenchmark/spectralnorm/results/spectralnorm"
import "sync"

func SpectralNorm_Worker(wg *sync.WaitGroup, roleChannels spectralnorm.Worker_Chan, inviteChannels invitations.SpectralNorm_Worker_InviteChan, env callbacks.SpectralNorm_Worker_Env) spectralnorm_2.Worker_Result {
	master_choice := <-roleChannels.Label_From_Master
	switch master_choice {
	case messages.TimesTask:
		ii := <-roleChannels.Int_From_Master
		n := <-roleChannels.Int_From_Master
		u := <-roleChannels.Vec_From_Master
		v := <-roleChannels.Vec_From_Master
		env.TimesTask_From_Master(ii, n, u, v)

		res := env.TimesResult_To_Master()
		roleChannels.Label_To_Master <- messages.TimesResult
		roleChannels.Vec_To_Master <- res

		<-roleChannels.Label_From_Master
		spectralnorm_worker_chan := <-inviteChannels.Master_Invite_To_SpectralNorm_Worker
		spectralnorm_worker_inviteChan := <-inviteChannels.Master_Invite_To_SpectralNorm_Worker_InviteChan
		spectralnorm_worker_env := env.To_SpectralNorm_Worker_Env()
		spectralnorm_worker_result := SpectralNorm_Worker(wg, spectralnorm_worker_chan, spectralnorm_worker_inviteChan, spectralnorm_worker_env)
		env.ResultFrom_SpectralNorm_Worker(spectralnorm_worker_result)

		return env.Done()
	case messages.Finish:
		env.Finish_From_Master()

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
}
