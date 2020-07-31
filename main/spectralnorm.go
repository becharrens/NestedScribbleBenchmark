package main

import (
	"ScribbleBenchmark/benchmark"
	"ScribbleBenchmark/spectralnorm/callbacks"
	"ScribbleBenchmark/spectralnorm/protocol"
	"ScribbleBenchmark/spectralnorm/results/spectralnorm"
	"ScribbleBenchmark/spectralnorm_base"
	"runtime"
	"time"
)

var nCPU = runtime.NumCPU()

var spectralNormParams = []int{
	100, 5500,
}

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

func TimeSpectralNorm(n int) time.Duration {
	env := NewSpectralNormEnv(n)
	start := time.Now()
	protocol.SpectralNorm(env)
	elapsed := time.Since(start)
	return elapsed
}

func TimeSpectralNormBase(n int) time.Duration {
	start := time.Now()
	spectralnorm_base.SpectralNorm(n)
	elapsed := time.Since(start)
	return elapsed
}

func SpectralNormBenchmark(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(spectralNormParams, repetitions, TimeSpectralNorm)
	base_results := benchmark.TimeImpl(spectralNormParams, repetitions, TimeSpectralNormBase)
	return scribble_results, base_results
}
