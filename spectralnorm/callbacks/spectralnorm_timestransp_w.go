package callbacks

type SpectralNorm_TimesTransp_W_Env interface {
	Finish_From_M()
	Done()
	TimesTranspResult_To_M() []float64
	TimesTranspTask_From_M(ii int, n int, u []float64, v []float64)
}

type SpectralNorm_TimesTranspWState struct {
	V []float64
}

func (s *SpectralNorm_TimesTranspWState) Finish_From_M() {
}

func (s *SpectralNorm_TimesTranspWState) Done() {
}

func (s *SpectralNorm_TimesTranspWState) TimesTranspResult_To_M() []float64 {
	return s.V
}

func (s *SpectralNorm_TimesTranspWState) TimesTranspTask_From_M(ii int, n int, u []float64, v []float64) {
	s.V = v
	TimesTransp(ii, n, s.V, u)
}

func New_SpectralNorm_TimesTransp_W_State() SpectralNorm_TimesTransp_W_Env {
	return &SpectralNorm_TimesTranspWState{}
}

func TimesTransp(ii, n int, v, u []float64) {
	ul := len(u)
	for i := ii; i < n; i++ {
		var vi float64
		for j := 0; j < ul; j++ {
			vi += u[j] / float64(A(j, i))
		}
		v[i] = vi
	}
}
