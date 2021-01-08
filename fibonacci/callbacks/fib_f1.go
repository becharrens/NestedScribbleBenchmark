package callbacks

import "NestedScribbleBenchmark/fibonacci/results/fib"

// TODO

type Fib_F1_Env interface {
	Done() fib.F1_Result
	Fib1_To_F3() int
}
