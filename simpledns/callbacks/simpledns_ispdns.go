package callbacks

import (
	"NestedScribbleBenchmark/simpledns/results/dnslookup"
	"NestedScribbleBenchmark/simpledns/results/simpledns"
	"NestedScribbleBenchmark/simpledns/results/simpledns_cached"
	"fmt"
)

type SimpleDNS_ispDNS_Choice_3 int

const (
	SimpleDNS_ispDNS_SimpleDNS_Cached_3 SimpleDNS_ispDNS_Choice_3 = iota
	SimpleDNS_ispDNS_DNSLookup_3
)

type SimpleDNS_ispDNS_Choice_2 int

const (
	SimpleDNS_ispDNS_SimpleDNS_Cached_2 SimpleDNS_ispDNS_Choice_2 = iota
	SimpleDNS_ispDNS_DNSLookup_2
)

type SimpleDNS_ispDNS_Choice int

const (
	SimpleDNS_ispDNS_SimpleDNS_Cached SimpleDNS_ispDNS_Choice = iota
	SimpleDNS_ispDNS_DNSLookup
)

type SimpleDNS_ispDNS_Env interface {
	IP_To_App_6() string
	ResultFrom_DNSLookup_res_3(result dnslookup.Res_Result)
	To_DNSLookup_res_Env_3() DNSLookup_res_Env
	DNSLookup_Setup_3()
	IP_To_App_5() string
	ResultFrom_SimpleDNS_Cached_res_3(result simpledns_cached.Res_Result)
	To_SimpleDNS_Cached_res_Env_3() SimpleDNS_Cached_res_Env
	SimpleDNS_Cached_Setup_3()
	IspDNS_Choice_3() SimpleDNS_ispDNS_Choice_3
	RecQuery_From_App_3(host string)
	Done_From_App_3()
	IP_To_App_4() string
	ResultFrom_DNSLookup_res_2(result dnslookup.Res_Result)
	To_DNSLookup_res_Env_2() DNSLookup_res_Env
	DNSLookup_Setup_2()
	IP_To_App_3() string
	ResultFrom_DNSLookup_res(result dnslookup.Res_Result)
	To_DNSLookup_res_Env() DNSLookup_res_Env
	DNSLookup_Setup()
	IP_To_App_2() string
	ResultFrom_SimpleDNS_Cached_res_2(result simpledns_cached.Res_Result)
	To_SimpleDNS_Cached_res_Env_2() SimpleDNS_Cached_res_Env
	SimpleDNS_Cached_Setup_2()
	IspDNS_Choice_2() SimpleDNS_ispDNS_Choice_2
	RecQuery_From_App_2(host string)
	Done_From_App_2()
	IP_To_App() string
	ResultFrom_SimpleDNS_Cached_res(result simpledns_cached.Res_Result)
	To_SimpleDNS_Cached_res_Env() SimpleDNS_Cached_res_Env
	SimpleDNS_Cached_Setup()
	IspDNS_Choice() SimpleDNS_ispDNS_Choice
	RecQuery_From_App(host string)
	Done() simpledns.IspDNS_Result
	Done_From_App()
}

type DNSISPDNSState struct {
	Cache   map[string]string
	Request string
}

func (d *DNSISPDNSState) IP_To_App_6() string {
	ip := d.Cache[d.Request]
	fmt.Println("dns_ispdns: returning looked-up IP '", ip, "'")
	return ip
}

func (d *DNSISPDNSState) ResultFrom_DNSLookup_res_3(result dnslookup.Res_Result) {
	d.Cache[d.Request] = result.IP
}

func (d *DNSISPDNSState) To_DNSLookup_res_Env_3() DNSLookup_res_Env {
	return &DNSLookupResState{Request: d.Request}
}

func (d *DNSISPDNSState) DNSLookup_Setup_3() {
	DNSIdx = 0
}

func (d *DNSISPDNSState) IP_To_App_5() string {
	ip := d.Cache[d.Request]
	fmt.Println("dns_ispdns: returning cached IP '", ip, "'")
	return ip
}

func (d *DNSISPDNSState) ResultFrom_SimpleDNS_Cached_res_3(result simpledns_cached.Res_Result) {
}

func (d *DNSISPDNSState) To_SimpleDNS_Cached_res_Env_3() SimpleDNS_Cached_res_Env {
	return &DNSCachedResState{}
}

