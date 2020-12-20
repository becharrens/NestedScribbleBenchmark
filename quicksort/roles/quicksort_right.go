package roles

import "NestedScribbleBenchmark/quicksort/messages"
import "NestedScribbleBenchmark/quicksort/channels/quicksort"
import "NestedScribbleBenchmark/quicksort/invitations"
import "NestedScribbleBenchmark/quicksort/callbacks"
import quicksort_2 "NestedScribbleBenchmark/quicksort/results/quicksort"
import "sync"

func QuickSort_Right(wg *sync.WaitGroup, roleChannels quicksort.Right_Chan, inviteChannels invitations.QuickSort_Right_InviteChan, env callbacks.QuickSort_Right_Env) quicksort_2.Right_Result {
	partition_choice := <-roleChannels.Label_From_Partition
	switch partition_choice {
	case messages.Done:
		env.Done_From_Partition()

		return env.Done()
	case messages.RightPartition:
		arr := <-roleChannels.IntArr_From_Partition
		env.RightPartition_From_Partition(arr)

		env.QuickSort2_Setup()

		quicksort2_rolechan := invitations.QuickSort2_RoleSetupChan{
			P_Chan: inviteChannels.Invite_Right_To_QuickSort2_P,
		}
		quicksort2_invitechan := invitations.QuickSort2_InviteSetupChan{
			P_InviteChan: inviteChannels.Invite_Right_To_QuickSort2_P_InviteChan,
		}
		QuickSort2_SendCommChannels(wg, quicksort2_rolechan, quicksort2_invitechan)

		quicksort2_p_chan := <-inviteChannels.Invite_Right_To_QuickSort2_P
		quicksort2_p_inviteChan := <-inviteChannels.Invite_Right_To_QuickSort2_P_InviteChan
		quicksort2_p_env := env.To_QuickSort2_P_Env()
		quicksort2_p_result := QuickSort2_P(wg, quicksort2_p_chan, quicksort2_p_inviteChan, quicksort2_p_env)
		env.ResultFrom_QuickSort2_P(quicksort2_p_result)

		arr_2 := env.SortedRight_To_Partition()
		roleChannels.Label_To_Partition <- messages.SortedRight
		roleChannels.IntArr_To_Partition <- arr_2

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
}
