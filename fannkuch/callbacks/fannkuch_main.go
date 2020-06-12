package callbacks

import "ScribbleBenchmark/fannkuch/messages/fannkuch"
import "ScribbleBenchmark/fannkuch/results/fannkuchrecursive"
import fannkuch_2 "ScribbleBenchmark/fannkuch/results/fannkuch"

type Fannkuch_Main_Env interface {
	Result_From_Worker_2(result fannkuch.Result)
	Done() fannkuch_2.Main_Result
	Result_From_Worker(result fannkuch.Result)
	ResultFrom_FannkuchRecursive_Source(result fannkuchrecursive.Source_Result)
	To_FannkuchRecursive_Source_Env() FannkuchRecursive_Source_Env
	Task_To_Worker() fannkuch.Task
}

type FannkuchMainState struct {
	Fact      []int
	N         int
	ChunkSize int
	Res       int
	Chk       int
}

func (f *FannkuchMainState) Result_From_Worker_2(result fannkuch.Result) {
	res, chk := aggregate_result(f.Res, f.Chk, result.MaxFlips, result.Checksum)
	f.Res = res
	f.Chk = chk
}

func (f *FannkuchMainState) Done() fannkuch_2.Main_Result {
	return fannkuch_2.Main_Result{
		Res: f.Res,
		Chk: f.Chk,
	}
}

func (f *FannkuchMainState) Result_From_Worker(result fannkuch.Result) {
	res, chk := aggregate_result(f.Res, f.Chk, result.MaxFlips, result.Checksum)
	f.Res = res
	f.Chk = chk
}

func (f *FannkuchMainState) ResultFrom_FannkuchRecursive_Source(result fannkuchrecursive.Source_Result) {
	f.Chk = result.Chk
	f.Res = result.Res
}

func (f *FannkuchMainState) To_FannkuchRecursive_Source_Env() FannkuchRecursive_Source_Env {
	return &FannkuchRecursiveSourceState{
		Res: f.Res,
		Chk: f.Chk,
	}
}

func (f *FannkuchMainState) Task_To_Worker() fannkuch.Task {
	return fannkuch.Task{
		Chunksz: f.ChunkSize,
		Fact:    f.Fact,
		IdxMin:  0,
		N:       f.N,
	}
}

func aggregate_result(currRes, currChk, maxFlips, checksum int) (int, int) {
	chk := currChk + checksum
	if currRes < maxFlips {
		return maxFlips, chk
	}
	return currRes, chk
}
