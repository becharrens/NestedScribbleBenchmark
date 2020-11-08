package callbacks

import (
	"NestedScribbleBenchmark/boundedfibonacci/messages/boundedfibonacci"
)
import boundedfibonacci_2 "NestedScribbleBenchmark/boundedfibonacci/results/boundedfibonacci"
import "NestedScribbleBenchmark/boundedfibonacci/results/boundedfib"

type BoundedFibonacci_F1_Env interface {
	Done() boundedfibonacci_2.F1_Result
	ResultFrom_BoundedFib_F1(result boundedfib.F1_Result)
	To_BoundedFib_F1_Env() BoundedFib_F1_Env
	StartFib1_From_Start(startfib1_msg boundedfibonacci.StartFib1)
}

type BoundedFibonnaciF1State struct {
	Ubound int
	Fib    int
	Idx    int
}

func (f *BoundedFibonnaciF1State) Done() boundedfibonacci_2.F1_Result {
	return boundedfibonacci_2.F1_Result{}
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

func (f *BoundedFibonnaciF1State) StartFib1_From_Start(startfib1 boundedfibonacci.StartFib1) {
	f.Idx = 1
	f.Fib = startfib1.Val
	f.Ubound = startfib1.N
}
