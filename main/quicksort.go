package main

import (
	"NestedScribbleBenchmark/benchmark"
	"NestedScribbleBenchmark/quicksort/callbacks"
	"NestedScribbleBenchmark/quicksort/protocol"
	"NestedScribbleBenchmark/quicksort/results/quicksort"
	"NestedScribbleBenchmark/quicksort_base"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

const (
	UBOUND = 1<<32 - 1
)

// var Seed int64 = 815267917
var Seed int64 = 274734990

var SEQ_THRESHOLD = 7500

type QuickSortEnv struct {
	Arr          []int
	SortedArr    []int
	SeqThreshold int
}

func (q *QuickSortEnv) New_Partition_Env() callbacks.QuickSort_Partition_Env {
	return &callbacks.QuickSortPartitionState{
		Arr:          q.Arr,
		SeqThreshold: SEQ_THRESHOLD,
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
	1000, 10000, 25000, 50000, 75000, 125000,
	100000, 250000, 500000, 1000000, 2000000, 3000000, 5000000, 7500000, 10000000, 15000000, 30000000, 45000000, 60000000,
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
	fmt.Println("Base done")
	// return base_results, base_results
	return scribble_results, base_results
}

func QuickSortBenchmark2(seqThreshold, repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	SEQ_THRESHOLD = seqThreshold
	quicksort_base.SEQ_THRESHOLD = seqThreshold
	rand.Seed(Seed)
	scribble_results := benchmark.TimeImpl(quickSortParams, repetitions, TimeQuickSort)
	fmt.Println("Scribble done")
	rand.Seed(Seed)
	base_results := benchmark.TimeImpl(quickSortParams, repetitions, TimeQuickSortBase)
	fmt.Println("Base done")
	return scribble_results, base_results
}

func QSThresholdSearch(repetitions int) {
	for i := 500; i < 10001; i += 500 {
		fmt.Println("Threshold =", i)
		scribble_results, base_results := QuickSortBenchmark2(i, repetitions)
		scrResStr := benchmark.ResultsToString("quicksort-scribble", scribble_results) + "\n;;"
		baseResStr := benchmark.ResultsToString("quicksort-base", base_results) + "\n;;"
		resultName := fmt.Sprintf("benchmark-results%d.txt", i)
		err := ioutil.WriteFile(resultName, []byte(scrResStr+"\n"+baseResStr), 0644)
		if err != nil {
			panic("Error while writing to file")
		}
	}
}
