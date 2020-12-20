package callbacks

import "NestedScribbleBenchmark/spectralnorm/results/spectralnorm"

type SpectralNorm_Worker_Env interface {
	Finish_From_Master()
	Done() spectralnorm.Worker_Result
	ResultFrom_SpectralNorm_Worker(result spectralnorm.Worker_Result)
	To_SpectralNorm_Worker_Env() SpectralNorm_Worker_Env
	TimesResult_To_Master() []float64
	TimesTask_From_Master(ii int, n int, u []float64, v []float64)
}

type SpectralNormWorkerState struct {
	V []float64
}

func (s *SpectralNormWorkerState) Finish_From_Master() {
}

func (s *SpectralNormWorkerState) Done() spectralnorm.Worker_Result {
	return spectralnorm.Worker_Result{}
}

func (s *SpectralNormWorkerState) ResultFrom_SpectralNorm_Worker(result spectralnorm.Worker_Result) {
}

func (s *SpectralNormWorkerState) To_SpectralNorm_Worker_Env() SpectralNorm_Worker_Env {
	return &SpectralNormWorkerState{}
}

func (s *SpectralNormWorkerState) TimesResult_To_Master() []float64 {
	return s.V
}

func (s *SpectralNormWorkerState) TimesTask_From_Master(ii int, n int, u []float64, v []float64) {
	s.V = v
	Times(ii, n, s.V, u)
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
