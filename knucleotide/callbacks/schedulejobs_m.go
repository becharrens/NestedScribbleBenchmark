package callbacks

import (
	"ScribbleBenchmark/knucleotide/messages/schedulejobs"
)
import schedulejobs_2 "ScribbleBenchmark/knucleotide/results/schedulejobs"

type ScheduleJobs_M_Choice int

const (
	ScheduleJobs_M_SequenceJob ScheduleJobs_M_Choice = iota
	ScheduleJobs_M_FrequencyJob
	ScheduleJobs_M_Finish
)

type ScheduleJobs_M_Env interface {
	Finish_To_W() schedulejobs.Finish
	FrequencyResult_From_W(frequencyresult_msg schedulejobs.FrequencyResult)
	ResultFrom_ScheduleJobs_M_2(result schedulejobs_2.M_Result)
	To_ScheduleJobs_M_Env_2() ScheduleJobs_M_Env
	ScheduleJobs_Setup_2()
	FrequencyJob_To_W() schedulejobs.FrequencyJob
	Done() schedulejobs_2.M_Result
	SequenceResult_From_W(sequenceresult_msg schedulejobs.SequenceResult)
	ResultFrom_ScheduleJobs_M(result schedulejobs_2.M_Result)
	To_ScheduleJobs_M_Env() ScheduleJobs_M_Env
	ScheduleJobs_Setup()
	SequenceJob_To_W() schedulejobs.SequenceJob
	M_Choice() ScheduleJobs_M_Choice
}

type ScheduleJobsMState struct {
	Dna        []byte
	PatternIdx int
	LengthIdx  int
	TaskType   TaskType
}

func (s *ScheduleJobsMState) Finish_To_W() schedulejobs.Finish {
	return schedulejobs.Finish{}
}

func (s *ScheduleJobsMState) FrequencyResult_From_W(frequencyresult_msg schedulejobs.FrequencyResult) {
	// TODO: Uncomment
	// fmt.Println(frequencyresult_msg.Res)
}

func (s *ScheduleJobsMState) ResultFrom_ScheduleJobs_M_2(result schedulejobs_2.M_Result) {
}

func (s *ScheduleJobsMState) To_ScheduleJobs_M_Env_2() ScheduleJobs_M_Env {
	jobType := nextTaskType(s.PatternIdx)
	patternIdx, lenIdx := nextTaskIndices(s.PatternIdx, s.LengthIdx, s.TaskType, jobType)
	return &ScheduleJobsMState{
		Dna:        s.Dna,
		PatternIdx: patternIdx,
		LengthIdx:  lenIdx,
		TaskType:   jobType,
	}
}

func (s *ScheduleJobsMState) ScheduleJobs_Setup_2() {
}

func (s *ScheduleJobsMState) FrequencyJob_To_W() schedulejobs.FrequencyJob {
	return schedulejobs.FrequencyJob{
		Dna: s.Dna,
		Len: Lengths[s.LengthIdx],
	}
}

func (s *ScheduleJobsMState) Done() schedulejobs_2.M_Result {
	return schedulejobs_2.M_Result{}
}

func (s *ScheduleJobsMState) SequenceResult_From_W(sequenceresult_msg schedulejobs.SequenceResult) {
	// TODO: Uncomment
	// fmt.Println(sequenceresult_msg.Res)
}

func (s *ScheduleJobsMState) ResultFrom_ScheduleJobs_M(result schedulejobs_2.M_Result) {
}

func (s *ScheduleJobsMState) To_ScheduleJobs_M_Env() ScheduleJobs_M_Env {
	jobType := nextTaskType(s.PatternIdx)
	patternIdx, lenIdx := nextTaskIndices(s.PatternIdx, s.LengthIdx, s.TaskType, jobType)
	return &ScheduleJobsMState{
		Dna:        s.Dna,
		PatternIdx: patternIdx,
		LengthIdx:  lenIdx,
		TaskType:   jobType,
	}
}

func (s *ScheduleJobsMState) ScheduleJobs_Setup() {
}

func (s *ScheduleJobsMState) SequenceJob_To_W() schedulejobs.SequenceJob {
	return schedulejobs.SequenceJob{
		Dna:      s.Dna,
		Sequence: Patterns[s.PatternIdx],
	}
}

func (s *ScheduleJobsMState) M_Choice() ScheduleJobs_M_Choice {
	if s.TaskType == SeqTask {
		return ScheduleJobs_M_SequenceJob
	}
	if s.LengthIdx >= 0 {
		return ScheduleJobs_M_FrequencyJob
	}
	return ScheduleJobs_M_Finish
}
