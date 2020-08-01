package callbacks

import (
	"NestedScribbleBenchmark/regexredux/messages/regexredux2"

	"github.com/GRbit/go-pcre"
)

type RegexRedux2_W_Env interface {
	Length_To_M() regexredux2.Length
	CalcLength_From_M(calclength_msg regexredux2.CalcLength)
	Done()
	NumMatches_To_M() regexredux2.NumMatches
	Task_From_M(task_msg regexredux2.Task)
}

type substitution struct {
	pattern     string
	replacement string
}

var substs = []substitution{
	{"tHa[Nt]", "<4>"},
	{"aND|caN|Ha[DS]|WaS", "<3>"},
	{"a[NSt]|BY", "<2>"},
	{"<[^>]*>", "|"},
	{"\\|[^|][^|]*\\|", "-"},
}

type RegexRedux2WState struct {
	Nmatches int
	Length   int
}

func (r *RegexRedux2WState) Length_To_M() regexredux2.Length {
	return regexredux2.Length{Len: r.Length}
}

func (r *RegexRedux2WState) CalcLength_From_M(calclength_msg regexredux2.CalcLength) {
	b := calclength_msg.B
	for i := 0; i < len(substs); i++ {
		b = pcre.
			MustCompileJIT(substs[i].pattern, 0, pcre.STUDY_JIT_COMPILE).
			ReplaceAll(b, []byte(substs[i].replacement), 0)
	}
	r.Length = len(b)
}

func (r *RegexRedux2WState) Done() {
}

func (r *RegexRedux2WState) NumMatches_To_M() regexredux2.NumMatches {
	return regexredux2.NumMatches{Nmatches: r.Nmatches}
}

func (r *RegexRedux2WState) Task_From_M(task_msg regexredux2.Task) {
	r.Nmatches = countMatches(task_msg.Pattern, task_msg.B)
}

func New_RegexRedux2_W_State() RegexRedux2_W_Env {
	return &RegexRedux2WState{}
}
