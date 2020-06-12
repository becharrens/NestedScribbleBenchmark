package callbacks

import "ScribbleBenchmark/fibonacci/messages/fibonacci"
import fibonacci_2 "ScribbleBenchmark/fibonacci/results/fibonacci"
import "ScribbleBenchmark/fibonacci/results/fib"

type Fibonacci_F2_Env interface {
	Done() fibonacci_2.F2_Result
	ResultFrom_Fib_F2(result fib.F2_Result)
	To_Fib_F2_Env() Fib_F2_Env
	StartFib2_From_Start(startfib2 fibonacci.StartFib2)
}

type FibonacciF2State struct {
	Ubound int
	Idx    int
	Fib    int
}

func (f *FibonacciF2State) Done() fibonacci_2.F2_Result {
	return fibonacci_2.F2_Result{}
}

func (f *FibonacciF2State) ResultFrom_Fib_F2(result fib.F2_Result) {
}

func (f *FibonacciF2State) To_Fib_F2_Env() Fib_F2_Env {
	return &FibF2State{
		Ubound: f.Ubound,
		Idx:    f.Idx,
		Fib:    f.Fib,
	}
}

func (f *FibonacciF2State) StartFib2_From_Start(startfib2 fibonacci.StartFib2) {
	f.Fib = startfib2.Val
	f.Ubound = startfib2.N
	f.Idx = 2
}
