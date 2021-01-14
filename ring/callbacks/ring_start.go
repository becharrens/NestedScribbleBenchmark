package callbacks

import (
	"NestedScribbleBenchmark/ring/results/forward"
	"NestedScribbleBenchmark/ring/results/ring"
	"fmt"
)

type Ring_Start_Choice int

const (
	Ring_Start_Forward Ring_Start_Choice = iota
	Ring_Start_Msg
)

type Ring_Start_Env interface {
	Msg_From_End_2(msg string, hops int)
	Msg_To_End() (string, int)
	Done() ring.Start_Result
	Msg_From_End(msg string, hops int)
	ResultFrom_Forward_S(result forward.S_Result)
	To_Forward_S_Env() Forward_S_Env
	Forward_Setup()
	Start_Choice() Ring_Start_Choice
}

type RingStartState struct {
	SendMsg  string
	RecvMsg  string
	RingSize int
}

func (r *RingStartState) Msg_From_End_2(msg string, hops int) {
	r.RecvMsg = msg
	fmt.Printf("ring start: received Msg{msg:%s, nhops left: %d} from end\n", msg, hops)
}

func (r *RingStartState) Msg_To_End() (string, int) {
	fmt.Printf("ring start: sending Msg{msg:%s, nhops:%d} to end\n", r.SendMsg, r.RingSize-1)
	return r.SendMsg, r.RingSize - 1
}

func (r *RingStartState) Done() ring.Start_Result {
	return ring.Start_Result{
		RecvMsg: r.RecvMsg,
	}
}

func (r *RingStartState) Msg_From_End(msg string, hops int) {
	r.RecvMsg = msg
	fmt.Printf("ring start: received Msg{msg:%s, nhops left: %d} from end\n", msg, hops)
}

func (r *RingStartState) ResultFrom_Forward_S(result forward.S_Result) {
}

func (r *RingStartState) To_Forward_S_Env() Forward_S_Env {
	return &ForwardSState{
		Msg:   r.SendMsg,
		NHops: r.RingSize - 1,
	}
}

func (r *RingStartState) Forward_Setup() {
}

func (r *RingStartState) Start_Choice() Ring_Start_Choice {
	fmt.Printf("ring start: sending msg \"%s\"\n", r.SendMsg)
	if r.RingSize <= 2 {
		return Ring_Start_Msg
	}
	return Ring_Start_Forward
}
