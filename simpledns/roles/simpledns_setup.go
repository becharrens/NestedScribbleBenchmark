package roles

import "NestedScribbleBenchmark/simpledns/messages"
import "NestedScribbleBenchmark/simpledns/channels/simpledns_cached"
import "NestedScribbleBenchmark/simpledns/channels/simpledns"
import "NestedScribbleBenchmark/simpledns/channels/dnslookup"
import "NestedScribbleBenchmark/simpledns/invitations"
import "sync"

func SimpleDNS_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.SimpleDNS_RoleSetupChan, inviteChannels invitations.SimpleDNS_InviteSetupChan)  {
	ispdns_invite_ispdns_2 := make(chan dnslookup.Res_Chan, 1)
	ispdns_invite_ispdns_invitechan_2 := make(chan invitations.DNSLookup_res_InviteChan, 1)
	ispdns_app_string := make(chan string, 1)
	ispdns_app_label := make(chan messages.SimpleDNS_Label, 1)
	ispdns_invite_ispdns := make(chan simpledns_cached.Res_Chan, 1)
	ispdns_invite_ispdns_invitechan := make(chan invitations.SimpleDNS_Cached_res_InviteChan, 1)
	app_ispdns_string := make(chan string, 1)
	app_ispdns_label := make(chan messages.SimpleDNS_Label, 1)

	ispdns_chan := simpledns.IspDNS_Chan{
		String_To_app: ispdns_app_string,
		String_From_app: app_ispdns_string,
		Label_To_app: ispdns_app_label,
		Label_From_app: app_ispdns_label,
	}
	app_chan := simpledns.App_Chan{
		String_To_ispDNS: app_ispdns_string,
		String_From_ispDNS: ispdns_app_string,
		Label_To_ispDNS: app_ispdns_label,
		Label_From_ispDNS: ispdns_app_label,
	}

	ispdns_inviteChan := invitations.SimpleDNS_ispDNS_InviteChan{
		Invite_IspDNS_To_SimpleDNS_Cached_res_InviteChan: ispdns_invite_ispdns_invitechan,
		Invite_IspDNS_To_SimpleDNS_Cached_res: ispdns_invite_ispdns,
		Invite_IspDNS_To_DNSLookup_res_InviteChan: ispdns_invite_ispdns_invitechan_2,
		Invite_IspDNS_To_DNSLookup_res: ispdns_invite_ispdns_2,
	}
	app_inviteChan := invitations.SimpleDNS_app_InviteChan{

	}

	roleChannels.App_Chan <- app_chan
	roleChannels.IspDNS_Chan <- ispdns_chan

	inviteChannels.App_InviteChan <- app_inviteChan
	inviteChannels.IspDNS_InviteChan <- ispdns_inviteChan
} 