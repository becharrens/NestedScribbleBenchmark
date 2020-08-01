package callbacks

import (
	"NestedScribbleBenchmark/fannkuch/messages/fannkuchrecursive"
)
import fannkuchrecursive_2 "NestedScribbleBenchmark/fannkuch/results/fannkuchrecursive"

type FannkuchRecursive_NewWorker_Choice int

const (
	FannkuchRecursive_NewWorker_FannkuchRecursive FannkuchRecursive_NewWorker_Choice = iota
	FannkuchRecursive_NewWorker_Result
)

type FannkuchRecursive_NewWorker_Env interface {
	Result_To_Source_2() fannkuchrecursive.Result
	Done()
	Result_To_Source() fannkuchrecursive.Result
	ResultFrom_FannkuchRecursive_Worker(result fannkuchrecursive_2.Worker_Result)
	To_FannkuchRecursive_Worker_Env() FannkuchRecursive_Worker_Env
	FannkuchRecursive_Setup()
	NewWorker_Choice() FannkuchRecursive_NewWorker_Choice
	Task_From_Worker(task fannkuchrecursive.Task)
}

type FannkuchRecursiveNewWorkerState struct {
	Fact    []int
	IdxMin  int
	IdxMax  int
	N       int
	Chunksz int
}

func (f *FannkuchRecursiveNewWorkerState) Result_To_Source_2() fannkuchrecursive.Result {
	maxFlips, checksum := fannkuchImpl(f.IdxMin, f.IdxMax, f.N, f.Fact)
	return fannkuchrecursive.Result{
		Checksum: checksum,
		MaxFlips: maxFlips,
	}
}

func (f *FannkuchRecursiveNewWorkerState) Done() {
}

func (f *FannkuchRecursiveNewWorkerState) Result_To_Source() fannkuchrecursive.Result {
	maxFlips, checksum := fannkuchImpl(f.IdxMin, f.IdxMax, f.N, f.Fact)
	return fannkuchrecursive.Result{
		Checksum: checksum,
		MaxFlips: maxFlips,
	}
}

func (f *FannkuchRecursiveNewWorkerState) ResultFrom_FannkuchRecursive_Worker(result fannkuchrecursive_2.Worker_Result) {
}

func (f *FannkuchRecursiveNewWorkerState) To_FannkuchRecursive_Worker_Env() FannkuchRecursive_Worker_Env {
	return &FannkuchRecursiveWorkerState{
		Fact:    f.Fact,
		IdxMin:  f.IdxMax,
		N:       f.N,
		Chunksz: f.Chunksz,
	}
}

func (f *FannkuchRecursiveNewWorkerState) FannkuchRecursive_Setup() {
}

func (f *FannkuchRecursiveNewWorkerState) NewWorker_Choice() FannkuchRecursive_NewWorker_Choice {
	if f.IdxMax < f.Fact[f.N] {
		return FannkuchRecursive_NewWorker_FannkuchRecursive
	}
	f.IdxMax = f.Fact[f.N]
	return FannkuchRecursive_NewWorker_Result
}

func (f *FannkuchRecursiveNewWorkerState) Task_From_Worker(task fannkuchrecursive.Task) {
	f.Fact = task.Fact
	f.IdxMin = task.IdxMin
	f.IdxMax = task.IdxMin + task.Chunksz
	f.N = task.N
	f.Chunksz = task.Chunksz
}

func New_FannkuchRecursive_NewWorker_State() FannkuchRecursive_NewWorker_Env {
	return &FannkuchRecursiveNewWorkerState{}
}
