package main

import (
	"ScribbleBenchmark/benchmark"
	"ScribbleBenchmark/fannkuch/callbacks"
	"ScribbleBenchmark/fannkuch/protocol"
	"ScribbleBenchmark/fannkuch/results/fannkuch"
	"ScribbleBenchmark/fannkuch_base"
	"time"
)

const NCHUNKS = 720

var fannkuchParams = []int{4, 5, 6, 7, 8, 9, 10, 11, 12}

type FannkuchEnv struct {
	N         int
	ChunkSize int
	Fact      []int
	Res       int
	Checksum  int
}

func (f *FannkuchEnv) New_Main_Env() callbacks.Fannkuch_Main_Env {
	return &callbacks.FannkuchMainState{
		Fact:      f.Fact,
		N:         f.N,
		ChunkSize: f.ChunkSize,
		Res:       0,
		Chk:       0,
	}
}

func (f *FannkuchEnv) New_Worker_Env() callbacks.Fannkuch_Worker_Env {
	return &callbacks.FannkuchWorkerState{}
}

func (f *FannkuchEnv) Main_Result(result fannkuch.Main_Result) {
	f.Checksum = result.Chk
	f.Res = result.Res
}

func (f *FannkuchEnv) Worker_Result(result fannkuch.Worker_Result) {
}

func NewFannkuchEnv(n int) *FannkuchEnv {
	fact := make([]int, n+1)
	fact[0] = 1
	for i := 1; i < len(fact); i++ {
		fact[i] = fact[i-1] * i
	}

	chunksz := (fact[n] + NCHUNKS - 1) / NCHUNKS
	chunksz += chunksz % 2
	return &FannkuchEnv{
		N:         n,
		ChunkSize: chunksz,
		Fact:      fact,
		Res:       0,
		Checksum:  0,
	}
}

func TimeFannkuch(n int) time.Duration {
	env := NewFannkuchEnv(n)
	start := time.Now()
	protocol.Fannkuch(env)
	elapsed := time.Since(start)
	return elapsed
}

func TimeFannkuchBase(n int) time.Duration {
	start := time.Now()
	_, _ = fannkuch_base.Fannkuch(n)
	return time.Since(start)
}

func FannkuchBenchmark(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(fannkuchParams, repetitions, TimeFannkuch)
	base_results := benchmark.TimeImpl(fannkuchParams, repetitions, TimeFannkuchBase)
	return scribble_results, base_results
}
