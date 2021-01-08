package callbacks

import "NestedScribbleBenchmark/fibonacci_prev/messages/fib"
import fib_2 "NestedScribbleBenchmark/fibonacci_prev/results/fib"

type Fib_F3_Env interface {
	Done()
	ResultFrom_Fib_F2(result fib_2.F2_Result)
	To_Fib_F2_Env() Fib_F2_Env
	Fib_Setup()
	NextFib_To_Res() fib.NextFib
	Fib2_From_F2(fib2_msg fib.Fib2)
	Fib1_From_F1(fib1_msg fib.Fib1)
}

type FibF3State struct {
	Fib  int
	Fib1 int
}

func (f *FibF3State) Done() {
}

func (f *FibF3State) ResultFrom_Fib_F2(result fib_2.F2_Result) {
}

func (f *FibF3State) To_Fib_F2_Env() Fib_F2_Env {
	return &FibF2State{
		Fib: f.Fib,
	}
}

func (f *FibF3State) Fib_Setup() {
}

func (f *FibF3State) NextFib_To_Res() fib.NextFib {
	return fib.NextFib{
		Val: f.Fib,
	}
}

func (f *FibF3State) Fib2_From_F2(fib2_msg fib.Fib2) {
	f.Fib = fib2_msg.Val + f.Fib1
}

func (f *FibF3State) Fib1_From_F1(fib1_msg fib.Fib1) {
	f.Fib1 = fib1_msg.Val
}

func New_Fib_F3_State() Fib_F3_Env {
	return &FibF3State{}
}
