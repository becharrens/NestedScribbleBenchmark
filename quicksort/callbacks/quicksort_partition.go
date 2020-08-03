package callbacks

import (
	"NestedScribbleBenchmark/quicksort/messages/quicksort"
)
import quicksort_2 "NestedScribbleBenchmark/quicksort/results/quicksort"

type QuickSort_Partition_Choice int

const (
	QuickSort_Partition_LeftParitition QuickSort_Partition_Choice = iota
	QuickSort_Partition_Done
)

const (
	// SEQ_THRESHOLD = 1024
	SEQ_THRESHOLD = 4000
)

type QuickSort_Partition_Env interface {
	Done_To_Right() quicksort.Done
	Done_To_Left() quicksort.Done
	Done() quicksort_2.Partition_Result
	SortedRight_From_Right(sortedright_msg quicksort.SortedRight)
	SortedLeft_From_Left(sortedleft_msg quicksort.SortedLeft)
	RightPartition_To_Right() quicksort.RightPartition
	LeftParitition_To_Left() quicksort.LeftParitition
	Partition_Choice() QuickSort_Partition_Choice
}

type QuickSortPartitionState struct {
	Arr       []int
	SortedArr []int
	Pivot     int
}

func (q *QuickSortPartitionState) Done_To_Right() quicksort.Done {
	return quicksort.Done{}
}

func (q *QuickSortPartitionState) Done_To_Left() quicksort.Done {
	return quicksort.Done{}
}

func (q *QuickSortPartitionState) Done() quicksort_2.Partition_Result {
	return quicksort_2.Partition_Result{SortedArr: q.SortedArr}
}

func (q *QuickSortPartitionState) SortedRight_From_Right(sortedright_msg quicksort.SortedRight) {
	// copy(q.SortedArr[q.Pivot + 1:], sortedright_msg.Arr)
}

func (q *QuickSortPartitionState) SortedLeft_From_Left(sortedleft_msg quicksort.SortedLeft) {
	// q.SortedArr = make([]int, len(q.Arr))
	// copy(q.SortedArr, sortedleft_msg.Arr)
}

func (q *QuickSortPartitionState) RightPartition_To_Right() quicksort.RightPartition {
	return quicksort.RightPartition{Arr: q.Arr[q.Pivot+1:]}
}

func (q *QuickSortPartitionState) LeftParitition_To_Left() quicksort.LeftParitition {
	return quicksort.LeftParitition{Arr: q.Arr[:q.Pivot+1]}
}

func (q *QuickSortPartitionState) Partition_Choice() QuickSort_Partition_Choice {
	q.SortedArr = q.Arr
	if len(q.Arr) < SEQ_THRESHOLD {
		seqQuickSort(q.Arr)
		return QuickSort_Partition_Done
	}
	q.Pivot = hoarePartition(q.Arr, 0, len(q.Arr)-1)
	return QuickSort_Partition_LeftParitition
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func hoarePartition(arr []int, low, high int) int {
	pivot := arr[low]
	i, j := low-1, high+1

	for {
		for {
			i++
			if arr[i] >= pivot {
				break
			}
		}

		for {
			j--
			if arr[j] <= pivot {
				break
			}
		}

		if i >= j {
			return j
		}
		swap(arr, i, j)
	}
}

func seqQuickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	pivot := hoarePartition(arr, 0, len(arr)-1)
	seqQuickSort(arr[:pivot+1])
	seqQuickSort(arr[pivot+1:])
}
