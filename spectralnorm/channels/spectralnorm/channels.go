package spectralnorm

import "NestedScribbleBenchmark/spectralnorm/messages/spectralnorm"

type Master_Chan struct {
	Worker_Finish      chan spectralnorm.Finish
	Worker_TimesResult chan spectralnorm.TimesResult
	Worker_TimesTask   chan spectralnorm.TimesTask
}

type Worker_Chan struct {
	Master_Finish      chan spectralnorm.Finish
	Master_TimesResult chan spectralnorm.TimesResult
	Master_TimesTask   chan spectralnorm.TimesTask
}
