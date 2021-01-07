package roles

import "NestedScribbleBenchmark/dns/messages"
import "NestedScribbleBenchmark/dns/channels/dnslookup"
import "NestedScribbleBenchmark/dns/channels/dns_cached"
import "NestedScribbleBenchmark/dns/channels/dns"
import "NestedScribbleBenchmark/dns/invitations"
import "sync"

func DNS_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.DNS_RoleSetupChan, inviteChannels invitations.DNS_InviteSetupChan)  {
	ispdns_invite_ispdns_2 := make(chan dnslookup.Res_Chan, 1)
	ispdns_invite_ispdns_invitechan_2 := make(chan invitations.DNSLookup_res_InviteChan, 1)
	dnsres_app_string := make(chan string, 1)
	dnsres_app_label := make(chan messages.DNS_Label, 1)
	ispdns_dnsres_string := make(chan string, 1)
	ispdns_dnsres_label := make(chan messages.DNS_Label, 1)
	ispdns_invite_ispdns := make(chan dns_cached.Res_Chan, 1)
	ispdns_invite_ispdns_invitechan := make(chan invitations.DNS_Cached_res_InviteChan, 1)
	dnsres_ispdns_string := make(chan string, 1)
	app_dnsres_string := make(chan string, 1)
	dnsres_ispdns_label := make(chan messages.DNS_Label, 1)
	app_dnsres_label := make(chan messages.DNS_Label, 1)

	ispdns_chan := dns.IspDNS_Chan{
		String_To_dnsRes: ispdns_dnsres_string,
		String_From_dnsRes: dnsres_ispdns_string,
		Label_To_dnsRes: ispdns_dnsres_label,
		Label_From_dnsRes: dnsres_ispdns_label,
	}
	dnsres_chan := dns.DnsRes_Chan{
		String_To_ispDNS: dnsres_ispdns_string,
		String_To_app: dnsres_app_string,
		String_From_ispDNS: ispdns_dnsres_string,
		String_From_app: app_dnsres_string,
		Label_To_ispDNS: dnsres_ispdns_label,
		Label_To_app: dnsres_app_label,
		Label_From_ispDNS: ispdns_dnsres_label,
		Label_From_app: app_dnsres_label,
	}
	app_chan := dns.App_Chan{
		String_To_dnsRes: app_dnsres_string,
		String_From_dnsRes: dnsres_app_string,
		Label_To_dnsRes: app_dnsres_label,
		Label_From_dnsRes: dnsres_app_label,
	}

	ispdns_inviteChan := invitations.DNS_ispDNS_InviteChan{
		Invite_IspDNS_To_DNS_Cached_res_InviteChan: ispdns_invite_ispdns_invitechan,
		Invite_IspDNS_To_DNS_Cached_res: ispdns_invite_ispdns,
		Invite_IspDNS_To_DNSLookup_res_InviteChan: ispdns_invite_ispdns_invitechan_2,
		Invite_IspDNS_To_DNSLookup_res: ispdns_invite_ispdns_2,
	}
	dnsres_inviteChan := invitations.DNS_dnsRes_InviteChan{

	}
	app_inviteChan := invitations.DNS_app_InviteChan{

	}

	roleChannels.App_Chan <- app_chan
	roleChannels.DnsRes_Chan <- dnsres_chan
	roleChannels.IspDNS_Chan <- ispdns_chan

	inviteChannels.App_InviteChan <- app_inviteChan
	inviteChannels.DnsRes_InviteChan <- dnsres_inviteChan
	inviteChannels.IspDNS_InviteChan <- ispdns_inviteChan
} 