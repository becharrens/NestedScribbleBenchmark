package main

import (
	"ScribbleBenchmark/benchmark"
	"ScribbleBenchmark/regexredux/callbacks"
	"ScribbleBenchmark/regexredux/protocol"
	"ScribbleBenchmark/regexredux/results/regexredux"
	"ScribbleBenchmark/regexredux_base"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/GRbit/go-pcre"
)

type RegexReduxEnv struct {
	B []byte
}

func (r *RegexReduxEnv) New_Master_Env() callbacks.RegexRedux_Master_Env {
	ilen := len(r.B)
	b := pcre.
		MustCompileJIT("(>[^\n]*)?\n", 0, pcre.STUDY_JIT_COMPILE).
		ReplaceAll(r.B, []byte{}, 0)
	clen := len(b)

	return &callbacks.RegexReduxMasterState{
		B:    b,
		Idx:  0,
		CLen: clen,
		ILen: ilen,
	}
}

func (r *RegexReduxEnv) New_Worker_Env() callbacks.RegexRedux_Worker_Env {
	return &callbacks.RegexReduxWorkerState{}
}

func (r *RegexReduxEnv) Master_Result(result regexredux.Master_Result) {
}

func (r *RegexReduxEnv) Worker_Result(result regexredux.Worker_Result) {
}

var regexreduxParams = []int{
	0, 1, 2, 3, 4, 5, 6, 7,
}

var regexreduxFiles = []string{
	"regexredux-input1000.txt",
	"regexredux-input10000.txt",
	"regexredux-input100000.txt",
	"regexredux-input500000.txt",
	"regexredux-input1000000.txt",
	"regexredux-input5000000.txt",
	"regexredux-input10000000.txt",
	"regexredux-input25000000.txt",
}

func NewRegexReduxEnv(b []byte) *RegexReduxEnv {
	return &RegexReduxEnv{
		B: b,
	}
}

func TimeRegexRedux(n int) time.Duration {
	b := readFile(n, regexreduxFiles)
	env := NewRegexReduxEnv(b)
	start := time.Now()
	protocol.RegexRedux(env)
	elapsed := time.Since(start)
	return elapsed
}

func TimeRegexReduxBase(n int) time.Duration {
	b := readFile(n, regexreduxFiles)
	start := time.Now()
	regexredux_base.RegexRedux(b)
	elapsed := time.Since(start)
	return elapsed
}

func readFile(n int, files []string) []byte {
	f, err := os.Open(fmt.Sprintf("testdata/%s", files[n]))
	if err != nil {
		panic("Can't open input file")
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		panic("Can't read input file")
	}
	err = f.Close()
	if err != nil {
		panic("Can't close input file")
	}
	return b
}

func RegexReduxBenchmark(repetitions int) (benchmark.BenchmarkTimes, benchmark.BenchmarkTimes) {
	scribble_results := benchmark.TimeImpl(regexreduxParams, repetitions, TimeRegexRedux)
	base_results := benchmark.TimeImpl(regexreduxParams, repetitions, TimeRegexReduxBase)
	return scribble_results, base_results
}
