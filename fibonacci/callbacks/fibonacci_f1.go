package callbacks

import (
	"NestedScribbleBenchmark/fibonacci/results/fibonacci"
)
import "NestedScribbleBenchmark/fibonacci/results/fib"

type Fibonacci_F1_Env interface {
	Done() fibonacci.F1_Result
	ResultFrom_Fib_F1(result fib.F1_Result)
	To_Fib_F1_Env() Fib_F1_Env
	StartFib1_From_Start(val int)
}

type FibonacciF1State struct {
	Fib int
}

func (f *FibonacciF1State) Done() fibonacci.F1_Result {
	return fibonacci.F1_Result{}
}

func (f *FibonacciF1State) ResultFrom_Fib_F1(result fib.F1_Result) {
}

func (f *FibonacciF1State) To_Fib_F1_Env() Fib_F1_Env {
	return &FibF1State{
		Fib: f.Fib,
	}
}

func (f *FibonacciF1State) StartFib1_From_Start(val int) {
	f.Fib = val
}
