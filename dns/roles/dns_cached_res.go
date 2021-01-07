package roles

import "NestedScribbleBenchmark/dns/channels/dns_cached"
import "NestedScribbleBenchmark/dns/invitations"
import "NestedScribbleBenchmark/dns/callbacks"
import dns_cached_2 "NestedScribbleBenchmark/dns/results/dns_cached"
import "sync"

func DNS_Cached_res(wg *sync.WaitGroup, roleChannels dns_cached.Res_Chan, inviteChannels invitations.DNS_Cached_res_InviteChan, env callbacks.DNS_Cached_res_Env) dns_cached_2.Res_Result {
	return env.Done()
} 