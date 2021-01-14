package invitations

import "NestedScribbleBenchmark/ring/channels/ring"
import "NestedScribbleBenchmark/ring/channels/forward"

type Ring_RoleSetupChan struct {
	End_Chan chan ring.End_Chan
	Start_Chan chan ring.Start_Chan
}

type Ring_InviteSetupChan struct {
	End_InviteChan chan Ring_End_InviteChan
	Start_InviteChan chan Ring_Start_InviteChan
}

type Ring_Start_InviteChan struct {
	Invite_End_To_Forward_E chan forward.E_Chan
	Invite_End_To_Forward_E_InviteChan chan Forward_E_InviteChan
	Invite_Start_To_Forward_S chan forward.S_Chan
	Invite_Start_To_Forward_S_InviteChan chan Forward_S_InviteChan
}

type Ring_End_InviteChan struct {
	Start_Invite_To_Forward_E chan forward.E_Chan
	Start_Invite_To_Forward_E_InviteChan chan Forward_E_InviteChan
}