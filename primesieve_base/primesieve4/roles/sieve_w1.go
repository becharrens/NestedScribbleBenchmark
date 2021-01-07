package roles

import (
	"NestedScribbleBenchmark/primesieve_base/primesieve4/channels/sieve_sendnums"
	"NestedScribbleBenchmark/primesieve_base/primesieve4/messages"
)
import "NestedScribbleBenchmark/primesieve_base/primesieve4/channels/sieve"
import "NestedScribbleBenchmark/primesieve_base/primesieve4/callbacks"
import sieve_2 "NestedScribbleBenchmark/primesieve_base/primesieve4/results/sieve"
import "sync"

func Sieve_W1(wg *sync.WaitGroup, roleChannels sieve.W1_Chan, env callbacks.Sieve_W1_Env) sieve_2.W1_Result {
	int := env.FilterPrime_To_W2()
	roleChannels.Label_To_W2 <- messages.FilterPrime
	roleChannels.Int_To_W2 <- int

	env.Sieve_SendNums_Setup()

	roleChannels.Label_To_W2 <- messages.Sieve_SendNums_W1_W2

	sieve_sendnums_s_chan := sieve_sendnums.S_Chan{
		Int_To_R:   roleChannels.Int_To_W2,
		Label_To_R: roleChannels.Label_To_W2,
	}
	sieve_sendnums_s_env := env.To_Sieve_SendNums_S_Env()
	sieve_sendnums_s_result := Sieve_SendNums_S(wg, sieve_sendnums_s_chan, sieve_sendnums_s_env)
	env.ResultFrom_Sieve_SendNums_S(sieve_sendnums_s_result)

	return env.Done()
}
