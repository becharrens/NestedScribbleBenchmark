package roles

import "NestedScribbleBenchmark/quicksort/messages"
import "NestedScribbleBenchmark/quicksort/channels/quicksort2"
import "NestedScribbleBenchmark/quicksort/invitations"
import "NestedScribbleBenchmark/quicksort/callbacks"
import "sync"

func QuickSort2_R(wg *sync.WaitGroup, roleChannels quicksort2.R_Chan, inviteChannels invitations.QuickSort2_R_InviteChan, env callbacks.QuickSort2_R_Env) {
	defer wg.Done()
	p_choice := <-roleChannels.Label_From_P
	switch p_choice {
	case messages.Done:
		env.Done_From_P()

		env.Done()
		return
	case messages.RightPartition:
		arr := <-roleChannels.IntArr_From_P
		env.RightPartition_From_P(arr)

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

		arr_2 := env.SortedRight_To_P()
		roleChannels.Label_To_P <- messages.SortedRight
		roleChannels.IntArr_To_P <- arr_2

		env.Done()
		return
	default:
		panic("Invalid choice was made")
	}
}
