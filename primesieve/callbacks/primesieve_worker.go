package callbacks

import (
	"NestedScribbleBenchmark/primesieve/messages/primesieve"
	"NestedScribbleBenchmark/primesieve/results/sieve"

	primesieve_2 "NestedScribbleBenchmark/primesieve/results/primesieve"
)

type PrimeSieve_Worker_Choice int

const (
	PrimeSieve_Worker_Prime PrimeSieve_Worker_Choice = iota
	PrimeSieve_Worker_Finish
)

type PrimeSieve_Worker_Env interface {
	Finish_To_Master() primesieve.Finish
	Done() primesieve_2.Worker_Result
	ResultFrom_Sieve_W1(result sieve.W1_Result)
	To_Sieve_W1_Env() Sieve_W1_Env
	Sieve_Setup()
	Prime_To_Master() primesieve.Prime
	Worker_Choice() PrimeSieve_Worker_Choice
	UBound_From_Master(ubound primesieve.UBound)
	FirstPrime_From_Master(firstprime primesieve.FirstPrime)
}

type PrimeSieveWorkerState struct {
	FirstPrime     int
	UBound         int
	PossiblePrimes []int
}

func (p *PrimeSieveWorkerState) FirstPrime_From_Master(firstprime primesieve.FirstPrime) {
	p.FirstPrime = firstprime.Prime
}

func (p *PrimeSieveWorkerState) Worker_Choice() PrimeSieve_Worker_Choice {
	if len(p.PossiblePrimes) == 0 {
		return PrimeSieve_Worker_Finish
	}
	return PrimeSieve_Worker_Prime
}

func (p *PrimeSieveWorkerState) UBound_From_Master(ubound primesieve.UBound) {
	p.UBound = ubound.N
	p.PossiblePrimes = initPossiblePrimes(p.FirstPrime, p.UBound)
}

func (p *PrimeSieveWorkerState) Prime_To_Master() primesieve.Prime {
	return primesieve.Prime{N: p.PossiblePrimes[0]}
}

func (p *PrimeSieveWorkerState) Sieve_Setup() {
}

func (p *PrimeSieveWorkerState) To_Sieve_W1_Env() Sieve_W1_Env {
	return &SieveW1State{
		FilterPrime:    p.PossiblePrimes[0],
		PossiblePrimes: p.PossiblePrimes[1:],
	}
}

func (p *PrimeSieveWorkerState) ResultFrom_Sieve_W1(result sieve.W1_Result) {
}

func (p *PrimeSieveWorkerState) Finish_To_Master() primesieve.Finish {
	return primesieve.Finish{}
}

func (p *PrimeSieveWorkerState) Done() primesieve_2.Worker_Result {
	return primesieve_2.Worker_Result{}
}

func initPossiblePrimes(firstPrime, ubound int) []int {
	var result []int
	for i := 2; i < ubound; i++ {
		if i%firstPrime > 0 {
			result = append(result, i)
		}
	}
	// fmt.Println("worker:", len(result), firstPrime, ubound)
	return result
}
