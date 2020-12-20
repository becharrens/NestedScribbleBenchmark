package callbacks

import "NestedScribbleBenchmark/quicksort/results/quicksort2"
import "NestedScribbleBenchmark/quicksort/results/quicksort"

type QuickSort_Right_Env interface {
	SortedRight_To_Partition() []int
	ResultFrom_QuickSort2_P(result quicksort2.P_Result)
	To_QuickSort2_P_Env() QuickSort2_P_Env
	QuickSort2_Setup()
	RightPartition_From_Partition(arr []int)
	Done() quicksort.Right_Result
	Done_From_Partition()
}

type QuickSortRightState struct {
	Arr       []int
	SortedArr []int
}

func (q *QuickSortRightState) SortedRight_To_Partition() []int {
	return q.SortedArr
}

func (q *QuickSortRightState) ResultFrom_QuickSort2_P(result quicksort2.P_Result) {
	q.SortedArr = result.SortedArr
}

func (q *QuickSortRightState) To_QuickSort2_P_Env() QuickSort2_P_Env {
	return &QuickSort2PState{
		Arr: q.Arr,
	}
}

func (q *QuickSortRightState) QuickSort2_Setup() {
}

func (q *QuickSortRightState) RightPartition_From_Partition(arr []int) {
	q.Arr = arr
}

func (q *QuickSortRightState) Done() quicksort.Right_Result {
	return quicksort.Right_Result{}
}

func (q *QuickSortRightState) Done_From_Partition() {
}
