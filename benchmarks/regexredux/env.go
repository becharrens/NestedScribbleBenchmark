package regexredux

import (
	"NestedScribbleBenchmark/regexredux/callbacks"
	"NestedScribbleBenchmark/regexredux/results/regexredux"

	"github.com/GRbit/go-pcre"
)

type RegexReduxEnv struct {
	B []byte
}

func (r *RegexReduxEnv) New_Master_Env() callbacks.RegexRedux_Master_Env {
	ilen := len(r.B)
	b := pcre.
		MustCompileJIT("(>[^\n]*)?\n", 0, pcre.STUDY_JIT_COMPILE).
		ReplaceAll(r.B, []byte{}, 0)
	clen := len(b)

	return &callbacks.RegexReduxMasterState{
		B:    b,
		Idx:  0,
		CLen: clen,
		ILen: ilen,
	}
}

func (r *RegexReduxEnv) New_Worker_Env() callbacks.RegexRedux_Worker_Env {
	return &callbacks.RegexReduxWorkerState{}
}

func (r *RegexReduxEnv) Master_Result(result regexredux.Master_Result) {
}

func (r *RegexReduxEnv) Worker_Result(result regexredux.Worker_Result) {
}

func NewRegexReduxEnv(b []byte) *RegexReduxEnv {
	return &RegexReduxEnv{
		B: b,
	}
}
