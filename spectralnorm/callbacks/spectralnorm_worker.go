package callbacks

import "NestedScribbleBenchmark/spectralnorm/messages/spectralnorm"
import spectralnorm_2 "NestedScribbleBenchmark/spectralnorm/results/spectralnorm"

type SpectralNorm_Worker_Env interface {
	Finish_From_Master(finish_msg spectralnorm.Finish)
	Done() spectralnorm_2.Worker_Result
	ResultFrom_SpectralNorm_Worker(result spectralnorm_2.Worker_Result)
	To_SpectralNorm_Worker_Env() SpectralNorm_Worker_Env
	TimesResult_To_Master() spectralnorm.TimesResult
	TimesTask_From_Master(timestask_msg spectralnorm.TimesTask)
}

type SpectralNormWorkerState struct {
	V []float64
}

func (s *SpectralNormWorkerState) Finish_From_Master(finish_msg spectralnorm.Finish) {
}

func (s *SpectralNormWorkerState) Done() spectralnorm_2.Worker_Result {
	return spectralnorm_2.Worker_Result{}
}

func (s *SpectralNormWorkerState) ResultFrom_SpectralNorm_Worker(result spectralnorm_2.Worker_Result) {
}

func (s *SpectralNormWorkerState) To_SpectralNorm_Worker_Env() SpectralNorm_Worker_Env {
	return &SpectralNormWorkerState{}
}

func (s *SpectralNormWorkerState) TimesResult_To_Master() spectralnorm.TimesResult {
	return spectralnorm.TimesResult{Res: s.V}
}

func (s *SpectralNormWorkerState) TimesTask_From_Master(timestask_msg spectralnorm.TimesTask) {
	s.V = timestask_msg.V
	Times(timestask_msg.Ii, timestask_msg.N, s.V, timestask_msg.U)
}

func Times(ii, n int, v, u []float64) {
	ul := len(u)
	for i := ii; i < n; i++ {
		var vi float64
		for j := 0; j < ul; j++ {
			vi += u[j] / float64(A(i, j))
		}
		v[i] = vi
	}
	return
}

func A(i, j int) int {
	return (i+j)*(i+j+1)/2 + i + 1
}
