package roles

import "NestedScribbleBenchmark/noughtsandcrosses/messages"
import "NestedScribbleBenchmark/noughtsandcrosses/channels/noughtsandcrosses"
import "NestedScribbleBenchmark/noughtsandcrosses/channels/calcmove"
import "NestedScribbleBenchmark/noughtsandcrosses/invitations"
import "sync"

func NoughtsAndCrosses_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.NoughtsAndCrosses_RoleSetupChan, inviteChannels invitations.NoughtsAndCrosses_InviteSetupChan)  {
	p2_p1_int := make(chan int, 1)
	p2_p1_label := make(chan messages.NoughtsAndCrosses_Label, 1)
	p2_invite_p2 := make(chan calcmove.P_Chan, 1)
	p2_invite_p2_invitechan := make(chan invitations.CalcMove_P_InviteChan, 1)
	p1_p2_int := make(chan int, 1)
	p1_p2_label := make(chan messages.NoughtsAndCrosses_Label, 1)
	p1_invite_p1 := make(chan calcmove.P_Chan, 1)
	p1_invite_p1_invitechan := make(chan invitations.CalcMove_P_InviteChan, 1)

	p2_chan := noughtsandcrosses.P2_Chan{
		Label_To_P1: p2_p1_label,
		Label_From_P1: p1_p2_label,
		Int_To_P1: p2_p1_int,
		Int_From_P1: p1_p2_int,
	}
	p1_chan := noughtsandcrosses.P1_Chan{
		Label_To_P2: p1_p2_label,
		Label_From_P2: p2_p1_label,
		Int_To_P2: p1_p2_int,
		Int_From_P2: p2_p1_int,
	}

	p2_inviteChan := invitations.NoughtsAndCrosses_P2_InviteChan{
		Invite_P2_To_CalcMove_P_InviteChan: p2_invite_p2_invitechan,
		Invite_P2_To_CalcMove_P: p2_invite_p2,
	}
	p1_inviteChan := invitations.NoughtsAndCrosses_P1_InviteChan{
		Invite_P1_To_CalcMove_P_InviteChan: p1_invite_p1_invitechan,
		Invite_P1_To_CalcMove_P: p1_invite_p1,
	}

	roleChannels.P1_Chan <- p1_chan
	roleChannels.P2_Chan <- p2_chan

	inviteChannels.P1_InviteChan <- p1_inviteChan
	inviteChannels.P2_InviteChan <- p2_inviteChan
} 