package roles

import "NestedScribbleBenchmark/quicksort/channels/quicksort2"
import "NestedScribbleBenchmark/quicksort/invitations"
import "NestedScribbleBenchmark/quicksort/callbacks"
import quicksort2_2 "NestedScribbleBenchmark/quicksort/results/quicksort2"
import "sync"

func QuickSort2_P(wg *sync.WaitGroup, roleChannels quicksort2.P_Chan, inviteChannels invitations.QuickSort2_P_InviteChan, env callbacks.QuickSort2_P_Env) quicksort2_2.P_Result {
	p_choice := env.P_Choice()
	switch p_choice {
	case callbacks.QuickSort2_P_LeftParitition:
		leftparitition_msg := env.LeftParitition_To_L()
		roleChannels.L_LeftParitition <- leftparitition_msg

		rightpartition_msg := env.RightPartition_To_R()
		roleChannels.R_RightPartition <- rightpartition_msg

		sortedleft_msg := <-roleChannels.L_SortedLeft
		env.SortedLeft_From_L(sortedleft_msg)

		sortedright_msg := <-roleChannels.R_SortedRight
		env.SortedRight_From_R(sortedright_msg)

		return env.Done()
	case callbacks.QuickSort2_P_Done:
		done_msg := env.Done_To_L()
		roleChannels.L_Done <- done_msg

		done_msg_2 := env.Done_To_R()
		roleChannels.R_Done <- done_msg_2

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
}
