package protocol

import "NestedScribbleBenchmark/generaldns/messages"
import "NestedScribbleBenchmark/generaldns/channels/recdnslookup"
import "NestedScribbleBenchmark/generaldns/channels/iterdnslookup"
import "NestedScribbleBenchmark/generaldns/channels/generaldns"
import "NestedScribbleBenchmark/generaldns/channels/cached"
import "NestedScribbleBenchmark/generaldns/invitations"
import "NestedScribbleBenchmark/generaldns/callbacks"
import generaldns_2 "NestedScribbleBenchmark/generaldns/results/generaldns"
import "NestedScribbleBenchmark/generaldns/roles"
import "sync"

type GeneralDNS_Env interface {
	New_App_Env() callbacks.GeneralDNS_app_Env
	New_DnsRes_Env() callbacks.GeneralDNS_dnsRes_Env
	App_Result(result generaldns_2.App_Result) 
	DnsRes_Result(result generaldns_2.DnsRes_Result) 
}

func Start_GeneralDNS_app(protocolEnv GeneralDNS_Env, wg *sync.WaitGroup, roleChannels generaldns.App_Chan, inviteChannels invitations.GeneralDNS_app_InviteChan, env callbacks.GeneralDNS_app_Env)  {
	defer wg.Done()
	result := roles.GeneralDNS_app(wg, roleChannels, inviteChannels, env)
	protocolEnv.App_Result(result)
} 

func Start_GeneralDNS_dnsRes(protocolEnv GeneralDNS_Env, wg *sync.WaitGroup, roleChannels generaldns.DnsRes_Chan, inviteChannels invitations.GeneralDNS_dnsRes_InviteChan, env callbacks.GeneralDNS_dnsRes_Env)  {
	defer wg.Done()
	result := roles.GeneralDNS_dnsRes(wg, roleChannels, inviteChannels, env)
	protocolEnv.DnsRes_Result(result)
} 

func GeneralDNS(protocolEnv GeneralDNS_Env)  {
	dnsres_invite_dnsres_3 := make(chan cached.Res_Chan, 1)
	dnsres_invite_dnsres_invitechan_3 := make(chan invitations.Cached_res_InviteChan, 1)
	dnsres_invite_dnsres_2 := make(chan iterdnslookup.Res_Chan, 1)
	dnsres_invite_dnsres_invitechan_2 := make(chan invitations.IterDNSLookup_res_InviteChan, 1)
	dnsres_app_string := make(chan string, 1)
	dnsres_app_label := make(chan messages.GeneralDNS_Label, 1)
	dnsres_invite_dnsres := make(chan recdnslookup.Res_Chan, 1)
	dnsres_invite_dnsres_invitechan := make(chan invitations.RecDNSLookup_res_InviteChan, 1)
	app_dnsres_string := make(chan string, 1)
	app_dnsres_label := make(chan messages.GeneralDNS_Label, 1)

	dnsres_chan := generaldns.DnsRes_Chan{
		String_To_app: dnsres_app_string,
		String_From_app: app_dnsres_string,
		Label_To_app: dnsres_app_label,
		Label_From_app: app_dnsres_label,
	}
	app_chan := generaldns.App_Chan{
		String_To_dnsRes: app_dnsres_string,
		String_From_dnsRes: dnsres_app_string,
		Label_To_dnsRes: app_dnsres_label,
		Label_From_dnsRes: dnsres_app_label,
	}

	dnsres_inviteChan := invitations.GeneralDNS_dnsRes_InviteChan{
		Invite_DnsRes_To_RecDNSLookup_res_InviteChan: dnsres_invite_dnsres_invitechan,
		Invite_DnsRes_To_RecDNSLookup_res: dnsres_invite_dnsres,
		Invite_DnsRes_To_IterDNSLookup_res_InviteChan: dnsres_invite_dnsres_invitechan_2,
		Invite_DnsRes_To_IterDNSLookup_res: dnsres_invite_dnsres_2,
		Invite_DnsRes_To_Cached_res_InviteChan: dnsres_invite_dnsres_invitechan_3,
		Invite_DnsRes_To_Cached_res: dnsres_invite_dnsres_3,
	}
	app_inviteChan := invitations.GeneralDNS_app_InviteChan{

	}

	var wg sync.WaitGroup

	wg.Add(2)

	app_env := protocolEnv.New_App_Env()
	dnsres_env := protocolEnv.New_DnsRes_Env()

	go Start_GeneralDNS_app(protocolEnv, &wg, app_chan, app_inviteChan, app_env)
	go Start_GeneralDNS_dnsRes(protocolEnv, &wg, dnsres_chan, dnsres_inviteChan, dnsres_env)

	wg.Wait()
} 