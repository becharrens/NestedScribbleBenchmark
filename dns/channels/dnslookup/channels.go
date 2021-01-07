package dnslookup

import "NestedScribbleBenchmark/dns/messages"

type Res_Chan struct {
	Label_From_dns chan messages.DNS_Label
	Label_To_dns chan messages.DNS_Label
	String_From_dns chan string
	String_To_dns chan string
}

type Dns_Chan struct {
	Label_From_res chan messages.DNS_Label
	Label_To_res chan messages.DNS_Label
	String_From_res chan string
	String_To_res chan string
}