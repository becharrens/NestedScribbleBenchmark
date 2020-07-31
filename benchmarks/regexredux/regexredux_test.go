package regexredux

import (
	"ScribbleBenchmark/benchmark"
	"ScribbleBenchmark/regexredux/protocol"
	"ScribbleBenchmark/regexredux_base"
	"testing"
)

func BenchmarkRegexRedux(b *testing.B) {
	benchmarks := []struct {
		name string
		file string
	}{
		{name: "1000", file: "regexredux-input1000.txt"},
		{name: "10000", file: "regexredux-input10000.txt"},
		{name: "100000", file: "regexredux-input100000.txt"},
		{name: "500000", file: "regexredux-input500000.txt"},
		{name: "1000000", file: "regexredux-input1000000.txt"},
		{name: "5000000", file: "regexredux-input5000000.txt"},
		{name: "10000000", file: "regexredux-input10000000.txt"},
		{name: "25000000", file: "regexredux-input25000000.txt"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				input := benchmark.ReadFile(bm.file)
				b.StartTimer()
				env := NewRegexReduxEnv(input)
				protocol.RegexRedux(env)
			}
		})
	}
}

func BenchmarkRegexReduxBase(b *testing.B) {
	benchmarks := []struct {
		name string
		file string
	}{
		{name: "1000", file: "regexredux-input1000.txt"},
		{name: "10000", file: "regexredux-input10000.txt"},
		{name: "100000", file: "regexredux-input100000.txt"},
		{name: "500000", file: "regexredux-input500000.txt"},
		{name: "1000000", file: "regexredux-input1000000.txt"},
		{name: "5000000", file: "regexredux-input5000000.txt"},
		{name: "10000000", file: "regexredux-input10000000.txt"},
		{name: "25000000", file: "regexredux-input25000000.txt"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				input := benchmark.ReadFile(bm.file)
				b.StartTimer()
				regexredux_base.RegexRedux(input)
			}
		})
	}
}
