package callbacks

import (
	"ScribbleBenchmark/regexredux/messages/regexredux"
	"ScribbleBenchmark/regexredux/results/regexredux2"
	"fmt"

	regexredux_2 "ScribbleBenchmark/regexredux/results/regexredux"
)

type RegexRedux_Master_Env interface {
	Done() regexredux_2.Master_Result
	NumMatches_From_Worker(nummatches_msg regexredux.NumMatches)
	ResultFrom_RegexRedux2_M(result regexredux2.M_Result)
	To_RegexRedux2_M_Env() RegexRedux2_M_Env
	RegexRedux2_Setup()
	Task_To_Worker() regexredux.Task
}

var variants = []string{
	"agggtaa[cgt]|[acg]ttaccct",
	"agggta[cgt]a|t[acg]taccct",
	"agggt[cgt]aa|tt[acg]accct",
	"aggg[acg]aaa|ttt[cgt]ccct",
	"agg[act]taaa|ttta[agt]cct",
	"ag[act]gtaaa|tttac[agt]ct",
	"a[act]ggtaaa|tttacc[agt]t",
	"[cgt]gggtaaa|tttaccc[acg]",
	"agggtaaa|tttaccct",
}

type RegexReduxMasterState struct {
	B    []byte
	Idx  int
	CLen int
	ILen int
	Len  int
}

func (r *RegexReduxMasterState) Done() regexredux_2.Master_Result {
	return regexredux_2.Master_Result{}
}

func (r *RegexReduxMasterState) NumMatches_From_Worker(nummatches_msg regexredux.NumMatches) {
	fmt.Printf("%s %d\n", variants[r.Idx], nummatches_msg.Nmatches)
	fmt.Printf("\n%d\n%d\n%d\n", r.ILen, r.CLen, r.Len)
}

func (r *RegexReduxMasterState) ResultFrom_RegexRedux2_M(result regexredux2.M_Result) {
	r.Len = result.Length
}

func (r *RegexReduxMasterState) To_RegexRedux2_M_Env() RegexRedux2_M_Env {
	return &RegexRedux2MState{
		B:          r.B,
		VariantIdx: r.Idx + 1,
		CLen:       r.CLen,
		ILen:       r.ILen,
	}
}

func (r *RegexReduxMasterState) RegexRedux2_Setup() {
}

func (r *RegexReduxMasterState) Task_To_Worker() regexredux.Task {
	return regexredux.Task{
		B:       r.B,
		Pattern: variants[r.Idx],
	}
}
