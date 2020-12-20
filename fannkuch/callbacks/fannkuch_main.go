package callbacks

import "NestedScribbleBenchmark/fannkuch/results/fannkuchrecursive"
import "NestedScribbleBenchmark/fannkuch/results/fannkuch"

var Fact []int

type Fannkuch_Main_Env interface {
	Result_From_Worker_2(maxflips int, checksum int)
	Done() fannkuch.Main_Result
	Result_From_Worker(maxflips int, checksum int)
	ResultFrom_FannkuchRecursive_Source(result fannkuchrecursive.Source_Result)
	To_FannkuchRecursive_Source_Env() FannkuchRecursive_Source_Env
	Task_To_Worker() (int, int, int)
}

type FannkuchMainState struct {
	N         int
	ChunkSize int
	Res       int
	Chk       int
}

func (f *FannkuchMainState) Result_From_Worker_2(maxflips int, checksum int) {
	res, chk := aggregate_result(f.Res, f.Chk, maxflips, checksum)
	f.Res = res
	f.Chk = chk
}

func (f *FannkuchMainState) Done() fannkuch.Main_Result {
	return fannkuch.Main_Result{
		Res: f.Res,
		Chk: f.Chk,
	}
}

func (f *FannkuchMainState) Result_From_Worker(maxflips int, checksum int) {
	res, chk := aggregate_result(f.Res, f.Chk, maxflips, checksum)
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

func (f *FannkuchMainState) Task_To_Worker() (int, int, int) {
	return 0, f.ChunkSize, f.N
}

func aggregate_result(currRes, currChk, maxFlips, checksum int) (int, int) {
	chk := currChk + checksum
	if currRes < maxFlips {
		return maxFlips, chk
	}
	return currRes, chk
}
