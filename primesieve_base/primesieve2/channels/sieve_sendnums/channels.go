package sieve_sendnums

import "NestedScribbleBenchmark/primesieve_base/primesieve2/messages"

type S_Chan struct {
	Int_To_R   chan int
	Label_To_R chan messages.PrimeSieve_Label
}

type R_Chan struct {
	Int_From_S   chan int
	Label_From_S chan messages.PrimeSieve_Label
}
