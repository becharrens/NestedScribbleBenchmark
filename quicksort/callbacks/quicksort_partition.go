package callbacks

import "NestedScribbleBenchmark/quicksort/results/quicksort"

type QuickSort_Partition_Choice int

const (
	QuickSort_Partition_LeftParitition QuickSort_Partition_Choice = iota
	QuickSort_Partition_Done
)

var SEQ_THRESHOLD = 3500

type QuickSort_Partition_Env interface {
	Done_To_Right()
	Done_To_Left()
	Done() quicksort.Partition_Result
	SortedRight_From_Right(arr []int)
	SortedLeft_From_Left(arr []int)
	RightPartition_To_Right() []int
	LeftParitition_To_Left() []int
	Partition_Choice() QuickSort_Partition_Choice
}

type QuickSortPartitionState struct {
	Arr          []int
	SortedArr    []int
	Pivot        int
	SeqThreshold int
}

func (q *QuickSortPartitionState) Done_To_Right() {
}

func (q *QuickSortPartitionState) Done_To_Left() {
}

func (q *QuickSortPartitionState) Done() quicksort.Partition_Result {
	return quicksort.Partition_Result{SortedArr: q.SortedArr}
}

func (q *QuickSortPartitionState) SortedRight_From_Right(arr []int) {
	// copy(q.SortedArr[q.Pivot + 1:], sortedright_msg.Arr)
}

func (q *QuickSortPartitionState) SortedLeft_From_Left(arr []int) {
	// q.SortedArr = make([]int, len(q.Arr))
	// copy(q.SortedArr, sortedleft_msg.Arr)
}

func (q *QuickSortPartitionState) RightPartition_To_Right() []int {
	return q.Arr[q.Pivot+1:]
}

func (q *QuickSortPartitionState) LeftParitition_To_Left() []int {
	return q.Arr[:q.Pivot+1]
}

func (q *QuickSortPartitionState) Partition_Choice() QuickSort_Partition_Choice {
	SEQ_THRESHOLD = q.SeqThreshold
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
