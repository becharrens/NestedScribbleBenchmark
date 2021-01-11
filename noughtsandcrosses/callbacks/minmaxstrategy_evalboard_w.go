package callbacks

import (
	"NestedScribbleBenchmark/noughtsandcrosses/impl"
	"NestedScribbleBenchmark/noughtsandcrosses/results/minmaxstrategy_evalboard"
)

type MinMaxStrategy_EvalBoard_W_Env interface {
	Done() minmaxstrategy_evalboard.W_Result
}

type MinMaxStrategyEvalBoardWState struct {
	Board impl.Board
	// CurrPlayer     impl.Player
	LastMovePlayer impl.Player
}

func (m *MinMaxStrategyEvalBoardWState) Done() minmaxstrategy_evalboard.W_Result {
	outcome := impl.BoardOutcome(m.Board, m.LastMovePlayer)
	score := impl.ScoreOutcome(outcome, m.LastMovePlayer)
	return minmaxstrategy_evalboard.W_Result{Score: score}
}
