package callbacks

import "NestedScribbleBenchmark/primesieve_base/primesieve4/results/sieve_sendnums"
import "NestedScribbleBenchmark/primesieve_base/primesieve4/results/sieve"

type Sieve_W2_Choice int

const (
	Sieve_W2_Prime Sieve_W2_Choice = iota
	Sieve_W2_Finish
)

type Sieve_W2_Env interface {
	Finish_To_M()
	Done()
	ResultFrom_Sieve_W1(result sieve.W1_Result)
	To_Sieve_W1_Env() Sieve_W1_Env
	Sieve_Setup()
	Prime_To_M() int
	W2_Choice() Sieve_W2_Choice
	ResultFrom_Sieve_SendNums_R(result sieve_sendnums.R_Result)
	To_Sieve_SendNums_R_Env() Sieve_SendNums_R_Env
	FilterPrime_From_W1(int int)
}

type SieveW2State struct {
	FilterPrime    int
	PossiblePrimes []int
}

func (s *SieveW2State) Finish_To_M() {
}

func (s *SieveW2State) Done() {
}

func (s *SieveW2State) ResultFrom_Sieve_W1(result sieve.W1_Result) {
}

func (s *SieveW2State) To_Sieve_W1_Env() Sieve_W1_Env {
	return &SieveW1State{
		FilterPrime:    s.PossiblePrimes[0],
		PossiblePrimes: s.PossiblePrimes[1:],
	}
}

func (s *SieveW2State) Sieve_Setup() {
}

func (s *SieveW2State) Prime_To_M() int {
	// fmt.Println("w2: prime to m", s.PossiblePrimes[0])
	return s.PossiblePrimes[0]
}

func (s *SieveW2State) W2_Choice() Sieve_W2_Choice {
	if len(s.PossiblePrimes) == 0 {
		return Sieve_W2_Finish
	}
	return Sieve_W2_Prime
}

func (s *SieveW2State) ResultFrom_Sieve_SendNums_R(result sieve_sendnums.R_Result) {
	s.PossiblePrimes = result.ReceivedNums
}

func (s *SieveW2State) To_Sieve_SendNums_R_Env() Sieve_SendNums_R_Env {
	return &SieveSendNumsRState{FilterPrime: s.FilterPrime}
}

func (s *SieveW2State) FilterPrime_From_W1(filterprime int) {
	s.FilterPrime = filterprime
}

func New_Sieve_W2_State() Sieve_W2_Env {
	return &SieveW2State{}
}
