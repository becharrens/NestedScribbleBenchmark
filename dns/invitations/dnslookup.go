package invitations

import "NestedScribbleBenchmark/dns/channels/dnslookup"

type DNSLookup_RoleSetupChan struct {
	Res_Chan chan dnslookup.Res_Chan
}

type DNSLookup_InviteSetupChan struct {
	Res_InviteChan chan DNSLookup_res_InviteChan
}

type DNSLookup_res_InviteChan struct {
	Invite_Res_To_DNSLookup_res chan dnslookup.Res_Chan
	Invite_Res_To_DNSLookup_res_InviteChan chan DNSLookup_res_InviteChan
}

type DNSLookup_dns_InviteChan struct {

}