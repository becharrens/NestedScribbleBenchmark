package callbacks

import (
	"NestedScribbleBenchmark/fibonacci/results/fib"
	"fmt"
	"strconv"
)

type Fib_Res_Env interface {
	Done() fib.Res_Result
	ResultFrom_Fib_Res(result fib.Res_Result)
	To_Fib_Res_Env() Fib_Res_Env
	NextFib_From_F3(val int)
}

type FibResState struct {
	Idx int
}

func (f *FibResState) Done() fib.Res_Result {
	return fib.Res_Result{}
}

func (f *FibResState) ResultFrom_Fib_Res(result fib.Res_Result) {
}

func (f *FibResState) To_Fib_Res_Env() Fib_Res_Env {
	return &FibResState{
		Idx: f.Idx + 1,
	}
}

func (f *FibResState) NextFib_From_F3(val int) {
	fmt.Println("Fib "+strconv.Itoa(f.Idx)+":", val)
}
