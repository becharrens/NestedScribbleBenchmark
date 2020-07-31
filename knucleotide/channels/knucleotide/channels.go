package knucleotide

import "ScribbleBenchmark/knucleotide/messages/knucleotide"

type Master_Chan struct {
	Worker_SequenceJob chan knucleotide.SequenceJob
	Worker_SequenceResult chan knucleotide.SequenceResult
}

type Worker_Chan struct {
	Master_SequenceJob chan knucleotide.SequenceJob
	Master_SequenceResult chan knucleotide.SequenceResult
}