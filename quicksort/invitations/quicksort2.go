package invitations

import "ScribbleBenchmark/quicksort/channels/quicksort2"

type QuickSort2_RoleSetupChan struct {
	P_Chan chan quicksort2.P_Chan
}

type QuickSort2_InviteSetupChan struct {
	P_InviteChan chan QuickSort2_P_InviteChan
}

type QuickSort2_P_InviteChan struct {

}

type QuickSort2_L_InviteChan struct {
	Invite_L_To_QuickSort2_P chan quicksort2.P_Chan
	Invite_L_To_QuickSort2_P_InviteChan chan QuickSort2_P_InviteChan
}

type QuickSort2_R_InviteChan struct {
	Invite_R_To_QuickSort2_P chan quicksort2.P_Chan
	Invite_R_To_QuickSort2_P_InviteChan chan QuickSort2_P_InviteChan
}