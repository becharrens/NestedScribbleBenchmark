package roles

import "NestedScribbleBenchmark/fannkuch/messages"
import "NestedScribbleBenchmark/fannkuch/channels/fannkuchrecursive"
import "NestedScribbleBenchmark/fannkuch/invitations"
import "NestedScribbleBenchmark/fannkuch/callbacks"
import fannkuchrecursive_2 "NestedScribbleBenchmark/fannkuch/results/fannkuchrecursive"
import "sync"

func FannkuchRecursive_Source(wg *sync.WaitGroup, roleChannels fannkuchrecursive.Source_Chan, inviteChannels invitations.FannkuchRecursive_Source_InviteChan, env callbacks.FannkuchRecursive_Source_Env) fannkuchrecursive_2.Source_Result {
	newworker_choice := <-roleChannels.Label_From_NewWorker
	switch newworker_choice {
	case messages.FannkuchRecursive_Source_NewWorker:
		fannkuchrecursive_source_chan := <-inviteChannels.NewWorker_Invite_To_FannkuchRecursive_Source
		fannkuchrecursive_source_inviteChan := <-inviteChannels.NewWorker_Invite_To_FannkuchRecursive_Source_InviteChan
		fannkuchrecursive_source_env := env.To_FannkuchRecursive_Source_Env()
		fannkuchrecursive_source_result := FannkuchRecursive_Source(wg, fannkuchrecursive_source_chan, fannkuchrecursive_source_inviteChan, fannkuchrecursive_source_env)
		env.ResultFrom_FannkuchRecursive_Source(fannkuchrecursive_source_result)

		<-roleChannels.Label_From_NewWorker
		MaxFlips := <-roleChannels.Int_From_NewWorker
		Checksum := <-roleChannels.Int_From_NewWorker
		env.Result_From_NewWorker(MaxFlips, Checksum)

		return env.Done()
	case messages.Result:
		MaxFlips_2 := <-roleChannels.Int_From_NewWorker
		Checksum_2 := <-roleChannels.Int_From_NewWorker
		env.Result_From_NewWorker_2(MaxFlips_2, Checksum_2)

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
}
