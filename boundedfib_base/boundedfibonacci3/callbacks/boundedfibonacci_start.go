package callbacks

import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/results/boundedfibonacci"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci3/results/boundedfib"

type BoundedFibonacci_Start_Env interface {
	Done() boundedfibonacci.Start_Result
	ResultFrom_BoundedFib_Res(result boundedfib.Res_Result)
	To_BoundedFib_Res_Env() BoundedFib_Res_Env
	BoundedFib_Setup()
	StartFib2_To_F2() (int, int)
	StartFib1_To_F1() (int, int)
}

type BoundedFibonacciStartState struct {
	N   int
	Fib int
}

func (f *BoundedFibonacciStartState) Done() boundedfibonacci.Start_Result {
	return boundedfibonacci.Start_Result{Fib: f.Fib}
}

func (f *BoundedFibonacciStartState) ResultFrom_BoundedFib_Res(result boundedfib.Res_Result) {
	f.Fib = result.Fib
}

func (f *BoundedFibonacciStartState) To_BoundedFib_Res_Env() BoundedFib_Res_Env {
	return &BoundedFibResState{}
}

func (f *BoundedFibonacciStartState) BoundedFib_Setup() {
}

func (f *BoundedFibonacciStartState) StartFib2_To_F2() (int, int) {
	return f.N, 1
}

func (f *BoundedFibonacciStartState) StartFib1_To_F1() (int, int) {
	return f.N, 1
}
