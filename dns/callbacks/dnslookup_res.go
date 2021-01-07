package callbacks

import (
	"NestedScribbleBenchmark/dns/results/dnslookup"
	"fmt"
)

type DNSLookup_res_Env interface {
	ResultFrom_DNSLookup_res(result dnslookup.Res_Result)
	To_DNSLookup_res_Env() DNSLookup_res_Env
	DNSLookup_Setup()
	DNSIP_From_Dns(ip string)
	Done() dnslookup.Res_Result
	IP_From_Dns(ip string)
	Req_To_Dns() string
}

type DNSLookupResState struct {
	Request string
	IP      string
}

func (d *DNSLookupResState) ResultFrom_DNSLookup_res(result dnslookup.Res_Result) {
	// d.Response = result.IP
}

func (d *DNSLookupResState) To_DNSLookup_res_Env() DNSLookup_res_Env {
	return d
}

func (d *DNSLookupResState) DNSLookup_Setup() {
}

func (d *DNSLookupResState) DNSIP_From_Dns(ip string) {
	fmt.Println("dnslookup_res: received dns ip '", ip, "'")
}

func (d *DNSLookupResState) Done() dnslookup.Res_Result {
	return dnslookup.Res_Result{
		IP: d.IP,
	}
}

func (d *DNSLookupResState) IP_From_Dns(ip string) {
	d.IP = ip
}

func (d *DNSLookupResState) Req_To_Dns() string {
	return d.Request
}
