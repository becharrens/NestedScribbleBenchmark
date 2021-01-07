package invitations

import "NestedScribbleBenchmark/dns/channels/dnslookup"
import "NestedScribbleBenchmark/dns/channels/dns_cached"
import "NestedScribbleBenchmark/dns/channels/dns"

type DNS_RoleSetupChan struct {
	App_Chan chan dns.App_Chan
	DnsRes_Chan chan dns.DnsRes_Chan
	IspDNS_Chan chan dns.IspDNS_Chan
}

type DNS_InviteSetupChan struct {
	App_InviteChan chan DNS_app_InviteChan
	DnsRes_InviteChan chan DNS_dnsRes_InviteChan
	IspDNS_InviteChan chan DNS_ispDNS_InviteChan
}

type DNS_app_InviteChan struct {

}

type DNS_dnsRes_InviteChan struct {

}

type DNS_ispDNS_InviteChan struct {
	Invite_IspDNS_To_DNSLookup_res chan dnslookup.Res_Chan
	Invite_IspDNS_To_DNSLookup_res_InviteChan chan DNSLookup_res_InviteChan
	Invite_IspDNS_To_DNS_Cached_res chan dns_cached.Res_Chan
	Invite_IspDNS_To_DNS_Cached_res_InviteChan chan DNS_Cached_res_InviteChan
}