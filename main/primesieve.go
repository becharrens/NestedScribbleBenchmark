package main

import (
	"NestedScribbleBenchmark/benchmark"
	"NestedScribbleBenchmark/primesieve/callbacks"
	"NestedScribbleBenchmark/primesieve/protocol"
	"NestedScribbleBenchmark/primesieve/results/primesieve"
	"NestedScribbleBenchmark/primesieve_base"
	callbacks2 "NestedScribbleBenchmark/primesieve_base/primesieve2/callbacks"
	protocol_2 "NestedScribbleBenchmark/primesieve_base/primesieve2/protocol"
	protocol_3 "NestedScribbleBenchmark/primesieve_base/primesieve3/protocol"
	callbacks_4 "NestedScribbleBenchmark/primesieve_base/primesieve4/callbacks"
	protocol_4 "NestedScribbleBenchmark/primesieve_base/primesieve4/protocol"
	primesieve_4 "NestedScribbleBenchmark/primesieve_base/primesieve4/results/primesieve"
	"fmt"
	"time"
)

type PrimeSieveEnv struct {
	N      int
	Primes []int
}

var primesieveParams = []int{
	100, 1100, 2100, 3100, 4100, 5100, 6100, 7100, 8100, 9100,
}

var primesieveBaselines = map[string]func(int) time.Duration{"primesieve-optimised": TimePrimeSieveBase,
	"primesieve-no-callbacks":    TimePrimeSieveBaseWithoutCallbacks,
	"primesieve-opt-invitations": TimePrimeSieveBaseOptimisedInvitations}

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

type PrimeSieveEnv2 struct {
	N      int
	Primes []int
}

func (p *PrimeSieveEnv2) New_Master_Env() callbacks2.PrimeSieve_Master_Env {
	return &callbacks2.PrimeSieveMasterState{
		N:      p.N,
		Primes: []int{2},
	}
}

func (p *PrimeSieveEnv2) New_Worker_Env() callbacks2.PrimeSieve_Worker_Env {
	return &callbacks2.PrimeSieveWorkerState{}
}

func (p *PrimeSieveEnv2) Master_Result(result []int) {
	p.Primes = result
}

func NewPrimeSieveEnv2(n int) *PrimeSieveEnv2 {
	return &PrimeSieveEnv2{
		N:      n,
		Primes: nil,
	}
}

type PrimeSieveEnv3 struct {
	N      int
	Primes []int
}

func (p *PrimeSieveEnv3) New_Master_Env() callbacks_4.PrimeSieve_Master_Env {
	return &callbacks_4.PrimeSieveMasterState{
		N:      p.N,
		Primes: []int{2},
	}
}

func (p *PrimeSieveEnv3) New_Worker_Env() callbacks_4.PrimeSieve_Worker_Env {
	return &callbacks_4.PrimeSieveWorkerState{}
}

func (p *PrimeSieveEnv3) Master_Result(result primesieve_4.Master_Result) {
	p.Primes = result.Primes
}

func (p *PrimeSieveEnv3) Worker_Result(result primesieve_4.Worker_Result) {
}

func NewPrimeSieveEnv4(n int) *PrimeSieveEnv3 {
	return &PrimeSieveEnv3{
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
	primesieve_base.PrimeSieve(n)
	elapsed := time.Since(start)
	return elapsed
}

func TimePrimeSieveBaseWithoutResultStructs(n int) time.Duration {
	// Remove result structs and cbs for empty results
	env := NewPrimeSieveEnv2(n)
	start := time.Now()
	protocol_2.PrimeSieve(env)
	elapsed := time.Since(start)
	// fmt.Println(env.Primes)
	return elapsed
}

func TimePrimeSieveBaseWithoutCallbacks(n int) time.Duration {
	// Remove result structs and cbs for empty results
	start := time.Now()
	protocol_3.PrimeSieve(n)
	elapsed := time.Since(start)
	// fmt.Println(res)
	return elapsed
}

func TimePrimeSieveBaseOptimisedInvitations(n int) time.Duration {
	// Remove result structs and cbs for empty results
	env := NewPrimeSieveEnv4(n)
	start := time.Now()
	protocol_4.PrimeSieve(env)
	elapsed := time.Since(start)
	// fmt.Println(env.Primes)
	return elapsed
}

func PrimeSieveBenchmark(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(primesieveParams, repetitions, TimePrimeSieve)
	base_results := benchmark.TimeImpl(primesieveParams, repetitions, TimePrimeSieveBase)
	return scribble_results, base_results
}

func PrimeSieveBenchmarkWithoutResultStructs(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(primesieveParams, repetitions, TimePrimeSieve)
	base_results := benchmark.TimeImpl(primesieveParams, repetitions, TimePrimeSieveBaseWithoutResultStructs)
	return scribble_results, base_results
}

func PrimeSieveBenchmarkWithoutCallbacks(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(primesieveParams, repetitions, TimePrimeSieve)
	base_results := benchmark.TimeImpl(primesieveParams, repetitions, TimePrimeSieveBaseWithoutCallbacks)
	return scribble_results, base_results
}

func PrimeSieveBenchmarkOptimisedInvitations(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(primesieveParams, repetitions, TimePrimeSieve)
	base_results := benchmark.TimeImpl(primesieveParams, repetitions, TimePrimeSieveBaseOptimisedInvitations)
	return scribble_results, base_results
}

func ComparePrimeSieveAgainstBaselines(repetitions int) (benchmark.BenchmarkTimes, map[string]benchmark.BenchmarkTimes) {
	fmt.Println("PrimeSieve")
	fmt.Println("Scribble")
	scribble_results := benchmark.TimeImpl(primesieveParams, repetitions, TimePrimeSieve)
	base_results := make(map[string]benchmark.BenchmarkTimes)
	for name, baseTimeFunc := range primesieveBaselines {
		fmt.Println("Baseline", name)
		base_results[name] = benchmark.TimeImpl(primesieveParams, repetitions, baseTimeFunc)
	}
	return scribble_results, base_results
}
