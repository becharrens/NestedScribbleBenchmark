package callbacks

import (
	"NestedScribbleBenchmark/ring/results/forward"
	"fmt"
)

type Forward_S_Env interface {
	Done() forward.S_Result
	Msg_To_RingNode() (string, int)
}

type ForwardSState struct {
	Msg   string
	NHops int
}

func (f *ForwardSState) Done() forward.S_Result {
	return forward.S_Result{}
}

func (f *ForwardSState) Msg_To_RingNode() (string, int) {
	fmt.Printf("forward s: sending Msg{msg:%s, nhops left:%d} to ringnode\n", f.Msg, f.NHops)
	return f.Msg, f.NHops
}
