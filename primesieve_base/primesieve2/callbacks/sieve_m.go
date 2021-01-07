package callbacks

type Sieve_M_Env interface {
	Finish_From_W2()
	Done() []int
	ResultFrom_Sieve_M(result []int)
	To_Sieve_M_Env() Sieve_M_Env
	Prime_From_W2(n int)
}

type SieveMState struct {
	Primes []int
	// StartNum int
	EndNum int
}

func (s *SieveMState) Prime_From_W2(n int) {
	// fmt.Println("m: StartNum ", s.StartNum)
	s.Primes = append(s.Primes, n)
}

func (s *SieveMState) ResultFrom_Sieve_M(result []int) {
	// s.Primes = result
}

func (s *SieveMState) To_Sieve_M_Env() Sieve_M_Env {
	return s
}

func (s *SieveMState) Finish_From_W2() {
	// fmt.Println("m: StartNum ", s.StartNum)
}

func (s *SieveMState) Done() []int {
	// fmt.Println("m: EndNum ", s.StartNum)
	return s.Primes
}
