package roles

import "ScribbleBenchmark/quicksort/channels/quicksort"
import "ScribbleBenchmark/quicksort/invitations"
import "ScribbleBenchmark/quicksort/callbacks"
import quicksort_2 "ScribbleBenchmark/quicksort/results/quicksort"
import "sync"

func QuickSort_Left(wg *sync.WaitGroup, roleChannels quicksort.Left_Chan, inviteChannels invitations.QuickSort_Left_InviteChan, env callbacks.QuickSort_Left_Env) quicksort_2.Left_Result {
	select {
		case leftparitition_msg := <-roleChannels.Partition_LeftParitition:
			env.LeftParitition_From_Partition(leftparitition_msg)

			env.QuickSort2_Setup()
			quicksort2_rolechan := invitations.QuickSort2_RoleSetupChan{
				P_Chan: inviteChannels.Invite_Left_To_QuickSort2_P,
			}
			quicksort2_invitechan := invitations.QuickSort2_InviteSetupChan{
				P_InviteChan: inviteChannels.Invite_Left_To_QuickSort2_P_InviteChan,
			}
			QuickSort2_SendCommChannels(wg, quicksort2_rolechan, quicksort2_invitechan)

			quicksort2_p_chan := <-inviteChannels.Invite_Left_To_QuickSort2_P
			quicksort2_p_inviteChan := <-inviteChannels.Invite_Left_To_QuickSort2_P_InviteChan
			quicksort2_p_env := env.To_QuickSort2_P_Env()
			quicksort2_p_result := QuickSort2_P(wg, quicksort2_p_chan, quicksort2_p_inviteChan, quicksort2_p_env)
			env.ResultFrom_QuickSort2_P(quicksort2_p_result)

			sortedleft_msg := env.SortedLeft_To_Partition()
			roleChannels.Partition_SortedLeft <- sortedleft_msg

			return env.Done()
		case done_msg := <-roleChannels.Partition_Done:
			env.Done_From_Partition(done_msg)

			return env.Done()
	}
} 