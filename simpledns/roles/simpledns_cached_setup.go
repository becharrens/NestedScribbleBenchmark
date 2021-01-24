package roles

import "NestedScribbleBenchmark/simpledns/channels/simpledns_cached"
import "NestedScribbleBenchmark/simpledns/invitations"
import "sync"

func SimpleDNS_Cached_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.SimpleDNS_Cached_RoleSetupChan, inviteChannels invitations.SimpleDNS_Cached_InviteSetupChan)  {
	res_chan := simpledns_cached.Res_Chan{

	}

	res_inviteChan := invitations.SimpleDNS_Cached_res_InviteChan{

	}

	roleChannels.Res_Chan <- res_chan

	inviteChannels.Res_InviteChan <- res_inviteChan
} 