package roles

import "ScribbleBenchmark/fannkuch/channels/fannkuch"
import "ScribbleBenchmark/fannkuch/invitations"
import "ScribbleBenchmark/fannkuch/callbacks"
import fannkuch_2 "ScribbleBenchmark/fannkuch/results/fannkuch"
import "sync"

func Fannkuch_Worker(wg *sync.WaitGroup, roleChannels fannkuch.Worker_Chan, inviteChannels invitations.Fannkuch_Worker_InviteChan, env callbacks.Fannkuch_Worker_Env) fannkuch_2.Worker_Result {
	task_msg := <-roleChannels.Main_Task
	env.Task_From_Main(task_msg)

	worker_choice := env.Worker_Choice()
	switch worker_choice {
	case callbacks.Fannkuch_Worker_FannkuchRecursive:
		env.FannkuchRecursive_Setup()
		fannkuchrecursive_rolechan := invitations.FannkuchRecursive_RoleSetupChan{
			Source_Chan: inviteChannels.Invite_Main_To_FannkuchRecursive_Source,
			Worker_Chan: inviteChannels.Invite_Worker_To_FannkuchRecursive_Worker,
		}
		fannkuchrecursive_invitechan := invitations.FannkuchRecursive_InviteSetupChan{
			Source_InviteChan: inviteChannels.Invite_Main_To_FannkuchRecursive_Source_InviteChan,
			Worker_InviteChan: inviteChannels.Invite_Worker_To_FannkuchRecursive_Worker_InviteChan,
		}
		FannkuchRecursive_SendCommChannels(wg, fannkuchrecursive_rolechan, fannkuchrecursive_invitechan)

		fannkuchrecursive_worker_chan := <-inviteChannels.Invite_Worker_To_FannkuchRecursive_Worker
		fannkuchrecursive_worker_inviteChan := <-inviteChannels.Invite_Worker_To_FannkuchRecursive_Worker_InviteChan
		fannkuchrecursive_worker_env := env.To_FannkuchRecursive_Worker_Env()
		fannkuchrecursive_worker_result := FannkuchRecursive_Worker(wg, fannkuchrecursive_worker_chan, fannkuchrecursive_worker_inviteChan, fannkuchrecursive_worker_env)
		env.ResultFrom_FannkuchRecursive_Worker(fannkuchrecursive_worker_result)

		result_msg := env.Result_To_Main()
		roleChannels.Main_Result <- result_msg

		return env.Done()
	case callbacks.Fannkuch_Worker_Result:
		result_msg_2 := env.Result_To_Main_2()
		roleChannels.Main_Result_2 <- result_msg_2

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
} 