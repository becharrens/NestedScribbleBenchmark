package roles

import "NestedScribbleBenchmark/noughtsandcrosses/channels/minmaxstrategy_evalboard"
import "NestedScribbleBenchmark/noughtsandcrosses/invitations"
import "sync"

func MinMaxStrategy_EvalBoard_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.MinMaxStrategy_EvalBoard_RoleSetupChan, inviteChannels invitations.MinMaxStrategy_EvalBoard_InviteSetupChan)  {
	w_chan := minmaxstrategy_evalboard.W_Chan{

	}

	w_inviteChan := invitations.MinMaxStrategy_EvalBoard_W_InviteChan{

	}

	roleChannels.W_Chan <- w_chan

	inviteChannels.W_InviteChan <- w_inviteChan
} 