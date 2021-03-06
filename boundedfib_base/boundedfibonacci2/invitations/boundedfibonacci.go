package invitations

import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/channels/boundedfibonacci"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/channels/boundedfib"

type BoundedFibonacci_RoleSetupChan struct {
	F1_Chan    chan boundedfibonacci.F1_Chan
	F2_Chan    chan boundedfibonacci.F2_Chan
	Start_Chan chan boundedfibonacci.Start_Chan
}

type BoundedFibonacci_InviteSetupChan struct {
	F1_InviteChan    chan BoundedFibonacci_F1_InviteChan
	F2_InviteChan    chan BoundedFibonacci_F2_InviteChan
	Start_InviteChan chan BoundedFibonacci_Start_InviteChan
}

type BoundedFibonacci_Start_InviteChan struct {
	Invite_F1_To_BoundedFib_F1 chan boundedfib.F1_Chan
	// Invite_F1_To_BoundedFib_F1_InviteChan     chan BoundedFib_F1_InviteChan
	Invite_F2_To_BoundedFib_F2                chan boundedfib.F2_Chan
	Invite_F2_To_BoundedFib_F2_InviteChan     chan BoundedFib_F2_InviteChan
	Invite_Start_To_BoundedFib_Res            chan boundedfib.Res_Chan
	Invite_Start_To_BoundedFib_Res_InviteChan chan BoundedFib_Res_InviteChan
}

type BoundedFibonacci_F1_InviteChan struct {
	Start_Invite_To_BoundedFib_F1 chan boundedfib.F1_Chan
	// Start_Invite_To_BoundedFib_F1_InviteChan chan BoundedFib_F1_InviteChan
}

type BoundedFibonacci_F2_InviteChan struct {
	Start_Invite_To_BoundedFib_F2            chan boundedfib.F2_Chan
	Start_Invite_To_BoundedFib_F2_InviteChan chan BoundedFib_F2_InviteChan
}
