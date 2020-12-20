package callbacks

import (
	"NestedScribbleBenchmark/spectralnorm/results/spectralnorm_timestransp"
	"math"
)
import "NestedScribbleBenchmark/spectralnorm/results/spectralnorm_times"
import "NestedScribbleBenchmark/spectralnorm/results/spectralnorm"

type SpectralNorm_Master_Choice int

const (
	SpectralNorm_Master_TimesTask SpectralNorm_Master_Choice = iota
	SpectralNorm_Master_Finish
)

type SpectralNorm_Master_Env interface {
	Finish_To_Worker()
	Done() spectralnorm.Master_Result
	ResultFrom_SpectralNorm_Master(result spectralnorm.Master_Result)
	To_SpectralNorm_Master_Env() SpectralNorm_Master_Env
	SpectralNorm_Setup()
	ResultFrom_SpectralNorm_TimesTransp_M_2(result spectralnorm_timestransp.M_Result)
	To_SpectralNorm_TimesTransp_M_Env_2() SpectralNorm_TimesTransp_M_Env
	SpectralNorm_TimesTransp_Setup_2()
	ResultFrom_SpectralNorm_Times_M_2(result spectralnorm_times.M_Result)
	To_SpectralNorm_Times_M_Env_2() SpectralNorm_Times_M_Env
	SpectralNorm_Times_Setup_2()
	ResultFrom_SpectralNorm_TimesTransp_M(result spectralnorm_timestransp.M_Result)
	To_SpectralNorm_TimesTransp_M_Env() SpectralNorm_TimesTransp_M_Env
	SpectralNorm_TimesTransp_Setup()
	TimesResult_From_Worker(res []float64)
	ResultFrom_SpectralNorm_Times_M(result spectralnorm_times.M_Result)
	To_SpectralNorm_Times_M_Env() SpectralNorm_Times_M_Env
	SpectralNorm_Times_Setup()
	TimesTask_To_Worker() (int, int, []float64, []float64)
	Master_Choice() SpectralNorm_Master_Choice
}

type SpectralNormMasterState struct {
	U            []float64
	V            []float64
	X            []float64
	I            int
	NCPU         int
	Iter         int
	SpectralNorm float64
}

func (s *SpectralNormMasterState) Finish_To_Worker() {
}

func (s *SpectralNormMasterState) Done() spectralnorm.Master_Result {
	return spectralnorm.Master_Result{
		SpectralNorm: s.SpectralNorm,
	}
}

func (s *SpectralNormMasterState) ResultFrom_SpectralNorm_Master(result spectralnorm.Master_Result) {
	s.SpectralNorm = result.SpectralNorm
}

func (s *SpectralNormMasterState) To_SpectralNorm_Master_Env() SpectralNorm_Master_Env {
	return &SpectralNormMasterState{
		U:    s.U,
		V:    s.V,
		I:    s.I,
		NCPU: s.NCPU,
		Iter: s.Iter + 1,
	}
}

func (s *SpectralNormMasterState) SpectralNorm_Setup() {
}

func (s *SpectralNormMasterState) ResultFrom_SpectralNorm_TimesTransp_M_2(result spectralnorm_timestransp.M_Result) {
	// s.U = result.Vec
}

func (s *SpectralNormMasterState) To_SpectralNorm_TimesTransp_M_Env_2() SpectralNorm_TimesTransp_M_Env {
	return &SpectralNorm_TimesTranspMState{
		U:    s.X,
		V:    s.U,
		I:    s.I,
		NCPU: s.NCPU,
	}
}

func (s *SpectralNormMasterState) SpectralNorm_TimesTransp_Setup_2() {
}

func (s *SpectralNormMasterState) ResultFrom_SpectralNorm_Times_M_2(result spectralnorm_times.M_Result) {
	// s.X = result.Vec
}

func (s *SpectralNormMasterState) To_SpectralNorm_Times_M_Env_2() SpectralNorm_Times_M_Env {
	return &SpectralNorm_TimesMState{
		U:    s.V,
		V:    s.X,
		I:    s.I,
		NCPU: s.NCPU,
	}
}

func (s *SpectralNormMasterState) SpectralNorm_Times_Setup_2() {
}

func (s *SpectralNormMasterState) ResultFrom_SpectralNorm_TimesTransp_M(result spectralnorm_timestransp.M_Result) {
	// s.V = result.Vec
	s.X = make([]float64, len(s.V))
}

func (s *SpectralNormMasterState) To_SpectralNorm_TimesTransp_M_Env() SpectralNorm_TimesTransp_M_Env {
	return &SpectralNorm_TimesTranspMState{
		U:    s.X,
		V:    s.V,
		I:    s.I,
		NCPU: s.NCPU,
	}
}

func (s *SpectralNormMasterState) SpectralNorm_TimesTransp_Setup() {
}

func (s *SpectralNormMasterState) TimesResult_From_Worker(res []float64) {
	// s.X = timesresult_msg.Res]
}

func (s *SpectralNormMasterState) ResultFrom_SpectralNorm_Times_M(result spectralnorm_times.M_Result) {
	// s.X = result.Vec
	return
}

func (s *SpectralNormMasterState) To_SpectralNorm_Times_M_Env() SpectralNorm_Times_M_Env {
	return &SpectralNorm_TimesMState{
		U:    s.U,
		V:    s.X,
		I:    s.I + 1,
		NCPU: s.NCPU,
	}
}

func (s *SpectralNormMasterState) SpectralNorm_Times_Setup() {
}

func (s *SpectralNormMasterState) TimesTask_To_Worker() (int, int, []float64, []float64) {
	s.X = make([]float64, len(s.U))
	return s.I * len(s.V) / s.NCPU, (s.I + 1) * len(s.V) / s.NCPU, s.U, s.X
}

func (s *SpectralNormMasterState) Master_Choice() SpectralNorm_Master_Choice {
	if s.Iter < 10 {
		return SpectralNorm_Master_TimesTask
	}
	var vBv, vv float64
	for i, vi := range s.V {
		vBv += s.U[i] * vi
		vv += vi * vi
	}
	s.SpectralNorm = math.Sqrt(vBv / vv)
	// TODO: Uncomment
	// fmt.Printf("%0.9f\n", s.SpectralNorm)
	return SpectralNorm_Master_Finish
}
