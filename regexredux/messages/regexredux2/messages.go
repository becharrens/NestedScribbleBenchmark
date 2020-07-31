package regexredux2

type CalcLength struct {
	B []byte
}

type Length struct {
	Len int
}

type NumMatches struct {
	Nmatches int
}

type Task struct {
	B       []byte
	Pattern string
}
