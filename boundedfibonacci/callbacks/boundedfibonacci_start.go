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

type BoundedFibonacciStartState struct {
	N   int
	Fib int
}

func (f *BoundedFibonacciStartState) Done() boundedfibonacci_2.Start_Result {
	return boundedfibonacci_2.Start_Result{Fib: f.Fib}
}

func (f *BoundedFibonacciStartState) ResultFrom_BoundedFib_Res(result boundedfib.Res_Result) {
	f.Fib = result.Fib
}

func (f *BoundedFibonacciStartState) To_BoundedFib_Res_Env() BoundedFib_Res_Env {
	return &BoundedFibResState{}
}

func (f *BoundedFibonacciStartState) BoundedFib_Setup() {
}

func (f *BoundedFibonacciStartState) StartFib2_To_F2() boundedfibonacci.StartFib2 {
	return boundedfibonacci.StartFib2{
		N:   f.N,
		Val: 1,
	}
}

func (f *BoundedFibonacciStartState) StartFib1_To_F1() boundedfibonacci.StartFib1 {
	return boundedfibonacci.StartFib1{
		N:   f.N,
		Val: 1,
	}
}
