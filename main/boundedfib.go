package main

import (
	"NestedScribbleBenchmark/benchmark"
	"NestedScribbleBenchmark/boundedfib_base"
	protocol_2 "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/protocol"
	callbacks_3 "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/callbacks"
	protocol_3 "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/protocol"
	boundedfibonacci_2 "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/results/boundedfibonacci"
	callbacks_4 "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci4/callbacks"
	protocol_4 "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci4/protocol"
	boundedfibonacci_4 "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci4/results/boundedfibonacci"
	"NestedScribbleBenchmark/boundedfibonacci/callbacks"
	"NestedScribbleBenchmark/boundedfibonacci/protocol"
	"NestedScribbleBenchmark/boundedfibonacci/results/boundedfibonacci"
	"fmt"
	"time"
)

var boundedfibParams = []int{
	5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90,
}

var fibBaselines = map[string]func(int) time.Duration{"fib-optimised": TimeBoundedFibonacciBase,
	"fib-no-callbacks":      TimeBoundedFibonacciWithoutCallbacks,
	"fib-opt-invitations":   TimeBoundedFibonacciOptimisedInvitations,
	"fib-opt-labelled-msgs": TimeBoundedFibonacciOptimisedLabelledMsgExchanges}

// var boundedfibParams = []int{
// 	10, 25, 40, 55, 70, 80, 90,
// }

type BoundedFibonacciEnv struct {
	N      int
	Result int
}

func (f *BoundedFibonacciEnv) New_Start_Env() callbacks.BoundedFibonacci_Start_Env {
	return &callbacks.BoundedFibonacciStartState{
		N:   f.N,
		Fib: 0,
	}
}

func (f *BoundedFibonacciEnv) New_F1_Env() callbacks.BoundedFibonacci_F1_Env {
	return &callbacks.BoundedFibonnaciF1State{}
}

func (f *BoundedFibonacciEnv) New_F2_Env() callbacks.BoundedFibonacci_F2_Env {
	return &callbacks.BoundedFibonacciF2State{}
}

func (f *BoundedFibonacciEnv) Start_Result(result boundedfibonacci.Start_Result) {
	f.Result = result.Fib
}

func (f *BoundedFibonacciEnv) F1_Result(result boundedfibonacci.F1_Result) {
}

func (f *BoundedFibonacciEnv) F2_Result(result boundedfibonacci.F2_Result) {
}

func NewBoundedFibonacciEnv(n int) *BoundedFibonacciEnv {
	return &BoundedFibonacciEnv{
		N:      n,
		Result: 0,
	}
}

type BoundedFibonacciEnv3 struct {
	N      int
	Result int
}

func (f *BoundedFibonacciEnv3) New_Start_Env() callbacks_3.BoundedFibonacci_Start_Env {
	return &callbacks_3.BoundedFibonacciStartState{
		N:   f.N,
		Fib: 0,
	}
}

func (f *BoundedFibonacciEnv3) New_F1_Env() callbacks_3.BoundedFibonacci_F1_Env {
	return &callbacks_3.BoundedFibonnaciF1State{}
}

func (f *BoundedFibonacciEnv3) New_F2_Env() callbacks_3.BoundedFibonacci_F2_Env {
	return &callbacks_3.BoundedFibonacciF2State{}
}

func (f *BoundedFibonacciEnv3) Start_Result(result boundedfibonacci_2.Start_Result) {
	f.Result = result.Fib
}

func (f *BoundedFibonacciEnv3) F1_Result(result boundedfibonacci_2.F1_Result) {
}

func (f *BoundedFibonacciEnv3) F2_Result(result boundedfibonacci_2.F2_Result) {
}

func NewBoundedFibonacciEnv3(n int) *BoundedFibonacciEnv3 {
	return &BoundedFibonacciEnv3{
		N:      n,
		Result: 0,
	}
}

type BoundedFibonacciEnv4 struct {
	N      int
	Result int
}

func (f *BoundedFibonacciEnv4) New_Start_Env() callbacks_4.BoundedFibonacci_Start_Env {
	return &callbacks_4.BoundedFibonacciStartState{
		N:   f.N,
		Fib: 0,
	}
}

func (f *BoundedFibonacciEnv4) New_F1_Env() callbacks_4.BoundedFibonacci_F1_Env {
	return &callbacks_4.BoundedFibonnaciF1State{}
}

func (f *BoundedFibonacciEnv4) New_F2_Env() callbacks_4.BoundedFibonacci_F2_Env {
	return &callbacks_4.BoundedFibonacciF2State{}
}

