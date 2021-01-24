package callbacks

import (
	"NestedScribbleBenchmark/generaldns/results/cached"
)

type Cached_res_Env interface {
	Done() cached.Res_Result
}

type CachedResState struct {
}

func (d *CachedResState) Done() cached.Res_Result {
	return cached.Res_Result{}
}
