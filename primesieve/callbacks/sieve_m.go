package callbacks

import "NestedScribbleBenchmark/primesieve/results/sieve"

type Sieve_M_Env interface {
	Finish_From_W2()
	Done() sieve.M_Result
	ResultFrom_Sieve_M(result sieve.M_Result)
	To_Sieve_M_Env() Sieve_M_Env
	Prime_From_W2(n int)
}

type SieveMState struct {
	Primes   []int
	StartNum int
	EndNum   int
}

func (s *SieveMState) Prime_From_W2(n int) {
	// fmt.Println("m: StartNum ", s.StartNum)
	s.Primes = append(s.Primes, n)
}

func (s *SieveMState) To_Sieve_M_Env() Sieve_M_Env {
	return &SieveMState{
		Primes:   s.Primes,
		StartNum: s.StartNum + 1,
	}
}

func (s *SieveMState) ResultFrom_Sieve_M(result sieve.M_Result) {
	s.Primes = result.Primes
}

func (s *SieveMState) Finish_From_W2() {
	// fmt.Println("m: StartNum ", s.StartNum)
}

func (s *SieveMState) Done() sieve.M_Result {
	// fmt.Println("m: EndNum ", s.StartNum)
	return sieve.M_Result{Primes: s.Primes}
}
