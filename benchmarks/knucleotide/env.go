package knucleotide

import (
	"NestedScribbleBenchmark/knucleotide/callbacks"
	"NestedScribbleBenchmark/knucleotide/results/knucleotide"
	"bufio"
	"bytes"
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

func (k *KNucleotideEnv) New_Worker_Env() callbacks.KNucleotide_Worker_Env {
	return &callbacks.KNucleotideWorkerState{}
}

func (k *KNucleotideEnv) Master_Result(result knucleotide.Master_Result) {
}

func (k *KNucleotideEnv) Worker_Result(result knucleotide.Worker_Result) {
}

func NewKNucleotideEnv(dna []byte) *KNucleotideEnv {
	return &KNucleotideEnv{B: dna}
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
