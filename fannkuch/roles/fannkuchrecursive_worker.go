package roles

import "NestedScribbleBenchmark/fannkuch/messages"
import "NestedScribbleBenchmark/fannkuch/channels/fannkuchrecursive"
import "NestedScribbleBenchmark/fannkuch/invitations"
import "NestedScribbleBenchmark/fannkuch/callbacks"
import fannkuchrecursive_2 "NestedScribbleBenchmark/fannkuch/results/fannkuchrecursive"
import "sync"

func FannkuchRecursive_Worker(wg *sync.WaitGroup, roleChannels fannkuchrecursive.Worker_Chan, inviteChannels invitations.FannkuchRecursive_Worker_InviteChan, env callbacks.FannkuchRecursive_Worker_Env) fannkuchrecursive_2.Worker_Result {
	IdxMin, Chunksz, N := env.Task_To_NewWorker()
	roleChannels.Label_To_NewWorker <- messages.Task
	roleChannels.Int_To_NewWorker <- IdxMin
	roleChannels.Int_To_NewWorker <- Chunksz
	roleChannels.Int_To_NewWorker <- N

	return env.Done()
}
