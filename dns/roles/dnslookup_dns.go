package roles

import "NestedScribbleBenchmark/dns/messages"
import "NestedScribbleBenchmark/dns/channels/dnslookup"
import "NestedScribbleBenchmark/dns/invitations"
import "NestedScribbleBenchmark/dns/callbacks"
import "sync"

func DNSLookup_dns(wg *sync.WaitGroup, roleChannels dnslookup.Dns_Chan, inviteChannels invitations.DNSLookup_dns_InviteChan, env callbacks.DNSLookup_dns_Env)  {
	defer wg.Done()
	<-roleChannels.Label_From_res
	host := <-roleChannels.String_From_res
	env.Req_From_Res(host)

	dns_choice := env.Dns_Choice()
	switch dns_choice {
	case callbacks.DNSLookup_dns_IP:
		ip := env.IP_To_Res()
		roleChannels.Label_To_res <- messages.IP
		roleChannels.String_To_res <- ip

		env.Done()
		return 
	case callbacks.DNSLookup_dns_DNSIP:
		ip_2 := env.DNSIP_To_Res()
		roleChannels.Label_To_res <- messages.DNSIP
		roleChannels.String_To_res <- ip_2

		env.Done()
		return 
	default:
		panic("Invalid choice was made")
	}
} 