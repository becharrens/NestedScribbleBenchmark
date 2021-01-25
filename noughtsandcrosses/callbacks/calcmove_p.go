package callbacks

import (
	"NestedScribbleBenchmark/noughtsandcrosses/impl"
	"NestedScribbleBenchmark/noughtsandcrosses/results/calcmove"
	"NestedScribbleBenchmark/noughtsandcrosses/results/minmaxstrategy"
	"NestedScribbleBenchmark/noughtsandcrosses/results/standardstrategy"
)

type CalcMove_P_Choice int

const (
	CalcMove_P_StandardStrategy CalcMove_P_Choice = iota
	CalcMove_P_MinMaxStrategy
)

type CalcMove_P_Env interface {
	ResultFrom_MinMaxStrategy_Master(result minmaxstrategy.Master_Result)
	To_MinMaxStrategy_Master_Env() MinMaxStrategy_Master_Env
	MinMaxStrategy_Setup()
	Done() calcmove.P_Result
	ResultFrom_StandardStrategy_P(result standardstrategy.P_Result)
	To_StandardStrategy_P_Env() StandardStrategy_P_Env
	StandardStrategy_Setup()
	P_Choice() CalcMove_P_Choice
}

type CalcMovePState struct {
	Board  impl.Board
	Player impl.Player
	IsAI   bool
	Move   int
}

func (c *CalcMovePState) ResultFrom_MinMaxStrategy_Master(result minmaxstrategy.Master_Result) {
	c.Move = result.Move
}

func (c *CalcMovePState) To_MinMaxStrategy_Master_Env() MinMaxStrategy_Master_Env {
	moves := make(impl.MoveStack, 0)
	return &MinMaxStrategyMasterState{
		Moves:        &moves,
		Scores:       make(map[int]int),
		BestScore:    impl.InitMoveScore(c.Player),
		NextMoveIdx:  0,
		Board:        c.Board,
		PlayerToMove: c.Player,
	}
}

func (c *CalcMovePState) MinMaxStrategy_Setup() {
}

func (c *CalcMovePState) Done() calcmove.P_Result {
	return calcmove.P_Result{
		Move: c.Move,
	}
}

func (c *CalcMovePState) ResultFrom_StandardStrategy_P(result standardstrategy.P_Result) {
	c.Move = result.Move
}

func (c *CalcMovePState) To_StandardStrategy_P_Env() StandardStrategy_P_Env {
	return &StandardStrategyPState{
		Board:  c.Board,
		Player: c.Player,
	}
}

func (c *CalcMovePState) StandardStrategy_Setup() {
}

func (c *CalcMovePState) P_Choice() CalcMove_P_Choice {
	if c.IsAI {
		return CalcMove_P_MinMaxStrategy
	} else {
		return CalcMove_P_StandardStrategy
	}
}
