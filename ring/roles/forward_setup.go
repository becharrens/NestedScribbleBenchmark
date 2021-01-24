package roles

import "NestedScribbleBenchmark/ring/messages"
import "NestedScribbleBenchmark/ring/channels/forward"
import "NestedScribbleBenchmark/ring/invitations"
import "NestedScribbleBenchmark/ring/callbacks"
import "sync"

func Forward_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.Forward_RoleSetupChan, inviteChannels invitations.Forward_InviteSetupChan) {
	ringnode_e_int := make(chan int, 1)
	ringnode_e_string := make(chan string, 1)
	ringnode_e_label := make(chan messages.Ring_Label, 1)
	ringnode_invite_e := make(chan forward.E_Chan, 1)
	ringnode_invite_e_invitechan := make(chan invitations.Forward_E_InviteChan, 1)
	ringnode_invite_ringnode := make(chan forward.S_Chan, 1)
	ringnode_invite_ringnode_invitechan := make(chan invitations.Forward_S_InviteChan, 1)
	s_ringnode_int := make(chan int, 1)
	s_ringnode_string := make(chan string, 1)
	s_ringnode_label := make(chan messages.Ring_Label, 1)

	s_chan := forward.S_Chan{
		String_To_RingNode: s_ringnode_string,
		Label_To_RingNode:  s_ringnode_label,
		Int_To_RingNode:    s_ringnode_int,
	}
	ringnode_chan := forward.RingNode_Chan{
		String_To_E:   ringnode_e_string,
		String_From_S: s_ringnode_string,
		Label_To_E:    ringnode_e_label,
		Label_From_S:  s_ringnode_label,
		Int_To_E:      ringnode_e_int,
		Int_From_S:    s_ringnode_int,
	}
	e_chan := forward.E_Chan{
		String_From_RingNode: ringnode_e_string,
		Label_From_RingNode:  ringnode_e_label,
		Int_From_RingNode:    ringnode_e_int,
	}

	s_inviteChan := invitations.Forward_S_InviteChan{}
	ringnode_inviteChan := invitations.Forward_RingNode_InviteChan{
		Invite_RingNode_To_Forward_S_InviteChan: ringnode_invite_ringnode_invitechan,
		Invite_RingNode_To_Forward_S:            ringnode_invite_ringnode,
		Invite_E_To_Forward_E_InviteChan:        ringnode_invite_e_invitechan,
		Invite_E_To_Forward_E:                   ringnode_invite_e,
	}
	e_inviteChan := invitations.Forward_E_InviteChan{
		RingNode_Invite_To_Forward_E_InviteChan: ringnode_invite_e_invitechan,
		RingNode_Invite_To_Forward_E:            ringnode_invite_e,
	}

	roleChannels.S_Chan <- s_chan
	roleChannels.E_Chan <- e_chan

	inviteChannels.S_InviteChan <- s_inviteChan
	inviteChannels.E_InviteChan <- e_inviteChan

	wg.Add(1)

	ringnode_env := callbacks.New_Forward_RingNode_State()
	go Forward_RingNode(wg, ringnode_chan, ringnode_inviteChan, ringnode_env)
}
