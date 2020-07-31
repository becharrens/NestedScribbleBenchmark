package spectralnorm

import (
	"ScribbleBenchmark/spectralnorm/callbacks"
	"ScribbleBenchmark/spectralnorm/results/spectralnorm"
	"runtime"
)

var nCPU = runtime.NumCPU()

type SpectralNormEnv struct {
	NCPU         int
	N            int
	SpectralNorm float64
}

func (s *SpectralNormEnv) New_Master_Env() callbacks.SpectralNorm_Master_Env {
	u := make([]float64, s.N)
	for i := range u {
		u[i] = 1
	}

	v := make([]float64, s.N)
	return &callbacks.SpectralNormMasterState{
		U:            u,
		V:            v,
		X:            nil,
		I:            0,
		NCPU:         nCPU,
		Iter:         0,
		SpectralNorm: 0,
	}
}

func (s *SpectralNormEnv) New_Worker_Env() callbacks.SpectralNorm_Worker_Env {
	return &callbacks.SpectralNormWorkerState{}
}

func (s *SpectralNormEnv) Master_Result(result spectralnorm.Master_Result) {
	s.SpectralNorm = result.SpectralNorm
}

func (s *SpectralNormEnv) Worker_Result(result spectralnorm.Worker_Result) {
}

func NewSpectralNormEnv(n int) *SpectralNormEnv {
	return &SpectralNormEnv{
		NCPU:         nCPU,
		N:            n,
		SpectralNorm: 0,
	}
}
