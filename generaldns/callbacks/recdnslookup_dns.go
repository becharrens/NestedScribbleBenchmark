package callbacks

import (
	"NestedScribbleBenchmark/generaldns/results/recdnslookup"
	"math/rand"
)
import "NestedScribbleBenchmark/generaldns/results/iterdnslookup"
import "NestedScribbleBenchmark/generaldns/results/cached"

type RecDNSLookup_dns_Choice int

const (
	RecDNSLookup_dns_RecDNSLookup RecDNSLookup_dns_Choice = iota
	RecDNSLookup_dns_IterDNSLookup
	RecDNSLookup_dns_Cached
)

type RecDNSLookup_dns_Env interface {
	IP_To_Res_3() string
	ResultFrom_Cached_res(result cached.Res_Result)
	To_Cached_res_Env() Cached_res_Env
	Cached_Setup()
	IP_To_Res_2() string
	ResultFrom_IterDNSLookup_res(result iterdnslookup.Res_Result)
	To_IterDNSLookup_res_Env() IterDNSLookup_res_Env
	IterDNSLookup_Setup()
	Done()
	IP_To_Res() string
	ResultFrom_RecDNSLookup_res(result recdnslookup.Res_Result)
	To_RecDNSLookup_res_Env() RecDNSLookup_res_Env
	RecDNSLookup_Setup()
	Dns_Choice() RecDNSLookup_dns_Choice
	RecReq_From_Res(host string)
}

var RecDNSIdx = 0

type RecDNSLookupDNSState struct {
	KnownIPS   map[string]string
	NextDNSIdx int
	Request    string
}

func (r *RecDNSLookupDNSState) IP_To_Res_3() string {
	return DNSIPs[r.NextDNSIdx]
}

func (r *RecDNSLookupDNSState) ResultFrom_Cached_res(result cached.Res_Result) {
}

func (r *RecDNSLookupDNSState) To_Cached_res_Env() Cached_res_Env {
	return &CachedResState{}
}

func (r *RecDNSLookupDNSState) Cached_Setup() {
}

func (r *RecDNSLookupDNSState) IP_To_Res_2() string {
	// If host is not in map, it will return ""
	return r.KnownIPS[r.Request]
}

func (r *RecDNSLookupDNSState) ResultFrom_IterDNSLookup_res(result iterdnslookup.Res_Result) {
	r.KnownIPS[r.Request] = result.IP
}

func (r *RecDNSLookupDNSState) To_IterDNSLookup_res_Env() IterDNSLookup_res_Env {
	return &IterDNSLookupResState{
		Request: r.Request,
	}
}

func (r *RecDNSLookupDNSState) IterDNSLookup_Setup() {
	DNSIdx = 0
}

func (r *RecDNSLookupDNSState) Done() {
}

func (r *RecDNSLookupDNSState) IP_To_Res() string {
	// If host is not in map, it will return ""
	return r.KnownIPS[r.Request]
}

func (r *RecDNSLookupDNSState) ResultFrom_RecDNSLookup_res(result recdnslookup.Res_Result) {
	r.KnownIPS[r.Request] = result.IP
}

func (r *RecDNSLookupDNSState) To_RecDNSLookup_res_Env() RecDNSLookup_res_Env {
	return &RecDNSLookupResState{
		Request: r.Request,
	}
}

func (r *RecDNSLookupDNSState) RecDNSLookup_Setup() {
}

func (r *RecDNSLookupDNSState) Dns_Choice() RecDNSLookup_dns_Choice {
	// If IP is known, or you have requested all DNS servers, stop search
	if _, ok := r.KnownIPS[r.Request]; ok || r.NextDNSIdx >= len(DNSIPs) {
		return RecDNSLookup_dns_Cached
	}

	// Decide whether to make recursive or iterative DNS request
	if rand.Intn(10) < 5 {
		return RecDNSLookup_dns_IterDNSLookup
	}
	return RecDNSLookup_dns_RecDNSLookup
}

func (r *RecDNSLookupDNSState) RecReq_From_Res(host string) {
	r.Request = host
}

func New_RecDNSLookup_dns_State() RecDNSLookup_dns_Env {
	nextDNSIdx := (RecDNSIdx + 1) % len(DNSIPs)
	dnsState := &RecDNSLookupDNSState{
		KnownIPS:   DNSCache[DNSIdx],
		NextDNSIdx: nextDNSIdx,
	}
	RecDNSIdx = nextDNSIdx
	return dnsState
}
