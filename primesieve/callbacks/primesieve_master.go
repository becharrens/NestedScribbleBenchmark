package callbacks

import "NestedScribbleBenchmark/primesieve/results/sieve"
import "NestedScribbleBenchmark/primesieve/results/primesieve"

type PrimeSieve_Master_Env interface {
	Finish_From_Worker()
	Done() primesieve.Master_Result
	ResultFrom_Sieve_M(result sieve.M_Result)
	To_Sieve_M_Env() Sieve_M_Env
	Prime_From_Worker(n int)
	UBound_To_Worker() int
	FirstPrime_To_Worker() int
}

type PrimeSieveMasterState struct {
	N      int
	Primes []int
}

func (p *PrimeSieveMasterState) FirstPrime_To_Worker() int {
	return 2
}

func (p *PrimeSieveMasterState) UBound_To_Worker() int {
	return p.N
}

func (p *PrimeSieveMasterState) Prime_From_Worker(n int) {
	// fmt.Println("master: prime from worker", prime.N)
	p.Primes = append(p.Primes, n)
}

func (p *PrimeSieveMasterState) To_Sieve_M_Env() Sieve_M_Env {
	return &SieveMState{
		Primes: p.Primes,
	}
}

func (p *PrimeSieveMasterState) ResultFrom_Sieve_M(result sieve.M_Result) {
	p.Primes = result.Primes
}

func (p *PrimeSieveMasterState) Finish_From_Worker() {
}

func (p *PrimeSieveMasterState) Done() primesieve.Master_Result {
	return primesieve.Master_Result{Primes: p.Primes}
}
