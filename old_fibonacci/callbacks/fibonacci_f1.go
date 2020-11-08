package callbacks

import "NestedScribbleBenchmark/old_fibonacci/messages/fibonacci"
import fibonacci_2 "NestedScribbleBenchmark/old_fibonacci/results/fibonacci"
import "NestedScribbleBenchmark/old_fibonacci/results/fib"

type Fibonacci_F1_Env interface {
	Done() fibonacci_2.F1_Result
	ResultFrom_Fib_F1(result fib.F1_Result)
	To_Fib_F1_Env() Fib_F1_Env
	StartFib1_From_Start(startfib1 fibonacci.StartFib1)
}

type FibonnaciF1State struct {
	Ubound int
	Fib    int
	Idx    int
}

func (f *FibonnaciF1State) Done() fibonacci_2.F1_Result {
	return fibonacci_2.F1_Result{}
}

func (f *FibonnaciF1State) ResultFrom_Fib_F1(result fib.F1_Result) {
}

func (f *FibonnaciF1State) To_Fib_F1_Env() Fib_F1_Env {
	return &FibF1State{
		Ubound: f.Ubound,
		Idx:    f.Idx,
		Fib:    f.Fib,
	}
}

func (f *FibonnaciF1State) StartFib1_From_Start(startfib1 fibonacci.StartFib1) {
	f.Idx = 1
	f.Fib = startfib1.Val
	f.Ubound = startfib1.N
}
