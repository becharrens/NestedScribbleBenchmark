package callbacks

import "NestedScribbleBenchmark/boundedfibonacci/results/boundedfib"

type BoundedFib_F1_Env interface {
	Done() boundedfib.F1_Result
	Fib1_To_F3() (int, int, int)
}

type BoundedFibF1State struct {
	Ubound int
	Idx    int
	Fib    int
}

func (f *BoundedFibF1State) Done() boundedfib.F1_Result {
	return boundedfib.F1_Result{}
}

func (f *BoundedFibF1State) Fib1_To_F3() (int, int, int) {
	return f.Ubound, f.Idx, f.Fib
}
