package roles

import "NestedScribbleBenchmark/quicksort/channels/quicksort2"
import "NestedScribbleBenchmark/quicksort/invitations"
import "NestedScribbleBenchmark/quicksort/callbacks"
import "sync"

func QuickSort2_R(wg *sync.WaitGroup, roleChannels quicksort2.R_Chan, inviteChannels invitations.QuickSort2_R_InviteChan, env callbacks.QuickSort2_R_Env) {
	defer wg.Done()
	select {
	case done_msg := <-roleChannels.P_Done:
		env.Done_From_P(done_msg)

		env.Done()
		return
	case rightpartition_msg := <-roleChannels.P_RightPartition:
		env.RightPartition_From_P(rightpartition_msg)

		env.QuickSort2_Setup()
		quicksort2_rolechan := invitations.QuickSort2_RoleSetupChan{
			P_Chan: inviteChannels.Invite_R_To_QuickSort2_P,
		}
		quicksort2_invitechan := invitations.QuickSort2_InviteSetupChan{
			P_InviteChan: inviteChannels.Invite_R_To_QuickSort2_P_InviteChan,
		}
		QuickSort2_SendCommChannels(wg, quicksort2_rolechan, quicksort2_invitechan)

		quicksort2_p_chan := <-inviteChannels.Invite_R_To_QuickSort2_P
		quicksort2_p_inviteChan := <-inviteChannels.Invite_R_To_QuickSort2_P_InviteChan
		quicksort2_p_env := env.To_QuickSort2_P_Env()
		quicksort2_p_result := QuickSort2_P(wg, quicksort2_p_chan, quicksort2_p_inviteChan, quicksort2_p_env)
		env.ResultFrom_QuickSort2_P(quicksort2_p_result)

		sortedright_msg := env.SortedRight_To_P()
		roleChannels.P_SortedRight <- sortedright_msg

		env.Done()
		return
	}
}
