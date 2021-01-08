package callbacks

import (
	"NestedScribbleBenchmark/fibonacci_prev/messages/fib"
	"fmt"
	"strconv"
)
import fib_2 "NestedScribbleBenchmark/fibonacci_prev/results/fib"

type Fib_Res_Env interface {
	Done() fib_2.Res_Result
	ResultFrom_Fib_Res(result fib_2.Res_Result)
	To_Fib_Res_Env() Fib_Res_Env
	NextFib_From_F3(nextfib_msg fib.NextFib)
}

type FibResState struct {
	Idx int
}

func (f *FibResState) Done() fib_2.Res_Result {
	return fib_2.Res_Result{}
}

func (f *FibResState) ResultFrom_Fib_Res(result fib_2.Res_Result) {
}

func (f *FibResState) To_Fib_Res_Env() Fib_Res_Env {
	return &FibResState{
		Idx: f.Idx + 1,
	}
}

func (f *FibResState) NextFib_From_F3(nextfib_msg fib.NextFib) {
	fmt.Println("Fib "+strconv.Itoa(f.Idx)+":", nextfib_msg.Val)
}
