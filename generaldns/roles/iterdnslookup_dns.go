package roles

import "NestedScribbleBenchmark/generaldns/messages"
import "NestedScribbleBenchmark/generaldns/channels/iterdnslookup"
import "NestedScribbleBenchmark/generaldns/invitations"
import "NestedScribbleBenchmark/generaldns/callbacks"
import "sync"

func IterDNSLookup_dns(wg *sync.WaitGroup, roleChannels iterdnslookup.Dns_Chan, inviteChannels invitations.IterDNSLookup_dns_InviteChan, env callbacks.IterDNSLookup_dns_Env)  {
	defer wg.Done()
	<-roleChannels.Label_From_res
	host := <-roleChannels.String_From_res
	env.IterReq_From_Res(host)

	dns_choice := env.Dns_Choice()
	switch dns_choice {
	case callbacks.IterDNSLookup_dns_IP:
		ip := env.IP_To_Res()
		roleChannels.Label_To_res <- messages.IP
		roleChannels.String_To_res <- ip

		env.Done()
		return 
	case callbacks.IterDNSLookup_dns_DNSIP:
		ip_2 := env.DNSIP_To_Res()
		roleChannels.Label_To_res <- messages.DNSIP
		roleChannels.String_To_res <- ip_2

		env.Done()
		return 
	default:
		panic("Invalid choice was made")
	}
} 