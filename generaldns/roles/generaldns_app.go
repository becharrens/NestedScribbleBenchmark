package roles

import "NestedScribbleBenchmark/generaldns/messages"
import "NestedScribbleBenchmark/generaldns/channels/generaldns"
import "NestedScribbleBenchmark/generaldns/invitations"
import "NestedScribbleBenchmark/generaldns/callbacks"
import generaldns_2 "NestedScribbleBenchmark/generaldns/results/generaldns"
import "sync"

func GeneralDNS_app(wg *sync.WaitGroup, roleChannels generaldns.App_Chan, inviteChannels invitations.GeneralDNS_app_InviteChan, env callbacks.GeneralDNS_app_Env) generaldns_2.App_Result {
	app_choice_2 := env.App_Choice()
	switch app_choice_2 {
	case callbacks.GeneralDNS_app_Done:
		env.Done_To_DnsRes()
		roleChannels.Label_To_dnsRes <- messages.Done

		return env.Done()
	case callbacks.GeneralDNS_app_Query:
		host := env.Query_To_DnsRes()
		roleChannels.Label_To_dnsRes <- messages.Query
		roleChannels.String_To_dnsRes <- host

		<-roleChannels.Label_From_dnsRes
		ip := <-roleChannels.String_From_dnsRes
		env.IP_From_DnsRes(ip)

REC:
		for {
			app_choice := env.App_Choice_2()
			switch app_choice {
			case callbacks.GeneralDNS_app_Done_2:
				env.Done_To_DnsRes_2()
				roleChannels.Label_To_dnsRes <- messages.Done

				return env.Done()
			case callbacks.GeneralDNS_app_Query_2:
				host_2 := env.Query_To_DnsRes_2()
				roleChannels.Label_To_dnsRes <- messages.Query
				roleChannels.String_To_dnsRes <- host_2

				<-roleChannels.Label_From_dnsRes
				ip_2 := <-roleChannels.String_From_dnsRes
				env.IP_From_DnsRes_2(ip_2)

				continue REC
			default:
				panic("Invalid choice was made")
			}
		}
	default:
		panic("Invalid choice was made")
	}
} 