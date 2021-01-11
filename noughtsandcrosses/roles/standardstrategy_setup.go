package roles

import "NestedScribbleBenchmark/noughtsandcrosses/channels/standardstrategy"
import "NestedScribbleBenchmark/noughtsandcrosses/invitations"
import "sync"

func StandardStrategy_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.StandardStrategy_RoleSetupChan, inviteChannels invitations.StandardStrategy_InviteSetupChan)  {
	p_chan := standardstrategy.P_Chan{

	}

	p_inviteChan := invitations.StandardStrategy_P_InviteChan{

	}

	roleChannels.P_Chan <- p_chan

	inviteChannels.P_InviteChan <- p_inviteChan
} 