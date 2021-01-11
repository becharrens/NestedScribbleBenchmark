package invitations

import "NestedScribbleBenchmark/noughtsandcrosses/channels/standardstrategy"
import "NestedScribbleBenchmark/noughtsandcrosses/channels/minmaxstrategy"
import "NestedScribbleBenchmark/noughtsandcrosses/channels/calcmove"

type CalcMove_RoleSetupChan struct {
	P_Chan chan calcmove.P_Chan
}

type CalcMove_InviteSetupChan struct {
	P_InviteChan chan CalcMove_P_InviteChan
}

type CalcMove_P_InviteChan struct {
	Invite_P_To_MinMaxStrategy_Master chan minmaxstrategy.Master_Chan
	Invite_P_To_MinMaxStrategy_Master_InviteChan chan MinMaxStrategy_Master_InviteChan
	Invite_P_To_StandardStrategy_P chan standardstrategy.P_Chan
	Invite_P_To_StandardStrategy_P_InviteChan chan StandardStrategy_P_InviteChan
}