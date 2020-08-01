package schedulejobs

import "NestedScribbleBenchmark/knucleotide/messages/schedulejobs"

type M_Chan struct {
	W_Finish          chan schedulejobs.Finish
	W_FrequencyJob    chan schedulejobs.FrequencyJob
	W_FrequencyResult chan schedulejobs.FrequencyResult
	W_SequenceJob     chan schedulejobs.SequenceJob
	W_SequenceResult  chan schedulejobs.SequenceResult
}

type W_Chan struct {
	M_Finish          chan schedulejobs.Finish
	M_FrequencyJob    chan schedulejobs.FrequencyJob
	M_FrequencyResult chan schedulejobs.FrequencyResult
	M_SequenceJob     chan schedulejobs.SequenceJob
	M_SequenceResult  chan schedulejobs.SequenceResult
}
