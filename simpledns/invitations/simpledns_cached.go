package invitations

import "NestedScribbleBenchmark/simpledns/channels/simpledns_cached"

type SimpleDNS_Cached_RoleSetupChan struct {
	Res_Chan chan simpledns_cached.Res_Chan
}

type SimpleDNS_Cached_InviteSetupChan struct {
	Res_InviteChan chan SimpleDNS_Cached_res_InviteChan
}

type SimpleDNS_Cached_res_InviteChan struct {

}