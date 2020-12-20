package callbacks

import (
	"bytes"
	"fmt"
	"sort"
)

type ScheduleJobs_W_Env interface {
	Finish_From_M()
	FrequencyResult_To_M() string
	FrequencyJob_From_M(len int, dna []byte)
	Done()
	SequenceResult_To_M() string
	SequenceJob_From_M(sequence string, dna []byte)
}

type ScheduleJobsWState struct {
	SeqReport  string
	FreqReport string
}

func (s *ScheduleJobsWState) Finish_From_M() {
}

func (s *ScheduleJobsWState) FrequencyResult_To_M() string {
	return s.FreqReport
}

func (s *ScheduleJobsWState) FrequencyJob_From_M(len int, dna []byte) {
	s.FreqReport = frequencyReport(dna, len)
}

func (s *ScheduleJobsWState) Done() {
}

func (s *ScheduleJobsWState) SequenceResult_To_M() string {
	return s.SeqReport
}

func (s *ScheduleJobsWState) SequenceJob_From_M(sequence string, dna []byte) {
	s.SeqReport = sequenceReport(dna, seqString(sequence))
}

func New_ScheduleJobs_W_State() ScheduleJobs_W_Env {
	return &ScheduleJobsWState{}
}

// seqString is a sequence of nucleotides as a string: "ACGT..."
type seqString string

// seqChars is a sequence of nucleotides as chars: 'A', 'C', 'G', 'T'...
type seqChars []byte

// seqBits is a sequence of nucleotides as 2 low bits per byte: 0, 1, 3, 2...
type seqBits []byte

// toBits converts *in-place*
func (seq seqChars) toBits() seqBits {
	for i := 0; i < len(seq); i++ {
		// 'A' => 0, 'C' => 1, 'T' => 2, 'G' => 3
		seq[i] = seq[i] >> 1 & 3
	}
	return seqBits(seq)
}

func (seq seqString) seqBits() seqBits {
	return seqChars(seq).toBits()
}

// seq32 is a short (<= 16) sequence of nucleotides in a compact form
// length is not embedded
type seq32 uint32

// seq64 is a short (17..32) sequence of nucleotides in a compact form
// length is not embedded
type seq64 uint64

// seq32 converts a seqBits to a seq32
func (seq seqBits) seq32() seq32 {
	var num seq32
	for _, char := range seq {
		num = num<<2 | seq32(char)
	}
	return num
}

// seq64 converts a seqBits to a seq64
func (seq seqBits) seq64() seq64 {
	var num seq64
	for _, char := range seq {
		num = num<<2 | seq64(char)
	}
	return num
}

// seqString converts a seq32 to a human readable string
func (num seq32) seqString(length int) seqString {
	sequence := make(seqChars, length)
	for i := 0; i < length; i++ {
		sequence[length-i-1] = "ACTG"[num&3]
		num = num >> 2
	}
	return seqString(sequence)
}

type counter uint32

func (dna seqBits) count32(length int) map[seq32]*counter {
	counts := make(map[seq32]*counter)
	key := dna[0 : length-1].seq32()
	mask := seq32(1)<<uint(2*length) - 1
	for index := length - 1; index < len(dna); index++ {
		key = key<<2&mask | seq32(dna[index])
		pointer := counts[key]
		if pointer == nil {
			n := counter(1)
			counts[key] = &n
		} else {
			*pointer++
		}
	}
	return counts
}

func (dna seqBits) count64(length int) map[seq64]*counter {
	counts := make(map[seq64]*counter)
	key := dna[0 : length-1].seq64()
	mask := seq64(1)<<uint(2*length) - 1
	for index := length - 1; index < len(dna); index++ {
		key = key<<2&mask | seq64(dna[index])
		pointer := counts[key]
		if pointer == nil {
			n := counter(1)
			counts[key] = &n
		} else {
			*pointer++
		}
	}
	return counts
}

type seqCount struct {
	seq   seqString
	count counter
}

type seqCounts []seqCount

func (ss seqCounts) Len() int { return len(ss) }

func (ss seqCounts) Swap(i, j int) { ss[i], ss[j] = ss[j], ss[i] }

// Less order descending by count then seq
func (ss seqCounts) Less(i, j int) bool {
	if ss[i].count == ss[j].count {
		return ss[i].seq > ss[j].seq
	}
	return ss[i].count > ss[j].count
}

func frequencyReport(dna seqBits, length int) string {
	counts := dna.count32(length)
	sortedSeqs := make(seqCounts, 0, len(counts))
	for num, pointer := range counts {
		sortedSeqs = append(
			sortedSeqs,
			seqCount{num.seqString(length), *pointer},
		)
	}
	sort.Sort(sortedSeqs)

	var buf bytes.Buffer
	buf.Grow((8 + length) * len(sortedSeqs))
	var scale float32 = 100.0 / float32(len(dna)-length+1)
	for _, sequence := range sortedSeqs {
		buf.WriteString(fmt.Sprintf(
			"%v %.3f\n", sequence.seq,
			float32(sequence.count)*scale),
		)
	}
	return buf.String()
}

func sequenceReport(dna seqBits, sequence seqString) string {
	var pointer *counter
	seq := sequence.seqBits()
	if len(sequence) <= 16 {
		counts := dna.count32(len(sequence))
		pointer = counts[seq.seq32()]
	} else {
		counts := dna.count64(len(sequence))
		pointer = counts[seq.seq64()]
	}
	var sequenceCount counter
	if pointer != nil {
		sequenceCount = *pointer
	}
	return fmt.Sprintf("%v\t%v", sequenceCount, sequence)
}
