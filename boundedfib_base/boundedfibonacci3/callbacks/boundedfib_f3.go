package callbacks

import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/results/boundedfib"

type BoundedFib_F3_Choice int

const (
	BoundedFib_F3_BoundedFib BoundedFib_F3_Choice = iota
	BoundedFib_F3_Result
)

type BoundedFib_F3_Env interface {
	End_To_F2()
	Result_To_Res() int
	Done()
	ResultFrom_BoundedFib_F2(result boundedfib.F2_Result)
	To_BoundedFib_F2_Env() BoundedFib_F2_Env
	BoundedFib_Setup()
	F3_Choice() BoundedFib_F3_Choice
	Fib2_From_F2(val int)
	// Fib2_From_F2(idx int, val int)
	Fib1_From_F1(ubound int, idx int, val int)
}

type BoundedFibF3State struct {
	Fib1   int
	Fib    int
	Ubound int
	Idx    int
}

func (f *BoundedFibF3State) End_To_F2() {
}

func (f *BoundedFibF3State) Result_To_Res() int {
	return f.Fib
}

func (f *BoundedFibF3State) Done() {
}

func (f *BoundedFibF3State) ResultFrom_BoundedFib_F2(result boundedfib.F2_Result) {
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

// func (f *BoundedFibF3State) Fib2_From_F2(idx int, val int) {
func (f *BoundedFibF3State) Fib2_From_F2(val int) {
	f.Fib = val + f.Fib1
}

func (f *BoundedFibF3State) Fib1_From_F1(ubound int, idx int, val int) {
	f.Fib1 = val
	f.Ubound = ubound
	f.Idx = idx + 2
}

func New_BoundedFib_F3_State() BoundedFib_F3_Env {
	return &BoundedFibF3State{}
}
