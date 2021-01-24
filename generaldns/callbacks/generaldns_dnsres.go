package callbacks

import (
	"NestedScribbleBenchmark/generaldns/results/cached"
	"NestedScribbleBenchmark/generaldns/results/generaldns"
	"NestedScribbleBenchmark/generaldns/results/iterdnslookup"
	"NestedScribbleBenchmark/generaldns/results/recdnslookup"
	"fmt"
	"math/rand"
)

type GeneralDNS_dnsRes_Choice_4 int

const (
	GeneralDNS_dnsRes_RecDNSLookup_4 GeneralDNS_dnsRes_Choice_4 = iota
	GeneralDNS_dnsRes_IterDNSLookup_4
	GeneralDNS_dnsRes_Cached_4
)

type GeneralDNS_dnsRes_Choice_3 int

const (
	GeneralDNS_dnsRes_RecDNSLookup_3 GeneralDNS_dnsRes_Choice_3 = iota
	GeneralDNS_dnsRes_IterDNSLookup_3
	GeneralDNS_dnsRes_Cached_3
)

type GeneralDNS_dnsRes_Choice_2 int

const (
	GeneralDNS_dnsRes_RecDNSLookup_2 GeneralDNS_dnsRes_Choice_2 = iota
	GeneralDNS_dnsRes_IterDNSLookup_2
	GeneralDNS_dnsRes_Cached_2
)

type GeneralDNS_dnsRes_Choice int

const (
	GeneralDNS_dnsRes_RecDNSLookup GeneralDNS_dnsRes_Choice = iota
	GeneralDNS_dnsRes_IterDNSLookup
	GeneralDNS_dnsRes_Cached
)

type GeneralDNS_dnsRes_Env interface {
	IP_To_App_12() string
	ResultFrom_Cached_res_4(result cached.Res_Result)
	To_Cached_res_Env_4() Cached_res_Env
	Cached_Setup_4()
	IP_To_App_11() string
	ResultFrom_IterDNSLookup_res_4(result iterdnslookup.Res_Result)
	To_IterDNSLookup_res_Env_4() IterDNSLookup_res_Env
	IterDNSLookup_Setup_4()
	IP_To_App_10() string
	ResultFrom_RecDNSLookup_res_4(result recdnslookup.Res_Result)
	To_RecDNSLookup_res_Env_4() RecDNSLookup_res_Env
	RecDNSLookup_Setup_4()
	DnsRes_Choice_4() GeneralDNS_dnsRes_Choice_4
	Query_From_App_4(host string)
	Done_From_App_4()
	IP_To_App_9() string
	ResultFrom_Cached_res_3(result cached.Res_Result)
	To_Cached_res_Env_3() Cached_res_Env
	Cached_Setup_3()
	IP_To_App_8() string
	ResultFrom_Cached_res_2(result cached.Res_Result)
	To_Cached_res_Env_2() Cached_res_Env
	Cached_Setup_2()
	IP_To_App_7() string
	ResultFrom_IterDNSLookup_res_3(result iterdnslookup.Res_Result)
	To_IterDNSLookup_res_Env_3() IterDNSLookup_res_Env
	IterDNSLookup_Setup_3()
	IP_To_App_6() string
	ResultFrom_RecDNSLookup_res_3(result recdnslookup.Res_Result)
	To_RecDNSLookup_res_Env_3() RecDNSLookup_res_Env
	RecDNSLookup_Setup_3()
	DnsRes_Choice_3() GeneralDNS_dnsRes_Choice_3
	Query_From_App_3(host string)
	Done_From_App_3()
	IP_To_App_5() string
	ResultFrom_IterDNSLookup_res_2(result iterdnslookup.Res_Result)
	To_IterDNSLookup_res_Env_2() IterDNSLookup_res_Env
	IterDNSLookup_Setup_2()
	IP_To_App_4() string
	ResultFrom_Cached_res(result cached.Res_Result)
	To_Cached_res_Env() Cached_res_Env
	Cached_Setup()
	IP_To_App_3() string
	ResultFrom_IterDNSLookup_res(result iterdnslookup.Res_Result)
	To_IterDNSLookup_res_Env() IterDNSLookup_res_Env
	IterDNSLookup_Setup()
	IP_To_App_2() string
	ResultFrom_RecDNSLookup_res_2(result recdnslookup.Res_Result)
	To_RecDNSLookup_res_Env_2() RecDNSLookup_res_Env
	RecDNSLookup_Setup_2()
	DnsRes_Choice_2() GeneralDNS_dnsRes_Choice_2
	Query_From_App_2(host string)
	Done_From_App_2()
	IP_To_App() string
	ResultFrom_RecDNSLookup_res(result recdnslookup.Res_Result)
	To_RecDNSLookup_res_Env() RecDNSLookup_res_Env
	RecDNSLookup_Setup()
	DnsRes_Choice() GeneralDNS_dnsRes_Choice
	Query_From_App(host string)
	Done() generaldns.DnsRes_Result
	Done_From_App()
}

