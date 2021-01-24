package callbacks

import (
	"NestedScribbleBenchmark/generaldns/results/recdnslookup"
	"fmt"
)

type RecDNSLookup_res_Env interface {
	Done() recdnslookup.Res_Result
	IP_From_Dns(ip string)
	RecReq_To_Dns() string
}

type RecDNSLookupResState struct {
	Request string
	IP      string
}

func (r *RecDNSLookupResState) Done() recdnslookup.Res_Result {
	return recdnslookup.Res_Result{
		IP: r.IP,
	}
}

func (r *RecDNSLookupResState) IP_From_Dns(ip string) {
	fmt.Println("recdnslookup_res: Received IP '", ip, "' from dns")
	r.IP = ip
}

func (r *RecDNSLookupResState) RecReq_To_Dns() string {
	return r.Request
}
