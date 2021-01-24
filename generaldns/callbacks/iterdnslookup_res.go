package callbacks

import (
	"NestedScribbleBenchmark/generaldns/results/iterdnslookup"
	"fmt"
)

type IterDNSLookup_res_Env interface {
	ResultFrom_IterDNSLookup_res(result iterdnslookup.Res_Result)
	To_IterDNSLookup_res_Env() IterDNSLookup_res_Env
	IterDNSLookup_Setup()
	DNSIP_From_Dns(ip string)
	Done() iterdnslookup.Res_Result
	IP_From_Dns(ip string)
	IterReq_To_Dns() string
}

type IterDNSLookupResState struct {
	Request string
	IP      string
}

func (d *IterDNSLookupResState) ResultFrom_IterDNSLookup_res(result iterdnslookup.Res_Result) {
	// d.Response = result.IP
}

func (d *IterDNSLookupResState) To_IterDNSLookup_res_Env() IterDNSLookup_res_Env {
	return d
}

func (d *IterDNSLookupResState) IterDNSLookup_Setup() {
}

func (d *IterDNSLookupResState) DNSIP_From_Dns(ip string) {
	fmt.Println("iterdnslookup_res: received dns ip '", ip, "'")
}

func (d *IterDNSLookupResState) Done() iterdnslookup.Res_Result {
	return iterdnslookup.Res_Result{
		IP: d.IP,
	}
}

func (d *IterDNSLookupResState) IP_From_Dns(ip string) {
	d.IP = ip
}

func (d *IterDNSLookupResState) IterReq_To_Dns() string {
	return d.Request
}