type GeneralDNSDNSResState struct {
	Cache   map[string]string
	Request string
}

func (g *GeneralDNSDNSResState) IP_To_App_12() string {
	ip := g.Cache[g.Request]
	fmt.Println("generaldns_dnsres: returning cached IP '", ip, "'")
	return ip
}

func (g *GeneralDNSDNSResState) ResultFrom_Cached_res_4(result cached.Res_Result) {
}

func (g *GeneralDNSDNSResState) To_Cached_res_Env_4() Cached_res_Env {
	return &CachedResState{}
}

func (g *GeneralDNSDNSResState) Cached_Setup_4() {
}

func (g *GeneralDNSDNSResState) IP_To_App_11() string {
	ip := g.Cache[g.Request]
	fmt.Println("generaldns_dnsres: returning looked-up IP '", ip, "'")
	return ip
}

func (g *GeneralDNSDNSResState) ResultFrom_IterDNSLookup_res_4(result iterdnslookup.Res_Result) {
	g.Cache[g.Request] = result.IP
}

func (g *GeneralDNSDNSResState) To_IterDNSLookup_res_Env_4() IterDNSLookup_res_Env {
	return &IterDNSLookupResState{
		Request: g.Request,
	}
}

func (g *GeneralDNSDNSResState) IterDNSLookup_Setup_4() {
	DNSIdx = 0
}

func (g *GeneralDNSDNSResState) IP_To_App_10() string {
	ip := g.Cache[g.Request]
	fmt.Println("generaldns_dnsres: returning looked-up IP '", ip, "'")
	return ip
}

func (g *GeneralDNSDNSResState) ResultFrom_RecDNSLookup_res_4(result recdnslookup.Res_Result) {
	g.Cache[g.Request] = result.IP
}

func (g *GeneralDNSDNSResState) To_RecDNSLookup_res_Env_4() RecDNSLookup_res_Env {
	return &RecDNSLookupResState{
		Request: g.Request,
	}
}

func (g *GeneralDNSDNSResState) RecDNSLookup_Setup_4() {
	RecDNSIdx = 0
}

func (g *GeneralDNSDNSResState) DnsRes_Choice_4() GeneralDNS_dnsRes_Choice_4 {
	// If IP is known return cached response
	if _, ok := g.Cache[g.Request]; ok {
		return GeneralDNS_dnsRes_Cached_4
	}

	// Decide whether to make recursive or iterative DNS request
	if rand.Intn(10) < 5 {
		return GeneralDNS_dnsRes_IterDNSLookup_4
	}
	return GeneralDNS_dnsRes_RecDNSLookup_4
}

func (g *GeneralDNSDNSResState) Query_From_App_4(host string) {
	g.Request = host
}

func (g *GeneralDNSDNSResState) Done_From_App_4() {
}

func (g *GeneralDNSDNSResState) IP_To_App_9() string {
	ip := g.Cache[g.Request]
	fmt.Println("generaldns_dnsres: returning cached IP '", ip, "'")
	return ip
}

