package callbacks

import (
	"ScribbleBenchmark/ring/messages/ring"
	"fmt"
)
import ring_2 "ScribbleBenchmark/ring/results/ring"
import "ScribbleBenchmark/ring/results/forward"

type Ring_End_Env interface {
	Msg_To_Start_2() ring.Msg
	Msg_From_Start(msg ring.Msg)
	Done() ring_2.End_Result
	Msg_To_Start() ring.Msg
	ResultFrom_Forward_E(result forward.E_Result)
	To_Forward_E_Env() Forward_E_Env
}

type RingEndState struct {
	Msg   string
	NHops int
}

func (r *RingEndState) Msg_To_Start_2() ring.Msg {
	fmt.Printf("ring end: sending Msg{msg:%s, nhops left:%d} to start\n", r.Msg, r.NHops)
	return ring.Msg{
		Hops: r.NHops,
		Msg:  r.Msg,
	}
}

func (r *RingEndState) Msg_From_Start(msg ring.Msg) {
	r.NHops = msg.Hops - 1
	r.Msg = msg.Msg
	fmt.Printf("ring end: received Msg{msg:%s, nhops left:%d} from start\n", r.Msg, msg.Hops)
}

func (r *RingEndState) Done() ring_2.End_Result {
	return ring_2.End_Result{}
}

func (r *RingEndState) Msg_To_Start() ring.Msg {
	fmt.Printf("ring end: sending Msg{msg:%s, nhops left:%d} to start\n", r.Msg, r.NHops)
	return ring.Msg{
		Hops: r.NHops,
		Msg:  r.Msg,
	}
}

func (r *RingEndState) ResultFrom_Forward_E(result forward.E_Result) {
	r.Msg = result.Msg
	r.NHops = result.NHops
	fmt.Printf("ring end: received Msg{msg:%s, nhops left:%d}\n", r.Msg, r.NHops)
}

func (r *RingEndState) To_Forward_E_Env() Forward_E_Env {
	return &ForwardEState{}
}
