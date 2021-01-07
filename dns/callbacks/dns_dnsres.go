package callbacks

import "NestedScribbleBenchmark/dns/results/dns"

type DNS_dnsRes_Env interface {
	IP_To_App_2() string
	IP_From_IspDNS_2(ip string)
	RecQuery_To_IspDNS_2() string
	Query_From_App_2(host string)
	Done_To_IspDNS_2()
	Done_From_App_2()
	IP_To_App() string
	IP_From_IspDNS(ip string)
	RecQuery_To_IspDNS() string
	Query_From_App(host string)
	Done() dns.DnsRes_Result
	Done_To_IspDNS()
	Done_From_App()
}

type DNSDNSResState struct {
	Request  string
	Response string
}

func (d *DNSDNSResState) IP_To_App_2() string {
	return d.Response
}

func (d *DNSDNSResState) IP_From_IspDNS_2(ip string) {
	d.Response = ip
}

func (d *DNSDNSResState) RecQuery_To_IspDNS_2() string {
	return d.Request
}

func (d *DNSDNSResState) Query_From_App_2(host string) {
	d.Request = host
}

func (d *DNSDNSResState) Done_To_IspDNS_2() {
}

func (d *DNSDNSResState) Done_From_App_2() {
}

func (d *DNSDNSResState) IP_To_App() string {
	return d.Response
}

func (d *DNSDNSResState) IP_From_IspDNS(ip string) {
	d.Response = ip
}

func (d *DNSDNSResState) RecQuery_To_IspDNS() string {
	return d.Request
}

func (d *DNSDNSResState) Query_From_App(host string) {
	d.Request = host
}

func (d *DNSDNSResState) Done() dns.DnsRes_Result {
	return dns.DnsRes_Result{}
}

func (d *DNSDNSResState) Done_To_IspDNS() {
}

func (d *DNSDNSResState) Done_From_App() {
}
