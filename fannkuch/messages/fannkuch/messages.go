package fannkuch

type Result struct {
	Checksum int
	MaxFlips int
}

type Task struct {
	Chunksz int
	Fact    []int
	IdxMin  int
	N       int
}
