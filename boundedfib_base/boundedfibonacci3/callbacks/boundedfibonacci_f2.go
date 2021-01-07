package callbacks

import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/results/boundedfibonacci"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/results/boundedfib"

type BoundedFibonacci_F2_Env interface {
	Done() boundedfibonacci.F2_Result
	ResultFrom_BoundedFib_F2(result boundedfib.F2_Result)
	To_BoundedFib_F2_Env() BoundedFib_F2_Env
	StartFib2_From_Start(n int, val int)
}

type BoundedFibonacciF2State struct {
	Ubound int
	Idx    int
	Fib    int
}

func (f *BoundedFibonacciF2State) Done() boundedfibonacci.F2_Result {
	return boundedfibonacci.F2_Result{}
}

func (f *BoundedFibonacciF2State) ResultFrom_BoundedFib_F2(result boundedfib.F2_Result) {
}

func (f *BoundedFibonacciF2State) To_BoundedFib_F2_Env() BoundedFib_F2_Env {
	return &BoundedFibF2State{
		Ubound: f.Ubound,
		Idx:    f.Idx,
		Fib:    f.Fib,
	}
}

func (f *BoundedFibonacciF2State) StartFib2_From_Start(n int, val int) {
	f.Fib = val
	f.Ubound = n
	f.Idx = 2
}
