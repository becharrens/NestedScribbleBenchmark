package invitations

import "NestedScribbleBenchmark/noughtsandcrosses/channels/noughtsandcrosses"
import "NestedScribbleBenchmark/noughtsandcrosses/channels/calcmove"

type NoughtsAndCrosses_RoleSetupChan struct {
	P1_Chan chan noughtsandcrosses.P1_Chan
	P2_Chan chan noughtsandcrosses.P2_Chan
}

type NoughtsAndCrosses_InviteSetupChan struct {
	P1_InviteChan chan NoughtsAndCrosses_P1_InviteChan
	P2_InviteChan chan NoughtsAndCrosses_P2_InviteChan
}

type NoughtsAndCrosses_P1_InviteChan struct {
	Invite_P1_To_CalcMove_P chan calcmove.P_Chan
	Invite_P1_To_CalcMove_P_InviteChan chan CalcMove_P_InviteChan
}

type NoughtsAndCrosses_P2_InviteChan struct {
	Invite_P2_To_CalcMove_P chan calcmove.P_Chan
	Invite_P2_To_CalcMove_P_InviteChan chan CalcMove_P_InviteChan
}