package callbacks

import (
	"NestedScribbleBenchmark/noughtsandcrosses/impl"
	"NestedScribbleBenchmark/noughtsandcrosses/results/calcmove"
	"NestedScribbleBenchmark/noughtsandcrosses/results/noughtsandcrosses"
	"fmt"
)

type NoughtsAndCrosses_P2_Choice_2 int

const (
	NoughtsAndCrosses_P2_Win_2 NoughtsAndCrosses_P2_Choice_2 = iota
	NoughtsAndCrosses_P2_Draw_2
	NoughtsAndCrosses_P2_Move_2
)

type NoughtsAndCrosses_P2_Choice int

const (
	NoughtsAndCrosses_P2_Win NoughtsAndCrosses_P2_Choice = iota
	NoughtsAndCrosses_P2_Draw
	NoughtsAndCrosses_P2_Move
)

type NoughtsAndCrosses_P2_Env interface {
	Move_To_P1_2() int
	Draw_To_P1_2() int
	Win_To_P1_2() int
	P2_Choice_2() NoughtsAndCrosses_P2_Choice_2
	ResultFrom_CalcMove_P_2(result calcmove.P_Result)
	To_CalcMove_P_Env_2() CalcMove_P_Env
	CalcMove_Setup_2()
	Move_From_P1_2(move int)
	Draw_From_P1_2(move int)
	Win_From_P1_2(move int)
	Move_To_P1() int
	Draw_To_P1() int
	Win_To_P1() int
	P2_Choice() NoughtsAndCrosses_P2_Choice
	ResultFrom_CalcMove_P(result calcmove.P_Result)
	To_CalcMove_P_Env() CalcMove_P_Env
	CalcMove_Setup()
	Move_From_P1(move int)
	Draw_From_P1(move int)
	Done() noughtsandcrosses.P2_Result
	Win_From_P1(move int)
}

type NoughtsAndCrossesP2State struct {
	Board    impl.Board
	NextMove int
	IsAI     bool
}

func (n *NoughtsAndCrossesP2State) Move_To_P1_2() int {
	fmt.Println(n.Board)
	fmt.Println()
	return n.NextMove
}

func (n *NoughtsAndCrossesP2State) Draw_To_P1_2() int {
	fmt.Println("Game ended in a draw:")
	fmt.Println(n.Board)
	return n.NextMove
}

func (n *NoughtsAndCrossesP2State) Win_To_P1_2() int {
	fmt.Println("P2 won:")
	fmt.Println(n.Board)
	return n.NextMove
}

func (n *NoughtsAndCrossesP2State) P2_Choice_2() NoughtsAndCrosses_P2_Choice_2 {
	switch impl.BoardOutcome(n.Board, impl.P2) {
	case impl.Continue:
		return NoughtsAndCrosses_P2_Move_2
	case impl.Win:
		return NoughtsAndCrosses_P2_Win_2
	case impl.Draw:
		return NoughtsAndCrosses_P2_Draw_2
	default:
		panic("Invalid board outcome")
	}
}

func (n *NoughtsAndCrossesP2State) ResultFrom_CalcMove_P_2(result calcmove.P_Result) {
	n.NextMove = result.Move
	n.Board[n.NextMove] = impl.P2Move
}

func (n *NoughtsAndCrossesP2State) To_CalcMove_P_Env_2() CalcMove_P_Env {
	return &CalcMovePState{
		Board:  n.Board,
		Player: impl.P2,
		IsAI:   n.IsAI,
	}
}

func (n *NoughtsAndCrossesP2State) CalcMove_Setup_2() {
}

func (n *NoughtsAndCrossesP2State) Move_From_P1_2(move int) {
	n.Board[move] = impl.P1Move
}

func (n *NoughtsAndCrossesP2State) Draw_From_P1_2(move int) {
	n.Board[move] = impl.P1Move
}

func (n *NoughtsAndCrossesP2State) Win_From_P1_2(move int) {
	n.Board[move] = impl.P1Move
}

func (n *NoughtsAndCrossesP2State) Move_To_P1() int {
	fmt.Println(n.Board)
	fmt.Println()
	return n.NextMove
}

func (n *NoughtsAndCrossesP2State) Draw_To_P1() int {
	fmt.Println("Game ended in a draw:")
	fmt.Println(n.Board)
	return n.NextMove
}

func (n *NoughtsAndCrossesP2State) Win_To_P1() int {
	fmt.Println("P2 won:")
	fmt.Println(n.Board)
	return n.NextMove
}

func (n *NoughtsAndCrossesP2State) P2_Choice() NoughtsAndCrosses_P2_Choice {
	switch impl.BoardOutcome(n.Board, impl.P2) {
	case impl.Continue:
		return NoughtsAndCrosses_P2_Move
	case impl.Win:
		return NoughtsAndCrosses_P2_Win
	case impl.Draw:
		return NoughtsAndCrosses_P2_Draw
	default:
		panic("Invalid board outcome")
	}
}

func (n *NoughtsAndCrossesP2State) ResultFrom_CalcMove_P(result calcmove.P_Result) {
	n.NextMove = result.Move
	n.Board[n.NextMove] = impl.P2Move
}

func (n *NoughtsAndCrossesP2State) To_CalcMove_P_Env() CalcMove_P_Env {
	return &CalcMovePState{
		Board:  n.Board,
		Player: impl.P2,
		IsAI:   n.IsAI,
	}
}

func (n *NoughtsAndCrossesP2State) CalcMove_Setup() {
}

func (n *NoughtsAndCrossesP2State) Move_From_P1(move int) {
	n.Board[move] = impl.P1Move
}

func (n *NoughtsAndCrossesP2State) Draw_From_P1(move int) {
	n.Board[move] = impl.P1Move
}

func (n *NoughtsAndCrossesP2State) Done() noughtsandcrosses.P2_Result {
	return noughtsandcrosses.P2_Result{}
}

func (n *NoughtsAndCrossesP2State) Win_From_P1(move int) {
	n.Board[move] = impl.P1Move
}
