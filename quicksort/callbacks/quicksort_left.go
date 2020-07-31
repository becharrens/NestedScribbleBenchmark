package callbacks

import "ScribbleBenchmark/quicksort/messages/quicksort"
import "ScribbleBenchmark/quicksort/results/quicksort2"
import quicksort_2 "ScribbleBenchmark/quicksort/results/quicksort"

type QuickSort_Left_Env interface {
	Done_From_Partition(done_msg quicksort.Done)
	Done() quicksort_2.Left_Result
	SortedLeft_To_Partition() quicksort.SortedLeft
	ResultFrom_QuickSort2_P(result quicksort2.P_Result)
	To_QuickSort2_P_Env() QuickSort2_P_Env
	QuickSort2_Setup()
	LeftParitition_From_Partition(leftparitition_msg quicksort.LeftParitition)
}

type QuickSortLeftState struct {
	Arr       []int
	SortedArr []int
}

func (q *QuickSortLeftState) Done_From_Partition(done_msg quicksort.Done) {
}

func (q *QuickSortLeftState) Done() quicksort_2.Left_Result {
	return quicksort_2.Left_Result{}
}

func (q *QuickSortLeftState) SortedLeft_To_Partition() quicksort.SortedLeft {
	return quicksort.SortedLeft{Arr: q.SortedArr}
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

func (q *QuickSortLeftState) LeftParitition_From_Partition(leftparitition_msg quicksort.LeftParitition) {
	q.Arr = leftparitition_msg.Arr
}
