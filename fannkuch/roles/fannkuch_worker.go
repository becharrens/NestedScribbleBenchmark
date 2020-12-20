package roles

import "NestedScribbleBenchmark/fannkuch/messages"
import "NestedScribbleBenchmark/fannkuch/channels/fannkuch"
import "NestedScribbleBenchmark/fannkuch/invitations"
import "NestedScribbleBenchmark/fannkuch/callbacks"
import fannkuch_2 "NestedScribbleBenchmark/fannkuch/results/fannkuch"
import "sync"

func Fannkuch_Worker(wg *sync.WaitGroup, roleChannels fannkuch.Worker_Chan, inviteChannels invitations.Fannkuch_Worker_InviteChan, env callbacks.Fannkuch_Worker_Env) fannkuch_2.Worker_Result {
	<-roleChannels.Label_From_Main
	IdxMin := <-roleChannels.Int_From_Main
	Chunksz := <-roleChannels.Int_From_Main
	N := <-roleChannels.Int_From_Main
	env.Task_From_Main(IdxMin, Chunksz, N)

	worker_choice := env.Worker_Choice()
	switch worker_choice {
	case callbacks.Fannkuch_Worker_FannkuchRecursive:
		env.FannkuchRecursive_Setup()
		roleChannels.Label_To_Main <- messages.FannkuchRecursive_Main_Worker

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

		MaxFlips, Checksum := env.Result_To_Main()
		roleChannels.Label_To_Main <- messages.Result
		roleChannels.Int_To_Main <- MaxFlips
		roleChannels.Int_To_Main <- Checksum

		return env.Done()
	case callbacks.Fannkuch_Worker_Result:
		MaxFlips_2, Checksum_2 := env.Result_To_Main_2()
		roleChannels.Label_To_Main <- messages.Result
		roleChannels.Int_To_Main <- MaxFlips_2
		roleChannels.Int_To_Main <- Checksum_2

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
}
