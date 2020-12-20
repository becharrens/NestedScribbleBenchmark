package callbacks

import "NestedScribbleBenchmark/boundedfibonacci/results/boundedfib"

type BoundedFib_F2_Env interface {
	ResultFrom_BoundedFib_F1(result boundedfib.F1_Result)
	To_BoundedFib_F1_Env() BoundedFib_F1_Env
	Done() boundedfib.F2_Result
	End_From_F3()
	Fib2_To_F3() (int, int)
}

type BoundedFibF2State struct {
	Ubound int
	Idx    int
	Fib    int
}

func (f *BoundedFibF2State) ResultFrom_BoundedFib_F1(result boundedfib.F1_Result) {
}

func (f *BoundedFibF2State) To_BoundedFib_F1_Env() BoundedFib_F1_Env {
	return &BoundedFibF1State{
		Ubound: f.Ubound,
		Idx:    f.Idx,
		Fib:    f.Fib,
	}
}

func (f *BoundedFibF2State) Done() boundedfib.F2_Result {
	return boundedfib.F2_Result{}
}

func (f *BoundedFibF2State) End_From_F3() {
}

func (f *BoundedFibF2State) Fib2_To_F3() (int, int) {
	return f.Idx, f.Fib
}
