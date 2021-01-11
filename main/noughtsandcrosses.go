package main

import (
	"NestedScribbleBenchmark/noughtsandcrosses/callbacks"
	"NestedScribbleBenchmark/noughtsandcrosses/impl"
	"NestedScribbleBenchmark/noughtsandcrosses/protocol"
	noughtsandcrosses_2 "NestedScribbleBenchmark/noughtsandcrosses/results/noughtsandcrosses"
)

type NoughtsAndCrossesEnv struct {
	P1AI bool
	P2AI bool
}

func (n *NoughtsAndCrossesEnv) New_P1_Env() callbacks.NoughtsAndCrosses_P1_Env {
	return &callbacks.NoughtsAndCrossesP1State{
		Board: make(impl.Board, impl.BoardSize),
		IsAI:  n.P1AI,
	}
}

func (n *NoughtsAndCrossesEnv) New_P2_Env() callbacks.NoughtsAndCrosses_P2_Env {
	return &callbacks.NoughtsAndCrossesP2State{
		Board: make(impl.Board, impl.BoardSize),
		IsAI:  n.P2AI,
	}
}

func (n *NoughtsAndCrossesEnv) P1_Result(result noughtsandcrosses_2.P1_Result) {
}

func (n *NoughtsAndCrossesEnv) P2_Result(result noughtsandcrosses_2.P2_Result) {
}

func NewNoughtsAndCrossesEnv(p1ai, p2ai bool) *NoughtsAndCrossesEnv {
	return &NoughtsAndCrossesEnv{
		P1AI: p1ai,
		P2AI: p2ai,
	}
}

func NoughtsAndCrosses(p1ai, p2ai bool) {
	env := NewNoughtsAndCrossesEnv(p1ai, p2ai)
	protocol.NoughtsAndCrosses(env)
}
