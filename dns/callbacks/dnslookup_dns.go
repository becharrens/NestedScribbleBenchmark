package callbacks

type DNSLookup_dns_Choice int

const (
	DNSLookup_dns_IP DNSLookup_dns_Choice = iota
	DNSLookup_dns_DNSIP
)

type DNSLookup_dns_Env interface {
	DNSIP_To_Res() string
	Done()
	IP_To_Res() string
	Dns_Choice() DNSLookup_dns_Choice
	Req_From_Res(host string)
}

var (
	DNSCache = []map[string]string{
		{"www.example.com": "31.312.12.12", "localhost": "127.0.0.1"},
		{"www.ecoop21.com": "12.34.56.78", "wikipedia.org": "98.76.54.32"},
		{"tour.golang.org": "142.250.68.17", "github.com": "32.5.3.123", "google.com": "1.1.2.2"},
	}
	DNSIPs = []string{"2.32.53.64", "54.96.83.53", "174.32.32.12"}
	DNSIdx = 0
)

type DNSLookupDNSState struct {
	KnownIPS   map[string]string
	NextDNSIdx int
	Request    string
}

func (d *DNSLookupDNSState) DNSIP_To_Res() string {
	return DNSIPs[d.NextDNSIdx]
}

func (d *DNSLookupDNSState) Done() {
}

func (d *DNSLookupDNSState) IP_To_Res() string {
	// If host is not in map, it will return ""
	return d.KnownIPS[d.Request]
}

func (d *DNSLookupDNSState) Dns_Choice() DNSLookup_dns_Choice {
	if _, ok := d.KnownIPS[d.Request]; ok || d.NextDNSIdx >= len(DNSIPs) {
		return DNSLookup_dns_IP
	}
	return DNSLookup_dns_DNSIP
}

func (d *DNSLookupDNSState) Req_From_Res(host string) {
	d.Request = host
}

func New_DNSLookup_dns_State() DNSLookup_dns_Env {
	nextDNSIdx := (DNSIdx + 1) % len(DNSIPs)
	dnsState := &DNSLookupDNSState{
		KnownIPS:   DNSCache[DNSIdx],
		NextDNSIdx: nextDNSIdx,
	}
	DNSIdx = nextDNSIdx
	return dnsState
}
