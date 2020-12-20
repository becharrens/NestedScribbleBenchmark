package main

import (
	"NestedScribbleBenchmark/benchmark"
	"NestedScribbleBenchmark/knucleotide/callbacks"
	"NestedScribbleBenchmark/knucleotide/protocol"
	"NestedScribbleBenchmark/knucleotide/results/knucleotide"
	"NestedScribbleBenchmark/knucleotide_base"
	"bufio"
	"bytes"
	"time"
)

type KNucleotideEnv struct {
	B []byte
}

func (k *KNucleotideEnv) New_Master_Env() callbacks.KNucleotide_Master_Env {
	return &callbacks.KNucleotideMasterState{
		Dna:        k.B,
		PatternIdx: len(callbacks.Patterns) - 1,
		LenIdx:     len(callbacks.Lengths) - 1,
	}
}

func (k *KNucleotideEnv) Master_Result(result knucleotide.Master_Result) {
}

var kNucleotideParams = []int{
	0, 1, 2, 3, 4, 5, 6, 7,
}
var knucleotideFiles = []string{
	"knucleotide-input1000.txt",
	"knucleotide-input10000.txt",
	"knucleotide-input50000.txt",
	"knucleotide-input100000.txt",
	"knucleotide-input500000.txt",
	"knucleotide-input1000000.txt",
	"knucleotide-input2500000.txt",
	"knucleotide-input5000000.txt",
}

func NewKNucleotideEnv(n int) *KNucleotideEnv {
	b := readFile(knucleotideFiles[n])
	dna := toBits(readSequence(">THREE", b))
	return &KNucleotideEnv{B: dna}
}

func TimeKNucleotide(n int) time.Duration {
	env := NewKNucleotideEnv(n)
	start := time.Now()
	protocol.KNucleotide(env)
	elapsed := time.Since(start)
	return elapsed
}

func TimeKNucleotideBase(n int) time.Duration {
	b := readFile(knucleotideFiles[n])
	dna := toBits(readSequence(">THREE", b))
	start := time.Now()
	knucleotide_base.KNucleotide(dna)
	elapsed := time.Since(start)
	return elapsed
}

func toBits(seq []byte) []byte {
	for i := 0; i < len(seq); i++ {
		// 'A' => 0, 'C' => 1, 'T' => 2, 'G' => 3
		seq[i] = seq[i] >> 1 & 3
	}
	return seq
}

func readSequence(prefix string, input []byte) (data []byte) {
	in, lineCount := findSequence(prefix, input)
	data = make([]byte, 0, lineCount*61)
	for {
		line, err := in.ReadSlice('\n')
		if len(line) <= 1 || line[0] == '>' {
			break
		}

		last := len(line) - 1
		if line[last] == '\n' {
			line = line[0:last]
		}
		data = append(data, line...)

		if err != nil {
			break
		}
	}
	return
}

func findSequence(prefix string, input []byte) (in *bufio.Reader, lineCount int) {
	pfx := []byte(prefix)
	in = bufio.NewReader(bytes.NewReader(input))
	for {
		line, err := in.ReadSlice('\n')
		if err != nil {
			panic("read error")
		}
		lineCount++
		if line[0] == '>' && bytes.HasPrefix(line, pfx) {
			break
		}
	}
	return
}

func KNucleotideBenchmark(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(kNucleotideParams, repetitions, TimeKNucleotide)
	base_results := benchmark.TimeImpl(kNucleotideParams, repetitions, TimeKNucleotideBase)
	return scribble_results, base_results
}
