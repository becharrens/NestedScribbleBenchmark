package callbacks

import "NestedScribbleBenchmark/fibonacci/results/fib"

type Fib_F3_Env interface {
	Done() 
	ResultFrom_Fib_F2(result fib.F2_Result) 
	To_Fib_F2_Env() Fib_F2_Env
	Fib_Setup() 
	NextFib_To_Res() int
	Fib2_From_F2(val int) 
	Fib1_From_F1(val int) 
}

func New_Fib_F3_State() Fib_F3_Env {
	panic("TODO: implement me")
} 