package callbacks

import "NestedScribbleBenchmark/primesieve_base/primesieve4/results/sieve_sendnums"

type Sieve_SendNums_S_Choice_2 int

const (
	Sieve_SendNums_S_Num_2 Sieve_SendNums_S_Choice_2 = iota
	Sieve_SendNums_S_End_2
)

type Sieve_SendNums_S_Choice int

const (
	Sieve_SendNums_S_Num Sieve_SendNums_S_Choice = iota
	Sieve_SendNums_S_End
)

type Sieve_SendNums_S_Env interface {
	End_To_R_2()
	Done() sieve_sendnums.S_Result
	End_To_R()
	Num_To_R_2() int
	S_Choice_2() Sieve_SendNums_S_Choice_2
	Num_To_R() int
	S_Choice() Sieve_SendNums_S_Choice
}

type SieveSendNumsSState struct {
	NumsToSend []int
	Idx        int
}

func (s *SieveSendNumsSState) End_To_R_2() {
}

func (s *SieveSendNumsSState) Done() sieve_sendnums.S_Result {
	return sieve_sendnums.S_Result{}
}

func (s *SieveSendNumsSState) End_To_R() {
}

func (s *SieveSendNumsSState) Num_To_R_2() int {
	result := s.NumsToSend[s.Idx]
	s.Idx++
	return result
}

func (s *SieveSendNumsSState) S_Choice_2() Sieve_SendNums_S_Choice_2 {
	// fmt.Println("s: len nums_to_send: ", len(s.NumsToSend))
	if s.Idx >= len(s.NumsToSend) {
		return Sieve_SendNums_S_End_2
	}
	return Sieve_SendNums_S_Num_2
}

func (s *SieveSendNumsSState) Num_To_R() int {
	result := s.NumsToSend[s.Idx]
	s.Idx++
	return result
}

func (s *SieveSendNumsSState) S_Choice() Sieve_SendNums_S_Choice {
	// fmt.Println("s: len nums_to_send: ", len(s.NumsToSend))
	if s.Idx >= len(s.NumsToSend) {
		return Sieve_SendNums_S_End
	}
	return Sieve_SendNums_S_Num
}
