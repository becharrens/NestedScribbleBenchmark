package quicksort

import "NestedScribbleBenchmark/quicksort/messages/quicksort"

type Partition_Chan struct {
	Left_Done            chan quicksort.Done
	Left_LeftParitition  chan quicksort.LeftParitition
	Left_SortedLeft      chan quicksort.SortedLeft
	Right_Done           chan quicksort.Done
	Right_RightPartition chan quicksort.RightPartition
	Right_SortedRight    chan quicksort.SortedRight
}

type Left_Chan struct {
	Partition_Done           chan quicksort.Done
	Partition_LeftParitition chan quicksort.LeftParitition
	Partition_SortedLeft     chan quicksort.SortedLeft
}

type Right_Chan struct {
	Partition_Done           chan quicksort.Done
	Partition_RightPartition chan quicksort.RightPartition
	Partition_SortedRight    chan quicksort.SortedRight
}
