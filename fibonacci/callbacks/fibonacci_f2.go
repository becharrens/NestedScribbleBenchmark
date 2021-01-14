package callbacks

import (
	"NestedScribbleBenchmark/fibonacci/results/fibonacci"
)
import "NestedScribbleBenchmark/fibonacci/results/fib"

type Fibonacci_F2_Env interface {
	Done() fibonacci.F2_Result
	ResultFrom_Fib_F2(result fib.F2_Result)
	To_Fib_F2_Env() Fib_F2_Env
	StartFib2_From_Start(val int)
}

type FibonacciF2State struct {
	Fib int
}

func (f *FibonacciF2State) Done() fibonacci.F2_Result {
	return fibonacci.F2_Result{}
}

func (f *FibonacciF2State) ResultFrom_Fib_F2(result fib.F2_Result) {
}

func (f *FibonacciF2State) To_Fib_F2_Env() Fib_F2_Env {
	return &FibF2State{
		Fib: f.Fib,
	}
}

func (f *FibonacciF2State) StartFib2_From_Start(val int) {
	f.Fib = val
}
