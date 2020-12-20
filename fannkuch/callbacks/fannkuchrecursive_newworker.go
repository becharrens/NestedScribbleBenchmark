package callbacks

import "NestedScribbleBenchmark/fannkuch/results/fannkuchrecursive"

type FannkuchRecursive_NewWorker_Choice int

const (
	FannkuchRecursive_NewWorker_FannkuchRecursive FannkuchRecursive_NewWorker_Choice = iota
	FannkuchRecursive_NewWorker_Result
)

type FannkuchRecursive_NewWorker_Env interface {
	Result_To_Source_2() (int, int)
	Done()
	Result_To_Source() (int, int)
	ResultFrom_FannkuchRecursive_Worker(result fannkuchrecursive.Worker_Result)
	To_FannkuchRecursive_Worker_Env() FannkuchRecursive_Worker_Env
	FannkuchRecursive_Setup()
	NewWorker_Choice() FannkuchRecursive_NewWorker_Choice
	Task_From_Worker(idxMin int, chunksz int, n int)
}

type FannkuchRecursiveNewWorkerState struct {
	IdxMin  int
	IdxMax  int
	N       int
	Chunksz int
}

func (f *FannkuchRecursiveNewWorkerState) Result_To_Source_2() (int, int) {
	maxFlips, checksum := fannkuchImpl(f.IdxMin, f.IdxMax, f.N)
	return maxFlips, checksum
}

func (f *FannkuchRecursiveNewWorkerState) Done() {
}

func (f *FannkuchRecursiveNewWorkerState) Result_To_Source() (int, int) {
	maxFlips, checksum := fannkuchImpl(f.IdxMin, f.IdxMax, f.N)
	return maxFlips, checksum
}

func (f *FannkuchRecursiveNewWorkerState) ResultFrom_FannkuchRecursive_Worker(result fannkuchrecursive.Worker_Result) {
}

func (f *FannkuchRecursiveNewWorkerState) To_FannkuchRecursive_Worker_Env() FannkuchRecursive_Worker_Env {
	return &FannkuchRecursiveWorkerState{
		IdxMin:  f.IdxMax,
		N:       f.N,
		Chunksz: f.Chunksz,
	}
}

func (f *FannkuchRecursiveNewWorkerState) FannkuchRecursive_Setup() {
}

func (f *FannkuchRecursiveNewWorkerState) NewWorker_Choice() FannkuchRecursive_NewWorker_Choice {
	if f.IdxMax < Fact[f.N] {
		return FannkuchRecursive_NewWorker_FannkuchRecursive
	}
	f.IdxMax = Fact[f.N]
	return FannkuchRecursive_NewWorker_Result
}

func (f *FannkuchRecursiveNewWorkerState) Task_From_Worker(idxMin int, chunksz int, n int) {
	f.IdxMin = idxMin
	f.IdxMax = idxMin + chunksz
	f.N = n
	f.Chunksz = chunksz
}

func New_FannkuchRecursive_NewWorker_State() FannkuchRecursive_NewWorker_Env {
	return &FannkuchRecursiveNewWorkerState{}
}
