package boundedfibonacci

import "NestedScribbleBenchmark/boundedfibonacci/messages/boundedfibonacci"

type Start_Chan struct {
	F1_StartFib1 chan boundedfibonacci.StartFib1
	F2_StartFib2 chan boundedfibonacci.StartFib2
}

type F1_Chan struct {
	Start_StartFib1 chan boundedfibonacci.StartFib1
}

type F2_Chan struct {
	Start_StartFib2 chan boundedfibonacci.StartFib2
}