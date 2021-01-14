package callbacks

import (
	"NestedScribbleBenchmark/fibonacci/results/fib"
)

type Fib_F1_Env interface {
	Done() fib.F1_Result
	Fib1_To_F3() int
}

type FibF1State struct {
	Fib int
}

func (f *FibF1State) Done() fib.F1_Result {
	return fib.F1_Result{}
}

func (f *FibF1State) Fib1_To_F3() int {
	return f.Fib
}
