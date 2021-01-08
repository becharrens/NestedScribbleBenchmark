package callbacks

import "NestedScribbleBenchmark/fibonacci/results/fib"

type Fib_Res_Env interface {
	Done() fib.Res_Result
	ResultFrom_Fib_Res(result fib.Res_Result) 
	To_Fib_Res_Env() Fib_Res_Env
	NextFib_From_F3(val int) 
}