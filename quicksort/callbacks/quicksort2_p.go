package callbacks

import "NestedScribbleBenchmark/quicksort/messages/quicksort2"
import quicksort2_2 "NestedScribbleBenchmark/quicksort/results/quicksort2"

type QuickSort2_P_Choice int

const (
	QuickSort2_P_LeftParitition QuickSort2_P_Choice = iota
	QuickSort2_P_Done
)

type QuickSort2_P_Env interface {
	Done_To_R() quicksort2.Done
	Done_To_L() quicksort2.Done
	Done() quicksort2_2.P_Result
	SortedRight_From_R(sortedright_msg quicksort2.SortedRight)
	SortedLeft_From_L(sortedleft_msg quicksort2.SortedLeft)
	RightPartition_To_R() quicksort2.RightPartition
	LeftParitition_To_L() quicksort2.LeftParitition
	P_Choice() QuickSort2_P_Choice
}

type QuickSort2PState struct {
	Arr       []int
	SortedArr []int
	Pivot     int
}

func (q *QuickSort2PState) Done_To_R() quicksort2.Done {
	return quicksort2.Done{}
}

func (q *QuickSort2PState) Done_To_L() quicksort2.Done {
	return quicksort2.Done{}
}

func (q *QuickSort2PState) Done() quicksort2_2.P_Result {
	return quicksort2_2.P_Result{SortedArr: q.SortedArr}
}

func (q *QuickSort2PState) SortedRight_From_R(sortedright_msg quicksort2.SortedRight) {
	copy(q.SortedArr[q.Pivot+1:], sortedright_msg.Arr)
}

func (q *QuickSort2PState) SortedLeft_From_L(sortedleft_msg quicksort2.SortedLeft) {
	// q.SortedArr = make([]int, len(q.Arr))
	// copy(q.SortedArr, sortedleft_msg.Arr)
}

func (q *QuickSort2PState) RightPartition_To_R() quicksort2.RightPartition {
	return quicksort2.RightPartition{Arr: q.Arr[q.Pivot+1:]}
}

func (q *QuickSort2PState) LeftParitition_To_L() quicksort2.LeftParitition {
	return quicksort2.LeftParitition{Arr: q.Arr[:q.Pivot+1]}
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
