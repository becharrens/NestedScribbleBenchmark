package callbacks

import (
	"NestedScribbleBenchmark/boundedfibonacci/messages/boundedfibonacci"
)
import boundedfibonacci_2 "NestedScribbleBenchmark/boundedfibonacci/results/boundedfibonacci"
import "NestedScribbleBenchmark/boundedfibonacci/results/boundedfib"

type BoundedFibonacci_Start_Env interface {
	Done() boundedfibonacci_2.Start_Result
	ResultFrom_BoundedFib_Res(result boundedfib.Res_Result)
	To_BoundedFib_Res_Env() BoundedFib_Res_Env
	BoundedFib_Setup()
	StartFib2_To_F2() boundedfibonacci.StartFib2
	StartFib1_To_F1() boundedfibonacci.StartFib1
}

type FibonacciStartState struct {
	N   int
	Fib int
}

func (f *FibonacciStartState) Done() boundedfibonacci_2.Start_Result {
	return boundedfibonacci_2.Start_Result{Fib: f.Fib}
}

func (f *FibonacciStartState) ResultFrom_BoundedFib_Res(result boundedfib.Res_Result) {
	f.Fib = result.Fib
}

func (f *FibonacciStartState) To_BoundedFib_Res_Env() BoundedFib_Res_Env {
	return &BoundedFibResState{}
}

func (f *FibonacciStartState) BoundedFib_Setup() {
}

func (f *FibonacciStartState) StartFib2_To_F2() boundedfibonacci.StartFib2 {
	return boundedfibonacci.StartFib2{
		N:   f.N,
		Val: 1,
	}
}

func (f *FibonacciStartState) StartFib1_To_F1() boundedfibonacci.StartFib1 {
	return boundedfibonacci.StartFib1{
		N:   f.N,
		Val: 1,
	}
}
