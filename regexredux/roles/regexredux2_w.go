package roles

import "ScribbleBenchmark/regexredux/channels/regexredux2"
import "ScribbleBenchmark/regexredux/invitations"
import "ScribbleBenchmark/regexredux/callbacks"
import "sync"

func RegexRedux2_W(wg *sync.WaitGroup, roleChannels regexredux2.W_Chan, inviteChannels invitations.RegexRedux2_W_InviteChan, env callbacks.RegexRedux2_W_Env)  {
	defer wg.Done()
	select {
		case task_msg := <-roleChannels.M_Task:
			env.Task_From_M(task_msg)

			nummatches_msg := env.NumMatches_To_M()
			roleChannels.M_NumMatches <- nummatches_msg

			env.Done()
			return 
		case calclength_msg := <-roleChannels.M_CalcLength:
			env.CalcLength_From_M(calclength_msg)

			length_msg := env.Length_To_M()
			roleChannels.M_Length <- length_msg

			env.Done()
			return 
	}
} 