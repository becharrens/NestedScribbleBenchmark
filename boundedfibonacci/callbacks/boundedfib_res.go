package callbacks

import (
	boundedfib_2 "NestedScribbleBenchmark/boundedfibonacci/messages/boundedfib"
)
import "NestedScribbleBenchmark/boundedfibonacci/results/boundedfib"

type BoundedFib_Res_Env interface {
	Result_From_F3(result_msg boundedfib_2.Result)
	Done() boundedfib.Res_Result
	ResultFrom_BoundedFib_Res(result boundedfib.Res_Result)
	To_BoundedFib_Res_Env() BoundedFib_Res_Env
}

type BoundedFibResState struct {
	Fib int
}

func (f *BoundedFibResState) Result_From_F3(result boundedfib_2.Result) {
	f.Fib = result.Fib
}

func (f *BoundedFibResState) Done() boundedfib.Res_Result {
	return boundedfib.Res_Result{Fib: f.Fib}
}

func (f *BoundedFibResState) ResultFrom_BoundedFib_Res(result boundedfib.Res_Result) {
	f.Fib = result.Fib
}

func (f *BoundedFibResState) To_BoundedFib_Res_Env() BoundedFib_Res_Env {
	return f
}
