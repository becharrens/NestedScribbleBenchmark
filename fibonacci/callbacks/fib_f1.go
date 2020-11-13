package callbacks

import "NestedScribbleBenchmark/fibonacci/messages/fib"
import fib_2 "NestedScribbleBenchmark/fibonacci/results/fib"

type Fib_F1_Env interface {
	Done() fib_2.F1_Result
	Fib1_To_F3() fib.Fib1
}

type FibF1State struct {
	Fib int
}

func (f *FibF1State) Done() fib_2.F1_Result {
	return fib_2.F1_Result{}
}

func (f *FibF1State) Fib1_To_F3() fib.Fib1 {
	return fib.Fib1{
		Val: f.Fib,
	}
}