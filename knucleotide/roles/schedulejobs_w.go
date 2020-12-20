package roles

import "NestedScribbleBenchmark/knucleotide/messages"
import "NestedScribbleBenchmark/knucleotide/channels/schedulejobs"
import "NestedScribbleBenchmark/knucleotide/invitations"
import "NestedScribbleBenchmark/knucleotide/callbacks"
import "sync"

func ScheduleJobs_W(wg *sync.WaitGroup, roleChannels schedulejobs.W_Chan, inviteChannels invitations.ScheduleJobs_W_InviteChan, env callbacks.ScheduleJobs_W_Env) {
	defer wg.Done()
	m_choice := <-roleChannels.Label_From_M
	switch m_choice {
	case messages.SequenceJob:
		sequence := <-roleChannels.String_From_M
		dna := <-roleChannels.ByteArr_From_M
		env.SequenceJob_From_M(sequence, dna)

		res := env.SequenceResult_To_M()
		roleChannels.Label_To_M <- messages.SequenceResult
		roleChannels.String_To_M <- res

		env.Done()
		return
	case messages.FrequencyJob:
		len := <-roleChannels.Int_From_M
		dna_2 := <-roleChannels.ByteArr_From_M
		env.FrequencyJob_From_M(len, dna_2)

		res_2 := env.FrequencyResult_To_M()
		roleChannels.Label_To_M <- messages.FrequencyResult
		roleChannels.String_To_M <- res_2

		env.Done()
		return
	case messages.Finish:
		env.Finish_From_M()

		env.Done()
		return
	default:
		panic("Invalid choice was made")
	}
}
