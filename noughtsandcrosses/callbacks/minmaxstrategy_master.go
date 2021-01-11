package callbacks

import (
	"NestedScribbleBenchmark/noughtsandcrosses/impl"
	"NestedScribbleBenchmark/noughtsandcrosses/results/minmaxstrategy"
)

type MinMaxStrategy_Master_Choice int

const (
	MinMaxStrategy_Master_CurrState MinMaxStrategy_Master_Choice = iota
	MinMaxStrategy_Master_FinalState
)

type MinMaxStrategy_Master_Env interface {
	Score_From_Worker_2(score int)
	FinalState_To_Worker() (impl.Board, impl.Player, impl.Player)
	Done() minmaxstrategy.Master_Result
	Score_From_Worker(score int)
	ResultFrom_MinMaxStrategy_Master(result minmaxstrategy.Master_Result)
	To_MinMaxStrategy_Master_Env() MinMaxStrategy_Master_Env
	MinMaxStrategy_Setup()
	CurrState_To_Worker() (impl.Board, impl.Player, impl.Player)
	Master_Choice() MinMaxStrategy_Master_Choice
}

type MinMaxStrategyMasterState struct {
	Moves        *impl.MoveStack
	Scores       map[int]int
	BestMove     int
	BestScore    int
	NextMoveIdx  int
	Board        impl.Board
	PlayerToMove impl.Player
	// CurrPlayer   impl.Player
}

func (m *MinMaxStrategyMasterState) Score_From_Worker_2(score int) {
	move := m.Moves.Pop()
	m.Scores[move] = score
	if m.PlayerToMove == impl.P1 {
		// Max
		if score > m.BestScore {
			m.BestMove = move
			m.BestScore = score
		}
	} else {
		// Min
		if score < m.BestScore {
			m.BestMove = move
			m.BestScore = score
		}
	}
}

func (m *MinMaxStrategyMasterState) FinalState_To_Worker() (impl.Board, impl.Player, impl.Player) {
	workerBoard := impl.MakeMove(m.Board, m.PlayerToMove, m.NextMoveIdx)
	m.Moves.Push(m.NextMoveIdx)
	m.NextMoveIdx++
	return workerBoard, m.PlayerToMove, m.PlayerToMove.Opponent()
}

func (m *MinMaxStrategyMasterState) Done() minmaxstrategy.Master_Result {
	return minmaxstrategy.Master_Result{
		Score: m.BestScore,
		Move:  m.BestMove,
	}
}

func (m *MinMaxStrategyMasterState) Score_From_Worker(score int) {
	move := m.Moves.Pop()
	m.Scores[move] = score
	if m.PlayerToMove == impl.P1 {
		// Max
		if score > m.BestScore {
			m.BestMove = move
			m.BestScore = score
		}
	} else {
		// Min
		if score < m.BestScore {
			m.BestMove = move
			m.BestScore = score
		}
	}
}

func (m *MinMaxStrategyMasterState) ResultFrom_MinMaxStrategy_Master(result minmaxstrategy.Master_Result) {
	// m.BestMove = result.Move
	// m.BestScore = result.Score
}

func (m *MinMaxStrategyMasterState) To_MinMaxStrategy_Master_Env() MinMaxStrategy_Master_Env {
	return m
}

func (m *MinMaxStrategyMasterState) MinMaxStrategy_Setup() {
}

func (m *MinMaxStrategyMasterState) CurrState_To_Worker() (impl.Board, impl.Player, impl.Player) {
	workerBoard := impl.MakeMove(m.Board, m.PlayerToMove, m.NextMoveIdx)
	m.Moves.Push(m.NextMoveIdx)
	m.NextMoveIdx++

	return workerBoard, m.PlayerToMove, m.PlayerToMove.Opponent()
}

func (m *MinMaxStrategyMasterState) Master_Choice() MinMaxStrategy_Master_Choice {
	m.NextMoveIdx = impl.NextMoveIdx(m.Board, m.NextMoveIdx)
	if impl.NextMoveIdx(m.Board, m.NextMoveIdx+1) == -1 {
		// Next move is final move
		return MinMaxStrategy_Master_FinalState
	}
	return MinMaxStrategy_Master_CurrState
}

func NEmpty(board impl.Board) int {
	count := 0
	for _, tile := range board {
		if tile == impl.Empty {
			count++
		}
	}
	return count
}
