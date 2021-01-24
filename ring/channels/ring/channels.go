package ring

import "NestedScribbleBenchmark/ring/messages"

type Start_Chan struct {
	Int_From_End    chan int
	Int_To_End      chan int
	Label_From_End  chan messages.Ring_Label
	Label_To_End    chan messages.Ring_Label
	String_From_End chan string
	String_To_End   chan string
}

type End_Chan struct {
	Int_From_Start    chan int
	Int_To_Start      chan int
	Label_From_Start  chan messages.Ring_Label
	Label_To_Start    chan messages.Ring_Label
	String_From_Start chan string
	String_To_Start   chan string
}
