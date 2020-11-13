package callbacks

import (
	"NestedScribbleBenchmark/ring/messages/forward"
	"fmt"
)
import forward_2 "NestedScribbleBenchmark/ring/results/forward"

type Forward_RingNode_Choice int

const (
	Forward_RingNode_Forward Forward_RingNode_Choice = iota
	Forward_RingNode_Msg
)

type Forward_RingNode_Env interface {
	Msg_To_E() forward.Msg
	Done()
	ResultFrom_Forward_S(result forward_2.S_Result)
	To_Forward_S_Env() Forward_S_Env
	Forward_Setup()
	RingNode_Choice() Forward_RingNode_Choice
	Msg_From_S(msg forward.Msg)
}

type ForwardRingNodeState struct {
	Msg   string
	NHops int
}

func (f *ForwardRingNodeState) Msg_To_E() forward.Msg {
	fmt.Printf("forward ringnode: sending Msg{msg:%s, nhops:%d} to e\n", f.Msg, f.NHops)
	return forward.Msg{
		Hops: f.NHops,
		Msg:  f.Msg,
	}
}

func (f *ForwardRingNodeState) Done() {
}

func (f *ForwardRingNodeState) ResultFrom_Forward_S(result forward_2.S_Result) {
}

func (f *ForwardRingNodeState) To_Forward_S_Env() Forward_S_Env {
	return &ForwardSState{
		Msg:   f.Msg,
		NHops: f.NHops,
	}
}

func (f *ForwardRingNodeState) Forward_Setup() {
}

func (f *ForwardRingNodeState) RingNode_Choice() Forward_RingNode_Choice {
	if f.NHops <= 1 {
		return Forward_RingNode_Msg
	}
	return Forward_RingNode_Forward
}

func (f *ForwardRingNodeState) Msg_From_S(msg forward.Msg) {
	f.NHops = msg.Hops
	f.Msg = msg.Msg
	fmt.Printf("forward ringnode: received Msg{msg:%s, nhops left:%d} from s\n", f.Msg, msg.Hops)
}

func New_Forward_RingNode_State() Forward_RingNode_Env {
	return &ForwardRingNodeState{}
}
