package callbacks

import (
	"NestedScribbleBenchmark/ring/messages/forward"
	"fmt"
)
import forward_2 "NestedScribbleBenchmark/ring/results/forward"

type Forward_S_Env interface {
	Done() forward_2.S_Result
	Msg_To_RingNode() forward.Msg
}

type ForwardSState struct {
	Msg   string
	NHops int
}

func (f *ForwardSState) Done() forward_2.S_Result {
	return forward_2.S_Result{}
}

func (f *ForwardSState) Msg_To_RingNode() forward.Msg {
	fmt.Printf("forward s: sending Msg{msg:%s, nhops left:%d} to ringnode\n", f.Msg, f.NHops-1)
	return forward.Msg{
		Hops: f.NHops - 1,
		Msg:  f.Msg,
	}
}
