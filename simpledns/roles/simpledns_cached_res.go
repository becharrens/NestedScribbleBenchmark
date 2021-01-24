package roles

import "NestedScribbleBenchmark/simpledns/channels/simpledns_cached"
import "NestedScribbleBenchmark/simpledns/invitations"
import "NestedScribbleBenchmark/simpledns/callbacks"
import simpledns_cached_2 "NestedScribbleBenchmark/simpledns/results/simpledns_cached"
import "sync"

func SimpleDNS_Cached_res(wg *sync.WaitGroup, roleChannels simpledns_cached.Res_Chan, inviteChannels invitations.SimpleDNS_Cached_res_InviteChan, env callbacks.SimpleDNS_Cached_res_Env) simpledns_cached_2.Res_Result {
	return env.Done()
} 