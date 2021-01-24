package callbacks

type IterDNSLookup_dns_Choice int

const (
	IterDNSLookup_dns_IP IterDNSLookup_dns_Choice = iota
	IterDNSLookup_dns_DNSIP
)

type IterDNSLookup_dns_Env interface {
	DNSIP_To_Res() string
	Done()
	IP_To_Res() string
	Dns_Choice() IterDNSLookup_dns_Choice
	IterReq_From_Res(host string)
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

type IterDNSLookupDNSState struct {
	KnownIPS   map[string]string
	NextDNSIdx int
	Request    string
}

func (d *IterDNSLookupDNSState) DNSIP_To_Res() string {
	return DNSIPs[d.NextDNSIdx]
}

func (d *IterDNSLookupDNSState) Done() {
}

func (d *IterDNSLookupDNSState) IP_To_Res() string {
	// If host is not in map, it will return ""
	return d.KnownIPS[d.Request]
}

func (d *IterDNSLookupDNSState) Dns_Choice() IterDNSLookup_dns_Choice {
	if _, ok := d.KnownIPS[d.Request]; ok || d.NextDNSIdx >= len(DNSIPs) {
		return IterDNSLookup_dns_IP
	}
	return IterDNSLookup_dns_DNSIP
}

func (d *IterDNSLookupDNSState) IterReq_From_Res(host string) {
	d.Request = host
}

func New_IterDNSLookup_dns_State() IterDNSLookup_dns_Env {
	nextDNSIdx := (DNSIdx + 1) % len(DNSIPs)
	dnsState := &IterDNSLookupDNSState{
		KnownIPS:   DNSCache[DNSIdx],
		NextDNSIdx: nextDNSIdx,
	}
	DNSIdx = nextDNSIdx
	return dnsState
}
