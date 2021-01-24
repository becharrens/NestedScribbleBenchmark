package invitations

import "NestedScribbleBenchmark/simpledns/channels/simpledns_cached"
import "NestedScribbleBenchmark/simpledns/channels/simpledns"
import "NestedScribbleBenchmark/simpledns/channels/dnslookup"

type SimpleDNS_RoleSetupChan struct {
	App_Chan chan simpledns.App_Chan
	IspDNS_Chan chan simpledns.IspDNS_Chan
}

type SimpleDNS_InviteSetupChan struct {
	App_InviteChan chan SimpleDNS_app_InviteChan
	IspDNS_InviteChan chan SimpleDNS_ispDNS_InviteChan
}

type SimpleDNS_app_InviteChan struct {

}

type SimpleDNS_ispDNS_InviteChan struct {
	Invite_IspDNS_To_DNSLookup_res chan dnslookup.Res_Chan
	Invite_IspDNS_To_DNSLookup_res_InviteChan chan DNSLookup_res_InviteChan
	Invite_IspDNS_To_SimpleDNS_Cached_res chan simpledns_cached.Res_Chan
	Invite_IspDNS_To_SimpleDNS_Cached_res_InviteChan chan SimpleDNS_Cached_res_InviteChan
}