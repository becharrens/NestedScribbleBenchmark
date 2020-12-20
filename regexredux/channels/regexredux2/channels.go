package regexredux2

import "NestedScribbleBenchmark/regexredux/messages"

type M_Chan struct {
	ByteArr_To_W chan []byte
	Int_From_W   chan int
	Label_From_W chan messages.RegexRedux_Label
	Label_To_W   chan messages.RegexRedux_Label
	String_To_W  chan string
}

type W_Chan struct {
	ByteArr_From_M chan []byte
	Int_To_M       chan int
	Label_From_M   chan messages.RegexRedux_Label
	Label_To_M     chan messages.RegexRedux_Label
	String_From_M  chan string
}
