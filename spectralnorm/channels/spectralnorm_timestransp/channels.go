package spectralnorm_timestransp

import "NestedScribbleBenchmark/spectralnorm/messages/spectralnorm_timestransp"

type M_Chan struct {
	W_Finish            chan spectralnorm_timestransp.Finish
	W_TimesTranspResult chan spectralnorm_timestransp.TimesTranspResult
	W_TimesTranspTask   chan spectralnorm_timestransp.TimesTranspTask
}

type W_Chan struct {
	M_Finish            chan spectralnorm_timestransp.Finish
	M_TimesTranspResult chan spectralnorm_timestransp.TimesTranspResult
	M_TimesTranspTask   chan spectralnorm_timestransp.TimesTranspTask
}
