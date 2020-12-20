package fannkuchrecursive

import "NestedScribbleBenchmark/fannkuch/messages"

type Source_Chan struct {
	Int_From_NewWorker   chan int
	Label_From_NewWorker chan messages.Fannkuch_Label
}

type Worker_Chan struct {
	Int_To_NewWorker   chan int
	Label_To_NewWorker chan messages.Fannkuch_Label
}

type NewWorker_Chan struct {
	Int_From_Worker   chan int
	Int_To_Source     chan int
	Label_From_Worker chan messages.Fannkuch_Label
	Label_To_Source   chan messages.Fannkuch_Label
}
