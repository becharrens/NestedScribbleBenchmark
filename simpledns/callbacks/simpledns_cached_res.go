package callbacks

import (
	"NestedScribbleBenchmark/simpledns/results/simpledns_cached"
)

type SimpleDNS_Cached_res_Env interface {
	Done() simpledns_cached.Res_Result
}

type DNSCachedResState struct {
}

func (d *DNSCachedResState) Done() simpledns_cached.Res_Result {
	return simpledns_cached.Res_Result{}
}
