package dyntaskgen

import "NestedScribbleBenchmark/clientserver/messages"

type S_Chan struct {
	Label_From_W chan messages.ClientServer_Label
	Label_To_W chan messages.ClientServer_Label
	String_From_W chan string
	String_To_W chan string
}

type W_Chan struct {
	Label_From_S chan messages.ClientServer_Label
	Label_To_S chan messages.ClientServer_Label
	String_From_S chan string
	String_To_S chan string
}