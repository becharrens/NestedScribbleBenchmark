package callbacks

import fib_2 "NestedScribbleBenchmark/old_fibonacci/messages/fib"
import "NestedScribbleBenchmark/old_fibonacci/results/fib"

type Fib_Res_Env interface {
	Result_From_F3(result fib_2.Result)
	Done() fib.Res_Result
	ResultFrom_Fib_Res(result fib.Res_Result)
	To_Fib_Res_Env() Fib_Res_Env
}

type FibResState struct {
	Fib int
}

func (f *FibResState) Result_From_F3(result fib_2.Result) {
	f.Fib = result.Fib
}

func (f *FibResState) Done() fib.Res_Result {
	return fib.Res_Result{Fib: f.Fib}
}

func (f *FibResState) ResultFrom_Fib_Res(result fib.Res_Result) {
	f.Fib = result.Fib
}

func (f *FibResState) To_Fib_Res_Env() Fib_Res_Env {
	return f
}
