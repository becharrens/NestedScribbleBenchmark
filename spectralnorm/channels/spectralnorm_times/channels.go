package spectralnorm_times

import "NestedScribbleBenchmark/spectralnorm/messages"

type M_Chan struct {
	Int_To_W     chan int
	Label_From_W chan messages.SpectralNorm_Label
	Label_To_W   chan messages.SpectralNorm_Label
	Vec_From_W   chan []float64
	Vec_To_W     chan []float64
}

type W_Chan struct {
	Int_From_M   chan int
	Label_From_M chan messages.SpectralNorm_Label
	Label_To_M   chan messages.SpectralNorm_Label
	Vec_From_M   chan []float64
	Vec_To_M     chan []float64
}
