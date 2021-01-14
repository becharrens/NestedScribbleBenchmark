package roles

import "NestedScribbleBenchmark/ring/messages"
import "NestedScribbleBenchmark/ring/channels/ring"
import "NestedScribbleBenchmark/ring/channels/forward"
import "NestedScribbleBenchmark/ring/invitations"
import "sync"

func Ring_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.Ring_RoleSetupChan, inviteChannels invitations.Ring_InviteSetupChan)  {
	start_end_int := make(chan int, 1)
	start_end_string := make(chan string, 1)
	end_start_int := make(chan int, 1)
	end_start_string := make(chan string, 1)
	end_start_label := make(chan messages.Ring_Label, 1)
	start_end_label := make(chan messages.Ring_Label, 1)
	start_invite_end := make(chan forward.E_Chan, 1)
	start_invite_end_invitechan := make(chan invitations.Forward_E_InviteChan, 1)
	start_invite_start := make(chan forward.S_Chan, 1)
	start_invite_start_invitechan := make(chan invitations.Forward_S_InviteChan, 1)

	start_chan := ring.Start_Chan{
		String_To_End: start_end_string,
		String_From_End: end_start_string,
		Label_To_End: start_end_label,
		Label_From_End: end_start_label,
		Int_To_End: start_end_int,
		Int_From_End: end_start_int,
	}
	end_chan := ring.End_Chan{
		String_To_Start: end_start_string,
		String_From_Start: start_end_string,
		Label_To_Start: end_start_label,
		Label_From_Start: start_end_label,
		Int_To_Start: end_start_int,
		Int_From_Start: start_end_int,
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