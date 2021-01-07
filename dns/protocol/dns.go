package protocol

import "NestedScribbleBenchmark/dns/messages"
import "NestedScribbleBenchmark/dns/channels/dnslookup"
import "NestedScribbleBenchmark/dns/channels/dns_cached"
import "NestedScribbleBenchmark/dns/channels/dns"
import "NestedScribbleBenchmark/dns/invitations"
import "NestedScribbleBenchmark/dns/callbacks"
import dns_2 "NestedScribbleBenchmark/dns/results/dns"
import "NestedScribbleBenchmark/dns/roles"
import "sync"

type DNS_Env interface {
	New_App_Env() callbacks.DNS_app_Env
	New_DnsRes_Env() callbacks.DNS_dnsRes_Env
	New_IspDNS_Env() callbacks.DNS_ispDNS_Env
	App_Result(result dns_2.App_Result) 
	DnsRes_Result(result dns_2.DnsRes_Result) 
	IspDNS_Result(result dns_2.IspDNS_Result) 
}

func Start_DNS_app(protocolEnv DNS_Env, wg *sync.WaitGroup, roleChannels dns.App_Chan, inviteChannels invitations.DNS_app_InviteChan, env callbacks.DNS_app_Env)  {
	defer wg.Done()
	result := roles.DNS_app(wg, roleChannels, inviteChannels, env)
	protocolEnv.App_Result(result)
} 

func Start_DNS_dnsRes(protocolEnv DNS_Env, wg *sync.WaitGroup, roleChannels dns.DnsRes_Chan, inviteChannels invitations.DNS_dnsRes_InviteChan, env callbacks.DNS_dnsRes_Env)  {
	defer wg.Done()
	result := roles.DNS_dnsRes(wg, roleChannels, inviteChannels, env)
	protocolEnv.DnsRes_Result(result)
} 

func Start_DNS_ispDNS(protocolEnv DNS_Env, wg *sync.WaitGroup, roleChannels dns.IspDNS_Chan, inviteChannels invitations.DNS_ispDNS_InviteChan, env callbacks.DNS_ispDNS_Env)  {
	defer wg.Done()
	result := roles.DNS_ispDNS(wg, roleChannels, inviteChannels, env)
	protocolEnv.IspDNS_Result(result)
} 

func DNS(protocolEnv DNS_Env)  {
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

	var wg sync.WaitGroup

	wg.Add(3)

	app_env := protocolEnv.New_App_Env()
	dnsres_env := protocolEnv.New_DnsRes_Env()
	ispdns_env := protocolEnv.New_IspDNS_Env()

	go Start_DNS_app(protocolEnv, &wg, app_chan, app_inviteChan, app_env)
	go Start_DNS_dnsRes(protocolEnv, &wg, dnsres_chan, dnsres_inviteChan, dnsres_env)
	go Start_DNS_ispDNS(protocolEnv, &wg, ispdns_chan, ispdns_inviteChan, ispdns_env)

	wg.Wait()
} 