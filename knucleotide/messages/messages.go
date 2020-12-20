package messages

type KNucleotide_Label int

const (
	Finish KNucleotide_Label = iota
	FrequencyJob
	FrequencyResult
	ScheduleJobs_M
	ScheduleJobs_Master
	SequenceJob
	SequenceResult
)
