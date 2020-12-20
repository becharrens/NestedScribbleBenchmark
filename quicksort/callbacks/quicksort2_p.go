package callbacks

import "NestedScribbleBenchmark/quicksort/results/quicksort2"

type QuickSort2_P_Choice int

const (
	QuickSort2_P_LeftParitition QuickSort2_P_Choice = iota
	QuickSort2_P_Done
)

type QuickSort2_P_Env interface {
	Done_To_R()
	Done_To_L()
	Done() quicksort2.P_Result
	SortedRight_From_R(arr []int)
	SortedLeft_From_L(arr []int)
	RightPartition_To_R() []int
	LeftParitition_To_L() []int
	P_Choice() QuickSort2_P_Choice
}

type QuickSort2PState struct {
	Arr       []int
	SortedArr []int
	Pivot     int
}

func (q *QuickSort2PState) Done_To_R() {
}

func (q *QuickSort2PState) Done_To_L() {
}

func (q *QuickSort2PState) Done() quicksort2.P_Result {
	return quicksort2.P_Result{SortedArr: q.SortedArr}
}

func (q *QuickSort2PState) SortedRight_From_R(arr []int) {
	// copy(q.SortedArr[q.Pivot+1:], arr)
}

func (q *QuickSort2PState) SortedLeft_From_L(arr []int) {
	// q.SortedArr = make([]int, len(q.Arr))
	// copy(q.SortedArr, sortedleft_msg.Arr)
}

func (q *QuickSort2PState) RightPartition_To_R() []int {
	return q.Arr[q.Pivot+1:]
}

func (q *QuickSort2PState) LeftParitition_To_L() []int {
	return q.Arr[:q.Pivot+1]
}

func (q *QuickSort2PState) P_Choice() QuickSort2_P_Choice {
	q.SortedArr = q.Arr
	if len(q.Arr) < SEQ_THRESHOLD {
		seqQuickSort(q.Arr)
		return QuickSort2_P_Done
	}
	q.Pivot = hoarePartition(q.Arr, 0, len(q.Arr)-1)
	return QuickSort2_P_LeftParitition
}
