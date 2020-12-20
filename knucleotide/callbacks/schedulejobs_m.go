package callbacks

import "NestedScribbleBenchmark/knucleotide/results/schedulejobs"

type ScheduleJobs_M_Choice int

const (
	ScheduleJobs_M_SequenceJob ScheduleJobs_M_Choice = iota
	ScheduleJobs_M_FrequencyJob
	ScheduleJobs_M_Finish
)

type ScheduleJobs_M_Env interface {
	Finish_To_W()
	FrequencyResult_From_W(res string)
	ResultFrom_ScheduleJobs_M_2(result schedulejobs.M_Result)
	To_ScheduleJobs_M_Env_2() ScheduleJobs_M_Env
	ScheduleJobs_Setup_2()
	FrequencyJob_To_W() (int, []byte)
	Done() schedulejobs.M_Result
	SequenceResult_From_W(res string)
	ResultFrom_ScheduleJobs_M(result schedulejobs.M_Result)
	To_ScheduleJobs_M_Env() ScheduleJobs_M_Env
	ScheduleJobs_Setup()
	SequenceJob_To_W() (string, []byte)
	M_Choice() ScheduleJobs_M_Choice
}

type ScheduleJobsMState struct {
	Dna        []byte
	PatternIdx int
	LengthIdx  int
	TaskType   TaskType
}

func (s *ScheduleJobsMState) Finish_To_W() {
}

func (s *ScheduleJobsMState) FrequencyResult_From_W(res string) {
	// TODO: Uncomment
	// fmt.Println(res)
}

func (s *ScheduleJobsMState) ResultFrom_ScheduleJobs_M_2(result schedulejobs.M_Result) {
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

func (s *ScheduleJobsMState) FrequencyJob_To_W() (int, []byte) {
	return Lengths[s.LengthIdx], s.Dna
}

func (s *ScheduleJobsMState) Done() schedulejobs.M_Result {
	return schedulejobs.M_Result{}
}

func (s *ScheduleJobsMState) SequenceResult_From_W(res string) {
	// TODO: Uncomment
	// fmt.Println(res)
}

func (s *ScheduleJobsMState) ResultFrom_ScheduleJobs_M(result schedulejobs.M_Result) {
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

func (s *ScheduleJobsMState) SequenceJob_To_W() (string, []byte) {
	return Patterns[s.PatternIdx], s.Dna
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
