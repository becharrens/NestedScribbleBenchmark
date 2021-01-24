package roles

import "NestedScribbleBenchmark/generaldns/messages"
import "NestedScribbleBenchmark/generaldns/channels/recdnslookup"
import "NestedScribbleBenchmark/generaldns/channels/iterdnslookup"
import "NestedScribbleBenchmark/generaldns/channels/generaldns"
import "NestedScribbleBenchmark/generaldns/channels/cached"
import "NestedScribbleBenchmark/generaldns/invitations"
import "sync"

func GeneralDNS_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.GeneralDNS_RoleSetupChan, inviteChannels invitations.GeneralDNS_InviteSetupChan)  {
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

	roleChannels.App_Chan <- app_chan
	roleChannels.DnsRes_Chan <- dnsres_chan

	inviteChannels.App_InviteChan <- app_inviteChan
	inviteChannels.DnsRes_InviteChan <- dnsres_inviteChan
} 