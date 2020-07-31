package callbacks

import "ScribbleBenchmark/quicksort/messages/quicksort2"
import quicksort2_2 "ScribbleBenchmark/quicksort/results/quicksort2"

type QuickSort2_R_Env interface {
	SortedRight_To_P() quicksort2.SortedRight
	ResultFrom_QuickSort2_P(result quicksort2_2.P_Result)
	To_QuickSort2_P_Env() QuickSort2_P_Env
	QuickSort2_Setup()
	RightPartition_From_P(rightpartition_msg quicksort2.RightPartition)
	Done()
	Done_From_P(done_msg quicksort2.Done)
}

type QuickSort2RState struct {
	Arr       []int
	SortedArr []int
}

func (q *QuickSort2RState) SortedRight_To_P() quicksort2.SortedRight {
	return quicksort2.SortedRight{Arr: q.SortedArr}
}

func (q *QuickSort2RState) ResultFrom_QuickSort2_P(result quicksort2_2.P_Result) {
	q.SortedArr = result.SortedArr
}

func (q *QuickSort2RState) To_QuickSort2_P_Env() QuickSort2_P_Env {
	return &QuickSort2PState{
		Arr: q.Arr,
	}
}

func (q *QuickSort2RState) QuickSort2_Setup() {
}

func (q *QuickSort2RState) RightPartition_From_P(rightpartition_msg quicksort2.RightPartition) {
	q.Arr = rightpartition_msg.Arr
}

func (q *QuickSort2RState) Done() {
}

func (q *QuickSort2RState) Done_From_P(done_msg quicksort2.Done) {
}

func New_QuickSort2_R_State() QuickSort2_R_Env {
	return &QuickSort2RState{}
}
