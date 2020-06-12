package callbacks

import "ScribbleBenchmark/fibonacci/messages/fib"
import fib_2 "ScribbleBenchmark/fibonacci/results/fib"

type Fib_F2_Env interface {
	ResultFrom_Fib_F1(result fib_2.F1_Result) 
	To_Fib_F1_Env() Fib_F1_Env
	Done() fib_2.F2_Result
	End_From_F3(end fib.End) 
	Fib2_To_F3() fib.Fib2
}

type FibF2State struct {
	Ubound int
	Idx int
	Fib int
}

func (f *FibF2State) ResultFrom_Fib_F1(result fib_2.F1_Result) {
}

func (f *FibF2State) To_Fib_F1_Env() Fib_F1_Env {
	return &FibF1State{
		Ubound: f.Ubound,
		Idx:    f.Idx,
		Fib:    f.Fib,
	}
}

func (f *FibF2State) Done() fib_2.F2_Result {
	return fib_2.F2_Result{}
}

func (f *FibF2State) End_From_F3(end fib.End) {
}

func (f *FibF2State) Fib2_To_F3() fib.Fib2 {
	return fib.Fib2{
		Idx: f.Idx,
		Val: f.Fib,
	}
}


