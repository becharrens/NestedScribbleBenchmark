package roles

import "NestedScribbleBenchmark/knucleotide/channels/schedulejobs"
import "NestedScribbleBenchmark/knucleotide/invitations"
import "NestedScribbleBenchmark/knucleotide/callbacks"
import "sync"

func ScheduleJobs_W(wg *sync.WaitGroup, roleChannels schedulejobs.W_Chan, inviteChannels invitations.ScheduleJobs_W_InviteChan, env callbacks.ScheduleJobs_W_Env) {
	defer wg.Done()
	select {
	case sequencejob_msg := <-roleChannels.M_SequenceJob:
		env.SequenceJob_From_M(sequencejob_msg)

		sequenceresult_msg := env.SequenceResult_To_M()
		roleChannels.M_SequenceResult <- sequenceresult_msg

		env.Done()
		return
	case frequencyjob_msg := <-roleChannels.M_FrequencyJob:
		env.FrequencyJob_From_M(frequencyjob_msg)

		frequencyresult_msg := env.FrequencyResult_To_M()
		roleChannels.M_FrequencyResult <- frequencyresult_msg

		env.Done()
		return
	case finish_msg := <-roleChannels.M_Finish:
		env.Finish_From_M(finish_msg)

		env.Done()
		return
	}
}
