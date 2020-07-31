package quicksort

import (
	"ScribbleBenchmark/quicksort/protocol"
	"ScribbleBenchmark/quicksort_base"
	"math/rand"
	"testing"
)

const SEED int64 = 274734990

func BenchmarkQuickSort(b *testing.B) {
	rand.Seed(SEED)
	benchmarks := []struct {
		name string
		arr  []int
	}{
		{name: "1000", arr: RandomArr(1000)},
		{name: "5500", arr: RandomArr(5500)},
		{name: "10000", arr: RandomArr(10000)},
		{name: "25000", arr: RandomArr(25000)},
		{name: "50000", arr: RandomArr(50000)},
		{name: "75000", arr: RandomArr(75000)},
		{name: "100000", arr: RandomArr(100000)},
		{name: "125000", arr: RandomArr(125000)},
		{name: "150000", arr: RandomArr(150000)},
		{name: "175000", arr: RandomArr(175000)},
		{name: "200000", arr: RandomArr(200000)},
	}
	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := make([]int, len(bm.arr))
				copy(arr, bm.arr)
				env := NewQuickSortEnv(arr)
				b.StartTimer()
				protocol.QuickSort(env)
			}
		})
	}
}

func BenchmarkQuickSortBase(b *testing.B) {
	rand.Seed(SEED)
	benchmarks := []struct {
		name string
		arr  []int
	}{
		{name: "1000", arr: RandomArr(1000)},
		{name: "5500", arr: RandomArr(5500)},
		{name: "10000", arr: RandomArr(10000)},
		{name: "25000", arr: RandomArr(25000)},
		{name: "50000", arr: RandomArr(50000)},
		{name: "75000", arr: RandomArr(75000)},
		{name: "100000", arr: RandomArr(100000)},
		{name: "125000", arr: RandomArr(125000)},
		{name: "150000", arr: RandomArr(150000)},
		{name: "175000", arr: RandomArr(175000)},
		{name: "200000", arr: RandomArr(200000)},
	}
	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				arr := make([]int, len(bm.arr))
				copy(arr, bm.arr)
				b.StartTimer()
				quicksort_base.QuickSort(arr)
			}
		})
	}
}

func RandomArr(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Int()
	}
	return arr
}
