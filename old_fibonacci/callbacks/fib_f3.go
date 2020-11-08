package callbacks

import "NestedScribbleBenchmark/old_fibonacci/messages/fib"
import fib_2 "NestedScribbleBenchmark/old_fibonacci/results/fib"

type Fib_F3_Choice int

const (
	Fib_F3_Fib Fib_F3_Choice = iota
	Fib_F3_Result
)

type Fib_F3_Env interface {
	End_To_F2() fib.End
	Result_To_Res() fib.Result
	Done()
	ResultFrom_Fib_F2(result fib_2.F2_Result)
	To_Fib_F2_Env() Fib_F2_Env
	Fib_Setup()
	F3_Choice() Fib_F3_Choice
	Fib2_From_F2(fib2 fib.Fib2)
	Fib1_From_F1(fib1 fib.Fib1)
}

type FibF3State struct {
	Fib1   int
	Fib    int
	Ubound int
	Idx    int
}

func (f *FibF3State) End_To_F2() fib.End {
	return fib.End{}
}

func (f *FibF3State) Result_To_Res() fib.Result {
	return fib.Result{Fib: f.Fib}
}

func (f *FibF3State) Done() {
}

func (f *FibF3State) ResultFrom_Fib_F2(result fib_2.F2_Result) {
}

func (f *FibF3State) To_Fib_F2_Env() Fib_F2_Env {
	return &FibF2State{
		Ubound: f.Ubound,
		Idx:    f.Idx,
		Fib:    f.Fib,
	}
}

func (f *FibF3State) Fib_Setup() {
}

func (f *FibF3State) F3_Choice() Fib_F3_Choice {
	if f.Idx == f.Ubound {
		return Fib_F3_Result
	}
	return Fib_F3_Fib
}

func (f *FibF3State) Fib2_From_F2(fib2 fib.Fib2) {
	f.Fib = fib2.Val + f.Fib1
}

func (f *FibF3State) Fib1_From_F1(fib1 fib.Fib1) {
	f.Fib1 = fib1.Val
	f.Ubound = fib1.Ubound
	f.Idx = fib1.Idx + 2
}

func New_Fib_F3_State() Fib_F3_Env {
	return &FibF3State{}
}
