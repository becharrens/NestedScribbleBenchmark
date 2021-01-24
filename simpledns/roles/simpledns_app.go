package roles

import "NestedScribbleBenchmark/simpledns/messages"
import "NestedScribbleBenchmark/simpledns/channels/simpledns"
import "NestedScribbleBenchmark/simpledns/invitations"
import "NestedScribbleBenchmark/simpledns/callbacks"
import simpledns_2 "NestedScribbleBenchmark/simpledns/results/simpledns"
import "sync"

func SimpleDNS_app(wg *sync.WaitGroup, roleChannels simpledns.App_Chan, inviteChannels invitations.SimpleDNS_app_InviteChan, env callbacks.SimpleDNS_app_Env) simpledns_2.App_Result {
	app_choice_2 := env.App_Choice()
	switch app_choice_2 {
	case callbacks.SimpleDNS_app_Done:
		env.Done_To_IspDNS()
		roleChannels.Label_To_ispDNS <- messages.Done

		return env.Done()
	case callbacks.SimpleDNS_app_RecQuery:
		host := env.RecQuery_To_IspDNS()
		roleChannels.Label_To_ispDNS <- messages.RecQuery
		roleChannels.String_To_ispDNS <- host

		<-roleChannels.Label_From_ispDNS
		ip := <-roleChannels.String_From_ispDNS
		env.IP_From_IspDNS(ip)

REC:
		for {
			app_choice := env.App_Choice_2()
			switch app_choice {
			case callbacks.SimpleDNS_app_Done_2:
				env.Done_To_IspDNS_2()
				roleChannels.Label_To_ispDNS <- messages.Done

				return env.Done()
			case callbacks.SimpleDNS_app_RecQuery_2:
				host_2 := env.RecQuery_To_IspDNS_2()
				roleChannels.Label_To_ispDNS <- messages.RecQuery
				roleChannels.String_To_ispDNS <- host_2

				<-roleChannels.Label_From_ispDNS
				ip_2 := <-roleChannels.String_From_ispDNS
				env.IP_From_IspDNS_2(ip_2)

				continue REC
			default:
				panic("Invalid choice was made")
			}
		}
	default:
		panic("Invalid choice was made")
	}
} 