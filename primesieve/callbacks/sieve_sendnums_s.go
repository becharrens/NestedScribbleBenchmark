package callbacks

import (
	"NestedScribbleBenchmark/primesieve/messages/sieve_sendnums"
	sieve_sendnums_2 "NestedScribbleBenchmark/primesieve/results/sieve_sendnums"
	"NestedScribbleBenchmark/primesieve/results/sieve_sendnums_send"
)

type Sieve_SendNums_S_Choice int

const (
	Sieve_SendNums_S_Num Sieve_SendNums_S_Choice = iota
	Sieve_SendNums_S_End
)

type Sieve_SendNums_S_Env interface {
	End_To_R() sieve_sendnums.End
	Done() sieve_sendnums_2.S_Result
	ResultFrom_Sieve_SendNums_SEND_S(result sieve_sendnums_send.S_Result)
	To_Sieve_SendNums_SEND_S_Env() Sieve_SendNums_SEND_S_Env
	Num_To_R() sieve_sendnums.Num
	S_Choice() Sieve_SendNums_S_Choice
}

type SieveSendNumsSState struct {
	NumsToSend []int
}

func (s *SieveSendNumsSState) ResultFrom_Sieve_SendNums_SEND_S(result sieve_sendnums_send.S_Result) {
}

func (s *SieveSendNumsSState) To_Sieve_SendNums_SEND_S_Env() Sieve_SendNums_SEND_S_Env {
	return &SieveSendNumsSSENDState{
		NumsToSend: s.NumsToSend[1:],
	}
}

func (s *SieveSendNumsSState) End_To_R_2() sieve_sendnums.End {
	return sieve_sendnums.End{}
}

func (s *SieveSendNumsSState) Done() sieve_sendnums_2.S_Result {
	return sieve_sendnums_2.S_Result{}
}

func (s *SieveSendNumsSState) End_To_R() sieve_sendnums.End {
	return sieve_sendnums.End{}
}

func (s *SieveSendNumsSState) Num_To_R() sieve_sendnums.Num {
	result := s.NumsToSend[0]
	return sieve_sendnums.Num{N: result}
}

func (s *SieveSendNumsSState) S_Choice() Sieve_SendNums_S_Choice {
	// fmt.Println("s: len nums_to_send: ", len(s.NumsToSend))
	if len(s.NumsToSend) < 1 {
		return Sieve_SendNums_S_End
	}
	return Sieve_SendNums_S_Num
}
