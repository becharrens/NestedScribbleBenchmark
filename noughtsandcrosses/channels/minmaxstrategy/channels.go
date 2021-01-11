package minmaxstrategy

import (
	"NestedScribbleBenchmark/noughtsandcrosses/impl"
	"NestedScribbleBenchmark/noughtsandcrosses/messages"
)

type Master_Chan struct {
	Board_To_Worker   chan impl.Board
	Int_From_Worker   chan int
	Label_From_Worker chan messages.NoughtsAndCrosses_Label
	Label_To_Worker   chan messages.NoughtsAndCrosses_Label
	Player_To_Worker  chan impl.Player
}

type Worker_Chan struct {
	Board_From_Master  chan impl.Board
	Int_To_Master      chan int
	Label_From_Master  chan messages.NoughtsAndCrosses_Label
	Label_To_Master    chan messages.NoughtsAndCrosses_Label
	Player_From_Master chan impl.Player
}
