package callbacks

import (
	"NestedScribbleBenchmark/primesieve/messages/sieve"
	sieve_2 "NestedScribbleBenchmark/primesieve/results/sieve"
)

type Sieve_M_Env interface {
	Finish_From_W2(finish sieve.Finish)
	Done() sieve_2.M_Result
	ResultFrom_Sieve_M(result sieve_2.M_Result)
	To_Sieve_M_Env() Sieve_M_Env
	Prime_From_W2(prime sieve.Prime)
}

type SieveMState struct {
	Primes   []int
	StartNum int
	EndNum   int
}

func (s *SieveMState) Prime_From_W2(prime sieve.Prime) {
	// fmt.Println("m: StartNum ", s.StartNum)
	s.Primes = append(s.Primes, prime.N)
}

func (s *SieveMState) To_Sieve_M_Env() Sieve_M_Env {
	return &SieveMState{
		Primes:   s.Primes,
		StartNum: s.StartNum + 1,
	}
}

func (s *SieveMState) ResultFrom_Sieve_M(result sieve_2.M_Result) {
	s.Primes = result.Primes
}

func (s *SieveMState) Finish_From_W2(finish sieve.Finish) {
	// fmt.Println("m: StartNum ", s.StartNum)
}

func (s *SieveMState) Done() sieve_2.M_Result {
	// fmt.Println("m: EndNum ", s.StartNum)
	return sieve_2.M_Result{Primes: s.Primes}
}
