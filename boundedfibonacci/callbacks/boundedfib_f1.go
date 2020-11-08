package callbacks

import (
	"NestedScribbleBenchmark/boundedfibonacci/messages/boundedfib"
)
import boundedfib_2 "NestedScribbleBenchmark/boundedfibonacci/results/boundedfib"

type BoundedFib_F1_Env interface {
	Done() boundedfib_2.F1_Result
	Fib1_To_F3() boundedfib.Fib1
}

type BoundedFibF1State struct {
	Ubound int
	Idx    int
	Fib    int
}

func (f *BoundedFibF1State) Done() boundedfib_2.F1_Result {
	return boundedfib_2.F1_Result{}
}

func (f *BoundedFibF1State) Fib1_To_F3() boundedfib.Fib1 {
	return boundedfib.Fib1{
		Idx:    f.Idx,
		Ubound: f.Ubound,
		Val:    f.Fib,
	}
}
