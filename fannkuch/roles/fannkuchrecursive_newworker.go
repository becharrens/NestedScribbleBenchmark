package roles

import "NestedScribbleBenchmark/fannkuch/channels/fannkuchrecursive"
import "NestedScribbleBenchmark/fannkuch/invitations"
import "NestedScribbleBenchmark/fannkuch/callbacks"
import "sync"

func FannkuchRecursive_NewWorker(wg *sync.WaitGroup, roleChannels fannkuchrecursive.NewWorker_Chan, inviteChannels invitations.FannkuchRecursive_NewWorker_InviteChan, env callbacks.FannkuchRecursive_NewWorker_Env) {
	defer wg.Done()
	task_msg := <-roleChannels.Worker_Task
	env.Task_From_Worker(task_msg)

	newworker_choice := env.NewWorker_Choice()
	switch newworker_choice {
	case callbacks.FannkuchRecursive_NewWorker_FannkuchRecursive:
		env.FannkuchRecursive_Setup()
		fannkuchrecursive_rolechan := invitations.FannkuchRecursive_RoleSetupChan{
			Source_Chan: inviteChannels.Invite_Source_To_FannkuchRecursive_Source,
			Worker_Chan: inviteChannels.Invite_NewWorker_To_FannkuchRecursive_Worker,
		}
		fannkuchrecursive_invitechan := invitations.FannkuchRecursive_InviteSetupChan{
			Source_InviteChan: inviteChannels.Invite_Source_To_FannkuchRecursive_Source_InviteChan,
			Worker_InviteChan: inviteChannels.Invite_NewWorker_To_FannkuchRecursive_Worker_InviteChan,
		}
		FannkuchRecursive_SendCommChannels(wg, fannkuchrecursive_rolechan, fannkuchrecursive_invitechan)

		fannkuchrecursive_worker_chan := <-inviteChannels.Invite_NewWorker_To_FannkuchRecursive_Worker
		fannkuchrecursive_worker_inviteChan := <-inviteChannels.Invite_NewWorker_To_FannkuchRecursive_Worker_InviteChan
		fannkuchrecursive_worker_env := env.To_FannkuchRecursive_Worker_Env()
		fannkuchrecursive_worker_result := FannkuchRecursive_Worker(wg, fannkuchrecursive_worker_chan, fannkuchrecursive_worker_inviteChan, fannkuchrecursive_worker_env)
		env.ResultFrom_FannkuchRecursive_Worker(fannkuchrecursive_worker_result)

		result_msg := env.Result_To_Source()
		roleChannels.Source_Result <- result_msg

		env.Done()
		return
	case callbacks.FannkuchRecursive_NewWorker_Result:
		result_msg_2 := env.Result_To_Source_2()
		roleChannels.Source_Result_2 <- result_msg_2

		env.Done()
		return
	default:
		panic("Invalid choice was made")
	}
}
