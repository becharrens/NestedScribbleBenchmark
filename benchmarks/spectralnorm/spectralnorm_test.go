package spectralnorm

import (
	"ScribbleBenchmark/spectralnorm/protocol"
	"ScribbleBenchmark/spectralnorm_base"
	"testing"
)

func BenchmarkSpectralNorm(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{name: "100", n: 100},
		{name: "500", n: 500},
		{name: "1500", n: 1500},
		{name: "2500", n: 2500},
		{name: "3500", n: 3500},
		{name: "4500", n: 4500},
		{name: "5500", n: 5500},
	}
	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				env := NewSpectralNormEnv(bm.n)
				b.StartTimer()
				protocol.SpectralNorm(env)
			}
		})
	}
}

func BenchmarkSpectralNormBase(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{name: "100", n: 100},
		{name: "500", n: 500},
		{name: "1500", n: 1500},
		{name: "2500", n: 2500},
		{name: "3500", n: 3500},
		{name: "4500", n: 4500},
		{name: "5500", n: 5500},
	}
	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				spectralnorm_base.SpectralNorm(bm.n)
			}
		})
	}
}
