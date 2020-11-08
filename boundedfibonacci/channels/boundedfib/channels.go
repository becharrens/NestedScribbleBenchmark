package boundedfib

import "NestedScribbleBenchmark/boundedfibonacci/messages/boundedfib"

type Res_Chan struct {
	F3_Result chan boundedfib.Result
}

type F1_Chan struct {
	F3_Fib1 chan boundedfib.Fib1
}

type F2_Chan struct {
	F3_End chan boundedfib.End
	F3_Fib2 chan boundedfib.Fib2
}

type F3_Chan struct {
	F1_Fib1 chan boundedfib.Fib1
	F2_End chan boundedfib.End
	F2_Fib2 chan boundedfib.Fib2
	Res_Result chan boundedfib.Result
}