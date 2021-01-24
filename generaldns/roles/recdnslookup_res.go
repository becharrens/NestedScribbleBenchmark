package roles

import "NestedScribbleBenchmark/generaldns/messages"
import "NestedScribbleBenchmark/generaldns/channels/recdnslookup"
import "NestedScribbleBenchmark/generaldns/invitations"
import "NestedScribbleBenchmark/generaldns/callbacks"
import recdnslookup_2 "NestedScribbleBenchmark/generaldns/results/recdnslookup"
import "sync"

func RecDNSLookup_res(wg *sync.WaitGroup, roleChannels recdnslookup.Res_Chan, inviteChannels invitations.RecDNSLookup_res_InviteChan, env callbacks.RecDNSLookup_res_Env) recdnslookup_2.Res_Result {
	host := env.RecReq_To_Dns()
	roleChannels.Label_To_dns <- messages.RecReq
	roleChannels.String_To_dns <- host

	<-roleChannels.Label_From_dns
	ip := <-roleChannels.String_From_dns
	env.IP_From_Dns(ip)

	return env.Done()
} 