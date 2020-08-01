package callbacks

import (
	forward_2 "NestedScribbleBenchmark/ring/messages/forward"
	"NestedScribbleBenchmark/ring/results/forward"
	"fmt"
)

type Forward_E_Env interface {
	Msg_From_RingNode(msg forward_2.Msg)
	Done() forward.E_Result
	ResultFrom_Forward_E(result forward.E_Result)
	To_Forward_E_Env() Forward_E_Env
}

type ForwardEState struct {
	NHops int
	Msg   string
}

func (f *ForwardEState) Msg_From_RingNode(msg forward_2.Msg) {
	f.NHops = msg.Hops
	f.Msg = msg.Msg
	fmt.Printf("forward e: received Msg{msg:%s, nhops left:%d} from ringnode\n", f.Msg, f.NHops)
}

func (f *ForwardEState) Done() forward.E_Result {
	return forward.E_Result{
		Msg:   f.Msg,
		NHops: f.NHops,
	}
}

func (f *ForwardEState) ResultFrom_Forward_E(result forward.E_Result) {
	f.NHops = result.NHops
	f.Msg = result.Msg
}

func (f *ForwardEState) To_Forward_E_Env() Forward_E_Env {
	return &ForwardEState{}
}
