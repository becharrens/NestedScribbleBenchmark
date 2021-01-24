package invitations

import "NestedScribbleBenchmark/generaldns/channels/recdnslookup"
import "NestedScribbleBenchmark/generaldns/channels/iterdnslookup"
import "NestedScribbleBenchmark/generaldns/channels/cached"

type RecDNSLookup_RoleSetupChan struct {
	Res_Chan chan recdnslookup.Res_Chan
}

type RecDNSLookup_InviteSetupChan struct {
	Res_InviteChan chan RecDNSLookup_res_InviteChan
}

type RecDNSLookup_res_InviteChan struct {

}

type RecDNSLookup_dns_InviteChan struct {
	Invite_Dns_To_Cached_res chan cached.Res_Chan
	Invite_Dns_To_Cached_res_InviteChan chan Cached_res_InviteChan
	Invite_Dns_To_IterDNSLookup_res chan iterdnslookup.Res_Chan
	Invite_Dns_To_IterDNSLookup_res_InviteChan chan IterDNSLookup_res_InviteChan
	Invite_Dns_To_RecDNSLookup_res chan recdnslookup.Res_Chan
	Invite_Dns_To_RecDNSLookup_res_InviteChan chan RecDNSLookup_res_InviteChan
}