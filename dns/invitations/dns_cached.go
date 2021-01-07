package invitations

import "NestedScribbleBenchmark/dns/channels/dns_cached"

type DNS_Cached_RoleSetupChan struct {
	Res_Chan chan dns_cached.Res_Chan
}

type DNS_Cached_InviteSetupChan struct {
	Res_InviteChan chan DNS_Cached_res_InviteChan
}

type DNS_Cached_res_InviteChan struct {

}