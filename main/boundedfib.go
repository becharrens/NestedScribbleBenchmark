package main

import (
	"NestedScribbleBenchmark/benchmark"
	"NestedScribbleBenchmark/bounded_fib_base"
	"NestedScribbleBenchmark/boundedfibonacci/results/boundedfibonacci"
	"NestedScribbleBenchmark/old_fibonacci/callbacks"
	"NestedScribbleBenchmark/old_fibonacci/protocol"
	"time"
)

var boundedfibParams = []int{
	5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90,
}

type BoundedFibonacciEnv struct {
	N      int
	Result int
}

func (f *BoundedFibonacciEnv) New_Start_Env() callbacks.Fibonacci_Start_Env {
	return &callbacks.FibonacciStartState{
		N:   f.N,
		Fib: 0,
	}
}

func (f *BoundedFibonacciEnv) New_F1_Env() callbacks.Fibonacci_F1_Env {
	return &callbacks.FibonnaciF1State{}
}

func (f *BoundedFibonacciEnv) New_F2_Env() callbacks.Fibonacci_F2_Env {
	return &callbacks.FibonacciF2State{}
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

func TimeBoundedFibonacci(n int) time.Duration {
	env := NewFibonacciEnv(n)
	start := time.Now()
	protocol.Fibonacci(env)
	elapsed := time.Since(start)
	return elapsed
}

func TimeBoundedFibonacciBase(n int) time.Duration {
	start := time.Now()
	_ = bounded_fib_base.Fibonacci(n)
	return time.Since(start)
}

func BoundedFibonacciBenchmark(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(boundedfibParams, repetitions, TimeBoundedFibonacci)
	base_results := benchmark.TimeImpl(boundedfibParams, repetitions, TimeBoundedFibonacciBase)
	return scribble_results, base_results
}
