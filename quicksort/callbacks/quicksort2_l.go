package callbacks

import "NestedScribbleBenchmark/quicksort/messages/quicksort2"
import quicksort2_2 "NestedScribbleBenchmark/quicksort/results/quicksort2"

type QuickSort2_L_Env interface {
	Done_From_P(done_msg quicksort2.Done)
	Done()
	SortedLeft_To_P() quicksort2.SortedLeft
	ResultFrom_QuickSort2_P(result quicksort2_2.P_Result)
	To_QuickSort2_P_Env() QuickSort2_P_Env
	QuickSort2_Setup()
	LeftParitition_From_P(leftparitition_msg quicksort2.LeftParitition)
}

type QuickSort2LState struct {
	Arr       []int
	SortedArr []int
}

func (q *QuickSort2LState) Done_From_P(done_msg quicksort2.Done) {
}

func (q *QuickSort2LState) Done() {
}

func (q *QuickSort2LState) SortedLeft_To_P() quicksort2.SortedLeft {
	return quicksort2.SortedLeft{Arr: q.SortedArr}
}

func (q *QuickSort2LState) ResultFrom_QuickSort2_P(result quicksort2_2.P_Result) {
	q.SortedArr = result.SortedArr
}

func (q *QuickSort2LState) To_QuickSort2_P_Env() QuickSort2_P_Env {
	return &QuickSort2PState{
		Arr: q.Arr,
	}
}

func (q *QuickSort2LState) QuickSort2_Setup() {
}

func (q *QuickSort2LState) LeftParitition_From_P(leftparitition_msg quicksort2.LeftParitition) {
	q.Arr = leftparitition_msg.Arr
}

func New_QuickSort2_L_State() QuickSort2_L_Env {
	return &QuickSort2LState{}
}
