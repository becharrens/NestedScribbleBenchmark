package primesieve

import "ScribbleBenchmark/primesieve/messages/primesieve"

type Master_Chan struct {
	Worker_Finish     chan primesieve.Finish
	Worker_FirstPrime chan primesieve.FirstPrime
	Worker_Prime      chan primesieve.Prime
	Worker_UBound     chan primesieve.UBound
}

type Worker_Chan struct {
	Master_Finish     chan primesieve.Finish
	Master_FirstPrime chan primesieve.FirstPrime
	Master_Prime      chan primesieve.Prime
	Master_UBound     chan primesieve.UBound
}
