package noughtsandcrosses

import "NestedScribbleBenchmark/noughtsandcrosses/messages"

type P1_Chan struct {
	Int_From_P2 chan int
	Int_To_P2 chan int
	Label_From_P2 chan messages.NoughtsAndCrosses_Label
	Label_To_P2 chan messages.NoughtsAndCrosses_Label
}

type P2_Chan struct {
	Int_From_P1 chan int
	Int_To_P1 chan int
	Label_From_P1 chan messages.NoughtsAndCrosses_Label
	Label_To_P1 chan messages.NoughtsAndCrosses_Label
}