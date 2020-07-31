package quicksort2

import "ScribbleBenchmark/quicksort/messages/quicksort2"

type P_Chan struct {
	L_Done chan quicksort2.Done
	L_LeftParitition chan quicksort2.LeftParitition
	L_SortedLeft chan quicksort2.SortedLeft
	R_Done chan quicksort2.Done
	R_RightPartition chan quicksort2.RightPartition
	R_SortedRight chan quicksort2.SortedRight
}

type L_Chan struct {
	P_Done chan quicksort2.Done
	P_LeftParitition chan quicksort2.LeftParitition
	P_SortedLeft chan quicksort2.SortedLeft
}

type R_Chan struct {
	P_Done chan quicksort2.Done
	P_RightPartition chan quicksort2.RightPartition
	P_SortedRight chan quicksort2.SortedRight
}