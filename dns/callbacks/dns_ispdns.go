package callbacks

import (
	"NestedScribbleBenchmark/dns/results/dns"
	"NestedScribbleBenchmark/dns/results/dns_cached"
	"NestedScribbleBenchmark/dns/results/dnslookup"
	"fmt"
)

type DNS_ispDNS_Choice_3 int

const (
	DNS_ispDNS_DNS_Cached_3 DNS_ispDNS_Choice_3 = iota
	DNS_ispDNS_DNSLookup_3
)

type DNS_ispDNS_Choice_2 int

const (
	DNS_ispDNS_DNS_Cached_2 DNS_ispDNS_Choice_2 = iota
	DNS_ispDNS_DNSLookup_2
)

type DNS_ispDNS_Choice int

const (
	DNS_ispDNS_DNS_Cached DNS_ispDNS_Choice = iota
	DNS_ispDNS_DNSLookup
)

type DNS_ispDNS_Env interface {
	Done_From_DnsRes_3()
	Done_From_DnsRes_2()
	IP_To_DnsRes_6() string
	ResultFrom_DNSLookup_res_3(result dnslookup.Res_Result)
	To_DNSLookup_res_Env_3() DNSLookup_res_Env
	DNSLookup_Setup_3()
	IP_To_DnsRes_5() string
	ResultFrom_DNS_Cached_res_3(result dns_cached.Res_Result)
	To_DNS_Cached_res_Env_3() DNS_Cached_res_Env
	DNS_Cached_Setup_3()
	IspDNS_Choice_3() DNS_ispDNS_Choice_3
	RecQuery_From_DnsRes_3(host string)
	IP_To_DnsRes_4() string
	ResultFrom_DNSLookup_res_2(result dnslookup.Res_Result)
	To_DNSLookup_res_Env_2() DNSLookup_res_Env
	DNSLookup_Setup_2()
	Done() dns.IspDNS_Result
	Done_From_DnsRes()
	IP_To_DnsRes_3() string
	ResultFrom_DNSLookup_res(result dnslookup.Res_Result)
	To_DNSLookup_res_Env() DNSLookup_res_Env
	DNSLookup_Setup()
	IP_To_DnsRes_2() string
	ResultFrom_DNS_Cached_res_2(result dns_cached.Res_Result)
	To_DNS_Cached_res_Env_2() DNS_Cached_res_Env
	DNS_Cached_Setup_2()
	IspDNS_Choice_2() DNS_ispDNS_Choice_2
	RecQuery_From_DnsRes_2(host string)
	IP_To_DnsRes() string
	ResultFrom_DNS_Cached_res(result dns_cached.Res_Result)
	To_DNS_Cached_res_Env() DNS_Cached_res_Env
	DNS_Cached_Setup()
	IspDNS_Choice() DNS_ispDNS_Choice
	RecQuery_From_DnsRes(host string)
}

type DNSISPDNSState struct {
	Cache   map[string]string
	Request string
}

func (d *DNSISPDNSState) Done_From_DnsRes_3() {
}

func (d *DNSISPDNSState) Done_From_DnsRes_2() {
}

func (d *DNSISPDNSState) IP_To_DnsRes_6() string {
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

func (d *DNSISPDNSState) IP_To_DnsRes_5() string {
	ip := d.Cache[d.Request]
	fmt.Println("dns_ispdns: returning cached IP '", ip, "'")
	return ip
}

func (d *DNSISPDNSState) ResultFrom_DNS_Cached_res_3(result dns_cached.Res_Result) {
}

func (d *DNSISPDNSState) To_DNS_Cached_res_Env_3() DNS_Cached_res_Env {
	return &DNSCachedResState{}
}

func (d *DNSISPDNSState) DNS_Cached_Setup_3() {
}

func (d *DNSISPDNSState) IspDNS_Choice_3() DNS_ispDNS_Choice_3 {
	if _, ok := d.Cache[d.Request]; ok {
		return DNS_ispDNS_DNS_Cached_3
	}
	return DNS_ispDNS_DNSLookup_3
}

func (d *DNSISPDNSState) RecQuery_From_DnsRes_3(host string) {
	d.Request = host
}

func (d *DNSISPDNSState) IP_To_DnsRes_4() string {
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

func (d *DNSISPDNSState) Done() dns.IspDNS_Result {
	return dns.IspDNS_Result{Cache: d.Cache}
}

func (d *DNSISPDNSState) Done_From_DnsRes() {
}

func (d *DNSISPDNSState) IP_To_DnsRes_3() string {
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

func (d *DNSISPDNSState) IP_To_DnsRes_2() string {
	ip := d.Cache[d.Request]
	fmt.Println("dns_ispdns: returning looked-up IP '", ip, "'")
	return ip
}

func (d *DNSISPDNSState) ResultFrom_DNS_Cached_res_2(result dns_cached.Res_Result) {
}

func (d *DNSISPDNSState) To_DNS_Cached_res_Env_2() DNS_Cached_res_Env {
	return &DNSCachedResState{}
}

func (d *DNSISPDNSState) DNS_Cached_Setup_2() {
}

func (d *DNSISPDNSState) IspDNS_Choice_2() DNS_ispDNS_Choice_2 {
	if _, ok := d.Cache[d.Request]; ok {
		return DNS_ispDNS_DNS_Cached_2
	}
	return DNS_ispDNS_DNSLookup_2
}

func (d *DNSISPDNSState) RecQuery_From_DnsRes_2(host string) {
	d.Request = host
}

func (d *DNSISPDNSState) IP_To_DnsRes() string {
	ip := d.Cache[d.Request]
	fmt.Println("dns_dnsres: returning cached IP '", ip, "'")
	return ip
}

func (d *DNSISPDNSState) ResultFrom_DNS_Cached_res(result dns_cached.Res_Result) {
}

func (d *DNSISPDNSState) To_DNS_Cached_res_Env() DNS_Cached_res_Env {
	return &DNSCachedResState{}
}

func (d *DNSISPDNSState) DNS_Cached_Setup() {
}

func (d *DNSISPDNSState) IspDNS_Choice() DNS_ispDNS_Choice {
	if _, ok := d.Cache[d.Request]; ok {
		return DNS_ispDNS_DNS_Cached
	}
	return DNS_ispDNS_DNSLookup
}

func (d *DNSISPDNSState) RecQuery_From_DnsRes(host string) {
	d.Request = host
}
