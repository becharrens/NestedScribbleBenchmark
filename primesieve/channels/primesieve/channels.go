package primesieve

import "NestedScribbleBenchmark/primesieve/messages"

type Master_Chan struct {
	Int_From_Worker   chan int
	Int_To_Worker     chan int
	Label_From_Worker chan messages.PrimeSieve_Label
	Label_To_Worker   chan messages.PrimeSieve_Label
}

type Worker_Chan struct {
	Int_From_Master   chan int
	Int_To_Master     chan int
	Label_From_Master chan messages.PrimeSieve_Label
	Label_To_Master   chan messages.PrimeSieve_Label
}
