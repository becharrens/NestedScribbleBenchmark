package callbacks

type Sieve_W1_Env interface {
	Done()
	To_Sieve_SendNums_S_Env() Sieve_SendNums_S_Env
	Sieve_SendNums_Setup()
	FilterPrime_To_W2() int
}

type SieveW1State struct {
	FilterPrime    int
	PossiblePrimes []int
}

func (s *SieveW1State) Done() {
}

func (s *SieveW1State) To_Sieve_SendNums_S_Env() Sieve_SendNums_S_Env {
	return &SieveSendNumsSState{
		NumsToSend: s.PossiblePrimes,
	}
}

func (s *SieveW1State) Sieve_SendNums_Setup() {
}

func (s *SieveW1State) FilterPrime_To_W2() int {
	return s.FilterPrime
}
