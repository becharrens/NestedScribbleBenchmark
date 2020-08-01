package callbacks

import "NestedScribbleBenchmark/knucleotide/messages/schedulejobs"

type ScheduleJobs_W_Env interface {
	Finish_From_M(finish_msg schedulejobs.Finish)
	FrequencyResult_To_M() schedulejobs.FrequencyResult
	FrequencyJob_From_M(frequencyjob_msg schedulejobs.FrequencyJob)
	Done()
	SequenceResult_To_M() schedulejobs.SequenceResult
	SequenceJob_From_M(sequencejob_msg schedulejobs.SequenceJob)
}

type ScheduleJobsWState struct {
	SeqReport  string
	FreqReport string
}

func (s *ScheduleJobsWState) Finish_From_M(finish_msg schedulejobs.Finish) {
}

func (s *ScheduleJobsWState) FrequencyResult_To_M() schedulejobs.FrequencyResult {
	return schedulejobs.FrequencyResult{Res: s.FreqReport}
}

func (s *ScheduleJobsWState) FrequencyJob_From_M(frequencyjob_msg schedulejobs.FrequencyJob) {
	s.FreqReport = frequencyReport(frequencyjob_msg.Dna, frequencyjob_msg.Len)
}

func (s *ScheduleJobsWState) Done() {
}

func (s *ScheduleJobsWState) SequenceResult_To_M() schedulejobs.SequenceResult {
	return schedulejobs.SequenceResult{Res: s.SeqReport}
}

func (s *ScheduleJobsWState) SequenceJob_From_M(sequencejob_msg schedulejobs.SequenceJob) {
	s.SeqReport = sequenceReport(sequencejob_msg.Dna, seqString(sequencejob_msg.Sequence))
}

func New_ScheduleJobs_W_State() ScheduleJobs_W_Env {
	return &ScheduleJobsWState{}
}