func (d *DNSISPDNSState) SimpleDNS_Cached_Setup_3() {
}

func (d *DNSISPDNSState) IspDNS_Choice_3() SimpleDNS_ispDNS_Choice_3 {
	if _, ok := d.Cache[d.Request]; ok {
		return SimpleDNS_ispDNS_SimpleDNS_Cached_3
	}
	return SimpleDNS_ispDNS_DNSLookup_3
}

func (d *DNSISPDNSState) RecQuery_From_App_3(host string) {
	d.Request = host
}

func (d *DNSISPDNSState) Done_From_App_3() {
}

func (d *DNSISPDNSState) IP_To_App_4() string {
	ip := d.Cache[d.Request]
	fmt.Println("dns_ispdns: returning looked-up IP '", ip, "'")
	return ip
}

func (d *DNSISPDNSState) ResultFrom_DNSLookup_res_2(result dnslookup.Res_Result) {
	d.Cache[d.Request] = result.IP
}

func (d *DNSISPDNSState) To_DNSLookup_res_Env_2() DNSLookup_res_Env {
	return &DNSLookupResState{Request: d.Request}
}

func (d *DNSISPDNSState) DNSLookup_Setup_2() {
	DNSIdx = 0
}

func (d *DNSISPDNSState) IP_To_App_3() string {
	ip := d.Cache[d.Request]
	fmt.Println("dns_ispdns: returning cached IP '", ip, "'")
	return ip
}

func (d *DNSISPDNSState) ResultFrom_DNSLookup_res(result dnslookup.Res_Result) {
	d.Cache[d.Request] = result.IP
}

func (d *DNSISPDNSState) To_DNSLookup_res_Env() DNSLookup_res_Env {
	return &DNSLookupResState{Request: d.Request}
}

func (d *DNSISPDNSState) DNSLookup_Setup() {
	DNSIdx = 0
}

func (d *DNSISPDNSState) IP_To_App_2() string {
	ip := d.Cache[d.Request]
	fmt.Println("dns_ispdns: returning looked-up IP '", ip, "'")
	return ip
}

func (d *DNSISPDNSState) ResultFrom_SimpleDNS_Cached_res_2(result simpledns_cached.Res_Result) {
}

func (d *DNSISPDNSState) To_SimpleDNS_Cached_res_Env_2() SimpleDNS_Cached_res_Env {
	return &DNSCachedResState{}
}

func (d *DNSISPDNSState) SimpleDNS_Cached_Setup_2() {
}

func (d *DNSISPDNSState) IspDNS_Choice_2() SimpleDNS_ispDNS_Choice_2 {
	if _, ok := d.Cache[d.Request]; ok {
		return SimpleDNS_ispDNS_SimpleDNS_Cached_2
	}
	return SimpleDNS_ispDNS_DNSLookup_2
}

func (d *DNSISPDNSState) RecQuery_From_App_2(host string) {
	d.Request = host
}

func (d *DNSISPDNSState) Done_From_App_2() {
}

func (d *DNSISPDNSState) IP_To_App() string {
	ip := d.Cache[d.Request]
	fmt.Println("dns_dnsres: returning cached IP '", ip, "'")
	return ip
}

func (d *DNSISPDNSState) ResultFrom_SimpleDNS_Cached_res(result simpledns_cached.Res_Result) {
}

func (d *DNSISPDNSState) To_SimpleDNS_Cached_res_Env() SimpleDNS_Cached_res_Env {
	return &DNSCachedResState{}
}

func (d *DNSISPDNSState) SimpleDNS_Cached_Setup() {
}

func (d *DNSISPDNSState) IspDNS_Choice() SimpleDNS_ispDNS_Choice {
	if _, ok := d.Cache[d.Request]; ok {
		return SimpleDNS_ispDNS_SimpleDNS_Cached
	}
	return SimpleDNS_ispDNS_DNSLookup
}

func (d *DNSISPDNSState) RecQuery_From_App(host string) {
	d.Request = host
}

func (d *DNSISPDNSState) Done() simpledns.IspDNS_Result {
	return simpledns.IspDNS_Result{Cache: d.Cache}
}

func (d *DNSISPDNSState) Done_From_App() {
}
