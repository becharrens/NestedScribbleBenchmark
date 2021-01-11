package callbacks

import "NestedScribbleBenchmark/fibonacci/messages/fibonacci"
import fibonacci_2 "NestedScribbleBenchmark/fibonacci/results/fibonacci"
import "NestedScribbleBenchmark/fibonacci/results/fib"

type Fibonacci_Start_Env interface {
	Done() fibonacci_2.Start_Result
	ResultFrom_Fib_Res(result fib.Res_Result)
	To_Fib_Res_Env() Fib_Res_Env
	Fib_Setup()
	StartFib2_To_F2() fibonacci.StartFib2
	StartFib1_To_F1() fibonacci.StartFib1
}

type FibonacciStartState struct {
}

func (f *FibonacciStartState) Done() fibonacci_2.Start_Result {
	return fibonacci_2.Start_Result{}
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

func (f *FibonacciStartState) StartFib2_To_F2() fibonacci.StartFib2 {
	return fibonacci.StartFib2{
		Val: 1,
	}
}

func (f *FibonacciStartState) StartFib1_To_F1() fibonacci.StartFib1 {
	return fibonacci.StartFib1{
		Val: 1,
	}
}
