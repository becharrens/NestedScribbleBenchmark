package roles

import "NestedScribbleBenchmark/fannkuch/messages"
import "NestedScribbleBenchmark/fannkuch/channels/fannkuch"
import "NestedScribbleBenchmark/fannkuch/invitations"
import "NestedScribbleBenchmark/fannkuch/callbacks"
import fannkuch_2 "NestedScribbleBenchmark/fannkuch/results/fannkuch"
import "sync"

func Fannkuch_Main(wg *sync.WaitGroup, roleChannels fannkuch.Main_Chan, inviteChannels invitations.Fannkuch_Main_InviteChan, env callbacks.Fannkuch_Main_Env) fannkuch_2.Main_Result {
	IdxMin, Chunksz, N := env.Task_To_Worker()
	roleChannels.Label_To_Worker <- messages.Task
	roleChannels.Int_To_Worker <- IdxMin
	roleChannels.Int_To_Worker <- Chunksz
	roleChannels.Int_To_Worker <- N

	worker_choice := <-roleChannels.Label_From_Worker
	switch worker_choice {
	case messages.FannkuchRecursive_Main_Worker:
		fannkuchrecursive_source_chan := <-inviteChannels.Worker_Invite_To_FannkuchRecursive_Source
		fannkuchrecursive_source_inviteChan := <-inviteChannels.Worker_Invite_To_FannkuchRecursive_Source_InviteChan
		fannkuchrecursive_source_env := env.To_FannkuchRecursive_Source_Env()
		fannkuchrecursive_source_result := FannkuchRecursive_Source(wg, fannkuchrecursive_source_chan, fannkuchrecursive_source_inviteChan, fannkuchrecursive_source_env)
		env.ResultFrom_FannkuchRecursive_Source(fannkuchrecursive_source_result)

		<-roleChannels.Label_From_Worker
		MaxFlips := <-roleChannels.Int_From_Worker
		Checksum := <-roleChannels.Int_From_Worker
		env.Result_From_Worker(MaxFlips, Checksum)

		return env.Done()
	case messages.Result:
		MaxFlips_2 := <-roleChannels.Int_From_Worker
		Checksum_2 := <-roleChannels.Int_From_Worker
		env.Result_From_Worker_2(MaxFlips_2, Checksum_2)

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
}
