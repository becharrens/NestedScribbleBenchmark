package fib

import "NestedScribbleBenchmark/fibonacci/messages/fib"

type Res_Chan struct {
	F3_NextFib chan fib.NextFib
}

type F1_Chan struct {
	F3_Fib1 chan fib.Fib1
}

type F2_Chan struct {
	F3_Fib2 chan fib.Fib2
}

type F3_Chan struct {
	F1_Fib1 chan fib.Fib1
	F2_Fib2 chan fib.Fib2
	Res_NextFib chan fib.NextFib
}