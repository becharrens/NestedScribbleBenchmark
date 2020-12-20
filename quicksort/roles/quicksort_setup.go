package roles

import "NestedScribbleBenchmark/quicksort/messages"
import "NestedScribbleBenchmark/quicksort/channels/quicksort2"
import "NestedScribbleBenchmark/quicksort/channels/quicksort"
import "NestedScribbleBenchmark/quicksort/invitations"
import "sync"

func QuickSort_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.QuickSort_RoleSetupChan, inviteChannels invitations.QuickSort_InviteSetupChan) {
	right_partition_intarr := make(chan []int, 1)
	right_partition_label := make(chan messages.QuickSort_Label, 1)
	left_partition_intarr := make(chan []int, 1)
	left_partition_label := make(chan messages.QuickSort_Label, 1)
	right_invite_right := make(chan quicksort2.P_Chan, 1)
	right_invite_right_invitechan := make(chan invitations.QuickSort2_P_InviteChan, 1)
	left_invite_left := make(chan quicksort2.P_Chan, 1)
	left_invite_left_invitechan := make(chan invitations.QuickSort2_P_InviteChan, 1)
	partition_right_intarr := make(chan []int, 1)
	partition_right_label := make(chan messages.QuickSort_Label, 1)
	partition_left_intarr := make(chan []int, 1)
	partition_left_label := make(chan messages.QuickSort_Label, 1)

	right_chan := quicksort.Right_Chan{
		Label_To_Partition:    right_partition_label,
		Label_From_Partition:  partition_right_label,
		IntArr_To_Partition:   right_partition_intarr,
		IntArr_From_Partition: partition_right_intarr,
	}
	partition_chan := quicksort.Partition_Chan{
		Label_To_Right:    partition_right_label,
		Label_To_Left:     partition_left_label,
		Label_From_Right:  right_partition_label,
		Label_From_Left:   left_partition_label,
		IntArr_To_Right:   partition_right_intarr,
		IntArr_To_Left:    partition_left_intarr,
		IntArr_From_Right: right_partition_intarr,
		IntArr_From_Left:  left_partition_intarr,
	}
	left_chan := quicksort.Left_Chan{
		Label_To_Partition:    left_partition_label,
		Label_From_Partition:  partition_left_label,
		IntArr_To_Partition:   left_partition_intarr,
		IntArr_From_Partition: partition_left_intarr,
	}

	right_inviteChan := invitations.QuickSort_Right_InviteChan{
		Invite_Right_To_QuickSort2_P_InviteChan: right_invite_right_invitechan,
		Invite_Right_To_QuickSort2_P:            right_invite_right,
	}
	partition_inviteChan := invitations.QuickSort_Partition_InviteChan{}
	left_inviteChan := invitations.QuickSort_Left_InviteChan{
		Invite_Left_To_QuickSort2_P_InviteChan: left_invite_left_invitechan,
		Invite_Left_To_QuickSort2_P:            left_invite_left,
	}

	roleChannels.Partition_Chan <- partition_chan
	roleChannels.Left_Chan <- left_chan
	roleChannels.Right_Chan <- right_chan

	inviteChannels.Partition_InviteChan <- partition_inviteChan
	inviteChannels.Left_InviteChan <- left_inviteChan
	inviteChannels.Right_InviteChan <- right_inviteChan
}