func (f *BoundedFibonacciEnv4) Start_Result(result boundedfibonacci_4.Start_Result) {
	f.Result = result.Fib
}

func (f *BoundedFibonacciEnv4) F1_Result(result boundedfibonacci_4.F1_Result) {
}

func (f *BoundedFibonacciEnv4) F2_Result(result boundedfibonacci_4.F2_Result) {
}

func NewBoundedFibonacciEnv4(n int) *BoundedFibonacciEnv4 {
	return &BoundedFibonacciEnv4{
		N:      n,
		Result: 0,
	}
}

func TimeBoundedFibonacci(n int) time.Duration {
	env := NewBoundedFibonacciEnv(n)
	start := time.Now()
	protocol.BoundedFibonacci(env)
	elapsed := time.Since(start)
	return elapsed
}

func TimeBoundedFibonacciBase(n int) time.Duration {
	start := time.Now()
	_ = boundedfib_base.Fibonacci(n)
	return time.Since(start)
}

func TimeBoundedFibonacciWithoutCallbacks(n int) time.Duration {
	// Without callbacks/envs, results, optimised empty invite struct, removed msg, no wg
	start := time.Now()
	protocol_2.BoundedFibonacci(n)
	elapsed := time.Since(start)
	return elapsed
}

func TimeBoundedFibonacciWithoutEmptyInviteStruct(n int) time.Duration {
	env := NewBoundedFibonacciEnv3(n)
	start := time.Now()
	protocol_3.BoundedFibonacci(env)
	elapsed := time.Since(start)
	return elapsed
}

func TimeBoundedFibonacciOptimisedLabelledMsgExchanges(n int) time.Duration {
	start := time.Now()
	boundedfib_base.Fib(n)
	elapsed := time.Since(start)
	// fmt.Println(res)
	return elapsed
}

func TimeBoundedFibonacciOptimisedInvitations(n int) time.Duration {
	// Removed extra message
	env := NewBoundedFibonacciEnv4(n)
	start := time.Now()
	protocol_4.BoundedFibonacci(env)
	elapsed := time.Since(start)
	// fmt.Println(env.Result)
	return elapsed
}

func BoundedFibonacciBenchmark(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(boundedfibParams, repetitions, TimeBoundedFibonacci)
	base_results := benchmark.TimeImpl(boundedfibParams, repetitions, TimeBoundedFibonacciBase)
	return scribble_results, base_results
}

func BoundedFibonacciBenchmarkWithoutCallbacks(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(boundedfibParams, repetitions, TimeBoundedFibonacci)
	base_results := benchmark.TimeImpl(boundedfibParams, repetitions, TimeBoundedFibonacciWithoutCallbacks)
	return scribble_results, base_results
}

func BoundedFibonacciBenchmarkTestWithoutEmptyInviteStruct(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(boundedfibParams, repetitions, TimeBoundedFibonacci)
	base_results := benchmark.TimeImpl(boundedfibParams, repetitions, TimeBoundedFibonacciWithoutEmptyInviteStruct)
	return scribble_results, base_results
}

func BoundedFibonacciBenchmarkOptimisedLabelledMsgExchanges(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(boundedfibParams, repetitions, TimeBoundedFibonacci)
	base_results := benchmark.TimeImpl(boundedfibParams, repetitions, TimeBoundedFibonacciOptimisedLabelledMsgExchanges)
	return scribble_results, base_results
}

func BoundedFibonacciBenchmarkOptimisedInvitations(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(boundedfibParams, repetitions, TimeBoundedFibonacci)
	base_results := benchmark.TimeImpl(boundedfibParams, repetitions, TimeBoundedFibonacciOptimisedInvitations)
	return scribble_results, base_results
}

func CompareFibonacciAgainstBaselines(repetitions int) (benchmark.BenchmarkTimes, map[string]benchmark.BenchmarkTimes) {
	fmt.Println("Bounded Fibonacci")
	fmt.Println("Scribble")
	scribble_results := benchmark.TimeImpl(boundedfibParams, repetitions, TimeBoundedFibonacci)
	base_results := make(map[string]benchmark.BenchmarkTimes)
	for name, baseTimeFunc := range fibBaselines {
		fmt.Println("Baseline", name)
		base_results[name] = benchmark.TimeImpl(boundedfibParams, repetitions, baseTimeFunc)
	}
	return scribble_results, base_results
}
