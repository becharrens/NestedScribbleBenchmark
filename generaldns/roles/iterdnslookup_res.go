package roles

import "NestedScribbleBenchmark/generaldns/messages"
import "NestedScribbleBenchmark/generaldns/channels/iterdnslookup"
import "NestedScribbleBenchmark/generaldns/invitations"
import "NestedScribbleBenchmark/generaldns/callbacks"
import iterdnslookup_2 "NestedScribbleBenchmark/generaldns/results/iterdnslookup"
import "sync"

func IterDNSLookup_res(wg *sync.WaitGroup, roleChannels iterdnslookup.Res_Chan, inviteChannels invitations.IterDNSLookup_res_InviteChan, env callbacks.IterDNSLookup_res_Env) iterdnslookup_2.Res_Result {
	host := env.IterReq_To_Dns()
	roleChannels.Label_To_dns <- messages.IterReq
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

		env.IterDNSLookup_Setup()
		
		iterdnslookup_rolechan := invitations.IterDNSLookup_RoleSetupChan{
			Res_Chan: inviteChannels.Invite_Res_To_IterDNSLookup_res,
		}
		iterdnslookup_invitechan := invitations.IterDNSLookup_InviteSetupChan{
			Res_InviteChan: inviteChannels.Invite_Res_To_IterDNSLookup_res_InviteChan,
		}
		IterDNSLookup_SendCommChannels(wg, iterdnslookup_rolechan, iterdnslookup_invitechan)

		iterdnslookup_res_chan := <-inviteChannels.Invite_Res_To_IterDNSLookup_res
		iterdnslookup_res_inviteChan := <-inviteChannels.Invite_Res_To_IterDNSLookup_res_InviteChan
		iterdnslookup_res_env := env.To_IterDNSLookup_res_Env()
		iterdnslookup_res_result := IterDNSLookup_res(wg, iterdnslookup_res_chan, iterdnslookup_res_inviteChan, iterdnslookup_res_env)
		env.ResultFrom_IterDNSLookup_res(iterdnslookup_res_result)

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
} 