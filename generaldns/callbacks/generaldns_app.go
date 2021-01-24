package callbacks

import (
	"NestedScribbleBenchmark/generaldns/results/generaldns"
	"fmt"
)

type GeneralDNS_app_Choice_2 int

const (
	GeneralDNS_app_Done_2 GeneralDNS_app_Choice_2 = iota
	GeneralDNS_app_Query_2
)

type GeneralDNS_app_Choice int

const (
	GeneralDNS_app_Done GeneralDNS_app_Choice = iota
	GeneralDNS_app_Query
)

type GeneralDNS_app_Env interface {
	IP_From_DnsRes_2(ip string)
	Query_To_DnsRes_2() string
	Done_To_DnsRes_2()
	App_Choice_2() GeneralDNS_app_Choice_2
	IP_From_DnsRes(ip string)
	Query_To_DnsRes() string
	Done() generaldns.App_Result
	Done_To_DnsRes()
	App_Choice() GeneralDNS_app_Choice
}

type GeneralDNSAppState struct {
	Queries   []string
	Idx       int
	Responses []string
}

func (d *GeneralDNSAppState) IP_From_DnsRes_2(ip string) {
	d.Responses = append(d.Responses, ip)
	fmt.Println("generaldns_app: Received IP: '", ip, "' for host: '", d.Queries[d.Idx-1], "'")
}

func (d *GeneralDNSAppState) Query_To_DnsRes_2() string {
	req := d.Queries[d.Idx]
	d.Idx++
	return req
}

func (d *GeneralDNSAppState) Done_To_DnsRes_2() {
}

func (d *GeneralDNSAppState) App_Choice_2() GeneralDNS_app_Choice_2 {
	if d.Idx < len(d.Queries) {
		return GeneralDNS_app_Query_2
	}
	return GeneralDNS_app_Done_2
}

func (d *GeneralDNSAppState) IP_From_DnsRes(ip string) {
	d.Responses = append(d.Responses, ip)
	fmt.Println("generaldns_app: Received IP: '", ip, "' for host: '", d.Queries[d.Idx-1], "'")
}

func (d *GeneralDNSAppState) Query_To_DnsRes() string {
	req := d.Queries[d.Idx]
	d.Idx++
	return req
}

func (d *GeneralDNSAppState) Done() generaldns.App_Result {
	return generaldns.App_Result{
		IPs: d.Responses,
	}
}

func (d *GeneralDNSAppState) Done_To_DnsRes() {
}

func (d *GeneralDNSAppState) App_Choice() GeneralDNS_app_Choice {
	if d.Idx < len(d.Queries) {
		return GeneralDNS_app_Query
	}
	return GeneralDNS_app_Done
}