func (g *GeneralDNSDNSResState) ResultFrom_Cached_res_3(result cached.Res_Result) {
}

func (g *GeneralDNSDNSResState) To_Cached_res_Env_3() Cached_res_Env {
	return &CachedResState{}
}

func (g *GeneralDNSDNSResState) Cached_Setup_3() {
}

func (g *GeneralDNSDNSResState) IP_To_App_8() string {
	ip := g.Cache[g.Request]
	fmt.Println("generaldns_dnsres: returning cached IP '", ip, "'")
	return ip
}

func (g *GeneralDNSDNSResState) ResultFrom_Cached_res_2(result cached.Res_Result) {
}

func (g *GeneralDNSDNSResState) To_Cached_res_Env_2() Cached_res_Env {
	return &CachedResState{}
}

func (g *GeneralDNSDNSResState) Cached_Setup_2() {
}

func (g *GeneralDNSDNSResState) IP_To_App_7() string {
	ip := g.Cache[g.Request]
	fmt.Println("generaldns_dnsres: returning looked-up IP '", ip, "'")
	return ip
}

func (g *GeneralDNSDNSResState) ResultFrom_IterDNSLookup_res_3(result iterdnslookup.Res_Result) {
	g.Cache[g.Request] = result.IP
}

func (g *GeneralDNSDNSResState) To_IterDNSLookup_res_Env_3() IterDNSLookup_res_Env {
	return &IterDNSLookupResState{
		Request: g.Request,
	}
}

func (g *GeneralDNSDNSResState) IterDNSLookup_Setup_3() {
	DNSIdx = 0
}

func (g *GeneralDNSDNSResState) IP_To_App_6() string {
	ip := g.Cache[g.Request]
	fmt.Println("generaldns_dnsres: returning looked-up IP '", ip, "'")
	return ip
}

func (g *GeneralDNSDNSResState) ResultFrom_RecDNSLookup_res_3(result recdnslookup.Res_Result) {
	g.Cache[g.Request] = result.IP
}

func (g *GeneralDNSDNSResState) To_RecDNSLookup_res_Env_3() RecDNSLookup_res_Env {
	return &RecDNSLookupResState{
		Request: g.Request,
	}
}

func (g *GeneralDNSDNSResState) RecDNSLookup_Setup_3() {
	RecDNSIdx = 0
}

func (g *GeneralDNSDNSResState) DnsRes_Choice_3() GeneralDNS_dnsRes_Choice_3 {
	// If IP is known return cached response
	if _, ok := g.Cache[g.Request]; ok {
		return GeneralDNS_dnsRes_Cached_3
	}

	// Decide whether to make recursive or iterative DNS request
	if rand.Intn(10) < 5 {
		return GeneralDNS_dnsRes_IterDNSLookup_3
	}
	return GeneralDNS_dnsRes_RecDNSLookup_3
}

func (g *GeneralDNSDNSResState) Query_From_App_3(host string) {
	g.Request = host
}

func (g *GeneralDNSDNSResState) Done_From_App_3() {
}

func (g *GeneralDNSDNSResState) IP_To_App_5() string {
	ip := g.Cache[g.Request]
	fmt.Println("generaldns_dnsres: returning looked-up IP '", ip, "'")
	return ip
}

func (g *GeneralDNSDNSResState) ResultFrom_IterDNSLookup_res_2(result iterdnslookup.Res_Result) {
	g.Cache[g.Request] = result.IP
}

func (g *GeneralDNSDNSResState) To_IterDNSLookup_res_Env_2() IterDNSLookup_res_Env {
	return &IterDNSLookupResState{
		Request: g.Request,
	}
}

func (g *GeneralDNSDNSResState) IterDNSLookup_Setup_2() {
	DNSIdx = 0
}

func (g *GeneralDNSDNSResState) IP_To_App_4() string {
	ip := g.Cache[g.Request]
	fmt.Println("generaldns_dnsres: returning cached IP '", ip, "'")
	return ip
}

