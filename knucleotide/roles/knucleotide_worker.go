package roles

import "ScribbleBenchmark/knucleotide/channels/knucleotide"
import "ScribbleBenchmark/knucleotide/invitations"
import "ScribbleBenchmark/knucleotide/callbacks"
import knucleotide_2 "ScribbleBenchmark/knucleotide/results/knucleotide"
import "sync"

func KNucleotide_Worker(wg *sync.WaitGroup, roleChannels knucleotide.Worker_Chan, inviteChannels invitations.KNucleotide_Worker_InviteChan, env callbacks.KNucleotide_Worker_Env) knucleotide_2.Worker_Result {
	sequencejob_msg := <-roleChannels.Master_SequenceJob
	env.SequenceJob_From_Master(sequencejob_msg)

	sequenceresult_msg := env.SequenceResult_To_Master()
	roleChannels.Master_SequenceResult <- sequenceresult_msg

	return env.Done()
} 