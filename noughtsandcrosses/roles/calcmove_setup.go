package roles

import "NestedScribbleBenchmark/noughtsandcrosses/channels/standardstrategy"
import "NestedScribbleBenchmark/noughtsandcrosses/channels/minmaxstrategy"
import "NestedScribbleBenchmark/noughtsandcrosses/channels/calcmove"
import "NestedScribbleBenchmark/noughtsandcrosses/invitations"
import "sync"

func CalcMove_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.CalcMove_RoleSetupChan, inviteChannels invitations.CalcMove_InviteSetupChan)  {
	p_invite_p_2 := make(chan minmaxstrategy.Master_Chan, 1)
	p_invite_p_invitechan_2 := make(chan invitations.MinMaxStrategy_Master_InviteChan, 1)
	p_invite_p := make(chan standardstrategy.P_Chan, 1)
	p_invite_p_invitechan := make(chan invitations.StandardStrategy_P_InviteChan, 1)

	p_chan := calcmove.P_Chan{

	}

	p_inviteChan := invitations.CalcMove_P_InviteChan{
		Invite_P_To_StandardStrategy_P_InviteChan: p_invite_p_invitechan,
		Invite_P_To_StandardStrategy_P: p_invite_p,
		Invite_P_To_MinMaxStrategy_Master_InviteChan: p_invite_p_invitechan_2,
		Invite_P_To_MinMaxStrategy_Master: p_invite_p_2,
	}

	roleChannels.P_Chan <- p_chan

	inviteChannels.P_InviteChan <- p_inviteChan
} 