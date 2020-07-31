package callbacks

import (
	"ScribbleBenchmark/knucleotide/messages/knucleotide"
)
import "ScribbleBenchmark/knucleotide/results/schedulejobs"
import knucleotide_2 "ScribbleBenchmark/knucleotide/results/knucleotide"

type KNucleotide_Master_Env interface {
	Done() knucleotide_2.Master_Result
	SequenceResult_From_Worker(sequenceresult_msg knucleotide.SequenceResult)
	ResultFrom_ScheduleJobs_M(result schedulejobs.M_Result)
	To_ScheduleJobs_M_Env() ScheduleJobs_M_Env
	ScheduleJobs_Setup()
	SequenceJob_To_Worker() knucleotide.SequenceJob
}

type TaskType int

const (
	LenTask TaskType = iota
	SeqTask
)

var Lengths = []int{1, 2}

var Patterns = []string{
	"GGT",
	"GGTA",
	"GGTATT",
	"GGTATTTTAATT",
	"GGTATTTTAATTTATAGT",
}

type KNucleotideMasterState struct {
	Dna        []byte
	PatternIdx int
	LenIdx     int
}

func (k *KNucleotideMasterState) Done() knucleotide_2.Master_Result {
	return knucleotide_2.Master_Result{}
}

func (k *KNucleotideMasterState) SequenceResult_From_Worker(sequenceresult_msg knucleotide.SequenceResult) {
	// TODO: Uncomment
	// fmt.Println(sequenceresult_msg.Res)
}

func (k *KNucleotideMasterState) ResultFrom_ScheduleJobs_M(result schedulejobs.M_Result) {
}

func (k *KNucleotideMasterState) To_ScheduleJobs_M_Env() ScheduleJobs_M_Env {
	return &ScheduleJobsMState{
		Dna:        k.Dna,
		PatternIdx: k.PatternIdx - 1,
		LengthIdx:  k.LenIdx,
		TaskType:   nextTaskType(k.PatternIdx),
	}
}

func (k *KNucleotideMasterState) ScheduleJobs_Setup() {
}

func (k *KNucleotideMasterState) SequenceJob_To_Worker() knucleotide.SequenceJob {
	return knucleotide.SequenceJob{
		Dna:      k.Dna,
		Sequence: Patterns[k.PatternIdx],
	}
}

func nextTaskType(patternIdx int) TaskType {
	if patternIdx == 0 {
		return LenTask
	}
	return SeqTask
}

func nextTaskIndices(patternIdx, lenIdx int, currTask, nextTask TaskType) (int, int) {
	if nextTask == SeqTask {
		return patternIdx - 1, lenIdx
	}
	if currTask == SeqTask && nextTask == LenTask {
		return patternIdx, lenIdx
	}
	return patternIdx, lenIdx - 1
}
