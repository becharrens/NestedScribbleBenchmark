package main

import (
	"NestedScribbleBenchmark/benchmark"
	"NestedScribbleBenchmark/fibonacci/callbacks"
	"NestedScribbleBenchmark/fibonacci/protocol"
	"NestedScribbleBenchmark/fibonacci/results/fibonacci"
	"NestedScribbleBenchmark/fibonacci_base"
	"time"
)

var fibonacciParams = []int{
	5, 10, 15, 20, 25, 30, 35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90,
}

type FibonacciEnv struct {
	N      int
	Result int
}

func (f *FibonacciEnv) New_Start_Env() callbacks.Fibonacci_Start_Env {
	return &callbacks.FibonacciStartState{
		N:   f.N,
		Fib: 0,
	}
}

func (f *FibonacciEnv) New_F1_Env() callbacks.Fibonacci_F1_Env {
	return &callbacks.FibonnaciF1State{}
}

func (f *FibonacciEnv) New_F2_Env() callbacks.Fibonacci_F2_Env {
	return &callbacks.FibonacciF2State{}
}

func (f *FibonacciEnv) Start_Result(result fibonacci.Start_Result) {
	f.Result = result.Fib
}

func (f *FibonacciEnv) F1_Result(result fibonacci.F1_Result) {
}

func (f *FibonacciEnv) F2_Result(result fibonacci.F2_Result) {
}

func NewFibonacciEnv(n int) *FibonacciEnv {
	return &FibonacciEnv{
		N:      n,
		Result: 0,
	}
}

func TimeFibonacci(n int) time.Duration {
	env := NewFibonacciEnv(n)
	start := time.Now()
	protocol.Fibonacci(env)
	elapsed := time.Since(start)
	return elapsed
}

func TimeFibonacciBase(n int) time.Duration {
	start := time.Now()
	_ = fibonacci_base.Fibonacci(n)
	return time.Since(start)
}

func FibonacciBenchmark(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(fibonacciParams, repetitions, TimeFibonacci)
	base_results := benchmark.TimeImpl(fibonacciParams, repetitions, TimeFibonacciBase)
	return scribble_results, base_results
}
