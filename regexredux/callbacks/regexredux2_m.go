package callbacks

import (
	"NestedScribbleBenchmark/regexredux/messages/regexredux2"
)
import regexredux2_2 "NestedScribbleBenchmark/regexredux/results/regexredux2"

type RegexRedux2_M_Choice int

const (
	RegexRedux2_M_Task RegexRedux2_M_Choice = iota
	RegexRedux2_M_CalcLength
)

type RegexRedux2_M_Env interface {
	Length_From_W(length_msg regexredux2.Length)
	CalcLength_To_W() regexredux2.CalcLength
	Done() regexredux2_2.M_Result
	NumMatches_From_W(nummatches_msg regexredux2.NumMatches)
	ResultFrom_RegexRedux2_M(result regexredux2_2.M_Result)
	To_RegexRedux2_M_Env() RegexRedux2_M_Env
	RegexRedux2_Setup()
	Task_To_W() regexredux2.Task
	M_Choice() RegexRedux2_M_Choice
}

type RegexRedux2MState struct {
	B          []byte
	VariantIdx int
	CLen       int
	ILen       int
	ResLength  int
}

func (r *RegexRedux2MState) Length_From_W(length_msg regexredux2.Length) {
	r.ResLength = length_msg.Len
}

func (r *RegexRedux2MState) CalcLength_To_W() regexredux2.CalcLength {
	return regexredux2.CalcLength{B: r.B}
}

func (r *RegexRedux2MState) Done() regexredux2_2.M_Result {
	return regexredux2_2.M_Result{Length: r.ResLength}
}

func (r *RegexRedux2MState) NumMatches_From_W(nummatches_msg regexredux2.NumMatches) {
	// TODO: Uncomment
	// fmt.Printf("%s %d\n", variants[r.VariantIdx], nummatches_msg.Nmatches)
}

func (r *RegexRedux2MState) ResultFrom_RegexRedux2_M(result regexredux2_2.M_Result) {
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

func (r *RegexRedux2MState) Task_To_W() regexredux2.Task {
	return regexredux2.Task{
		B:       r.B,
		Pattern: variants[r.VariantIdx],
	}
}

func (r *RegexRedux2MState) M_Choice() RegexRedux2_M_Choice {
	if r.VariantIdx < len(variants) {
		return RegexRedux2_M_Task
	}
	return RegexRedux2_M_CalcLength
}
