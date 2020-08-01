package main

import (
	"NestedScribbleBenchmark/benchmark"
	"NestedScribbleBenchmark/primesieve/callbacks"
	"NestedScribbleBenchmark/primesieve/protocol"
	"NestedScribbleBenchmark/primesieve/results/primesieve"
	"NestedScribbleBenchmark/primesieve_base"
	"time"
)

type PrimeSieveEnv struct {
	N      int
	Primes []int
}

var primesieveParams = []int{
	100, 1100, 2100, 3100, 4100, 5100, 6100, 7100, 8100, 9100,
}

func (p *PrimeSieveEnv) New_Master_Env() callbacks.PrimeSieve_Master_Env {
	return &callbacks.PrimeSieveMasterState{
		N:      p.N,
		Primes: []int{2},
	}
}

func (p *PrimeSieveEnv) New_Worker_Env() callbacks.PrimeSieve_Worker_Env {
	return &callbacks.PrimeSieveWorkerState{}
}

func (p *PrimeSieveEnv) Master_Result(result primesieve.Master_Result) {
	p.Primes = result.Primes
}

func (p *PrimeSieveEnv) Worker_Result(result primesieve.Worker_Result) {
}

func NewPrimeSieveEnv(n int) *PrimeSieveEnv {
	return &PrimeSieveEnv{
		N:      n,
		Primes: nil,
	}
}

func TimePrimeSieve(n int) time.Duration {
	env := NewPrimeSieveEnv(n)
	start := time.Now()
	protocol.PrimeSieve(env)
	elapsed := time.Since(start)
	return elapsed
}

func TimePrimeSieveBase(n int) time.Duration {
	start := time.Now()
	_ = primesieve_base.PrimeSieve(n)
	return time.Since(start)
}

func PrimeSieveBenchmark(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(primesieveParams, repetitions, TimePrimeSieve)
	base_results := benchmark.TimeImpl(primesieveParams, repetitions, TimePrimeSieveBase)
	return scribble_results, base_results
}
