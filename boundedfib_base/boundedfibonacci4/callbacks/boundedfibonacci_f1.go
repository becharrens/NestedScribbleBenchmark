package callbacks

import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci4/results/boundedfibonacci"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci4/results/boundedfib"

type BoundedFibonacci_F1_Env interface {
	Done() boundedfibonacci.F1_Result
	ResultFrom_BoundedFib_F1(result boundedfib.F1_Result)
	To_BoundedFib_F1_Env() BoundedFib_F1_Env
	StartFib1_From_Start(n int, val int)
}

type BoundedFibonnaciF1State struct {
	Ubound int
	Fib    int
	Idx    int
}

func (f *BoundedFibonnaciF1State) Done() boundedfibonacci.F1_Result {
	return boundedfibonacci.F1_Result{}
}

func (f *BoundedFibonnaciF1State) ResultFrom_BoundedFib_F1(result boundedfib.F1_Result) {
}

func (f *BoundedFibonnaciF1State) To_BoundedFib_F1_Env() BoundedFib_F1_Env {
	return &BoundedFibF1State{
		Ubound: f.Ubound,
		Idx:    f.Idx,
		Fib:    f.Fib,
	}
}

func (f *BoundedFibonnaciF1State) StartFib1_From_Start(n int, val int) {
	f.Idx = 1
	f.Fib = val
	f.Ubound = n
}
