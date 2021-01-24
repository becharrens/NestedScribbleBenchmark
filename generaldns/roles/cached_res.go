package roles

import "NestedScribbleBenchmark/generaldns/channels/cached"
import "NestedScribbleBenchmark/generaldns/invitations"
import "NestedScribbleBenchmark/generaldns/callbacks"
import cached_2 "NestedScribbleBenchmark/generaldns/results/cached"
import "sync"

func Cached_res(wg *sync.WaitGroup, roleChannels cached.Res_Chan, inviteChannels invitations.Cached_res_InviteChan, env callbacks.Cached_res_Env) cached_2.Res_Result {
	return env.Done()
} 