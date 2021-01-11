package callbacks

import (
	"NestedScribbleBenchmark/noughtsandcrosses/impl"
	"NestedScribbleBenchmark/noughtsandcrosses/results/calcmove"
	"NestedScribbleBenchmark/noughtsandcrosses/results/noughtsandcrosses"
	"fmt"
)

type NoughtsAndCrosses_P1_Choice_2 int

const (
	NoughtsAndCrosses_P1_Win_2 NoughtsAndCrosses_P1_Choice_2 = iota
	NoughtsAndCrosses_P1_Draw_2
	NoughtsAndCrosses_P1_Move_2
)

type NoughtsAndCrosses_P1_Choice int

const (
	NoughtsAndCrosses_P1_Win NoughtsAndCrosses_P1_Choice = iota
	NoughtsAndCrosses_P1_Draw
	NoughtsAndCrosses_P1_Move
)

type NoughtsAndCrosses_P1_Env interface {
	Move_From_P2_2(move int)
	Draw_From_P2_2(move int)
	Win_From_P2_2(move int)
	Move_To_P2_2() int
	Draw_To_P2_2() int
	Win_To_P2_2() int
	P1_Choice_2() NoughtsAndCrosses_P1_Choice_2
	ResultFrom_CalcMove_P_2(result calcmove.P_Result)
	To_CalcMove_P_Env_2() CalcMove_P_Env
	CalcMove_Setup_2()
	Move_From_P2(move int)
	Draw_From_P2(move int)
	Win_From_P2(move int)
	Move_To_P2() int
	Draw_To_P2() int
	Done() noughtsandcrosses.P1_Result
	Win_To_P2() int
	P1_Choice() NoughtsAndCrosses_P1_Choice
	ResultFrom_CalcMove_P(result calcmove.P_Result)
	To_CalcMove_P_Env() CalcMove_P_Env
	CalcMove_Setup()
}

type NoughtsAndCrossesP1State struct {
	Board    impl.Board
	NextMove int
	IsAI     bool
}

func (n *NoughtsAndCrossesP1State) Move_From_P2_2(move int) {
	n.Board[move] = impl.P2Move
}

func (n *NoughtsAndCrossesP1State) Draw_From_P2_2(move int) {
	n.Board[move] = impl.P2Move
}

func (n *NoughtsAndCrossesP1State) Win_From_P2_2(move int) {
	n.Board[move] = impl.P2Move
}

func (n *NoughtsAndCrossesP1State) Move_To_P2_2() int {
	fmt.Println(n.Board)
	fmt.Println()
	return n.NextMove
}

func (n *NoughtsAndCrossesP1State) Draw_To_P2_2() int {
	fmt.Println("Game ended in a draw:")
	fmt.Println(n.Board)
	return n.NextMove
}

func (n *NoughtsAndCrossesP1State) Win_To_P2_2() int {
	fmt.Println("P1 won:")
	fmt.Println(n.Board)
	return n.NextMove
}

func (n *NoughtsAndCrossesP1State) P1_Choice_2() NoughtsAndCrosses_P1_Choice_2 {
	switch impl.BoardOutcome(n.Board, impl.P1) {
	case impl.Continue:
		return NoughtsAndCrosses_P1_Move_2
	case impl.Win:
		return NoughtsAndCrosses_P1_Win_2
	case impl.Draw:
		return NoughtsAndCrosses_P1_Draw_2
	default:
		panic("Invalid board outcome")
	}
}

func (n *NoughtsAndCrossesP1State) ResultFrom_CalcMove_P_2(result calcmove.P_Result) {
	n.NextMove = result.Move
	n.Board[n.NextMove] = impl.P1Move
}

func (n *NoughtsAndCrossesP1State) To_CalcMove_P_Env_2() CalcMove_P_Env {
	return &CalcMovePState{
		Board:  n.Board,
		Player: impl.P1,
		IsAI:   n.IsAI,
	}
}

func (n *NoughtsAndCrossesP1State) CalcMove_Setup_2() {
}

func (n *NoughtsAndCrossesP1State) Move_From_P2(move int) {
	n.Board[move] = impl.P2Move
}

func (n *NoughtsAndCrossesP1State) Draw_From_P2(move int) {
	n.Board[move] = impl.P2Move
}

func (n *NoughtsAndCrossesP1State) Win_From_P2(move int) {
	n.Board[move] = impl.P2Move
}

func (n *NoughtsAndCrossesP1State) Move_To_P2() int {
	fmt.Println(n.Board)
	fmt.Println()
	return n.NextMove
}

func (n *NoughtsAndCrossesP1State) Draw_To_P2() int {
	fmt.Println("Game ended in a draw:")
	fmt.Println(n.Board)
	return n.NextMove
}

func (n *NoughtsAndCrossesP1State) Done() noughtsandcrosses.P1_Result {
	return noughtsandcrosses.P1_Result{}
}

func (n *NoughtsAndCrossesP1State) Win_To_P2() int {
	fmt.Println("P1 won:")
	fmt.Println(n.Board)
	return n.NextMove
}

func (n *NoughtsAndCrossesP1State) P1_Choice() NoughtsAndCrosses_P1_Choice {
	switch impl.BoardOutcome(n.Board, impl.P1) {
	case impl.Continue:
		return NoughtsAndCrosses_P1_Move
	case impl.Win:
		return NoughtsAndCrosses_P1_Win
	case impl.Draw:
		return NoughtsAndCrosses_P1_Draw
	default:
		panic("Invalid board outcome")
	}
}

func (n *NoughtsAndCrossesP1State) ResultFrom_CalcMove_P(result calcmove.P_Result) {
	n.NextMove = result.Move
	n.Board[n.NextMove] = impl.P1Move
}

func (n *NoughtsAndCrossesP1State) To_CalcMove_P_Env() CalcMove_P_Env {
	return &CalcMovePState{
		Board:  n.Board,
		Player: impl.P1,
		IsAI:   n.IsAI,
	}
}

func (n *NoughtsAndCrossesP1State) CalcMove_Setup() {
}
