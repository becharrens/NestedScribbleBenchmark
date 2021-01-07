package messages

type DNS_Label int
const (
	DNSIP DNS_Label = iota
	DNSLookup_IspDNS
	DNSLookup_Res
	DNS_Cached_IspDNS
	Done
	IP
	Query
	RecQuery
	Req
)