package roles

import "NestedScribbleBenchmark/fannkuch/channels/fannkuch"
import "NestedScribbleBenchmark/fannkuch/invitations"
import "NestedScribbleBenchmark/fannkuch/callbacks"
import fannkuch_2 "NestedScribbleBenchmark/fannkuch/results/fannkuch"
import "sync"

func Fannkuch_Main(wg *sync.WaitGroup, roleChannels fannkuch.Main_Chan, inviteChannels invitations.Fannkuch_Main_InviteChan, env callbacks.Fannkuch_Main_Env) fannkuch_2.Main_Result {
	task_msg := env.Task_To_Worker()
	roleChannels.Worker_Task <- task_msg

	select {
	case fannkuchrecursive_source_chan := <-inviteChannels.Worker_Invite_To_FannkuchRecursive_Source:
		fannkuchrecursive_source_inviteChan := <-inviteChannels.Worker_Invite_To_FannkuchRecursive_Source_InviteChan
		fannkuchrecursive_source_env := env.To_FannkuchRecursive_Source_Env()
		fannkuchrecursive_source_result := FannkuchRecursive_Source(wg, fannkuchrecursive_source_chan, fannkuchrecursive_source_inviteChan, fannkuchrecursive_source_env)
		env.ResultFrom_FannkuchRecursive_Source(fannkuchrecursive_source_result)

		result_msg := <-roleChannels.Worker_Result
		env.Result_From_Worker(result_msg)

		return env.Done()
	case result_msg_2 := <-roleChannels.Worker_Result_2:
		env.Result_From_Worker_2(result_msg_2)

		return env.Done()
	}
}
