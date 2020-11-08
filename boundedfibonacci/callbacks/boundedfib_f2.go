package callbacks

import (
	"NestedScribbleBenchmark/boundedfibonacci/messages/boundedfib"
)
import boundedfib_2 "NestedScribbleBenchmark/boundedfibonacci/results/boundedfib"

type BoundedFib_F2_Env interface {
	ResultFrom_BoundedFib_F1(result boundedfib_2.F1_Result)
	To_BoundedFib_F1_Env() BoundedFib_F1_Env
	Done() boundedfib_2.F2_Result
	End_From_F3(end_msg boundedfib.End)
	Fib2_To_F3() boundedfib.Fib2
}

type BoundedFibF2State struct {
	Ubound int
	Idx    int
	Fib    int
}

func (f *BoundedFibF2State) ResultFrom_BoundedFib_F1(result boundedfib_2.F1_Result) {
}

func (f *BoundedFibF2State) To_BoundedFib_F1_Env() BoundedFib_F1_Env {
	return &BoundedFibF1State{
		Ubound: f.Ubound,
		Idx:    f.Idx,
		Fib:    f.Fib,
	}
}

func (f *BoundedFibF2State) Done() boundedfib_2.F2_Result {
	return boundedfib_2.F2_Result{}
}

func (f *BoundedFibF2State) End_From_F3(end boundedfib.End) {
}

func (f *BoundedFibF2State) Fib2_To_F3() boundedfib.Fib2 {
	return boundedfib.Fib2{
		Idx: f.Idx,
		Val: f.Fib,
	}
}
