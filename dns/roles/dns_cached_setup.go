package roles

import "NestedScribbleBenchmark/dns/channels/dns_cached"
import "NestedScribbleBenchmark/dns/invitations"
import "sync"

func DNS_Cached_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.DNS_Cached_RoleSetupChan, inviteChannels invitations.DNS_Cached_InviteSetupChan)  {
	res_chan := dns_cached.Res_Chan{

	}

	res_inviteChan := invitations.DNS_Cached_res_InviteChan{

	}

	roleChannels.Res_Chan <- res_chan

	inviteChannels.Res_InviteChan <- res_inviteChan
} 