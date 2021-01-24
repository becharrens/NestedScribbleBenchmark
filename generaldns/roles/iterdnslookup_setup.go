package roles

import "NestedScribbleBenchmark/generaldns/messages"
import "NestedScribbleBenchmark/generaldns/channels/iterdnslookup"
import "NestedScribbleBenchmark/generaldns/invitations"
import "NestedScribbleBenchmark/generaldns/callbacks"
import "sync"

func IterDNSLookup_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.IterDNSLookup_RoleSetupChan, inviteChannels invitations.IterDNSLookup_InviteSetupChan)  {
	res_invite_res := make(chan iterdnslookup.Res_Chan, 1)
	res_invite_res_invitechan := make(chan invitations.IterDNSLookup_res_InviteChan, 1)
	dns_res_string := make(chan string, 1)
	dns_res_label := make(chan messages.GeneralDNS_Label, 1)
	res_dns_string := make(chan string, 1)
	res_dns_label := make(chan messages.GeneralDNS_Label, 1)

	res_chan := iterdnslookup.Res_Chan{
		String_To_dns: res_dns_string,
		String_From_dns: dns_res_string,
		Label_To_dns: res_dns_label,
		Label_From_dns: dns_res_label,
	}
	dns_chan := iterdnslookup.Dns_Chan{
		String_To_res: dns_res_string,
		String_From_res: res_dns_string,
		Label_To_res: dns_res_label,
		Label_From_res: res_dns_label,
	}

	res_inviteChan := invitations.IterDNSLookup_res_InviteChan{
		Invite_Res_To_IterDNSLookup_res_InviteChan: res_invite_res_invitechan,
		Invite_Res_To_IterDNSLookup_res: res_invite_res,
	}
	dns_inviteChan := invitations.IterDNSLookup_dns_InviteChan{

	}

	roleChannels.Res_Chan <- res_chan

	inviteChannels.Res_InviteChan <- res_inviteChan

	wg.Add(1)

	dns_env := callbacks.New_IterDNSLookup_dns_State()
	go IterDNSLookup_dns(wg, dns_chan, dns_inviteChan, dns_env)
} 