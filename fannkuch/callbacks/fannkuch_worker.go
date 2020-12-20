package callbacks

import "NestedScribbleBenchmark/fannkuch/results/fannkuchrecursive"
import "NestedScribbleBenchmark/fannkuch/results/fannkuch"

type Fannkuch_Worker_Choice int

const (
	Fannkuch_Worker_FannkuchRecursive Fannkuch_Worker_Choice = iota
	Fannkuch_Worker_Result
)

type Fannkuch_Worker_Env interface {
	Result_To_Main_2() (int, int)
	Done() fannkuch.Worker_Result
	Result_To_Main() (int, int)
	ResultFrom_FannkuchRecursive_Worker(result fannkuchrecursive.Worker_Result)
	To_FannkuchRecursive_Worker_Env() FannkuchRecursive_Worker_Env
	FannkuchRecursive_Setup()
	Worker_Choice() Fannkuch_Worker_Choice
	Task_From_Main(idxmin int, chunksz int, n int)
}

type FannkuchWorkerState struct {
	IdxMin  int
	IdxMax  int
	N       int
	Chunksz int
}

func (f *FannkuchWorkerState) Result_To_Main_2() (int, int) {
	maxFlips, checksum := fannkuchImpl(f.IdxMin, f.IdxMax, f.N)
	return maxFlips, checksum
}

func (f *FannkuchWorkerState) Done() fannkuch.Worker_Result {
	return fannkuch.Worker_Result{}
}

func (f *FannkuchWorkerState) Result_To_Main() (int, int) {
	maxFlips, checksum := fannkuchImpl(f.IdxMin, f.IdxMax, f.N)
	return maxFlips, checksum
}

func (f *FannkuchWorkerState) ResultFrom_FannkuchRecursive_Worker(result fannkuchrecursive.Worker_Result) {
}

func (f *FannkuchWorkerState) To_FannkuchRecursive_Worker_Env() FannkuchRecursive_Worker_Env {
	return &FannkuchRecursiveWorkerState{
		IdxMin:  f.IdxMax,
		N:       f.N,
		Chunksz: f.Chunksz,
	}
}

func (f *FannkuchWorkerState) FannkuchRecursive_Setup() {
}

func (f *FannkuchWorkerState) Worker_Choice() Fannkuch_Worker_Choice {
	if f.IdxMax < Fact[f.N] {
		return Fannkuch_Worker_FannkuchRecursive
	}
	f.IdxMax = Fact[f.N]
	return Fannkuch_Worker_Result
}

func (f *FannkuchWorkerState) Task_From_Main(idxmin int, chunksz int, n int) {
	f.N = n
	f.IdxMin = idxmin
	f.IdxMax = idxmin + chunksz
	f.N = n
	f.Chunksz = chunksz
}

// ------

func fannkuchImpl(idxMin int, idxMax int, n int) (int, int) {
	p := make([]int, n)
	pp := make([]int, n)
	count := make([]int, n)

	// first permutation
	for i := 0; i < n; i++ {
		p[i] = i
	}
	for i, idx := n-1, idxMin; i > 0; i-- {
		d := idx / Fact[i]
		count[i] = d
		idx = idx % Fact[i]

		copy(pp, p)
		for j := 0; j <= i; j++ {
			if j+d <= i {
				p[j] = pp[j+d]
			} else {
				p[j] = pp[j+d-i-1]
			}
		}
	}

	maxFlips := 1
	checkSum := 0

	for idx, sign := idxMin, true; ; sign = !sign {

		// count flips
		first := p[0]
		if first != 0 {
			flips := 1
			if p[first] != 0 {
				copy(pp, p)
				p0 := first
				for {
					flips++
					for i, j := 1, p0-1; i < j; i, j = i+1, j-1 {
						pp[i], pp[j] = pp[j], pp[i]
					}
					t := pp[p0]
					pp[p0] = p0
					p0 = t
					if pp[p0] == 0 {
						break
					}
				}
			}
			if maxFlips < flips {
				maxFlips = flips
			}
			if sign {
				checkSum += flips
			} else {
				checkSum -= flips
			}
		}

		if idx++; idx == idxMax {
			break
		}

		// next permutation
		if sign {
			p[0], p[1] = p[1], first
		} else {
			p[1], p[2] = p[2], p[1]
			for k := 2; ; k++ {
				if count[k]++; count[k] <= k {
					break
				}
				count[k] = 0
				for j := 0; j <= k; j++ {
					p[j] = p[j+1]
				}
				p[k+1] = first
				first = p[0]
			}
		}
	}
	return maxFlips, checkSum
}
