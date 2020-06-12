package roles

import "ScribbleBenchmark/fannkuch/channels/fannkuchrecursive"
import "ScribbleBenchmark/fannkuch/invitations"
import "ScribbleBenchmark/fannkuch/callbacks"
import fannkuchrecursive_2 "ScribbleBenchmark/fannkuch/results/fannkuchrecursive"
import "sync"

func FannkuchRecursive_Worker(wg *sync.WaitGroup, roleChannels fannkuchrecursive.Worker_Chan, inviteChannels invitations.FannkuchRecursive_Worker_InviteChan, env callbacks.FannkuchRecursive_Worker_Env) fannkuchrecursive_2.Worker_Result {
	task_msg := env.Task_To_NewWorker()
	roleChannels.NewWorker_Task <- task_msg

	return env.Done()
} 