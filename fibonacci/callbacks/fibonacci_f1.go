package callbacks

import "NestedScribbleBenchmark/fibonacci/results/fibonacci"
import "NestedScribbleBenchmark/fibonacci/results/fib"

type Fibonacci_F1_Env interface {
	Done() fibonacci.F1_Result
	ResultFrom_Fib_F1(result fib.F1_Result) 
	To_Fib_F1_Env() Fib_F1_Env
	StartFib1_From_Start(val int) 
}