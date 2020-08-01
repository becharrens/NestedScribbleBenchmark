package fibonacci

import (
	"NestedScribbleBenchmark/fibonacci/callbacks"
	"NestedScribbleBenchmark/fibonacci/results/fibonacci"
)

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
