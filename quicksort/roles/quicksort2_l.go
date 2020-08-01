package roles

import "NestedScribbleBenchmark/quicksort/channels/quicksort2"
import "NestedScribbleBenchmark/quicksort/invitations"
import "NestedScribbleBenchmark/quicksort/callbacks"
import "sync"

func QuickSort2_L(wg *sync.WaitGroup, roleChannels quicksort2.L_Chan, inviteChannels invitations.QuickSort2_L_InviteChan, env callbacks.QuickSort2_L_Env) {
	defer wg.Done()
	select {
	case leftparitition_msg := <-roleChannels.P_LeftParitition:
		env.LeftParitition_From_P(leftparitition_msg)

		env.QuickSort2_Setup()
		quicksort2_rolechan := invitations.QuickSort2_RoleSetupChan{
			P_Chan: inviteChannels.Invite_L_To_QuickSort2_P,
		}
		quicksort2_invitechan := invitations.QuickSort2_InviteSetupChan{
			P_InviteChan: inviteChannels.Invite_L_To_QuickSort2_P_InviteChan,
		}
		QuickSort2_SendCommChannels(wg, quicksort2_rolechan, quicksort2_invitechan)

		quicksort2_p_chan := <-inviteChannels.Invite_L_To_QuickSort2_P
		quicksort2_p_inviteChan := <-inviteChannels.Invite_L_To_QuickSort2_P_InviteChan
		quicksort2_p_env := env.To_QuickSort2_P_Env()
		quicksort2_p_result := QuickSort2_P(wg, quicksort2_p_chan, quicksort2_p_inviteChan, quicksort2_p_env)
		env.ResultFrom_QuickSort2_P(quicksort2_p_result)

		sortedleft_msg := env.SortedLeft_To_P()
		roleChannels.P_SortedLeft <- sortedleft_msg

		env.Done()
		return
	case done_msg := <-roleChannels.P_Done:
		env.Done_From_P(done_msg)

		env.Done()
		return
	}
}
