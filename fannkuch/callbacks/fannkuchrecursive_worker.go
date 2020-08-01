package callbacks

import "NestedScribbleBenchmark/fannkuch/messages/fannkuchrecursive"
import fannkuchrecursive_2 "NestedScribbleBenchmark/fannkuch/results/fannkuchrecursive"

type FannkuchRecursive_Worker_Env interface {
	Done() fannkuchrecursive_2.Worker_Result
	Task_To_NewWorker() fannkuchrecursive.Task
}

type FannkuchRecursiveWorkerState struct {
	Fact    []int
	IdxMin  int
	N       int
	Chunksz int
}

func (f *FannkuchRecursiveWorkerState) Done() fannkuchrecursive_2.Worker_Result {
	return fannkuchrecursive_2.Worker_Result{}
}

func (f *FannkuchRecursiveWorkerState) Task_To_NewWorker() fannkuchrecursive.Task {
	return fannkuchrecursive.Task{
		Chunksz: f.Chunksz,
		Fact:    f.Fact,
		IdxMin:  f.IdxMin,
		N:       f.N,
	}
}
