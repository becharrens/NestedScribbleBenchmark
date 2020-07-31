package roles

import "ScribbleBenchmark/quicksort/channels/quicksort"
import "ScribbleBenchmark/quicksort/invitations"
import "ScribbleBenchmark/quicksort/callbacks"
import quicksort_2 "ScribbleBenchmark/quicksort/results/quicksort"
import "sync"

func QuickSort_Partition(wg *sync.WaitGroup, roleChannels quicksort.Partition_Chan, inviteChannels invitations.QuickSort_Partition_InviteChan, env callbacks.QuickSort_Partition_Env) quicksort_2.Partition_Result {
	partition_choice := env.Partition_Choice()
	switch partition_choice {
	case callbacks.QuickSort_Partition_LeftParitition:
		leftparitition_msg := env.LeftParitition_To_Left()
		roleChannels.Left_LeftParitition <- leftparitition_msg

		rightpartition_msg := env.RightPartition_To_Right()
		roleChannels.Right_RightPartition <- rightpartition_msg

		sortedleft_msg := <-roleChannels.Left_SortedLeft
		env.SortedLeft_From_Left(sortedleft_msg)

		sortedright_msg := <-roleChannels.Right_SortedRight
		env.SortedRight_From_Right(sortedright_msg)

		return env.Done()
	case callbacks.QuickSort_Partition_Done:
		done_msg := env.Done_To_Left()
		roleChannels.Left_Done <- done_msg

		done_msg_2 := env.Done_To_Right()
		roleChannels.Right_Done <- done_msg_2

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
} 