package roles

import "NestedScribbleBenchmark/regexredux/channels/regexredux"
import "NestedScribbleBenchmark/regexredux/invitations"
import "NestedScribbleBenchmark/regexredux/callbacks"
import regexredux_2 "NestedScribbleBenchmark/regexredux/results/regexredux"
import "sync"

func RegexRedux_Worker(wg *sync.WaitGroup, roleChannels regexredux.Worker_Chan, inviteChannels invitations.RegexRedux_Worker_InviteChan, env callbacks.RegexRedux_Worker_Env) regexredux_2.Worker_Result {
	task_msg := <-roleChannels.Master_Task
	env.Task_From_Master(task_msg)

	nummatches_msg := env.NumMatches_To_Master()
	roleChannels.Master_NumMatches <- nummatches_msg

	return env.Done()
}
