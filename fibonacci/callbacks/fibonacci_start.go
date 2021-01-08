package callbacks

import "NestedScribbleBenchmark/fibonacci/results/fibonacci"
import "NestedScribbleBenchmark/fibonacci/results/fib"

type Fibonacci_Start_Env interface {
	Done() fibonacci.Start_Result
	ResultFrom_Fib_Res(result fib.Res_Result) 
	To_Fib_Res_Env() Fib_Res_Env
	Fib_Setup() 
	StartFib2_To_F2() int
	StartFib1_To_F1() int
}