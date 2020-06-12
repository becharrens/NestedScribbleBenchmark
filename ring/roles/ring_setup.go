package roles

import "ScribbleBenchmark/ring/messages/ring"
import ring_2 "ScribbleBenchmark/ring/channels/ring"
import "ScribbleBenchmark/ring/channels/forward"
import "ScribbleBenchmark/ring/invitations"
import "sync"

func Ring_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.Ring_RoleSetupChan, inviteChannels invitations.Ring_InviteSetupChan)  {
	end_start_msg_2 := make(chan ring.Msg, 1)
	start_end_msg := make(chan ring.Msg, 1)
	end_start_msg := make(chan ring.Msg, 1)
	start_invite_end := make(chan forward.E_Chan, 1)
	start_invite_end_invitechan := make(chan invitations.Forward_E_InviteChan, 1)
	start_invite_start := make(chan forward.S_Chan, 1)
	start_invite_start_invitechan := make(chan invitations.Forward_S_InviteChan, 1)

	start_chan := ring_2.Start_Chan{
		End_Msg_3: end_start_msg_2,
		End_Msg_2: start_end_msg,
		End_Msg: end_start_msg,
	}
	end_chan := ring_2.End_Chan{
		Start_Msg_3: end_start_msg_2,
		Start_Msg_2: start_end_msg,
		Start_Msg: end_start_msg,
	}

	start_inviteChan := invitations.Ring_Start_InviteChan{
		Invite_Start_To_Forward_S_InviteChan: start_invite_start_invitechan,
		Invite_Start_To_Forward_S: start_invite_start,
		Invite_End_To_Forward_E_InviteChan: start_invite_end_invitechan,
		Invite_End_To_Forward_E: start_invite_end,
	}
	end_inviteChan := invitations.Ring_End_InviteChan{
		Start_Invite_To_Forward_E_InviteChan: start_invite_end_invitechan,
		Start_Invite_To_Forward_E: start_invite_end,
	}

	roleChannels.Start_Chan <- start_chan
	roleChannels.End_Chan <- end_chan

	inviteChannels.Start_InviteChan <- start_inviteChan
	inviteChannels.End_InviteChan <- end_inviteChan
} 