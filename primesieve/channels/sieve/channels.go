package sieve

import "NestedScribbleBenchmark/primesieve/messages"

type M_Chan struct {
	Int_From_W2   chan int
	Label_From_W2 chan messages.PrimeSieve_Label
}

type W1_Chan struct {
	Int_To_W2   chan int
	Label_To_W2 chan messages.PrimeSieve_Label
}

type W2_Chan struct {
	Int_From_W1   chan int
	Int_To_M      chan int
	Label_From_W1 chan messages.PrimeSieve_Label
	Label_To_M    chan messages.PrimeSieve_Label
}
