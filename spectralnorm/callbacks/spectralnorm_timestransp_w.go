package callbacks

import "NestedScribbleBenchmark/spectralnorm/messages/spectralnorm_timestransp"

type SpectralNorm_TimesTransp_W_Env interface {
	Finish_From_M(finish_msg spectralnorm_timestransp.Finish)
	Done()
	TimesTranspResult_To_M() spectralnorm_timestransp.TimesTranspResult
	TimesTranspTask_From_M(timestransptask_msg spectralnorm_timestransp.TimesTranspTask)
}

type SpectralNorm_TimesTranspWState struct {
	V []float64
}

func (s *SpectralNorm_TimesTranspWState) Finish_From_M(finish_msg spectralnorm_timestransp.Finish) {
}

func (s *SpectralNorm_TimesTranspWState) Done() {
}

func (s *SpectralNorm_TimesTranspWState) TimesTranspResult_To_M() spectralnorm_timestransp.TimesTranspResult {
	return spectralnorm_timestransp.TimesTranspResult{Res: s.V}
}

func (s *SpectralNorm_TimesTranspWState) TimesTranspTask_From_M(timestransptask_msg spectralnorm_timestransp.TimesTranspTask) {
	s.V = timestransptask_msg.V
	TimesTransp(timestransptask_msg.Ii, timestransptask_msg.N, s.V, timestransptask_msg.U)
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
