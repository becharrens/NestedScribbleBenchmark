package roles

import "NestedScribbleBenchmark/generaldns/messages"
import "NestedScribbleBenchmark/generaldns/channels/recdnslookup"
import "NestedScribbleBenchmark/generaldns/channels/iterdnslookup"
import "NestedScribbleBenchmark/generaldns/channels/cached"
import "NestedScribbleBenchmark/generaldns/invitations"
import "NestedScribbleBenchmark/generaldns/callbacks"
import "sync"

func RecDNSLookup_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.RecDNSLookup_RoleSetupChan, inviteChannels invitations.RecDNSLookup_InviteSetupChan)  {
	dns_invite_dns_3 := make(chan cached.Res_Chan, 1)
	dns_invite_dns_invitechan_3 := make(chan invitations.Cached_res_InviteChan, 1)
	dns_invite_dns_2 := make(chan iterdnslookup.Res_Chan, 1)
	dns_invite_dns_invitechan_2 := make(chan invitations.IterDNSLookup_res_InviteChan, 1)
	dns_res_string := make(chan string, 1)
	dns_res_label := make(chan messages.GeneralDNS_Label, 1)
	dns_invite_dns := make(chan recdnslookup.Res_Chan, 1)
	dns_invite_dns_invitechan := make(chan invitations.RecDNSLookup_res_InviteChan, 1)
	res_dns_string := make(chan string, 1)
	res_dns_label := make(chan messages.GeneralDNS_Label, 1)

	res_chan := recdnslookup.Res_Chan{
		String_To_dns: res_dns_string,
		String_From_dns: dns_res_string,
		Label_To_dns: res_dns_label,
		Label_From_dns: dns_res_label,
	}
	dns_chan := recdnslookup.Dns_Chan{
		String_To_res: dns_res_string,
		String_From_res: res_dns_string,
		Label_To_res: dns_res_label,
		Label_From_res: res_dns_label,
	}

	res_inviteChan := invitations.RecDNSLookup_res_InviteChan{

	}
	dns_inviteChan := invitations.RecDNSLookup_dns_InviteChan{
		Invite_Dns_To_RecDNSLookup_res_InviteChan: dns_invite_dns_invitechan,
		Invite_Dns_To_RecDNSLookup_res: dns_invite_dns,
		Invite_Dns_To_IterDNSLookup_res_InviteChan: dns_invite_dns_invitechan_2,
		Invite_Dns_To_IterDNSLookup_res: dns_invite_dns_2,
		Invite_Dns_To_Cached_res_InviteChan: dns_invite_dns_invitechan_3,
		Invite_Dns_To_Cached_res: dns_invite_dns_3,
	}

	roleChannels.Res_Chan <- res_chan

	inviteChannels.Res_InviteChan <- res_inviteChan

	wg.Add(1)

	dns_env := callbacks.New_RecDNSLookup_dns_State()
	go RecDNSLookup_dns(wg, dns_chan, dns_inviteChan, dns_env)
} 