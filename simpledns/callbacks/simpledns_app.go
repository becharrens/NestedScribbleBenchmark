package callbacks

import (
	"NestedScribbleBenchmark/simpledns/results/simpledns"
	"fmt"
)

type SimpleDNS_app_Choice_2 int

const (
	SimpleDNS_app_Done_2 SimpleDNS_app_Choice_2 = iota
	SimpleDNS_app_RecQuery_2
)

type SimpleDNS_app_Choice int

const (
	SimpleDNS_app_Done SimpleDNS_app_Choice = iota
	SimpleDNS_app_RecQuery
)

type SimpleDNS_app_Env interface {
	IP_From_IspDNS_2(ip string)
	RecQuery_To_IspDNS_2() string
	Done_To_IspDNS_2()
	App_Choice_2() SimpleDNS_app_Choice_2
	IP_From_IspDNS(ip string)
	RecQuery_To_IspDNS() string
	Done() simpledns.App_Result
	Done_To_IspDNS()
	App_Choice() SimpleDNS_app_Choice
}

type DNSAppState struct {
	Queries   []string
	Idx       int
	Responses []string
}

func (d *DNSAppState) IP_From_IspDNS_2(ip string) {
	d.Responses = append(d.Responses, ip)
	fmt.Println("dns_app: Received IP: '", ip, "' for host: '", d.Queries[d.Idx-1], "'")
}

func (d *DNSAppState) RecQuery_To_IspDNS_2() string {
	req := d.Queries[d.Idx]
	d.Idx++
	return req
}

func (d *DNSAppState) Done_To_IspDNS_2() {
}

func (d *DNSAppState) App_Choice_2() SimpleDNS_app_Choice_2 {
	if d.Idx < len(d.Queries) {
		return SimpleDNS_app_RecQuery_2
	}
	return SimpleDNS_app_Done_2
}

func (d *DNSAppState) IP_From_IspDNS(ip string) {
	d.Responses = append(d.Responses, ip)
	fmt.Println("dns_app: Received IP: '", ip, "' for host: '", d.Queries[d.Idx-1], "'")
}

func (d *DNSAppState) RecQuery_To_IspDNS() string {
	req := d.Queries[d.Idx]
	d.Idx++
	return req
}

func (d *DNSAppState) Done() simpledns.App_Result {
	return simpledns.App_Result{
		IPs: d.Responses,
	}
}

func (d *DNSAppState) Done_To_IspDNS() {
}

func (d *DNSAppState) App_Choice() SimpleDNS_app_Choice {
	if d.Idx < len(d.Queries) {
		return SimpleDNS_app_RecQuery
	}
	return SimpleDNS_app_Done
}
