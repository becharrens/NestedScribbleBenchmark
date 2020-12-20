package quicksort

import "NestedScribbleBenchmark/quicksort/messages"

type Partition_Chan struct {
	IntArr_From_Left  chan []int
	IntArr_From_Right chan []int
	IntArr_To_Left    chan []int
	IntArr_To_Right   chan []int
	Label_From_Left   chan messages.QuickSort_Label
	Label_From_Right  chan messages.QuickSort_Label
	Label_To_Left     chan messages.QuickSort_Label
	Label_To_Right    chan messages.QuickSort_Label
}

type Left_Chan struct {
	IntArr_From_Partition chan []int
	IntArr_To_Partition   chan []int
	Label_From_Partition  chan messages.QuickSort_Label
	Label_To_Partition    chan messages.QuickSort_Label
}

type Right_Chan struct {
	IntArr_From_Partition chan []int
	IntArr_To_Partition   chan []int
	Label_From_Partition  chan messages.QuickSort_Label
	Label_To_Partition    chan messages.QuickSort_Label
}