func (g *GeneralDNSDNSResState) ResultFrom_Cached_res(result cached.Res_Result) {
}

func (g *GeneralDNSDNSResState) To_Cached_res_Env() Cached_res_Env {
	return &CachedResState{}
}

func (g *GeneralDNSDNSResState) Cached_Setup() {
}

func (g *GeneralDNSDNSResState) IP_To_App_3() string {
	ip := g.Cache[g.Request]
	fmt.Println("generaldns_dnsres: returning looked-up IP '", ip, "'")
	return ip
}

func (g *GeneralDNSDNSResState) ResultFrom_IterDNSLookup_res(result iterdnslookup.Res_Result) {
	g.Cache[g.Request] = result.IP
}

func (g *GeneralDNSDNSResState) To_IterDNSLookup_res_Env() IterDNSLookup_res_Env {
	return &IterDNSLookupResState{
		Request: g.Request,
	}
}

func (g *GeneralDNSDNSResState) IterDNSLookup_Setup() {
	DNSIdx = 0
}

func (g *GeneralDNSDNSResState) IP_To_App_2() string {
	ip := g.Cache[g.Request]
	fmt.Println("generaldns_dnsres: returning looked-up IP '", ip, "'")
	return ip
}

func (g *GeneralDNSDNSResState) ResultFrom_RecDNSLookup_res_2(result recdnslookup.Res_Result) {
	g.Cache[g.Request] = result.IP
}

func (g *GeneralDNSDNSResState) To_RecDNSLookup_res_Env_2() RecDNSLookup_res_Env {
	return &RecDNSLookupResState{
		Request: g.Request,
	}
}

func (g *GeneralDNSDNSResState) RecDNSLookup_Setup_2() {
	RecDNSIdx = 0
}

func (g *GeneralDNSDNSResState) DnsRes_Choice_2() GeneralDNS_dnsRes_Choice_2 {
	// If IP is known return cached response
	if _, ok := g.Cache[g.Request]; ok {
		return GeneralDNS_dnsRes_Cached_2
	}

	// Decide whether to make recursive or iterative DNS request
	if rand.Intn(10) < 5 {
		return GeneralDNS_dnsRes_IterDNSLookup_2
	}
	return GeneralDNS_dnsRes_RecDNSLookup_2
}

func (g *GeneralDNSDNSResState) Query_From_App_2(host string) {
	g.Request = host
}

func (g *GeneralDNSDNSResState) Done_From_App_2() {
}

func (g *GeneralDNSDNSResState) IP_To_App() string {
	ip := g.Cache[g.Request]
	fmt.Println("generaldns_dnsres: returning looked-up IP '", ip, "'")
	return ip
}

func (g *GeneralDNSDNSResState) ResultFrom_RecDNSLookup_res(result recdnslookup.Res_Result) {
	g.Cache[g.Request] = result.IP
}

func (g *GeneralDNSDNSResState) To_RecDNSLookup_res_Env() RecDNSLookup_res_Env {
	return &RecDNSLookupResState{
		Request: g.Request,
	}
}

func (g *GeneralDNSDNSResState) RecDNSLookup_Setup() {
	RecDNSIdx = 0
}

func (g *GeneralDNSDNSResState) DnsRes_Choice() GeneralDNS_dnsRes_Choice {
	// If IP is known return cached response
	if _, ok := g.Cache[g.Request]; ok {
		return GeneralDNS_dnsRes_Cached
	}

	// Decide whether to make recursive or iterative DNS request
	if rand.Intn(10) < 5 {
		return GeneralDNS_dnsRes_IterDNSLookup
	}
	return GeneralDNS_dnsRes_RecDNSLookup
}

func (g *GeneralDNSDNSResState) Query_From_App(host string) {
	g.Request = host
}

func (g *GeneralDNSDNSResState) Done() generaldns.DnsRes_Result {
	return generaldns.DnsRes_Result{
		Cache: g.Cache,
	}
}

func (g *GeneralDNSDNSResState) Done_From_App() {
}
