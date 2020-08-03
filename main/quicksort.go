package main

import (
	"NestedScribbleBenchmark/benchmark"
	"NestedScribbleBenchmark/quicksort/callbacks"
	"NestedScribbleBenchmark/quicksort/protocol"
	"NestedScribbleBenchmark/quicksort/results/quicksort"
	"NestedScribbleBenchmark/quicksort_base"
	"fmt"
	"math/rand"
	"time"
)

const (
	UBOUND = 1<<32 - 1
)

// var Seed int64 = 815267917
var Seed int64 = 274734990

type QuickSortEnv struct {
	Arr       []int
	SortedArr []int
}

func (q *QuickSortEnv) New_Partition_Env() callbacks.QuickSort_Partition_Env {
	return &callbacks.QuickSortPartitionState{
		Arr: q.Arr,
	}
}

func (q *QuickSortEnv) New_Left_Env() callbacks.QuickSort_Left_Env {
	return &callbacks.QuickSortLeftState{}
}

func (q *QuickSortEnv) New_Right_Env() callbacks.QuickSort_Right_Env {
	return &callbacks.QuickSortRightState{}
}

func (q *QuickSortEnv) Partition_Result(result quicksort.Partition_Result) {
	q.SortedArr = result.SortedArr
}

func (q *QuickSortEnv) Left_Result(result quicksort.Left_Result) {
}

func (q *QuickSortEnv) Right_Result(result quicksort.Right_Result) {
}

var quickSortParams = []int{
	// 1000,
	100000, 500000, 1000000, 2000000, 5000000, 10000000, 50000000, 100000000,
	// 1000, 10000, 25000, 50000, 75000, 100000, 125000
}

func NewQuickSortEnv(n int) *QuickSortEnv {
	return &QuickSortEnv{
		Arr: RandomArr(n),
	}
}

func RandomArr(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Int()
	}
	return arr
}

func TimeQuickSort(n int) time.Duration {
	env := NewQuickSortEnv(n)
	start := time.Now()
	protocol.QuickSort(env)
	elapsed := time.Since(start)
	// DEBUG
	// fmt.Println(sort.IntsAreSorted(env.SortedArr))
	return elapsed
}

func TimeQuickSortBase(n int) time.Duration {
	arr := RandomArr(n)
	start := time.Now()
	quicksort_base.QuickSort(arr)
	elapsed := time.Since(start)
	// DEBUG
	// fmt.Println(sort.IntsAreSorted(arr))
	return elapsed
}

var arrays [][]int

func GenArrays() {
	rand.Seed(Seed)
	for _, v := range quickSortParams {
		arrays = append(arrays, RandomArr(v))
	}
}

func QuickSortBenchmark(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	rand.Seed(Seed)
	scribble_results := benchmark.TimeImpl(quickSortParams, repetitions, TimeQuickSort)
	fmt.Println("Scribble done")
	rand.Seed(Seed)
	base_results := benchmark.TimeImpl(quickSortParams, repetitions, TimeQuickSortBase)
	// return base_results, base_results
	return scribble_results, base_results
}
