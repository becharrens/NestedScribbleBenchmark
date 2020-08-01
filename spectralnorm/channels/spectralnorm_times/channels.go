package spectralnorm_times

import "NestedScribbleBenchmark/spectralnorm/messages/spectralnorm_times"

type M_Chan struct {
	W_Finish      chan spectralnorm_times.Finish
	W_TimesResult chan spectralnorm_times.TimesResult
	W_TimesTask   chan spectralnorm_times.TimesTask
}

type W_Chan struct {
	M_Finish      chan spectralnorm_times.Finish
	M_TimesResult chan spectralnorm_times.TimesResult
	M_TimesTask   chan spectralnorm_times.TimesTask
}
