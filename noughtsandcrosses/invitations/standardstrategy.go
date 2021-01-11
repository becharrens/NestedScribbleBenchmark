package invitations

import "NestedScribbleBenchmark/noughtsandcrosses/channels/standardstrategy"

type StandardStrategy_RoleSetupChan struct {
	P_Chan chan standardstrategy.P_Chan
}

type StandardStrategy_InviteSetupChan struct {
	P_InviteChan chan StandardStrategy_P_InviteChan
}

type StandardStrategy_P_InviteChan struct {

}