package callbacks

import (
	"NestedScribbleBenchmark/fibonacci/results/fib"
)

type Fib_F2_Env interface {
	Done() fib.F2_Result
	ResultFrom_Fib_F1(result fib.F1_Result)
	To_Fib_F1_Env() Fib_F1_Env
	Fib2_To_F3() int
}

type FibF2State struct {
	Fib int
}

func (f *FibF2State) Done() fib.F2_Result {
	return fib.F2_Result{}
}

func (f *FibF2State) ResultFrom_Fib_F1(result fib.F1_Result) {
}

func (f *FibF2State) To_Fib_F1_Env() Fib_F1_Env {
	return &FibF1State{
		Fib: f.Fib,
	}
}

func (f *FibF2State) Fib2_To_F3() int {
	return f.Fib
}
