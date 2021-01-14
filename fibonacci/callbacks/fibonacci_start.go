package callbacks

import (
	"NestedScribbleBenchmark/fibonacci/results/fibonacci"
)
import "NestedScribbleBenchmark/fibonacci/results/fib"

type Fibonacci_Start_Env interface {
	Done() fibonacci.Start_Result
	ResultFrom_Fib_Res(result fib.Res_Result)
	To_Fib_Res_Env() Fib_Res_Env
	Fib_Setup()
	StartFib2_To_F2() int
	StartFib1_To_F1() int
}

type FibonacciStartState struct {
}

func (f *FibonacciStartState) Done() fibonacci.Start_Result {
	return fibonacci.Start_Result{}
}

func (f *FibonacciStartState) ResultFrom_Fib_Res(result fib.Res_Result) {
}

func (f *FibonacciStartState) To_Fib_Res_Env() Fib_Res_Env {
	return &FibResState{
		Idx: 3,
	}
}

func (f *FibonacciStartState) Fib_Setup() {
}

func (f *FibonacciStartState) StartFib2_To_F2() int {
	return 1
}

func (f *FibonacciStartState) StartFib1_To_F1() int {
	return 1
}
