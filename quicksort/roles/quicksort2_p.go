package roles

import "NestedScribbleBenchmark/quicksort/messages"
import "NestedScribbleBenchmark/quicksort/channels/quicksort2"
import "NestedScribbleBenchmark/quicksort/invitations"
import "NestedScribbleBenchmark/quicksort/callbacks"
import quicksort2_2 "NestedScribbleBenchmark/quicksort/results/quicksort2"
import "sync"

func QuickSort2_P(wg *sync.WaitGroup, roleChannels quicksort2.P_Chan, inviteChannels invitations.QuickSort2_P_InviteChan, env callbacks.QuickSort2_P_Env) quicksort2_2.P_Result {
	p_choice := env.P_Choice()
	switch p_choice {
	case callbacks.QuickSort2_P_LeftParitition:
		arr := env.LeftParitition_To_L()
		roleChannels.Label_To_L <- messages.LeftParitition
		roleChannels.IntArr_To_L <- arr

		arr_2 := env.RightPartition_To_R()
		roleChannels.Label_To_R <- messages.RightPartition
		roleChannels.IntArr_To_R <- arr_2

		<-roleChannels.Label_From_L
		arr_3 := <-roleChannels.IntArr_From_L
		env.SortedLeft_From_L(arr_3)

		<-roleChannels.Label_From_R
		arr_4 := <-roleChannels.IntArr_From_R
		env.SortedRight_From_R(arr_4)

		return env.Done()
	case callbacks.QuickSort2_P_Done:
		env.Done_To_L()
		roleChannels.Label_To_L <- messages.Done

		env.Done_To_R()
		roleChannels.Label_To_R <- messages.Done

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
}
