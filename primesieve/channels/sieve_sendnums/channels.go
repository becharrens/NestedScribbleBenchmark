package sieve_sendnums

import "ScribbleBenchmark/primesieve/messages/sieve_sendnums"

type S_Chan struct {
	R_End chan sieve_sendnums.End
	R_Num chan sieve_sendnums.Num
}

type R_Chan struct {
	S_End chan sieve_sendnums.End
	S_Num chan sieve_sendnums.Num
}
