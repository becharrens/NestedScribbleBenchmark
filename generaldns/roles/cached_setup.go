package roles

import "NestedScribbleBenchmark/generaldns/channels/cached"
import "NestedScribbleBenchmark/generaldns/invitations"
import "sync"

func Cached_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.Cached_RoleSetupChan, inviteChannels invitations.Cached_InviteSetupChan)  {
	res_chan := cached.Res_Chan{

	}

	res_inviteChan := invitations.Cached_res_InviteChan{

	}

	roleChannels.Res_Chan <- res_chan

	inviteChannels.Res_InviteChan <- res_inviteChan
} 