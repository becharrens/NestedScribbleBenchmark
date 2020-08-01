package primesieve

import (
	"NestedScribbleBenchmark/primesieve/callbacks"
	"NestedScribbleBenchmark/primesieve/results/primesieve"
)

type PrimeSieveEnv struct {
	N      int
	Primes []int
}

func (p *PrimeSieveEnv) New_Master_Env() callbacks.PrimeSieve_Master_Env {
	return &callbacks.PrimeSieveMasterState{
		N:      p.N,
		Primes: []int{2},
	}
}

func (p *PrimeSieveEnv) New_Worker_Env() callbacks.PrimeSieve_Worker_Env {
	return &callbacks.PrimeSieveWorkerState{}
}

func (p *PrimeSieveEnv) Master_Result(result primesieve.Master_Result) {
	p.Primes = result.Primes
}

func (p *PrimeSieveEnv) Worker_Result(result primesieve.Worker_Result) {
}

func NewPrimeSieveEnv(n int) *PrimeSieveEnv {
	return &PrimeSieveEnv{
		N:      n,
		Primes: nil,
	}
}
