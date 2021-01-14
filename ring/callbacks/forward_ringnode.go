package callbacks

import (
	"NestedScribbleBenchmark/ring/results/forward"
	"fmt"
)

type Forward_RingNode_Choice int

const (
	Forward_RingNode_Forward Forward_RingNode_Choice = iota
	Forward_RingNode_Msg
)

type Forward_RingNode_Env interface {
	Msg_To_E() (string, int)
	Done()
	ResultFrom_Forward_S(result forward.S_Result)
	To_Forward_S_Env() Forward_S_Env
	Forward_Setup()
	RingNode_Choice() Forward_RingNode_Choice
	Msg_From_S(msg string, hops int)
}

type ForwardRingNodeState struct {
	Msg   string
	NHops int
}

func (f *ForwardRingNodeState) Msg_To_E() (string, int) {
	fmt.Printf("forward ringnode: sending Msg{msg:%s, nhops:%d} to e\n", f.Msg, f.NHops)
	return f.Msg, f.NHops
}

func (f *ForwardRingNodeState) Done() {
}

func (f *ForwardRingNodeState) ResultFrom_Forward_S(result forward.S_Result) {
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
	if f.NHops < 2 {
		return Forward_RingNode_Msg
	}
	return Forward_RingNode_Forward
}

func (f *ForwardRingNodeState) Msg_From_S(msg string, hops int) {
	fmt.Printf("forward ringnode: received Msg{msg:%s, nhops left:%d} from s\n", msg, hops)
	f.NHops = hops - 1
	f.Msg = msg
}

func New_Forward_RingNode_State() Forward_RingNode_Env {
	return &ForwardRingNodeState{}
}
