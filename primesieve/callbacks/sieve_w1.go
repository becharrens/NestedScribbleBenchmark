package callbacks

import "NestedScribbleBenchmark/primesieve/messages/sieve"
import "NestedScribbleBenchmark/primesieve/results/sieve_sendnums"
import sieve_2 "NestedScribbleBenchmark/primesieve/results/sieve"

type Sieve_W1_Env interface {
	Done() sieve_2.W1_Result
	ResultFrom_Sieve_SendNums_S(result sieve_sendnums.S_Result)
	To_Sieve_SendNums_S_Env() Sieve_SendNums_S_Env
	Sieve_SendNums_Setup()
	FilterPrime_To_W2() sieve.FilterPrime
}

type SieveW1State struct {
	FilterPrime    int
	PossiblePrimes []int
}

func (s *SieveW1State) Done() sieve_2.W1_Result {
	return sieve_2.W1_Result{}
}

func (s *SieveW1State) ResultFrom_Sieve_SendNums_S(result sieve_sendnums.S_Result) {
}

func (s *SieveW1State) To_Sieve_SendNums_S_Env() Sieve_SendNums_S_Env {
	return &SieveSendNumsSState{
		NumsToSend: s.PossiblePrimes,
	}
}

func (s *SieveW1State) Sieve_SendNums_Setup() {
}

func (s *SieveW1State) FilterPrime_To_W2() sieve.FilterPrime {
	return sieve.FilterPrime{Int: s.FilterPrime}
}
