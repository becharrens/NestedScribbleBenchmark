package callbacks

import (
	"NestedScribbleBenchmark/fibonacci/results/fib"
)

type Fib_F3_Env interface {
	Done()
	ResultFrom_Fib_F2(result fib.F2_Result)
	To_Fib_F2_Env() Fib_F2_Env
	Fib_Setup()
	NextFib_To_Res() int
	Fib2_From_F2(val int)
	Fib1_From_F1(val int)
}

type FibF3State struct {
	Fib  int
	Fib1 int
}

func (f *FibF3State) Done() {
}

func (f *FibF3State) ResultFrom_Fib_F2(result fib.F2_Result) {
}

func (f *FibF3State) To_Fib_F2_Env() Fib_F2_Env {
	return &FibF2State{
		Fib: f.Fib,
	}
}

func (f *FibF3State) Fib_Setup() {
}

func (f *FibF3State) NextFib_To_Res() int {
	return f.Fib
}

func (f *FibF3State) Fib2_From_F2(val int) {
	f.Fib = val + f.Fib1
}

func (f *FibF3State) Fib1_From_F1(val int) {
	f.Fib1 = val
}

func New_Fib_F3_State() Fib_F3_Env {
	return &FibF3State{}
}
