package main

import (
	"NestedScribbleBenchmark/ring/callbacks"
	"NestedScribbleBenchmark/ring/results/ring"
)

type RingEnv struct {
	SendMsg string
	NHops   int
	RecvMsg string
}

func (r *RingEnv) New_Start_Env() callbacks.Ring_Start_Env {
	return &callbacks.RingStartState{
		SendMsg: r.SendMsg,
		RecvMsg: "",
		NHops:   r.NHops,
	}
}

func (r *RingEnv) New_End_Env() callbacks.Ring_End_Env {
	return &callbacks.RingEndState{}
}

func (r *RingEnv) Start_Result(result ring.Start_Result) {
	r.RecvMsg = result.RecvMsg
}

func (r *RingEnv) End_Result(result ring.End_Result) {
}

func NewRingEnv(msg string, nhops int) *RingEnv {
	return &RingEnv{
		SendMsg: msg,
		NHops:   nhops,
		RecvMsg: "",
	}
}
