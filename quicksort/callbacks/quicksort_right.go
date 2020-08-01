package callbacks

import "NestedScribbleBenchmark/quicksort/messages/quicksort"
import "NestedScribbleBenchmark/quicksort/results/quicksort2"
import quicksort_2 "NestedScribbleBenchmark/quicksort/results/quicksort"

type QuickSort_Right_Env interface {
	SortedRight_To_Partition() quicksort.SortedRight
	ResultFrom_QuickSort2_P(result quicksort2.P_Result)
	To_QuickSort2_P_Env() QuickSort2_P_Env
	QuickSort2_Setup()
	RightPartition_From_Partition(rightpartition_msg quicksort.RightPartition)
	Done() quicksort_2.Right_Result
	Done_From_Partition(done_msg quicksort.Done)
}

type QuickSortRightState struct {
	Arr       []int
	SortedArr []int
}

func (q *QuickSortRightState) SortedRight_To_Partition() quicksort.SortedRight {
	return quicksort.SortedRight{Arr: q.SortedArr}
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

func (q *QuickSortRightState) RightPartition_From_Partition(rightpartition_msg quicksort.RightPartition) {
	q.Arr = rightpartition_msg.Arr
}

func (q *QuickSortRightState) Done() quicksort_2.Right_Result {
	return quicksort_2.Right_Result{}
}

func (q *QuickSortRightState) Done_From_Partition(done_msg quicksort.Done) {
}
