package sieve

import "ScribbleBenchmark/primesieve/messages/sieve"

type M_Chan struct {
	W2_Finish chan sieve.Finish
	W2_Prime  chan sieve.Prime
}

type W1_Chan struct {
	W2_FilterPrime chan sieve.FilterPrime
}

type W2_Chan struct {
	M_Finish       chan sieve.Finish
	M_Prime        chan sieve.Prime
	W1_FilterPrime chan sieve.FilterPrime
}
