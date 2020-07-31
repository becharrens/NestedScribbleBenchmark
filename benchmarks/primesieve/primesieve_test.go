package primesieve

import (
	"ScribbleBenchmark/primesieve/protocol"
	"ScribbleBenchmark/primesieve_base"
	"testing"
)

func BenchmarkPrimeSieve(b *testing.B) {
	benchmarks := []struct {
		name   string
		ubound int
	}{
		{name: "100", ubound: 100},
		{name: "1100", ubound: 1100},
		{name: "2100", ubound: 2100},
		{name: "3100", ubound: 3100},
		{name: "4100", ubound: 4100},
		{name: "5100", ubound: 5100},
		{name: "6100", ubound: 6100},
		{name: "7100", ubound: 7100},
		{name: "8100", ubound: 8100},
		{name: "9100", ubound: 9100},
	}
	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				env := NewPrimeSieveEnv(bm.ubound)
				b.StartTimer()
				protocol.PrimeSieve(env)
			}
		})
	}
}

func BenchmarkPrimeSeiveBase(b *testing.B) {
	benchmarks := []struct {
		name   string
		ubound int
	}{
		{name: "100", ubound: 100},
		{name: "1100", ubound: 1100},
		{name: "2100", ubound: 2100},
		{name: "3100", ubound: 3100},
		{name: "4100", ubound: 4100},
		{name: "5100", ubound: 5100},
		{name: "6100", ubound: 6100},
		{name: "7100", ubound: 7100},
		{name: "8100", ubound: 8100},
		{name: "9100", ubound: 9100},
	}
	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				primesieve_base.PrimeSieve(bm.ubound)
			}
		})
	}
}
