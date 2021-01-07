package callbacks

import (
	"NestedScribbleBenchmark/regexredux/results/regexredux"
	"NestedScribbleBenchmark/regexredux/results/regexredux2"
)

type RegexRedux_Master_Env interface {
	Done() regexredux.Master_Result
	ResultFrom_RegexRedux2_M(result regexredux2.M_Result)
	To_RegexRedux2_M_Env() RegexRedux2_M_Env
	RegexRedux2_Setup()
}

type RegexReduxMasterState struct {
	B    []byte
	Idx  int
	CLen int
	ILen int
	Len  int
}

func (r *RegexReduxMasterState) Done() regexredux.Master_Result {
	return regexredux.Master_Result{}
}

func (r *RegexReduxMasterState) ResultFrom_RegexRedux2_M(result regexredux2.M_Result) {
	r.Len = result.Length
	// fmt.Printf("\n%d\n%d\n%d\n", r.ILen, r.CLen, r.Len)
}

func (r *RegexReduxMasterState) To_RegexRedux2_M_Env() RegexRedux2_M_Env {
	return &RegexRedux2MState{
		B:          r.B,
		VariantIdx: r.Idx,
		CLen:       r.CLen,
		ILen:       r.ILen,
	}
}

func (r *RegexReduxMasterState) RegexRedux2_Setup() {
}
