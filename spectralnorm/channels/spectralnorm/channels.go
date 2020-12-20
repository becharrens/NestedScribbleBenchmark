package spectralnorm

import "NestedScribbleBenchmark/spectralnorm/messages"

type Master_Chan struct {
	Int_To_Worker     chan int
	Label_From_Worker chan messages.SpectralNorm_Label
	Label_To_Worker   chan messages.SpectralNorm_Label
	Vec_From_Worker   chan []float64
	Vec_To_Worker     chan []float64
}

type Worker_Chan struct {
	Int_From_Master   chan int
	Label_From_Master chan messages.SpectralNorm_Label
	Label_To_Master   chan messages.SpectralNorm_Label
	Vec_From_Master   chan []float64
	Vec_To_Master     chan []float64
}
