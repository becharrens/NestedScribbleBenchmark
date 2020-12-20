package roles

import "NestedScribbleBenchmark/quicksort/messages"
import "NestedScribbleBenchmark/quicksort/channels/quicksort"
import "NestedScribbleBenchmark/quicksort/invitations"
import "NestedScribbleBenchmark/quicksort/callbacks"
import quicksort_2 "NestedScribbleBenchmark/quicksort/results/quicksort"
import "sync"

func QuickSort_Partition(wg *sync.WaitGroup, roleChannels quicksort.Partition_Chan, inviteChannels invitations.QuickSort_Partition_InviteChan, env callbacks.QuickSort_Partition_Env) quicksort_2.Partition_Result {
	partition_choice := env.Partition_Choice()
	switch partition_choice {
	case callbacks.QuickSort_Partition_LeftParitition:
		arr := env.LeftParitition_To_Left()
		roleChannels.Label_To_Left <- messages.LeftParitition
		roleChannels.IntArr_To_Left <- arr

		arr_2 := env.RightPartition_To_Right()
		roleChannels.Label_To_Right <- messages.RightPartition
		roleChannels.IntArr_To_Right <- arr_2

		<-roleChannels.Label_From_Left
		arr_3 := <-roleChannels.IntArr_From_Left
		env.SortedLeft_From_Left(arr_3)

		<-roleChannels.Label_From_Right
		arr_4 := <-roleChannels.IntArr_From_Right
		env.SortedRight_From_Right(arr_4)

		return env.Done()
	case callbacks.QuickSort_Partition_Done:
		env.Done_To_Left()
		roleChannels.Label_To_Left <- messages.Done

		env.Done_To_Right()
		roleChannels.Label_To_Right <- messages.Done

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
}
