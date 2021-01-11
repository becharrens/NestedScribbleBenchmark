package roles

import "NestedScribbleBenchmark/noughtsandcrosses/channels/minmaxstrategy_evalboard"
import "NestedScribbleBenchmark/noughtsandcrosses/invitations"
import "NestedScribbleBenchmark/noughtsandcrosses/callbacks"
import minmaxstrategy_evalboard_2 "NestedScribbleBenchmark/noughtsandcrosses/results/minmaxstrategy_evalboard"
import "sync"

func MinMaxStrategy_EvalBoard_W(wg *sync.WaitGroup, roleChannels minmaxstrategy_evalboard.W_Chan, inviteChannels invitations.MinMaxStrategy_EvalBoard_W_InviteChan, env callbacks.MinMaxStrategy_EvalBoard_W_Env) minmaxstrategy_evalboard_2.W_Result {
	return env.Done()
} 