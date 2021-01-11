package callbacks

import (
	"NestedScribbleBenchmark/noughtsandcrosses/impl"
	"NestedScribbleBenchmark/noughtsandcrosses/results/minmaxstrategy"
	"NestedScribbleBenchmark/noughtsandcrosses/results/minmaxstrategy_evalboard"
)

type MinMaxStrategy_Worker_Choice_2 int

const (
	MinMaxStrategy_Worker_MinMaxStrategy_2 MinMaxStrategy_Worker_Choice_2 = iota
	MinMaxStrategy_Worker_MinMaxStrategy_EvalBoard_2
)

type MinMaxStrategy_Worker_Choice int

const (
	MinMaxStrategy_Worker_MinMaxStrategy MinMaxStrategy_Worker_Choice = iota
	MinMaxStrategy_Worker_MinMaxStrategy_EvalBoard
)

type MinMaxStrategy_Worker_Env interface {
	Score_To_Master_4() int
	ResultFrom_MinMaxStrategy_EvalBoard_W_2(result minmaxstrategy_evalboard.W_Result)
	To_MinMaxStrategy_EvalBoard_W_Env_2() MinMaxStrategy_EvalBoard_W_Env
	MinMaxStrategy_EvalBoard_Setup_2()
	Score_To_Master_3() int
	ResultFrom_MinMaxStrategy_Master_2(result minmaxstrategy.Master_Result)
	To_MinMaxStrategy_Master_Env_2() MinMaxStrategy_Master_Env
	MinMaxStrategy_Setup_2()
	Worker_Choice_2() MinMaxStrategy_Worker_Choice_2
	FinalState_From_Master(board impl.Board, currPlayer impl.Player, toMove impl.Player)
	Score_To_Master_2() int
	ResultFrom_MinMaxStrategy_EvalBoard_W(result minmaxstrategy_evalboard.W_Result)
	To_MinMaxStrategy_EvalBoard_W_Env() MinMaxStrategy_EvalBoard_W_Env
	MinMaxStrategy_EvalBoard_Setup()
	Done()
	Score_To_Master() int
	ResultFrom_MinMaxStrategy_Master(result minmaxstrategy.Master_Result)
	To_MinMaxStrategy_Master_Env() MinMaxStrategy_Master_Env
	MinMaxStrategy_Setup()
	Worker_Choice() MinMaxStrategy_Worker_Choice
	CurrState_From_Master(board impl.Board, currPlayer impl.Player, toMove impl.Player)
}

// CurrPlayer = Player who made the last move
// PlayerToMove = Player who will make the next move
type MinMaxStrategyWorkerState struct {
	Score        int
	Board        impl.Board
	PlayerToMove impl.Player
	CurrPlayer   impl.Player
}

func (m *MinMaxStrategyWorkerState) Score_To_Master_4() int {
	return m.Score
}

func (m *MinMaxStrategyWorkerState) ResultFrom_MinMaxStrategy_EvalBoard_W_2(result minmaxstrategy_evalboard.W_Result) {
	m.Score = result.Score
}

func (m *MinMaxStrategyWorkerState) To_MinMaxStrategy_EvalBoard_W_Env_2() MinMaxStrategy_EvalBoard_W_Env {
	return &MinMaxStrategyEvalBoardWState{
		Board:          m.Board,
		LastMovePlayer: m.CurrPlayer,
	}
}

func (m *MinMaxStrategyWorkerState) MinMaxStrategy_EvalBoard_Setup_2() {
}

func (m *MinMaxStrategyWorkerState) Score_To_Master_3() int {
	return m.Score
}

func (m *MinMaxStrategyWorkerState) ResultFrom_MinMaxStrategy_Master_2(result minmaxstrategy.Master_Result) {
	m.Score = result.Score
}

func (m *MinMaxStrategyWorkerState) To_MinMaxStrategy_Master_Env_2() MinMaxStrategy_Master_Env {
	moves := make(impl.MoveStack, 0)
	return &MinMaxStrategyMasterState{
		Moves:        &moves,
		BestMove:     0,
		Scores:       make(map[int]int),
		BestScore:    impl.InitMoveScore(m.PlayerToMove),
		NextMoveIdx:  0,
		Board:        m.Board,
		PlayerToMove: m.PlayerToMove,
	}
}

func (m *MinMaxStrategyWorkerState) MinMaxStrategy_Setup_2() {
}

func (m *MinMaxStrategyWorkerState) Worker_Choice_2() MinMaxStrategy_Worker_Choice_2 {
	// Eval last move of other player
	switch impl.BoardOutcome(m.Board, m.CurrPlayer) {
	case impl.Continue:
		return MinMaxStrategy_Worker_MinMaxStrategy_2
	default:
		return MinMaxStrategy_Worker_MinMaxStrategy_EvalBoard_2
	}
}

func (m *MinMaxStrategyWorkerState) FinalState_From_Master(board impl.Board, currPlayer impl.Player, toMove impl.Player) {
	m.Board = board
	m.CurrPlayer = currPlayer
	m.PlayerToMove = toMove
}

func (m *MinMaxStrategyWorkerState) Score_To_Master_2() int {
	return m.Score
}

func (m *MinMaxStrategyWorkerState) ResultFrom_MinMaxStrategy_EvalBoard_W(result minmaxstrategy_evalboard.W_Result) {
	m.Score = result.Score
}

func (m *MinMaxStrategyWorkerState) To_MinMaxStrategy_EvalBoard_W_Env() MinMaxStrategy_EvalBoard_W_Env {
	return &MinMaxStrategyEvalBoardWState{
		Board:          m.Board,
		LastMovePlayer: m.CurrPlayer,
	}
}

func (m *MinMaxStrategyWorkerState) MinMaxStrategy_EvalBoard_Setup() {
}

func (m *MinMaxStrategyWorkerState) Done() {
}

func (m *MinMaxStrategyWorkerState) Score_To_Master() int {
	return m.Score
}

func (m *MinMaxStrategyWorkerState) ResultFrom_MinMaxStrategy_Master(result minmaxstrategy.Master_Result) {
	m.Score = result.Score
}

func (m *MinMaxStrategyWorkerState) To_MinMaxStrategy_Master_Env() MinMaxStrategy_Master_Env {
	moves := make(impl.MoveStack, 0)
	return &MinMaxStrategyMasterState{
		Moves:        &moves,
		BestMove:     0,
		Scores:       make(map[int]int),
		BestScore:    impl.InitMoveScore(m.PlayerToMove),
		NextMoveIdx:  0,
		Board:        m.Board,
		PlayerToMove: m.PlayerToMove,
	}
}

func (m *MinMaxStrategyWorkerState) MinMaxStrategy_Setup() {
}

func (m *MinMaxStrategyWorkerState) Worker_Choice() MinMaxStrategy_Worker_Choice {
	switch impl.BoardOutcome(m.Board, m.CurrPlayer) {
	case impl.Continue:
		return MinMaxStrategy_Worker_MinMaxStrategy
	default:
		return MinMaxStrategy_Worker_MinMaxStrategy_EvalBoard
	}
}

func (m *MinMaxStrategyWorkerState) CurrState_From_Master(board impl.Board, currPlayer impl.Player, toMove impl.Player) {
	m.Board = board
	m.CurrPlayer = currPlayer
	m.PlayerToMove = toMove
}

func New_MinMaxStrategy_Worker_State() MinMaxStrategy_Worker_Env {
	return &MinMaxStrategyWorkerState{}
}
