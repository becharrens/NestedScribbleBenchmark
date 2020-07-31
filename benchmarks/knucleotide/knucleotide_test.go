package knucleotide

import (
	"ScribbleBenchmark/benchmark"
	"ScribbleBenchmark/knucleotide/protocol"
	"ScribbleBenchmark/knucleotide_base"
	"testing"
)

func BenchmarkKNucleotide(b *testing.B) {
	benchmarks := []struct {
		name string
		file string
	}{
		{name: "1000", file: "knucleotide-input1000.txt"},
		{name: "10000", file: "knucleotide-input10000.txt"},
		{name: "50000", file: "knucleotide-input50000.txt"},
		{name: "100000", file: "knucleotide-input100000.txt"},
		{name: "500000", file: "knucleotide-input500000.txt"},
		{name: "1000000", file: "knucleotide-input1000000.txt"},
		{name: "2500000", file: "knucleotide-input2500000.txt"},
		{name: "5000000", file: "knucleotide-input5000000.txt"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {

			input := benchmark.ReadFile(bm.file)
			dna := toBits(readSequence(">THREE", input))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				env := NewKNucleotideEnv(dna)
				protocol.KNucleotide(env)
			}
		})
	}
}

func BenchmarkKNucleotideBase(b *testing.B) {
	benchmarks := []struct {
		name string
		file string
	}{
		{name: "1000", file: "knucleotide-input1000.txt"},
		{name: "10000", file: "knucleotide-input10000.txt"},
		{name: "50000", file: "knucleotide-input50000.txt"},
		{name: "100000", file: "knucleotide-input100000.txt"},
		{name: "500000", file: "knucleotide-input500000.txt"},
		{name: "1000000", file: "knucleotide-input1000000.txt"},
		{name: "2500000", file: "knucleotide-input2500000.txt"},
		{name: "5000000", file: "knucleotide-input5000000.txt"},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {

			input := benchmark.ReadFile(bm.file)
			dna := toBits(readSequence(">THREE", input))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				knucleotide_base.KNucleotide(dna)
			}
		})
	}
}
