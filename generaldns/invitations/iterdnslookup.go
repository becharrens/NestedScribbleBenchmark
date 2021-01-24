package invitations

import "NestedScribbleBenchmark/generaldns/channels/iterdnslookup"

type IterDNSLookup_RoleSetupChan struct {
	Res_Chan chan iterdnslookup.Res_Chan
}

type IterDNSLookup_InviteSetupChan struct {
	Res_InviteChan chan IterDNSLookup_res_InviteChan
}

type IterDNSLookup_res_InviteChan struct {
	Invite_Res_To_IterDNSLookup_res chan iterdnslookup.Res_Chan
	Invite_Res_To_IterDNSLookup_res_InviteChan chan IterDNSLookup_res_InviteChan
}

type IterDNSLookup_dns_InviteChan struct {

}