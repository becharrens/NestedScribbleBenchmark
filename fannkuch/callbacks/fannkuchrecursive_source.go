package callbacks

import "NestedScribbleBenchmark/fannkuch/results/fannkuchrecursive"

type FannkuchRecursive_Source_Env interface {
	Result_From_NewWorker_2(maxflips int, checksum int)
	Done() fannkuchrecursive.Source_Result
	Result_From_NewWorker(maxFlips int, checksum int)
	ResultFrom_FannkuchRecursive_Source(result fannkuchrecursive.Source_Result)
	To_FannkuchRecursive_Source_Env() FannkuchRecursive_Source_Env
}

type FannkuchRecursiveSourceState struct {
	Res int
	Chk int
}

func (f *FannkuchRecursiveSourceState) Result_From_NewWorker_2(maxflips int, checksum int) {
	res, chk := aggregate_result(f.Res, f.Chk, maxflips, checksum)
	f.Res = res
	f.Chk = chk
}

func (f *FannkuchRecursiveSourceState) Done() fannkuchrecursive.Source_Result {
	return fannkuchrecursive.Source_Result{
		Res: f.Res,
		Chk: f.Chk,
	}
}

func (f *FannkuchRecursiveSourceState) Result_From_NewWorker(maxflips int, checksum int) {
	res, chk := aggregate_result(f.Res, f.Chk, maxflips, checksum)
	f.Res = res
	f.Chk = chk
}

func (f *FannkuchRecursiveSourceState) ResultFrom_FannkuchRecursive_Source(result fannkuchrecursive.Source_Result) {
	f.Res = result.Res
	f.Chk = result.Chk
}

func (f *FannkuchRecursiveSourceState) To_FannkuchRecursive_Source_Env() FannkuchRecursive_Source_Env {
	return &FannkuchRecursiveSourceState{
		Res: f.Res,
		Chk: f.Chk,
	}
}
