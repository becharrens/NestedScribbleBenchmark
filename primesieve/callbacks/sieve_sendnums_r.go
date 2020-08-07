package callbacks

import (
	"NestedScribbleBenchmark/primesieve/messages/sieve_sendnums"
	sieve_sendnums_2 "NestedScribbleBenchmark/primesieve/results/sieve_sendnums"
	"NestedScribbleBenchmark/primesieve/results/sieve_sendnums_send"
)

type Sieve_SendNums_R_Env interface {
	End_From_S(end sieve_sendnums.End)
	Done() sieve_sendnums_2.R_Result
	ResultFrom_Sieve_SendNums_SEND_R(result sieve_sendnums_send.R_Result)
	To_Sieve_SendNums_SEND_R_Env() Sieve_SendNums_SEND_R_Env
	Sieve_SendNums_SEND_Setup()
	Num_From_S(num sieve_sendnums.Num)
}

type SieveSendNumsRState struct {
	FilterPrime  int
	NumsReceived []int
}

func (s *SieveSendNumsRState) Done() sieve_sendnums_2.R_Result {
	return sieve_sendnums_2.R_Result{ReceivedNums: s.NumsReceived}
}

func (s *SieveSendNumsRState) End_From_S(end sieve_sendnums.End) {
}

func (s *SieveSendNumsRState) Num_From_S(num sieve_sendnums.Num) {
	// s.Num = num.N
	s.FilterPrime = num.N
	s.NumsReceived = append(s.NumsReceived, num.N)
	// fmt.Println("r: len nums_received", len(s.NumsReceived))

}

func (s *SieveSendNumsRState) ResultFrom_Sieve_SendNums_SEND_R(result sieve_sendnums_send.R_Result) {
	// s.NumsReceived = prepend(s.Num, result.ReceivedNums)
	s.NumsReceived = result.ReceivedNums
	// fmt.Println("r: len receivedNums", len(s.NumsReceived))
}

func (s *SieveSendNumsRState) To_Sieve_SendNums_SEND_R_Env() Sieve_SendNums_SEND_R_Env {
	return &SieveSendNumsSENDRState{
		FilterPrime:  s.FilterPrime,
		NumsReceived: s.NumsReceived,
	}
}

func (s *SieveSendNumsRState) Sieve_SendNums_SEND_Setup() {
}
