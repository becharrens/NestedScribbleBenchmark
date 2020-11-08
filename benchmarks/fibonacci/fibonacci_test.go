package fibonacci

import (
	"NestedScribbleBenchmark/bounded_fib_base"
	"NestedScribbleBenchmark/old_fibonacci/protocol"
	"testing"
)

func BenchmarkFibonacci(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{name: "5", n: 5},
		{name: "10", n: 10},
		{name: "15", n: 15},
		{name: "20", n: 20},
		{name: "25", n: 25},
		{name: "30", n: 30},
		{name: "35", n: 35},
		{name: "40", n: 40},
		{name: "45", n: 45},
		{name: "50", n: 50},
		{name: "55", n: 55},
		{name: "60", n: 60},
		{name: "65", n: 65},
		{name: "70", n: 70},
		{name: "75", n: 75},
		{name: "80", n: 80},
		{name: "85", n: 85},
		{name: "90", n: 90},
	}
	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				env := NewFibonacciEnv(bm.n)
				b.StartTimer()
				protocol.Fibonacci(env)
			}
		})
	}
}

func BenchmarkFibonacciBase(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{name: "5", n: 5},
		{name: "10", n: 10},
		{name: "15", n: 15},
		{name: "20", n: 20},
		{name: "25", n: 25},
		{name: "30", n: 30},
		{name: "35", n: 35},
		{name: "40", n: 40},
		{name: "45", n: 45},
		{name: "50", n: 50},
		{name: "55", n: 55},
		{name: "60", n: 60},
		{name: "65", n: 65},
		{name: "70", n: 70},
		{name: "75", n: 75},
		{name: "80", n: 80},
		{name: "85", n: 85},
		{name: "90", n: 90},
	}
	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bounded_fib_base.Fibonacci(bm.n)
			}
		})
	}
}
