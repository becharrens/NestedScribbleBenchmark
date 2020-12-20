package callbacks

import "NestedScribbleBenchmark/spectralnorm/results/spectralnorm_timestransp"

type SpectralNorm_TimesTransp_M_Choice int

const (
	SpectralNorm_TimesTransp_M_TimesTranspTask SpectralNorm_TimesTransp_M_Choice = iota
	SpectralNorm_TimesTransp_M_Finish
)

type SpectralNorm_TimesTransp_M_Env interface {
	Finish_To_W()
	Done() spectralnorm_timestransp.M_Result
	TimesTranspResult_From_W(res []float64)
	ResultFrom_SpectralNorm_TimesTransp_M(result spectralnorm_timestransp.M_Result)
	To_SpectralNorm_TimesTransp_M_Env() SpectralNorm_TimesTransp_M_Env
	SpectralNorm_TimesTransp_Setup()
	TimesTranspTask_To_W() (int, int, []float64, []float64)
	M_Choice() SpectralNorm_TimesTransp_M_Choice
}

type SpectralNorm_TimesTranspMState struct {
	U    []float64
	V    []float64
	I    int
	NCPU int
}

func (s *SpectralNorm_TimesTranspMState) Finish_To_W() {

}

func (s *SpectralNorm_TimesTranspMState) Done() spectralnorm_timestransp.M_Result {
	return spectralnorm_timestransp.M_Result{Vec: s.V}
}

func (s *SpectralNorm_TimesTranspMState) TimesTranspResult_From_W(res []float64) {
	// s.V = timestranspresult_msg.Res
}

func (s *SpectralNorm_TimesTranspMState) ResultFrom_SpectralNorm_TimesTransp_M(result spectralnorm_timestransp.M_Result) {
	// s.V = result.Vec
}

func (s *SpectralNorm_TimesTranspMState) To_SpectralNorm_TimesTransp_M_Env() SpectralNorm_TimesTransp_M_Env {
	return &SpectralNorm_TimesTranspMState{
		U:    s.U,
		V:    s.V,
		I:    s.I + 1,
		NCPU: s.NCPU,
	}
}

func (s *SpectralNorm_TimesTranspMState) SpectralNorm_TimesTransp_Setup() {
}

func (s *SpectralNorm_TimesTranspMState) TimesTranspTask_To_W() (int, int, []float64, []float64) {
	return s.I * len(s.V) / s.NCPU, (s.I + 1) * len(s.V) / s.NCPU, s.U, s.V
}

func (s *SpectralNorm_TimesTranspMState) M_Choice() SpectralNorm_TimesTransp_M_Choice {
	if s.I < s.NCPU {
		return SpectralNorm_TimesTransp_M_TimesTranspTask
	}
	return SpectralNorm_TimesTransp_M_Finish
}
