package fannkuch

import "NestedScribbleBenchmark/fannkuch/messages"

type Main_Chan struct {
	Int_From_Worker   chan int
	Int_To_Worker     chan int
	Label_From_Worker chan messages.Fannkuch_Label
	Label_To_Worker   chan messages.Fannkuch_Label
}

type Worker_Chan struct {
	Int_From_Main   chan int
	Int_To_Main     chan int
	Label_From_Main chan messages.Fannkuch_Label
	Label_To_Main   chan messages.Fannkuch_Label
}
