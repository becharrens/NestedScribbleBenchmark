package messages

type SimpleDNS_Label int
const (
	DNSIP SimpleDNS_Label = iota
	DNSLookup_IspDNS
	DNSLookup_Res
	Done
	IP
	RecQuery
	Req
	SimpleDNS_Cached_IspDNS
)