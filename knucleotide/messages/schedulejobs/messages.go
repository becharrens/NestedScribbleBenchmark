package schedulejobs

type Finish struct {
}

type FrequencyJob struct {
	Dna []byte
	Len int
}

type FrequencyResult struct {
	Res string
}

type SequenceJob struct {
	Dna      []byte
	Sequence string
}

type SequenceResult struct {
	Res string
}
