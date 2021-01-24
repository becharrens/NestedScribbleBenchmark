package roles

import "NestedScribbleBenchmark/simpledns/messages"
import "NestedScribbleBenchmark/simpledns/channels/dnslookup"
import "NestedScribbleBenchmark/simpledns/invitations"
import "NestedScribbleBenchmark/simpledns/callbacks"
import "sync"

func DNSLookup_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.DNSLookup_RoleSetupChan, inviteChannels invitations.DNSLookup_InviteSetupChan)  {
	res_invite_res := make(chan dnslookup.Res_Chan, 1)
	res_invite_res_invitechan := make(chan invitations.DNSLookup_res_InviteChan, 1)
	dns_res_string := make(chan string, 1)
	dns_res_label := make(chan messages.SimpleDNS_Label, 1)
	res_dns_string := make(chan string, 1)
	res_dns_label := make(chan messages.SimpleDNS_Label, 1)

	res_chan := dnslookup.Res_Chan{
		String_To_dns: res_dns_string,
		String_From_dns: dns_res_string,
		Label_To_dns: res_dns_label,
		Label_From_dns: dns_res_label,
	}
	dns_chan := dnslookup.Dns_Chan{
		String_To_res: dns_res_string,
		String_From_res: res_dns_string,
		Label_To_res: dns_res_label,
		Label_From_res: res_dns_label,
	}

	res_inviteChan := invitations.DNSLookup_res_InviteChan{
		Invite_Res_To_DNSLookup_res_InviteChan: res_invite_res_invitechan,
		Invite_Res_To_DNSLookup_res: res_invite_res,
	}
	dns_inviteChan := invitations.DNSLookup_dns_InviteChan{

	}

	roleChannels.Res_Chan <- res_chan

	inviteChannels.Res_InviteChan <- res_inviteChan

	wg.Add(1)

	dns_env := callbacks.New_DNSLookup_dns_State()
	go DNSLookup_dns(wg, dns_chan, dns_inviteChan, dns_env)
} 