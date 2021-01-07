package callbacks

type Sieve_SendNums_R_Env interface {
	End_From_S_2()
	Done() []int
	End_From_S()
	Num_From_S_2(n int)
	Num_From_S(n int)
}

type SieveSendNumsRState struct {
	FilterPrime  int
	NumsReceived []int
}

func (s *SieveSendNumsRState) End_From_S_2() {
}

func (s *SieveSendNumsRState) Done() []int {
	return s.NumsReceived
}

func (s *SieveSendNumsRState) End_From_S() {
}

func (s *SieveSendNumsRState) Num_From_S_2(n int) {
	if n%s.FilterPrime != 0 {
		s.NumsReceived = append(s.NumsReceived, n)
	}
}

func (s *SieveSendNumsRState) Num_From_S(n int) {
	// s.Num = num.N
	if n%s.FilterPrime != 0 {
		s.NumsReceived = append(s.NumsReceived, n)
	}
	// fmt.Println("r: len nums_received", len(s.NumsReceived))

}
