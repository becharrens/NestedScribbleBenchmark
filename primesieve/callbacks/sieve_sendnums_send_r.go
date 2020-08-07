package callbacks

import (
	"NestedScribbleBenchmark/primesieve/messages/sieve_sendnums_send"
	sieve_sendnums_send_2 "NestedScribbleBenchmark/primesieve/results/sieve_sendnums_send"
)

type Sieve_SendNums_SEND_R_Env interface {
	End_From_S(end sieve_sendnums_send.End)
	Done() sieve_sendnums_send_2.R_Result
	ResultFrom_Sieve_SendNums_SEND_R(result sieve_sendnums_send_2.R_Result)
	To_Sieve_SendNums_SEND_R_Env() Sieve_SendNums_SEND_R_Env
	Sieve_SendNums_SEND_Setup()
	Num_From_S(num sieve_sendnums_send.Num)
}

type SieveSendNumsSENDRState struct {
	FilterPrime  int
	NumsReceived []int
}

func (s *SieveSendNumsSENDRState) End_From_S(end sieve_sendnums_send.End) {
}

func (s *SieveSendNumsSENDRState) Done() sieve_sendnums_send_2.R_Result {
	return sieve_sendnums_send_2.R_Result{ReceivedNums: s.NumsReceived}
}

func (s *SieveSendNumsSENDRState) ResultFrom_Sieve_SendNums_SEND_R(result sieve_sendnums_send_2.R_Result) {
	s.NumsReceived = result.ReceivedNums
}

func (s *SieveSendNumsSENDRState) To_Sieve_SendNums_SEND_R_Env() Sieve_SendNums_SEND_R_Env {
	return s
}

func (s *SieveSendNumsSENDRState) Sieve_SendNums_SEND_Setup() {
}

func (s *SieveSendNumsSENDRState) Num_From_S(num sieve_sendnums_send.Num) {
	if num.N%s.FilterPrime != 0 {
		s.NumsReceived = append(s.NumsReceived, num.N)
	}
	// fmt.Println("send_r: len nums_received", len(s.NumsReceived))
}
