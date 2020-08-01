package callbacks

import "NestedScribbleBenchmark/spectralnorm/messages/spectralnorm_timestransp"
import spectralnorm_timestransp_2 "NestedScribbleBenchmark/spectralnorm/results/spectralnorm_timestransp"

type SpectralNorm_TimesTransp_M_Choice int

const (
	SpectralNorm_TimesTransp_M_TimesTranspTask SpectralNorm_TimesTransp_M_Choice = iota
	SpectralNorm_TimesTransp_M_Finish
)

type SpectralNorm_TimesTransp_M_Env interface {
	Finish_To_W() spectralnorm_timestransp.Finish
	Done() spectralnorm_timestransp_2.M_Result
	TimesTranspResult_From_W(timestranspresult_msg spectralnorm_timestransp.TimesTranspResult)
	ResultFrom_SpectralNorm_TimesTransp_M(result spectralnorm_timestransp_2.M_Result)
	To_SpectralNorm_TimesTransp_M_Env() SpectralNorm_TimesTransp_M_Env
	SpectralNorm_TimesTransp_Setup()
	TimesTranspTask_To_W() spectralnorm_timestransp.TimesTranspTask
	M_Choice() SpectralNorm_TimesTransp_M_Choice
}

type SpectralNorm_TimesTranspMState struct {
	U    []float64
	V    []float64
	I    int
	NCPU int
}

func (s *SpectralNorm_TimesTranspMState) Finish_To_W() spectralnorm_timestransp.Finish {
	return spectralnorm_timestransp.Finish{}
}

func (s *SpectralNorm_TimesTranspMState) Done() spectralnorm_timestransp_2.M_Result {
	return spectralnorm_timestransp_2.M_Result{Vec: s.V}
}

func (s *SpectralNorm_TimesTranspMState) TimesTranspResult_From_W(timestranspresult_msg spectralnorm_timestransp.TimesTranspResult) {
	// s.V = timestranspresult_msg.Res
}

func (s *SpectralNorm_TimesTranspMState) ResultFrom_SpectralNorm_TimesTransp_M(result spectralnorm_timestransp_2.M_Result) {
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

func (s *SpectralNorm_TimesTranspMState) TimesTranspTask_To_W() spectralnorm_timestransp.TimesTranspTask {
	return spectralnorm_timestransp.TimesTranspTask{
		Ii: s.I * len(s.V) / s.NCPU,
		N:  (s.I + 1) * len(s.V) / s.NCPU,
		U:  s.U,
		V:  s.V,
	}
}

func (s *SpectralNorm_TimesTranspMState) M_Choice() SpectralNorm_TimesTransp_M_Choice {
	if s.I < s.NCPU {
		return SpectralNorm_TimesTransp_M_TimesTranspTask
	}
	return SpectralNorm_TimesTransp_M_Finish
}
