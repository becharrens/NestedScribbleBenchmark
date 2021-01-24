package invitations

import "NestedScribbleBenchmark/generaldns/channels/recdnslookup"
import "NestedScribbleBenchmark/generaldns/channels/iterdnslookup"
import "NestedScribbleBenchmark/generaldns/channels/generaldns"
import "NestedScribbleBenchmark/generaldns/channels/cached"

type GeneralDNS_RoleSetupChan struct {
	App_Chan chan generaldns.App_Chan
	DnsRes_Chan chan generaldns.DnsRes_Chan
}

type GeneralDNS_InviteSetupChan struct {
	App_InviteChan chan GeneralDNS_app_InviteChan
	DnsRes_InviteChan chan GeneralDNS_dnsRes_InviteChan
}

type GeneralDNS_app_InviteChan struct {

}

type GeneralDNS_dnsRes_InviteChan struct {
	Invite_DnsRes_To_Cached_res chan cached.Res_Chan
	Invite_DnsRes_To_Cached_res_InviteChan chan Cached_res_InviteChan
	Invite_DnsRes_To_IterDNSLookup_res chan iterdnslookup.Res_Chan
	Invite_DnsRes_To_IterDNSLookup_res_InviteChan chan IterDNSLookup_res_InviteChan
	Invite_DnsRes_To_RecDNSLookup_res chan recdnslookup.Res_Chan
	Invite_DnsRes_To_RecDNSLookup_res_InviteChan chan RecDNSLookup_res_InviteChan
}