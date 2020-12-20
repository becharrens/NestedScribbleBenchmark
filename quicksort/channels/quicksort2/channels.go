package quicksort2

import "NestedScribbleBenchmark/quicksort/messages"

type P_Chan struct {
	IntArr_From_L chan []int
	IntArr_From_R chan []int
	IntArr_To_L   chan []int
	IntArr_To_R   chan []int
	Label_From_L  chan messages.QuickSort_Label
	Label_From_R  chan messages.QuickSort_Label
	Label_To_L    chan messages.QuickSort_Label
	Label_To_R    chan messages.QuickSort_Label
}

type L_Chan struct {
	IntArr_From_P chan []int
	IntArr_To_P   chan []int
	Label_From_P  chan messages.QuickSort_Label
	Label_To_P    chan messages.QuickSort_Label
}

type R_Chan struct {
	IntArr_From_P chan []int
	IntArr_To_P   chan []int
	Label_From_P  chan messages.QuickSort_Label
	Label_To_P    chan messages.QuickSort_Label
}
