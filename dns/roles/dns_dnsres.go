package roles

import "NestedScribbleBenchmark/dns/messages"
import "NestedScribbleBenchmark/dns/channels/dns"
import "NestedScribbleBenchmark/dns/invitations"
import "NestedScribbleBenchmark/dns/callbacks"
import dns_2 "NestedScribbleBenchmark/dns/results/dns"
import "sync"

func DNS_dnsRes(wg *sync.WaitGroup, roleChannels dns.DnsRes_Chan, inviteChannels invitations.DNS_dnsRes_InviteChan, env callbacks.DNS_dnsRes_Env) dns_2.DnsRes_Result {
	app_choice_2 := <-roleChannels.Label_From_app
	switch app_choice_2 {
	case messages.Done:
		env.Done_From_App()

		env.Done_To_IspDNS()
		roleChannels.Label_To_ispDNS <- messages.Done

		return env.Done()
	case messages.Query:
		host := <-roleChannels.String_From_app
		env.Query_From_App(host)

		host_2 := env.RecQuery_To_IspDNS()
		roleChannels.Label_To_ispDNS <- messages.RecQuery
		roleChannels.String_To_ispDNS <- host_2

		<-roleChannels.Label_From_ispDNS
		ip := <-roleChannels.String_From_ispDNS
		env.IP_From_IspDNS(ip)

		ip_2 := env.IP_To_App()
		roleChannels.Label_To_app <- messages.IP
		roleChannels.String_To_app <- ip_2

REC:
		for {
			app_choice := <-roleChannels.Label_From_app
			switch app_choice {
			case messages.Done:
				env.Done_From_App_2()

				env.Done_To_IspDNS_2()
				roleChannels.Label_To_ispDNS <- messages.Done

				return env.Done()
			case messages.Query:
				host_3 := <-roleChannels.String_From_app
				env.Query_From_App_2(host_3)

				host_4 := env.RecQuery_To_IspDNS_2()
				roleChannels.Label_To_ispDNS <- messages.RecQuery
				roleChannels.String_To_ispDNS <- host_4

				<-roleChannels.Label_From_ispDNS
				ip_3 := <-roleChannels.String_From_ispDNS
				env.IP_From_IspDNS_2(ip_3)

				ip_4 := env.IP_To_App_2()
				roleChannels.Label_To_app <- messages.IP
				roleChannels.String_To_app <- ip_4

				continue REC
			default:
				panic("Invalid choice was made")
			}
		}
	default:
		panic("Invalid choice was made")
	}
} 