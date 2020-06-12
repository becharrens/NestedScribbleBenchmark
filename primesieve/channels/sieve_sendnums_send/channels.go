package sieve_sendnums_send

import "ScribbleBenchmark/primesieve/messages/sieve_sendnums_send"

type R_Chan struct {
	S_End chan sieve_sendnums_send.End
	S_Num chan sieve_sendnums_send.Num
}

type S_Chan struct {
	R_End chan sieve_sendnums_send.End
	R_Num chan sieve_sendnums_send.Num
}
