package main

import (
	callbacks2 "NestedScribbleBenchmark/generaldns/callbacks"
	"NestedScribbleBenchmark/generaldns/protocol"
	generaldns_2 "NestedScribbleBenchmark/generaldns/results/generaldns"
	"fmt"
)

type GeneralDNSEnv struct {
	DNSCache     map[string]string
	QueryResults []string
}

func (d *GeneralDNSEnv) New_App_Env() callbacks2.GeneralDNS_app_Env {
	return &callbacks2.GeneralDNSAppState{
		Queries:   dnsQueries,
		Idx:       0,
		Responses: nil,
	}
}

func (d *GeneralDNSEnv) New_DnsRes_Env() callbacks2.GeneralDNS_dnsRes_Env {
	return &callbacks2.GeneralDNSDNSResState{
		Cache: make(map[string]string),
	}
}

func (d *GeneralDNSEnv) App_Result(result generaldns_2.App_Result) {
	d.QueryResults = result.IPs
}

func (d *GeneralDNSEnv) DnsRes_Result(result generaldns_2.DnsRes_Result) {
	d.DNSCache = result.Cache
}

func NewGeneralDNSEnv() *GeneralDNSEnv {
	return &GeneralDNSEnv{}
}

func RunGeneralDNS() {
	fmt.Println("Running Client-Server protocol")
	env := NewGeneralDNSEnv()
	protocol.GeneralDNS(env)
}
