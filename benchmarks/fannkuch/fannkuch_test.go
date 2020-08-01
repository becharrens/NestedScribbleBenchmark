package fannkuch

import (
	"NestedScribbleBenchmark/fannkuch/protocol"
	"NestedScribbleBenchmark/fannkuch_base"
	"testing"
)

func BenchmarkFannkuch(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{name: "4", n: 4},
		{name: "5", n: 5},
		{name: "6", n: 6},
		{name: "7", n: 7},
		{name: "8", n: 8},
		{name: "9", n: 9},
		{name: "10", n: 10},
		{name: "11", n: 11},
		{name: "12", n: 12},
	}
	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				env := NewFannkuchEnv(bm.n)
				b.StartTimer()
				protocol.Fannkuch(env)
			}
		})
	}
}

func BenchmarkFannkuchBase(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{name: "4", n: 4},
		{name: "5", n: 5},
		{name: "6", n: 6},
		{name: "7", n: 7},
		{name: "8", n: 8},
		{name: "9", n: 9},
		{name: "10", n: 10},
		{name: "11", n: 11},
		{name: "12", n: 12},
	}
	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fannkuch_base.Fannkuch(bm.n)
			}
		})
	}
}
