package main

import (
	"NestedScribbleBenchmark/ring/callbacks"
	"NestedScribbleBenchmark/ring/protocol"
	"NestedScribbleBenchmark/ring/results/ring"
	"fmt"
)

const RING_MSG = "PASS IT ON"

type RingEnv struct {
	SendMsg string
	NHops   int
	RecvMsg string
}

func (r *RingEnv) New_Start_Env() callbacks.Ring_Start_Env {
	return &callbacks.RingStartState{
		SendMsg:  r.SendMsg,
		RecvMsg:  "",
		RingSize: r.NHops,
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

func NewRingEnv(msg string, ringSize int) *RingEnv {
	return &RingEnv{
		SendMsg: msg,
		NHops:   ringSize,
		RecvMsg: "",
	}
}

func RunRing(ringSize int) {
	env := NewRingEnv(RING_MSG, ringSize)
	protocol.Ring(env)
	fmt.Println("Message which went around the ring: '" + env.RecvMsg + "'")
}
