package callbacks

import (
	"NestedScribbleBenchmark/ring/messages/ring"
	"fmt"
)
import ring_2 "NestedScribbleBenchmark/ring/results/ring"
import "NestedScribbleBenchmark/ring/results/forward"

type Ring_Start_Choice int

const (
	Ring_Start_Forward Ring_Start_Choice = iota
	Ring_Start_Msg
)

type Ring_Start_Env interface {
	Msg_From_End_2(msg ring.Msg)
	Msg_To_End() ring.Msg
	Done() ring_2.Start_Result
	Msg_From_End(msg ring.Msg)
	ResultFrom_Forward_S(result forward.S_Result)
	To_Forward_S_Env() Forward_S_Env
	Forward_Setup()
	Start_Choice() Ring_Start_Choice
}

type RingStartState struct {
	SendMsg string
	RecvMsg string
	NHops   int
}

func (r *RingStartState) Msg_From_End_2(msg ring.Msg) {
	r.RecvMsg = msg.Msg
	fmt.Printf("ring start: received Msg{msg:%s, nhops left: %d} from end\n", msg.Msg, msg.Hops)

}

func (r *RingStartState) Msg_To_End() ring.Msg {
	fmt.Printf("ring start: sending Msg{msg:%s, nhops:%d} to end\n", r.SendMsg, r.NHops)
	return ring.Msg{
		Hops: r.NHops,
		Msg:  r.SendMsg,
	}
}

func (r *RingStartState) Done() ring_2.Start_Result {
	return ring_2.Start_Result{
		RecvMsg: r.RecvMsg,
	}
}

func (r *RingStartState) Msg_From_End(msg ring.Msg) {
	r.RecvMsg = msg.Msg
	fmt.Printf("ring start: received Msg{msg:%s, nhops left: %d} from end\n", msg.Msg, msg.Hops)
}

func (r *RingStartState) ResultFrom_Forward_S(result forward.S_Result) {
}

func (r *RingStartState) To_Forward_S_Env() Forward_S_Env {
	return &ForwardSState{
		Msg:   r.SendMsg,
		NHops: r.NHops,
	}
}

func (r *RingStartState) Forward_Setup() {
}

func (r *RingStartState) Start_Choice() Ring_Start_Choice {
	fmt.Printf("ring start: sending msg \"%s\"\n", r.SendMsg)
	if r.NHops <= 1 {
		return Ring_Start_Msg
	}
	return Ring_Start_Forward
}
