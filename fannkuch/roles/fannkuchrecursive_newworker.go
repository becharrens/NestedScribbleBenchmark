package roles

import "NestedScribbleBenchmark/fannkuch/messages"
import "NestedScribbleBenchmark/fannkuch/channels/fannkuchrecursive"
import "NestedScribbleBenchmark/fannkuch/invitations"
import "NestedScribbleBenchmark/fannkuch/callbacks"
import "sync"

func FannkuchRecursive_NewWorker(wg *sync.WaitGroup, roleChannels fannkuchrecursive.NewWorker_Chan, inviteChannels invitations.FannkuchRecursive_NewWorker_InviteChan, env callbacks.FannkuchRecursive_NewWorker_Env) {
	defer wg.Done()
	<-roleChannels.Label_From_Worker
	IdxMin := <-roleChannels.Int_From_Worker
	Chunksz := <-roleChannels.Int_From_Worker
	N := <-roleChannels.Int_From_Worker
	env.Task_From_Worker(IdxMin, Chunksz, N)

	newworker_choice := env.NewWorker_Choice()
	switch newworker_choice {
	case callbacks.FannkuchRecursive_NewWorker_FannkuchRecursive:
		env.FannkuchRecursive_Setup()
		roleChannels.Label_To_Source <- messages.FannkuchRecursive_Source_NewWorker

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

		MaxFlips, Checksum := env.Result_To_Source()
		roleChannels.Label_To_Source <- messages.Result
		roleChannels.Int_To_Source <- MaxFlips
		roleChannels.Int_To_Source <- Checksum

		env.Done()
		return
	case callbacks.FannkuchRecursive_NewWorker_Result:
		MaxFlips_2, Checksum_2 := env.Result_To_Source_2()
		roleChannels.Label_To_Source <- messages.Result
		roleChannels.Int_To_Source <- MaxFlips_2
		roleChannels.Int_To_Source <- Checksum_2

		env.Done()
		return
	default:
		panic("Invalid choice was made")
	}
}
