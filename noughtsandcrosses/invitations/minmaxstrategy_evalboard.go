package invitations

import "NestedScribbleBenchmark/noughtsandcrosses/channels/minmaxstrategy_evalboard"

type MinMaxStrategy_EvalBoard_RoleSetupChan struct {
	W_Chan chan minmaxstrategy_evalboard.W_Chan
}

type MinMaxStrategy_EvalBoard_InviteSetupChan struct {
	W_InviteChan chan MinMaxStrategy_EvalBoard_W_InviteChan
}

type MinMaxStrategy_EvalBoard_W_InviteChan struct {

}