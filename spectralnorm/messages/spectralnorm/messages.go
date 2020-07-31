package spectralnorm

type Finish struct {
}

type TimesResult struct {
	Res []float64
}

type TimesTask struct {
	Ii int
	N  int
	U  []float64
	V  []float64
}
