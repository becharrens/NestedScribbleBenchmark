package messages

type PrimeSieve_Label int

const (
	End PrimeSieve_Label = iota
	FilterPrime
	Finish
	FirstPrime
	Num
	Prime
	Sieve_M_W2
	Sieve_Master_Worker
	Sieve_SendNums_W1_W2
	UBound
)
