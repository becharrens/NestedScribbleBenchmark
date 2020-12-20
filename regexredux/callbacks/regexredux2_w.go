package callbacks

import (
	"github.com/GRbit/go-pcre"
)

type RegexRedux2_W_Env interface {
	Length_To_M() int
	CalcLength_From_M(b []byte)
	Done()
	NumMatches_To_M() int
	Task_From_M(pattern string, b []byte)
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

func (r *RegexRedux2WState) Length_To_M() int {
	return r.Length
}

func (r *RegexRedux2WState) CalcLength_From_M(b []byte) {
	for i := 0; i < len(substs); i++ {
		b = pcre.
			MustCompileJIT(substs[i].pattern, 0, pcre.STUDY_JIT_COMPILE).
			ReplaceAll(b, []byte(substs[i].replacement), 0)
	}
	r.Length = len(b)
}

func (r *RegexRedux2WState) Done() {
}

func (r *RegexRedux2WState) NumMatches_To_M() int {
	return r.Nmatches
}

func (r *RegexRedux2WState) Task_From_M(pattern string, b []byte) {
	r.Nmatches = countMatches(pattern, b)
}

func New_RegexRedux2_W_State() RegexRedux2_W_Env {
	return &RegexRedux2WState{}
}

func countMatches(pat string, b []byte) int {
	m := pcre.MustCompileJIT(pat, 0, pcre.STUDY_JIT_COMPILE).Matcher(b, 0)
	n := 0

	for f := m.Matches; f; f = m.Match(b, 0) {
		n++

		b = b[m.Index()[1]:]
	}

	return n
}
