package roles

import "NestedScribbleBenchmark/generaldns/messages"
import "NestedScribbleBenchmark/generaldns/channels/recdnslookup"
import "NestedScribbleBenchmark/generaldns/invitations"
import "NestedScribbleBenchmark/generaldns/callbacks"
import "sync"

func RecDNSLookup_dns(wg *sync.WaitGroup, roleChannels recdnslookup.Dns_Chan, inviteChannels invitations.RecDNSLookup_dns_InviteChan, env callbacks.RecDNSLookup_dns_Env)  {
	defer wg.Done()
	<-roleChannels.Label_From_res
	host := <-roleChannels.String_From_res
	env.RecReq_From_Res(host)

	dns_choice := env.Dns_Choice()
	switch dns_choice {
	case callbacks.RecDNSLookup_dns_RecDNSLookup:
		env.RecDNSLookup_Setup()
		
		recdnslookup_rolechan := invitations.RecDNSLookup_RoleSetupChan{
			Res_Chan: inviteChannels.Invite_Dns_To_RecDNSLookup_res,
		}
		recdnslookup_invitechan := invitations.RecDNSLookup_InviteSetupChan{
			Res_InviteChan: inviteChannels.Invite_Dns_To_RecDNSLookup_res_InviteChan,
		}
		RecDNSLookup_SendCommChannels(wg, recdnslookup_rolechan, recdnslookup_invitechan)

		recdnslookup_res_chan := <-inviteChannels.Invite_Dns_To_RecDNSLookup_res
		recdnslookup_res_inviteChan := <-inviteChannels.Invite_Dns_To_RecDNSLookup_res_InviteChan
		recdnslookup_res_env := env.To_RecDNSLookup_res_Env()
		recdnslookup_res_result := RecDNSLookup_res(wg, recdnslookup_res_chan, recdnslookup_res_inviteChan, recdnslookup_res_env)
		env.ResultFrom_RecDNSLookup_res(recdnslookup_res_result)

		ip := env.IP_To_Res()
		roleChannels.Label_To_res <- messages.IP
		roleChannels.String_To_res <- ip

		env.Done()
		return 
	case callbacks.RecDNSLookup_dns_IterDNSLookup:
		env.IterDNSLookup_Setup()
		
		iterdnslookup_rolechan := invitations.IterDNSLookup_RoleSetupChan{
			Res_Chan: inviteChannels.Invite_Dns_To_IterDNSLookup_res,
		}
		iterdnslookup_invitechan := invitations.IterDNSLookup_InviteSetupChan{
			Res_InviteChan: inviteChannels.Invite_Dns_To_IterDNSLookup_res_InviteChan,
		}
		IterDNSLookup_SendCommChannels(wg, iterdnslookup_rolechan, iterdnslookup_invitechan)

		iterdnslookup_res_chan := <-inviteChannels.Invite_Dns_To_IterDNSLookup_res
		iterdnslookup_res_inviteChan := <-inviteChannels.Invite_Dns_To_IterDNSLookup_res_InviteChan
		iterdnslookup_res_env := env.To_IterDNSLookup_res_Env()
		iterdnslookup_res_result := IterDNSLookup_res(wg, iterdnslookup_res_chan, iterdnslookup_res_inviteChan, iterdnslookup_res_env)
		env.ResultFrom_IterDNSLookup_res(iterdnslookup_res_result)

		ip_2 := env.IP_To_Res_2()
		roleChannels.Label_To_res <- messages.IP
		roleChannels.String_To_res <- ip_2

		env.Done()
		return 
	case callbacks.RecDNSLookup_dns_Cached:
		env.Cached_Setup()
		
		cached_rolechan := invitations.Cached_RoleSetupChan{
			Res_Chan: inviteChannels.Invite_Dns_To_Cached_res,
		}
		cached_invitechan := invitations.Cached_InviteSetupChan{
			Res_InviteChan: inviteChannels.Invite_Dns_To_Cached_res_InviteChan,
		}
		Cached_SendCommChannels(wg, cached_rolechan, cached_invitechan)

		cached_res_chan := <-inviteChannels.Invite_Dns_To_Cached_res
		cached_res_inviteChan := <-inviteChannels.Invite_Dns_To_Cached_res_InviteChan
		cached_res_env := env.To_Cached_res_Env()
		cached_res_result := Cached_res(wg, cached_res_chan, cached_res_inviteChan, cached_res_env)
		env.ResultFrom_Cached_res(cached_res_result)

		ip_3 := env.IP_To_Res_3()
		roleChannels.Label_To_res <- messages.IP
		roleChannels.String_To_res <- ip_3

		env.Done()
		return 
	default:
		panic("Invalid choice was made")
	}
} 