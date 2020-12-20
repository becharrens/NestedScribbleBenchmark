package callbacks

type SpectralNorm_Times_W_Env interface {
	Finish_From_M()
	Done()
	TimesResult_To_M() []float64
	TimesTask_From_M(ii int, n int, u []float64, v []float64)
}

type SpectralNorm_TimesWState struct {
	V []float64
}

func (s *SpectralNorm_TimesWState) Finish_From_M() {
}

func (s *SpectralNorm_TimesWState) Done() {
}

func (s *SpectralNorm_TimesWState) TimesResult_To_M() []float64 {
	return s.V
}

func (s *SpectralNorm_TimesWState) TimesTask_From_M(ii int, n int, u []float64, v []float64) {
	s.V = v
	Times(ii, n, s.V, u)
}

func New_SpectralNorm_Times_W_State() SpectralNorm_Times_W_Env {
	return &SpectralNorm_TimesWState{}
}
