package callbacks

import "NestedScribbleBenchmark/dns/results/dns_cached"

type DNS_Cached_res_Env interface {
	Done() dns_cached.Res_Result
}

type DNSCachedResState struct {
}

func (d *DNSCachedResState) Done() dns_cached.Res_Result {
	return dns_cached.Res_Result{}
}
