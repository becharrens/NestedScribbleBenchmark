package callbacks

import (
	"NestedScribbleBenchmark/regexredux/messages/regexredux"

	"github.com/GRbit/go-pcre"
)
import regexredux_2 "NestedScribbleBenchmark/regexredux/results/regexredux"

type RegexRedux_Worker_Env interface {
	Done() regexredux_2.Worker_Result
	NumMatches_To_Master() regexredux.NumMatches
	Task_From_Master(task_msg regexredux.Task)
}

type RegexReduxWorkerState struct {
	Nmatches int
}

func (r *RegexReduxWorkerState) Done() regexredux_2.Worker_Result {
	return regexredux_2.Worker_Result{}
}

func (r *RegexReduxWorkerState) NumMatches_To_Master() regexredux.NumMatches {
	return regexredux.NumMatches{Nmatches: r.Nmatches}
}

func (r *RegexReduxWorkerState) Task_From_Master(task_msg regexredux.Task) {
	r.Nmatches = countMatches(task_msg.Pattern, task_msg.B)
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
