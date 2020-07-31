package quicksort

import (
	"ScribbleBenchmark/quicksort/callbacks"
	"ScribbleBenchmark/quicksort/results/quicksort"
)

type QuickSortEnv struct {
	Arr       []int
	SortedArr []int
}

func (q *QuickSortEnv) New_Partition_Env() callbacks.QuickSort_Partition_Env {
	return &callbacks.QuickSortPartitionState{
		Arr: q.Arr,
	}
}

func (q *QuickSortEnv) New_Left_Env() callbacks.QuickSort_Left_Env {
	return &callbacks.QuickSortLeftState{}
}

func (q *QuickSortEnv) New_Right_Env() callbacks.QuickSort_Right_Env {
	return &callbacks.QuickSortRightState{}
}

func (q *QuickSortEnv) Partition_Result(result quicksort.Partition_Result) {
	q.SortedArr = result.SortedArr
}

func (q *QuickSortEnv) Left_Result(result quicksort.Left_Result) {
}

func (q *QuickSortEnv) Right_Result(result quicksort.Right_Result) {
}

var quickSortParams = []int{
	1000, 10000, 25000, 50000, 75000, 100000, 125000,
}

func NewQuickSortEnv(arr []int) *QuickSortEnv {
	return &QuickSortEnv{
		Arr: arr,
	}
}
