package messages

type Fannkuch_Label int

const (
	FannkuchRecursive_Main_Worker Fannkuch_Label = iota
	FannkuchRecursive_Source_NewWorker
	Result
	Task
)
