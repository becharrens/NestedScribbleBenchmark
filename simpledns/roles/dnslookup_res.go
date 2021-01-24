package roles

import "NestedScribbleBenchmark/simpledns/messages"
import "NestedScribbleBenchmark/simpledns/channels/dnslookup"
import "NestedScribbleBenchmark/simpledns/invitations"
import "NestedScribbleBenchmark/simpledns/callbacks"
import dnslookup_2 "NestedScribbleBenchmark/simpledns/results/dnslookup"
import "sync"

func DNSLookup_res(wg *sync.WaitGroup, roleChannels dnslookup.Res_Chan, inviteChannels invitations.DNSLookup_res_InviteChan, env callbacks.DNSLookup_res_Env) dnslookup_2.Res_Result {
	host := env.Req_To_Dns()
	roleChannels.Label_To_dns <- messages.Req
	roleChannels.String_To_dns <- host

	dns_choice := <-roleChannels.Label_From_dns
	switch dns_choice {
	case messages.IP:
		ip := <-roleChannels.String_From_dns
		env.IP_From_Dns(ip)

		return env.Done()
	case messages.DNSIP:
		ip_2 := <-roleChannels.String_From_dns
		env.DNSIP_From_Dns(ip_2)

		env.DNSLookup_Setup()
		
		dnslookup_rolechan := invitations.DNSLookup_RoleSetupChan{
			Res_Chan: inviteChannels.Invite_Res_To_DNSLookup_res,
		}
		dnslookup_invitechan := invitations.DNSLookup_InviteSetupChan{
			Res_InviteChan: inviteChannels.Invite_Res_To_DNSLookup_res_InviteChan,
		}
		DNSLookup_SendCommChannels(wg, dnslookup_rolechan, dnslookup_invitechan)

		dnslookup_res_chan := <-inviteChannels.Invite_Res_To_DNSLookup_res
		dnslookup_res_inviteChan := <-inviteChannels.Invite_Res_To_DNSLookup_res_InviteChan
		dnslookup_res_env := env.To_DNSLookup_res_Env()
		dnslookup_res_result := DNSLookup_res(wg, dnslookup_res_chan, dnslookup_res_inviteChan, dnslookup_res_env)
		env.ResultFrom_DNSLookup_res(dnslookup_res_result)

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
} 