package dns

import "NestedScribbleBenchmark/dns/messages"

type App_Chan struct {
	Label_From_dnsRes chan messages.DNS_Label
	Label_To_dnsRes chan messages.DNS_Label
	String_From_dnsRes chan string
	String_To_dnsRes chan string
}

type DnsRes_Chan struct {
	Label_From_app chan messages.DNS_Label
	Label_From_ispDNS chan messages.DNS_Label
	Label_To_app chan messages.DNS_Label
	Label_To_ispDNS chan messages.DNS_Label
	String_From_app chan string
	String_From_ispDNS chan string
	String_To_app chan string
	String_To_ispDNS chan string
}

type IspDNS_Chan struct {
	Label_From_dnsRes chan messages.DNS_Label
	Label_To_dnsRes chan messages.DNS_Label
	String_From_dnsRes chan string
	String_To_dnsRes chan string
}