package callbacks

import "NestedScribbleBenchmark/quicksort/results/quicksort2"
import "NestedScribbleBenchmark/quicksort/results/quicksort"

type QuickSort_Left_Env interface {
	Done_From_Partition()
	Done() quicksort.Left_Result
	SortedLeft_To_Partition() []int
	ResultFrom_QuickSort2_P(result quicksort2.P_Result)
	To_QuickSort2_P_Env() QuickSort2_P_Env
	QuickSort2_Setup()
	LeftParitition_From_Partition(arr []int)
}

type QuickSortLeftState struct {
	Arr       []int
	SortedArr []int
}

func (q *QuickSortLeftState) Done_From_Partition() {
}

func (q *QuickSortLeftState) Done() quicksort.Left_Result {
	return quicksort.Left_Result{}
}

func (q *QuickSortLeftState) SortedLeft_To_Partition() []int {
	return q.SortedArr
}

func (q *QuickSortLeftState) ResultFrom_QuickSort2_P(result quicksort2.P_Result) {
	q.SortedArr = result.SortedArr
}

func (q *QuickSortLeftState) To_QuickSort2_P_Env() QuickSort2_P_Env {
	return &QuickSort2PState{
		Arr: q.Arr,
	}
}

func (q *QuickSortLeftState) QuickSort2_Setup() {
}

func (q *QuickSortLeftState) LeftParitition_From_Partition(arr []int) {
	q.Arr = arr
}
