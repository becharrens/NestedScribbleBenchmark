package callbacks

import (
	"NestedScribbleBenchmark/primesieve/messages/primesieve"
	primesieve_2 "NestedScribbleBenchmark/primesieve/results/primesieve"
	"NestedScribbleBenchmark/primesieve/results/sieve"
)

type PrimeSieve_Master_Env interface {
	Finish_From_Worker(finish primesieve.Finish)
	Done() primesieve_2.Master_Result
	ResultFrom_Sieve_M(result sieve.M_Result)
	To_Sieve_M_Env() Sieve_M_Env
	Prime_From_Worker(prime primesieve.Prime)
	UBound_To_Worker() primesieve.UBound
	FirstPrime_To_Worker() primesieve.FirstPrime
}

type PrimeSieveMasterState struct {
	N      int
	Primes []int
}

func (p *PrimeSieveMasterState) FirstPrime_To_Worker() primesieve.FirstPrime {
	return primesieve.FirstPrime{
		Prime: 2,
	}
}

func (p *PrimeSieveMasterState) UBound_To_Worker() primesieve.UBound {
	return primesieve.UBound{N: p.N}
}

func (p *PrimeSieveMasterState) Prime_From_Worker(prime primesieve.Prime) {
	// fmt.Println("master: prime from worker", prime.N)
	p.Primes = append(p.Primes, prime.N)
}

func (p *PrimeSieveMasterState) To_Sieve_M_Env() Sieve_M_Env {
	return &SieveMState{
		Primes: p.Primes,
	}
}

func (p *PrimeSieveMasterState) ResultFrom_Sieve_M(result sieve.M_Result) {
	p.Primes = result.Primes
}

func (p *PrimeSieveMasterState) Finish_From_Worker(finish primesieve.Finish) {
}

func (p *PrimeSieveMasterState) Done() primesieve_2.Master_Result {
	return primesieve_2.Master_Result{Primes: p.Primes}
}
