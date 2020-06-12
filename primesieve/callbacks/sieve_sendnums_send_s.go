package callbacks

import (
	"ScribbleBenchmark/primesieve/messages/sieve_sendnums_send"
	sieve_sendnums_send_2 "ScribbleBenchmark/primesieve/results/sieve_sendnums_send"
)

type Sieve_SendNums_SEND_S_Choice int

const (
	Sieve_SendNums_SEND_S_Num Sieve_SendNums_SEND_S_Choice = iota
	Sieve_SendNums_SEND_S_End
)

type Sieve_SendNums_SEND_S_Env interface {
	End_To_R() sieve_sendnums_send.End
	Done() sieve_sendnums_send_2.S_Result
	ResultFrom_Sieve_SendNums_SEND_S(result sieve_sendnums_send_2.S_Result)
	To_Sieve_SendNums_SEND_S_Env() Sieve_SendNums_SEND_S_Env
	Num_To_R() sieve_sendnums_send.Num
	S_Choice() Sieve_SendNums_SEND_S_Choice
}

type SieveSendNumsSSENDState struct {
	NumsToSend []int
}

func (s *SieveSendNumsSSENDState) End_To_R() sieve_sendnums_send.End {
	return sieve_sendnums_send.End{}
}

func (s *SieveSendNumsSSENDState) Done() sieve_sendnums_send_2.S_Result {
	return sieve_sendnums_send_2.S_Result{}
}

func (s *SieveSendNumsSSENDState) ResultFrom_Sieve_SendNums_SEND_S(result sieve_sendnums_send_2.S_Result) {
}

func (s *SieveSendNumsSSENDState) To_Sieve_SendNums_SEND_S_Env() Sieve_SendNums_SEND_S_Env {
	return &SieveSendNumsSSENDState{NumsToSend: s.NumsToSend[1:]}
}

func (s *SieveSendNumsSSENDState) Num_To_R() sieve_sendnums_send.Num {
	return sieve_sendnums_send.Num{N: s.NumsToSend[0]}
}

func (s *SieveSendNumsSSENDState) S_Choice() Sieve_SendNums_SEND_S_Choice {
	// fmt.Println("send_s: len nums_to_send: ", len(s.NumsToSend))
	if len(s.NumsToSend) < 1 {
		return Sieve_SendNums_SEND_S_End
	}
	return Sieve_SendNums_SEND_S_Num
}
