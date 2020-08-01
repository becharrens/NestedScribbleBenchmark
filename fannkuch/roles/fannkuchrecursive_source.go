package roles

import "NestedScribbleBenchmark/fannkuch/channels/fannkuchrecursive"
import "NestedScribbleBenchmark/fannkuch/invitations"
import "NestedScribbleBenchmark/fannkuch/callbacks"
import fannkuchrecursive_2 "NestedScribbleBenchmark/fannkuch/results/fannkuchrecursive"
import "sync"

func FannkuchRecursive_Source(wg *sync.WaitGroup, roleChannels fannkuchrecursive.Source_Chan, inviteChannels invitations.FannkuchRecursive_Source_InviteChan, env callbacks.FannkuchRecursive_Source_Env) fannkuchrecursive_2.Source_Result {
	select {
	case fannkuchrecursive_source_chan := <-inviteChannels.NewWorker_Invite_To_FannkuchRecursive_Source:
		fannkuchrecursive_source_inviteChan := <-inviteChannels.NewWorker_Invite_To_FannkuchRecursive_Source_InviteChan
		fannkuchrecursive_source_env := env.To_FannkuchRecursive_Source_Env()
		fannkuchrecursive_source_result := FannkuchRecursive_Source(wg, fannkuchrecursive_source_chan, fannkuchrecursive_source_inviteChan, fannkuchrecursive_source_env)
		env.ResultFrom_FannkuchRecursive_Source(fannkuchrecursive_source_result)

		result_msg := <-roleChannels.NewWorker_Result
		env.Result_From_NewWorker(result_msg)

		return env.Done()
	case result_msg_2 := <-roleChannels.NewWorker_Result_2:
		env.Result_From_NewWorker_2(result_msg_2)

		return env.Done()
	}
}
