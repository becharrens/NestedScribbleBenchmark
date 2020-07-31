package invitations

import "ScribbleBenchmark/quicksort/channels/quicksort2"
import "ScribbleBenchmark/quicksort/channels/quicksort"

type QuickSort_RoleSetupChan struct {
	Left_Chan chan quicksort.Left_Chan
	Partition_Chan chan quicksort.Partition_Chan
	Right_Chan chan quicksort.Right_Chan
}

type QuickSort_InviteSetupChan struct {
	Left_InviteChan chan QuickSort_Left_InviteChan
	Partition_InviteChan chan QuickSort_Partition_InviteChan
	Right_InviteChan chan QuickSort_Right_InviteChan
}

type QuickSort_Partition_InviteChan struct {

}

type QuickSort_Left_InviteChan struct {
	Invite_Left_To_QuickSort2_P chan quicksort2.P_Chan
	Invite_Left_To_QuickSort2_P_InviteChan chan QuickSort2_P_InviteChan
}

type QuickSort_Right_InviteChan struct {
	Invite_Right_To_QuickSort2_P chan quicksort2.P_Chan
	Invite_Right_To_QuickSort2_P_InviteChan chan QuickSort2_P_InviteChan
}