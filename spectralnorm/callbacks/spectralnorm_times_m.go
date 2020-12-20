package callbacks

import "NestedScribbleBenchmark/spectralnorm/results/spectralnorm_times"

type SpectralNorm_Times_M_Choice int

const (
	SpectralNorm_Times_M_TimesTask SpectralNorm_Times_M_Choice = iota
	SpectralNorm_Times_M_Finish
)

type SpectralNorm_Times_M_Env interface {
	Finish_To_W()
	Done() spectralnorm_times.M_Result
	TimesResult_From_W(res []float64)
	ResultFrom_SpectralNorm_Times_M(result spectralnorm_times.M_Result)
	To_SpectralNorm_Times_M_Env() SpectralNorm_Times_M_Env
	SpectralNorm_Times_Setup()
	TimesTask_To_W() (int, int, []float64, []float64)
	M_Choice() SpectralNorm_Times_M_Choice
}

type SpectralNorm_TimesMState struct {
	U    []float64
	V    []float64
	I    int
	NCPU int
}

func (s *SpectralNorm_TimesMState) Finish_To_W() {
}

func (s *SpectralNorm_TimesMState) Done() spectralnorm_times.M_Result {
	return spectralnorm_times.M_Result{Vec: s.V}
}

func (s *SpectralNorm_TimesMState) TimesResult_From_W(res []float64) {
	// s.V = timesresult_msg.Res
}

func (s *SpectralNorm_TimesMState) ResultFrom_SpectralNorm_Times_M(result spectralnorm_times.M_Result) {
	// s.V = result.Vec
}

func (s *SpectralNorm_TimesMState) To_SpectralNorm_Times_M_Env() SpectralNorm_Times_M_Env {
	return &SpectralNorm_TimesMState{
		U:    s.U,
		V:    s.V,
		I:    s.I + 1,
		NCPU: s.NCPU,
	}
}

func (s *SpectralNorm_TimesMState) SpectralNorm_Times_Setup() {
}

func (s *SpectralNorm_TimesMState) TimesTask_To_W() (int, int, []float64, []float64) {
	return s.I * len(s.V) / s.NCPU, (s.I + 1) * len(s.V) / s.NCPU, s.U, s.V
}

func (s *SpectralNorm_TimesMState) M_Choice() SpectralNorm_Times_M_Choice {
	if s.I < s.NCPU {
		return SpectralNorm_Times_M_TimesTask
	}
	return SpectralNorm_Times_M_Finish
}
