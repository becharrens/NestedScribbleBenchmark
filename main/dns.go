package main

import (
	"NestedScribbleBenchmark/dns/callbacks"
	"NestedScribbleBenchmark/dns/protocol"
	dns_2 "NestedScribbleBenchmark/dns/results/dns"
	"fmt"
)

var dnsQueries = []string{"www.example.com", "www.ecoop21.com", "tour.golang.org", "localhost", "wikipedia.org", "github.com", "google.com", "wikipedia.org"}

type DNSEnv struct {
	DNSCache     map[string]string
	QueryResults []string
}

func (d *DNSEnv) New_IspDNS_Env() callbacks.DNS_ispDNS_Env {
	return &callbacks.DNSISPDNSState{
		Cache: make(map[string]string),
	}
}

func (d *DNSEnv) IspDNS_Result(result dns_2.IspDNS_Result) {
	d.DNSCache = result.Cache
}

func (d *DNSEnv) New_App_Env() callbacks.DNS_app_Env {
	return &callbacks.DNSAppState{
		Queries:   dnsQueries,
		Idx:       0,
		Responses: nil,
	}
}

func (d *DNSEnv) New_DnsRes_Env() callbacks.DNS_dnsRes_Env {
	return &callbacks.DNSDNSResState{}
}

func (d *DNSEnv) App_Result(result dns_2.App_Result) {
	d.QueryResults = result.IPs
}

func (d *DNSEnv) DnsRes_Result(result dns_2.DnsRes_Result) {
}

func NewDNSEnv() *DNSEnv {
	return &DNSEnv{}
}

func RunDNS() {
	fmt.Println("Running Client-Server protocol")
	env := NewDNSEnv()
	protocol.DNS(env)
}
