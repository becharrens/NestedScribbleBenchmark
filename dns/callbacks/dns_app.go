package callbacks

import (
	"NestedScribbleBenchmark/dns/results/dns"
	"fmt"
)

type DNS_app_Choice_2 int

const (
	DNS_app_Done_2 DNS_app_Choice_2 = iota
	DNS_app_Query_2
)

type DNS_app_Choice int

const (
	DNS_app_Done DNS_app_Choice = iota
	DNS_app_Query
)

type DNS_app_Env interface {
	IP_From_DnsRes_2(ip string)
	Query_To_DnsRes_2() string
	Done_To_DnsRes_2()
	App_Choice_2() DNS_app_Choice_2
	IP_From_DnsRes(ip string)
	Query_To_DnsRes() string
	Done() dns.App_Result
	Done_To_DnsRes()
	App_Choice() DNS_app_Choice
}

type DNSAppState struct {
	Queries   []string
	Idx       int
	Responses []string
}

func (d *DNSAppState) IP_From_DnsRes_2(ip string) {
	d.Responses = append(d.Responses, ip)
	fmt.Println("dns_app: Received IP: '", ip, "' for host: '", d.Queries[d.Idx-1], "'")
}

func (d *DNSAppState) Query_To_DnsRes_2() string {
	req := d.Queries[d.Idx]
	d.Idx++
	return req
}

func (d *DNSAppState) Done_To_DnsRes_2() {
}

func (d *DNSAppState) App_Choice_2() DNS_app_Choice_2 {
	if d.Idx < len(d.Queries) {
		return DNS_app_Query_2
	}
	return DNS_app_Done_2
}

func (d *DNSAppState) IP_From_DnsRes(ip string) {
	d.Responses = append(d.Responses, ip)
	fmt.Println("dns_app: Received IP: '", ip, "' for host: '", d.Queries[d.Idx-1], "'")
}

func (d *DNSAppState) Query_To_DnsRes() string {
	req := d.Queries[d.Idx]
	d.Idx++
	return req
}

func (d *DNSAppState) Done() dns.App_Result {
	return dns.App_Result{
		IPs: d.Responses,
	}
}

func (d *DNSAppState) Done_To_DnsRes() {
}

func (d *DNSAppState) App_Choice() DNS_app_Choice {
	if d.Idx < len(d.Queries) {
		return DNS_app_Query
	}
	return DNS_app_Done
}
