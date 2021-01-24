package generaldns

import "NestedScribbleBenchmark/generaldns/messages"

type App_Chan struct {
	Label_From_dnsRes chan messages.GeneralDNS_Label
	Label_To_dnsRes chan messages.GeneralDNS_Label
	String_From_dnsRes chan string
	String_To_dnsRes chan string
}

type DnsRes_Chan struct {
	Label_From_app chan messages.GeneralDNS_Label
	Label_To_app chan messages.GeneralDNS_Label
	String_From_app chan string
	String_To_app chan string
}