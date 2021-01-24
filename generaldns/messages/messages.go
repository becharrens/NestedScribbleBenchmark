package messages

type GeneralDNS_Label int
const (
	Cached_Dns GeneralDNS_Label = iota
	Cached_DnsRes
	DNSIP
	Done
	IP
	IterDNSLookup_Dns
	IterDNSLookup_DnsRes
	IterDNSLookup_Res
	IterReq
	Query
	RecDNSLookup_Dns
	RecDNSLookup_DnsRes
	RecReq
)