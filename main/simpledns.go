package main

import (
	"NestedScribbleBenchmark/simpledns/callbacks"
	"NestedScribbleBenchmark/simpledns/protocol"
	simpledns_2 "NestedScribbleBenchmark/simpledns/results/simpledns"
	"fmt"
)

type SimpleDNSEnv struct {
	DNSCache     map[string]string
	QueryResults []string
}

func (d *SimpleDNSEnv) New_IspDNS_Env() callbacks.SimpleDNS_ispDNS_Env {
	return &callbacks.DNSISPDNSState{
		Cache: make(map[string]string),
	}
}

func (d *SimpleDNSEnv) IspDNS_Result(result simpledns_2.IspDNS_Result) {
	d.DNSCache = result.Cache
}

func (d *SimpleDNSEnv) New_App_Env() callbacks.SimpleDNS_app_Env {
	return &callbacks.DNSAppState{
		Queries:   dnsQueries,
		Idx:       0,
		Responses: nil,
	}
}

func (d *SimpleDNSEnv) App_Result(result simpledns_2.App_Result) {
	d.QueryResults = result.IPs
}

func NewSimpleDNSEnv() *SimpleDNSEnv {
	return &SimpleDNSEnv{}
}

func RunSimpleDNS() {
	fmt.Println("Running Client-Server protocol")
	env := NewSimpleDNSEnv()
	protocol.SimpleDNS(env)
}
