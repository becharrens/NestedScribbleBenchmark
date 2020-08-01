package fannkuch

import (
	"NestedScribbleBenchmark/fannkuch/callbacks"
	"NestedScribbleBenchmark/fannkuch/results/fannkuch"
)

const NCHUNKS = 720

type FannkuchEnv struct {
	N         int
	ChunkSize int
	Fact      []int
	Res       int
	Checksum  int
}

func (f *FannkuchEnv) New_Main_Env() callbacks.Fannkuch_Main_Env {
	return &callbacks.FannkuchMainState{
		Fact:      f.Fact,
		N:         f.N,
		ChunkSize: f.ChunkSize,
		Res:       0,
		Chk:       0,
	}
}

func (f *FannkuchEnv) New_Worker_Env() callbacks.Fannkuch_Worker_Env {
	return &callbacks.FannkuchWorkerState{}
}

func (f *FannkuchEnv) Main_Result(result fannkuch.Main_Result) {
	f.Checksum = result.Chk
	f.Res = result.Res
}

func (f *FannkuchEnv) Worker_Result(result fannkuch.Worker_Result) {
}

func NewFannkuchEnv(n int) *FannkuchEnv {
	fact := make([]int, n+1)
	fact[0] = 1
	for i := 1; i < len(fact); i++ {
		fact[i] = fact[i-1] * i
	}

	chunksz := (fact[n] + NCHUNKS - 1) / NCHUNKS
	chunksz += chunksz % 2
	return &FannkuchEnv{
		N:         n,
		ChunkSize: chunksz,
		Fact:      fact,
		Res:       0,
		Checksum:  0,
	}
}
