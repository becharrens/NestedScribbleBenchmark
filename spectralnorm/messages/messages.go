package messages

type SpectralNorm_Label int

const (
	Finish SpectralNorm_Label = iota
	SpectralNorm_Master_Worker
	SpectralNorm_Times_M
	SpectralNorm_Times_Master
	SpectralNorm_TimesTransp_M
	SpectralNorm_TimesTransp_Master
	TimesResult
	TimesTask
	TimesTranspResult
	TimesTranspTask
)
