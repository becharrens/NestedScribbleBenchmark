package callbacks

import "NestedScribbleBenchmark/quicksort/results/quicksort2"

type QuickSort2_R_Env interface {
	SortedRight_To_P() []int
	ResultFrom_QuickSort2_P(result quicksort2.P_Result)
	To_QuickSort2_P_Env() QuickSort2_P_Env
	QuickSort2_Setup()
	RightPartition_From_P(arr []int)
	Done()
	Done_From_P()
}

type QuickSort2RState struct {
	Arr       []int
	SortedArr []int
}

func (q *QuickSort2RState) SortedRight_To_P() []int {
	return q.SortedArr
}

func (q *QuickSort2RState) ResultFrom_QuickSort2_P(result quicksort2.P_Result) {
	q.SortedArr = result.SortedArr
}

func (q *QuickSort2RState) To_QuickSort2_P_Env() QuickSort2_P_Env {
	return &QuickSort2PState{
		Arr: q.Arr,
	}
}

func (q *QuickSort2RState) QuickSort2_Setup() {
}

func (q *QuickSort2RState) RightPartition_From_P(arr []int) {
	q.Arr = arr
}

func (q *QuickSort2RState) Done() {
}

func (q *QuickSort2RState) Done_From_P() {
}

func New_QuickSort2_R_State() QuickSort2_R_Env {
	return &QuickSort2RState{}
}
