package spectralnorm_timestransp

type Finish struct {
}

type TimesTranspResult struct {
	Res []float64
}

type TimesTranspTask struct {
	Ii int
	N  int
	U  []float64
	V  []float64
}
