package main

import (
	"NestedScribbleBenchmark/benchmark"
	"NestedScribbleBenchmark/fannkuch/callbacks"
	"NestedScribbleBenchmark/fannkuch/protocol"
	"NestedScribbleBenchmark/fannkuch/results/fannkuch"
	"NestedScribbleBenchmark/fannkuch_base"
	"time"
)

const NCHUNKS = 720

var fannkuchParams = []int{4, 5, 6, 7, 8, 9, 10, 11, 12}

type FannkuchEnv struct {
	N         int
	ChunkSize int
	Res       int
	Checksum  int
}

func (f *FannkuchEnv) New_Main_Env() callbacks.Fannkuch_Main_Env {
	return &callbacks.FannkuchMainState{
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
	callbacks.Fact = make([]int, n+1)
	callbacks.Fact[0] = 1
	for i := 1; i < len(callbacks.Fact); i++ {
		callbacks.Fact[i] = callbacks.Fact[i-1] * i
	}

	chunksz := (callbacks.Fact[n] + NCHUNKS - 1) / NCHUNKS
	chunksz += chunksz % 2
	return &FannkuchEnv{
		N:         n,
		ChunkSize: chunksz,
		Res:       0,
		Checksum:  0,
	}
}

func TimeFannkuch(n int) time.Duration {
	start := time.Now()
	env := NewFannkuchEnv(n)
	protocol.Fannkuch(env)
	elapsed := time.Since(start)
	// fmt.Println(env.Res, env.Checksum)
	return elapsed
}

func TimeFannkuchBase(n int) time.Duration {
	start := time.Now()
	fannkuch_base.Fannkuch(n)
	// fmt.Println(res, chk)
	return time.Since(start)
}

func FannkuchBenchmark(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(fannkuchParams, repetitions, TimeFannkuch)
	base_results := benchmark.TimeImpl(fannkuchParams, repetitions, TimeFannkuchBase)
	return scribble_results, base_results
}
