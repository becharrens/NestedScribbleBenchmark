package callbacks

import (
	"NestedScribbleBenchmark/regexredux/results/regexredux2"
)

type RegexRedux2_M_Choice int

const (
	RegexRedux2_M_Task RegexRedux2_M_Choice = iota
	RegexRedux2_M_CalcLength
)

type RegexRedux2_M_Env interface {
	Length_From_W(len int)
	CalcLength_To_W() []byte
	Done() regexredux2.M_Result
	NumMatches_From_W(nmatches int)
	ResultFrom_RegexRedux2_M(result regexredux2.M_Result)
	To_RegexRedux2_M_Env() RegexRedux2_M_Env
	RegexRedux2_Setup()
	Task_To_W() (string, []byte)
	M_Choice() RegexRedux2_M_Choice
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

type RegexRedux2MState struct {
	B          []byte
	VariantIdx int
	CLen       int
	ILen       int
	ResLength  int
}

func (r *RegexRedux2MState) Length_From_W(len int) {
	r.ResLength = len
}

func (r *RegexRedux2MState) CalcLength_To_W() []byte {
	return r.B
}

func (r *RegexRedux2MState) Done() regexredux2.M_Result {
	return regexredux2.M_Result{Length: r.ResLength}
}

func (r *RegexRedux2MState) NumMatches_From_W(nmatches int) {
	// TODO: Uncomment
	// fmt.Printf("%s %d\n", variants[r.VariantIdx], nummatches_msg.Nmatches)
}

func (r *RegexRedux2MState) ResultFrom_RegexRedux2_M(result regexredux2.M_Result) {
	r.ResLength = result.Length
}

func (r *RegexRedux2MState) To_RegexRedux2_M_Env() RegexRedux2_M_Env {
	return &RegexRedux2MState{
		B:          r.B,
		VariantIdx: r.VariantIdx + 1,
		CLen:       r.CLen,
		ILen:       r.ILen,
	}
}

func (r *RegexRedux2MState) RegexRedux2_Setup() {
}

func (r *RegexRedux2MState) Task_To_W() (string, []byte) {
	return variants[r.VariantIdx], r.B
}

func (r *RegexRedux2MState) M_Choice() RegexRedux2_M_Choice {
	if r.VariantIdx < len(variants) {
		return RegexRedux2_M_Task
	}
	return RegexRedux2_M_CalcLength
}
