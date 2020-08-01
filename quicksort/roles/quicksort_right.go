package roles

import "NestedScribbleBenchmark/quicksort/channels/quicksort"
import "NestedScribbleBenchmark/quicksort/invitations"
import "NestedScribbleBenchmark/quicksort/callbacks"
import quicksort_2 "NestedScribbleBenchmark/quicksort/results/quicksort"
import "sync"

func QuickSort_Right(wg *sync.WaitGroup, roleChannels quicksort.Right_Chan, inviteChannels invitations.QuickSort_Right_InviteChan, env callbacks.QuickSort_Right_Env) quicksort_2.Right_Result {
	select {
	case done_msg := <-roleChannels.Partition_Done:
		env.Done_From_Partition(done_msg)

		return env.Done()
	case rightpartition_msg := <-roleChannels.Partition_RightPartition:
		env.RightPartition_From_Partition(rightpartition_msg)

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

		sortedright_msg := env.SortedRight_To_Partition()
		roleChannels.Partition_SortedRight <- sortedright_msg

		return env.Done()
	}
}
