package callbacks

import (
	"NestedScribbleBenchmark/ring/results/forward"
	"NestedScribbleBenchmark/ring/results/ring"
	"fmt"
)

type Ring_End_Env interface {
	Msg_To_Start_2() (string, int)
	Msg_From_Start(msg string, hops int)
	Done() ring.End_Result
	Msg_To_Start() (string, int)
	ResultFrom_Forward_E(result forward.E_Result)
	To_Forward_E_Env() Forward_E_Env
}

type RingEndState struct {
	Msg   string
	NHops int
}

func (r *RingEndState) Msg_To_Start_2() (string, int) {
	fmt.Printf("ring end: sending Msg{msg:%s, nhops left:%d} to start\n", r.Msg, r.NHops)
	return r.Msg, r.NHops
}

func (r *RingEndState) Msg_From_Start(msg string, hops int) {
	r.NHops = hops - 1
	r.Msg = msg
	fmt.Printf("ring end: received Msg{msg:%s, nhops left:%d} from start\n", msg, hops)
}

func (r *RingEndState) Done() ring.End_Result {
	return ring.End_Result{}
}

func (r *RingEndState) Msg_To_Start() (string, int) {
	fmt.Printf("ring end: sending Msg{msg:%s, nhops left:%d} to start\n", r.Msg, r.NHops)
	return r.Msg, r.NHops
}

func (r *RingEndState) ResultFrom_Forward_E(result forward.E_Result) {
	fmt.Printf("ring end: received Msg{msg:%s, nhops left:%d}\n", result.Msg, result.NHops)
	r.Msg = result.Msg
	r.NHops = result.NHops - 1
}

func (r *RingEndState) To_Forward_E_Env() Forward_E_Env {
	return &ForwardEState{}
}
