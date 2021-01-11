package roles

import "NestedScribbleBenchmark/noughtsandcrosses/channels/standardstrategy"
import "NestedScribbleBenchmark/noughtsandcrosses/invitations"
import "NestedScribbleBenchmark/noughtsandcrosses/callbacks"
import standardstrategy_2 "NestedScribbleBenchmark/noughtsandcrosses/results/standardstrategy"
import "sync"

func StandardStrategy_P(wg *sync.WaitGroup, roleChannels standardstrategy.P_Chan, inviteChannels invitations.StandardStrategy_P_InviteChan, env callbacks.StandardStrategy_P_Env) standardstrategy_2.P_Result {
	return env.Done()
} 