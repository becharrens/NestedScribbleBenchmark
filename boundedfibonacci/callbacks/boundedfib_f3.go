package callbacks

import (
	"NestedScribbleBenchmark/boundedfibonacci/messages/boundedfib"
)
import boundedfib_2 "NestedScribbleBenchmark/boundedfibonacci/results/boundedfib"

type BoundedFib_F3_Choice int

const (
	BoundedFib_F3_BoundedFib BoundedFib_F3_Choice = iota
	BoundedFib_F3_Result
)

type BoundedFib_F3_Env interface {
	End_To_F2() boundedfib.End
	Result_To_Res() boundedfib.Result
	Done()
	ResultFrom_BoundedFib_F2(result boundedfib_2.F2_Result)
	To_BoundedFib_F2_Env() BoundedFib_F2_Env
	BoundedFib_Setup()
	F3_Choice() BoundedFib_F3_Choice
	Fib2_From_F2(fib2_msg boundedfib.Fib2)
	Fib1_From_F1(fib1_msg boundedfib.Fib1)
}

type BoundedFibF3State struct {
	Fib1   int
	Fib    int
	Ubound int
	Idx    int
}

func (f *BoundedFibF3State) End_To_F2() boundedfib.End {
	return boundedfib.End{}
}

func (f *BoundedFibF3State) Result_To_Res() boundedfib.Result {
	return boundedfib.Result{Fib: f.Fib}
}

func (f *BoundedFibF3State) Done() {
}

func (f *BoundedFibF3State) ResultFrom_BoundedFib_F2(result boundedfib_2.F2_Result) {
}

func (f *BoundedFibF3State) To_BoundedFib_F2_Env() BoundedFib_F2_Env {
	return &BoundedFibF2State{
		Ubound: f.Ubound,
		Idx:    f.Idx,
		Fib:    f.Fib,
	}
}

func (f *BoundedFibF3State) BoundedFib_Setup() {
}

func (f *BoundedFibF3State) F3_Choice() BoundedFib_F3_Choice {
	if f.Idx == f.Ubound {
		return BoundedFib_F3_Result
	}
	return BoundedFib_F3_BoundedFib
}

func (f *BoundedFibF3State) Fib2_From_F2(fib2 boundedfib.Fib2) {
	f.Fib = fib2.Val + f.Fib1
}

func (f *BoundedFibF3State) Fib1_From_F1(fib1 boundedfib.Fib1) {
	f.Fib1 = fib1.Val
	f.Ubound = fib1.Ubound
	f.Idx = fib1.Idx + 2
}

func New_BoundedFib_F3_State() BoundedFib_F3_Env {
	return &BoundedFibF3State{}
}
