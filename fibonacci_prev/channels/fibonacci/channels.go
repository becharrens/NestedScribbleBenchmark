package fibonacci

import "NestedScribbleBenchmark/fibonacci_prev/messages/fibonacci"

type Start_Chan struct {
	F1_StartFib1 chan fibonacci.StartFib1
	F2_StartFib2 chan fibonacci.StartFib2
}

type F1_Chan struct {
	Start_StartFib1 chan fibonacci.StartFib1
}

type F2_Chan struct {
	Start_StartFib2 chan fibonacci.StartFib2
}
