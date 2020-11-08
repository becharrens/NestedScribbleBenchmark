package callbacks

import (
	"NestedScribbleBenchmark/boundedfibonacci/messages/boundedfibonacci"
)
import boundedfibonacci_2 "NestedScribbleBenchmark/boundedfibonacci/results/boundedfibonacci"
import "NestedScribbleBenchmark/boundedfibonacci/results/boundedfib"

type BoundedFibonacci_F2_Env interface {
	Done() boundedfibonacci_2.F2_Result
	ResultFrom_BoundedFib_F2(result boundedfib.F2_Result)
	To_BoundedFib_F2_Env() BoundedFib_F2_Env
	StartFib2_From_Start(startfib2_msg boundedfibonacci.StartFib2)
}

type BoundedFibonacciF2State struct {
	Ubound int
	Idx    int
	Fib    int
}

func (f *BoundedFibonacciF2State) Done() boundedfibonacci_2.F2_Result {
	return boundedfibonacci_2.F2_Result{}
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

func (f *BoundedFibonacciF2State) StartFib2_From_Start(startfib2 boundedfibonacci.StartFib2) {
	f.Fib = startfib2.Val
	f.Ubound = startfib2.N
	f.Idx = 2
}
