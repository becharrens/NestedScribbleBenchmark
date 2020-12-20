package callbacks

import "NestedScribbleBenchmark/knucleotide/results/schedulejobs"
import "NestedScribbleBenchmark/knucleotide/results/knucleotide"

type KNucleotide_Master_Env interface {
	Done() knucleotide.Master_Result
	ResultFrom_ScheduleJobs_M(result schedulejobs.M_Result)
	To_ScheduleJobs_M_Env() ScheduleJobs_M_Env
	ScheduleJobs_Setup()
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

func (k *KNucleotideMasterState) Done() knucleotide.Master_Result {
	return knucleotide.Master_Result{}
}

func (k *KNucleotideMasterState) ResultFrom_ScheduleJobs_M(result schedulejobs.M_Result) {
}

func (k *KNucleotideMasterState) To_ScheduleJobs_M_Env() ScheduleJobs_M_Env {
	return &ScheduleJobsMState{
		Dna:        k.Dna,
		PatternIdx: k.PatternIdx,
		LengthIdx:  k.LenIdx,
		TaskType:   SeqTask,
	}
}

func (k *KNucleotideMasterState) ScheduleJobs_Setup() {
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
