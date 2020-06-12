package callbacks

import "ScribbleBenchmark/fibonacci/messages/fibonacci"
import fibonacci_2 "ScribbleBenchmark/fibonacci/results/fibonacci"
import "ScribbleBenchmark/fibonacci/results/fib"

type Fibonacci_Start_Env interface {
	Done() fibonacci_2.Start_Result
	ResultFrom_Fib_Res(result fib.Res_Result)
	To_Fib_Res_Env() Fib_Res_Env
	Fib_Setup()
	StartFib2_To_F2() fibonacci.StartFib2
	StartFib1_To_F1() fibonacci.StartFib1
}

type FibonacciStartState struct {
	N   int
	Fib int
}

func (f *FibonacciStartState) Done() fibonacci_2.Start_Result {
	return fibonacci_2.Start_Result{Fib: f.Fib}
}

func (f *FibonacciStartState) ResultFrom_Fib_Res(result fib.Res_Result) {
	f.Fib = result.Fib
}

func (f *FibonacciStartState) To_Fib_Res_Env() Fib_Res_Env {
	return &FibResState{}
}

func (f *FibonacciStartState) Fib_Setup() {
}

func (f *FibonacciStartState) StartFib2_To_F2() fibonacci.StartFib2 {
	return fibonacci.StartFib2{
		N:   f.N,
		Val: 1,
	}
}

func (f *FibonacciStartState) StartFib1_To_F1() fibonacci.StartFib1 {
	return fibonacci.StartFib1{
		N:   f.N,
		Val: 1,
	}
}
