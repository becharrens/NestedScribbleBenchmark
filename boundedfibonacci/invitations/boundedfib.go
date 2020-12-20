package invitations

import "NestedScribbleBenchmark/boundedfibonacci/channels/boundedfib"

type BoundedFib_RoleSetupChan struct {
	F1_Chan  chan boundedfib.F1_Chan
	F2_Chan  chan boundedfib.F2_Chan
	Res_Chan chan boundedfib.Res_Chan
}

type BoundedFib_InviteSetupChan struct {
	F1_InviteChan  chan BoundedFib_F1_InviteChan
	F2_InviteChan  chan BoundedFib_F2_InviteChan
	Res_InviteChan chan BoundedFib_Res_InviteChan
}

type BoundedFib_Res_InviteChan struct {
	F3_Invite_To_BoundedFib_Res            chan boundedfib.Res_Chan
	F3_Invite_To_BoundedFib_Res_InviteChan chan BoundedFib_Res_InviteChan
}

type BoundedFib_F1_InviteChan struct {
}

type BoundedFib_F2_InviteChan struct {
	F3_Invite_To_BoundedFib_F1            chan boundedfib.F1_Chan
	F3_Invite_To_BoundedFib_F1_InviteChan chan BoundedFib_F1_InviteChan
}

type BoundedFib_F3_InviteChan struct {
	Invite_F2_To_BoundedFib_F1              chan boundedfib.F1_Chan
	Invite_F2_To_BoundedFib_F1_InviteChan   chan BoundedFib_F1_InviteChan
	Invite_F3_To_BoundedFib_F2              chan boundedfib.F2_Chan
	Invite_F3_To_BoundedFib_F2_InviteChan   chan BoundedFib_F2_InviteChan
	Invite_Res_To_BoundedFib_Res            chan boundedfib.Res_Chan
	Invite_Res_To_BoundedFib_Res_InviteChan chan BoundedFib_Res_InviteChan
}
