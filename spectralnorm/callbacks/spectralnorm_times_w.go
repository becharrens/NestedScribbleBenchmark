package callbacks

import "NestedScribbleBenchmark/spectralnorm/messages/spectralnorm_times"

type SpectralNorm_Times_W_Env interface {
	Finish_From_M(finish_msg spectralnorm_times.Finish)
	Done()
	TimesResult_To_M() spectralnorm_times.TimesResult
	TimesTask_From_M(timestask_msg spectralnorm_times.TimesTask)
}

type SpectralNorm_TimesWState struct {
	V []float64
}

func (s *SpectralNorm_TimesWState) Finish_From_M(finish_msg spectralnorm_times.Finish) {
}

func (s *SpectralNorm_TimesWState) Done() {
}

func (s *SpectralNorm_TimesWState) TimesResult_To_M() spectralnorm_times.TimesResult {
	return spectralnorm_times.TimesResult{Res: s.V}
}

func (s *SpectralNorm_TimesWState) TimesTask_From_M(timestask_msg spectralnorm_times.TimesTask) {
	s.V = timestask_msg.V
	Times(timestask_msg.Ii, timestask_msg.N, s.V, timestask_msg.U)
}

func New_SpectralNorm_Times_W_State() SpectralNorm_Times_W_Env {
	return &SpectralNorm_TimesWState{}
}
