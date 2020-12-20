package callbacks

import "NestedScribbleBenchmark/fannkuch/results/fannkuchrecursive"

type FannkuchRecursive_Worker_Env interface {
	Done() fannkuchrecursive.Worker_Result
	Task_To_NewWorker() (int, int, int)
}

type FannkuchRecursiveWorkerState struct {
	IdxMin  int
	N       int
	Chunksz int
}

func (f *FannkuchRecursiveWorkerState) Done() fannkuchrecursive.Worker_Result {
	return fannkuchrecursive.Worker_Result{}
}

func (f *FannkuchRecursiveWorkerState) Task_To_NewWorker() (int, int, int) {
	return f.IdxMin, f.Chunksz, f.N
}
