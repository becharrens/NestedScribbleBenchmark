package forward

import "NestedScribbleBenchmark/ring/messages"

type S_Chan struct {
	Int_To_RingNode    chan int
	Label_To_RingNode  chan messages.Ring_Label
	String_To_RingNode chan string
}

type E_Chan struct {
	Int_From_RingNode    chan int
	Label_From_RingNode  chan messages.Ring_Label
	String_From_RingNode chan string
}

type RingNode_Chan struct {
	Int_From_S    chan int
	Int_To_E      chan int
	Label_From_S  chan messages.Ring_Label
	Label_To_E    chan messages.Ring_Label
	String_From_S chan string
	String_To_E   chan string
}
