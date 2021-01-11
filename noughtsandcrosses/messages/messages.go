package messages

type NoughtsAndCrosses_Label int
const (
	CalcMove_P1 NoughtsAndCrosses_Label = iota
	CalcMove_P2
	CurrState
	Draw
	FinalState
	MinMaxStrategy_Master
	MinMaxStrategy_P
	MinMaxStrategy_Worker
	MinMaxStrategy_EvalBoard_Worker
	Move
	Score
	StandardStrategy_P
	Win
)