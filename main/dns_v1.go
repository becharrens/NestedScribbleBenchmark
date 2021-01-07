package main

// import (
// 	"NestedScribbleBenchmark/dns_v1/callbacks"
// 	"NestedScribbleBenchmark/dns_v1/protocol"
// 	dns_2 "NestedScribbleBenchmark/dns_v1/results/dns"
// 	"fmt"
// )
//
// type DNSEnvV1 struct {
// 	DNSCache     map[string]string
// 	QueryResults []string
// }
//
// func (d *DNSEnvV1) New_App_Env() callbacks.DNS_app_Env {
// 	return &callbacks.DNSAppState{
// 		Queries:   dnsQueries,
// 		Idx:       0,
// 		Responses: nil,
// 	}
// }
//
// func (d *DNSEnvV1) New_DnsRes_Env() callbacks.DNS_dnsRes_Env {
// 	return &callbacks.DNSDNSResState{
// 		Cache: make(map[string]string),
// 	}
// }
//
// func (d *DNSEnvV1) App_Result(result dns_2.App_Result) {
// 	d.QueryResults = result.IPs
// }
//
// func (d *DNSEnvV1) DnsRes_Result(result dns_2.DnsRes_Result) {
// 	d.DNSCache = result.Cache
// }
//
// func NewDNSEnvV1() *DNSEnvV1 {
// 	return &DNSEnvV1{}
// }
//
// func RunDNSV1() {
// 	fmt.Println("Running Client-Server protocol")
// 	env := NewDNSEnvV1()
// 	protocol.DNS(env)
// }
