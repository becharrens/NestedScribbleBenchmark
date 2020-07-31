package roles

import "ScribbleBenchmark/quicksort/messages/quicksort"
import "ScribbleBenchmark/quicksort/channels/quicksort2"
import quicksort_2 "ScribbleBenchmark/quicksort/channels/quicksort"
import "ScribbleBenchmark/quicksort/invitations"
import "sync"

func QuickSort_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.QuickSort_RoleSetupChan, inviteChannels invitations.QuickSort_InviteSetupChan)  {
	partition_right_done := make(chan quicksort.Done, 1)
	partition_left_done := make(chan quicksort.Done, 1)
	right_partition_sortedright := make(chan quicksort.SortedRight, 1)
	left_partition_sortedleft := make(chan quicksort.SortedLeft, 1)
	right_invite_right := make(chan quicksort2.P_Chan, 1)
	right_invite_right_invitechan := make(chan invitations.QuickSort2_P_InviteChan, 1)
	left_invite_left := make(chan quicksort2.P_Chan, 1)
	left_invite_left_invitechan := make(chan invitations.QuickSort2_P_InviteChan, 1)
	partition_right_rightpartition := make(chan quicksort.RightPartition, 1)
	partition_left_leftparitition := make(chan quicksort.LeftParitition, 1)

	right_chan := quicksort_2.Right_Chan{
		Partition_SortedRight: right_partition_sortedright,
		Partition_RightPartition: partition_right_rightpartition,
		Partition_Done: partition_right_done,
	}
	partition_chan := quicksort_2.Partition_Chan{
		Right_SortedRight: right_partition_sortedright,
		Right_RightPartition: partition_right_rightpartition,
		Right_Done: partition_right_done,
		Left_SortedLeft: left_partition_sortedleft,
		Left_LeftParitition: partition_left_leftparitition,
		Left_Done: partition_left_done,
	}
	left_chan := quicksort_2.Left_Chan{
		Partition_SortedLeft: left_partition_sortedleft,
		Partition_LeftParitition: partition_left_leftparitition,
		Partition_Done: partition_left_done,
	}

	right_inviteChan := invitations.QuickSort_Right_InviteChan{
		Invite_Right_To_QuickSort2_P_InviteChan: right_invite_right_invitechan,
		Invite_Right_To_QuickSort2_P: right_invite_right,
	}
	partition_inviteChan := invitations.QuickSort_Partition_InviteChan{

	}
	left_inviteChan := invitations.QuickSort_Left_InviteChan{
		Invite_Left_To_QuickSort2_P_InviteChan: left_invite_left_invitechan,
		Invite_Left_To_QuickSort2_P: left_invite_left,
	}

	roleChannels.Partition_Chan <- partition_chan
	roleChannels.Left_Chan <- left_chan
	roleChannels.Right_Chan <- right_chan

	inviteChannels.Partition_InviteChan <- partition_inviteChan
	inviteChannels.Left_InviteChan <- left_inviteChan
	inviteChannels.Right_InviteChan <- right_inviteChan
} 