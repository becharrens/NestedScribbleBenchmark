package callbacks

import "NestedScribbleBenchmark/fibonacci/results/fib"

type Fib_F2_Env interface {
	Done() fib.F2_Result
	ResultFrom_Fib_F1(result fib.F1_Result) 
	To_Fib_F1_Env() Fib_F1_Env
	Fib2_To_F3() int
}