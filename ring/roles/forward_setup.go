package roles

import "ScribbleBenchmark/ring/messages/forward"
import forward_2 "ScribbleBenchmark/ring/channels/forward"
import "ScribbleBenchmark/ring/invitations"
import "ScribbleBenchmark/ring/callbacks"
import "sync"

func Forward_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.Forward_RoleSetupChan, inviteChannels invitations.Forward_InviteSetupChan)  {
	ringnode_e_msg := make(chan forward.Msg, 1)
	ringnode_invite_e := make(chan forward_2.E_Chan, 1)
	ringnode_invite_e_invitechan := make(chan invitations.Forward_E_InviteChan, 1)
	ringnode_invite_ringnode := make(chan forward_2.S_Chan, 1)
	ringnode_invite_ringnode_invitechan := make(chan invitations.Forward_S_InviteChan, 1)
	s_ringnode_msg := make(chan forward.Msg, 1)

	s_chan := forward_2.S_Chan{
		RingNode_Msg: s_ringnode_msg,
	}
	ringnode_chan := forward_2.RingNode_Chan{
		S_Msg: s_ringnode_msg,
		E_Msg: ringnode_e_msg,
	}
	e_chan := forward_2.E_Chan{
		RingNode_Msg: ringnode_e_msg,
	}

	s_inviteChan := invitations.Forward_S_InviteChan{

	}
	ringnode_inviteChan := invitations.Forward_RingNode_InviteChan{
		Invite_RingNode_To_Forward_S_InviteChan: ringnode_invite_ringnode_invitechan,
		Invite_RingNode_To_Forward_S: ringnode_invite_ringnode,
		Invite_E_To_Forward_E_InviteChan: ringnode_invite_e_invitechan,
		Invite_E_To_Forward_E: ringnode_invite_e,
	}
	e_inviteChan := invitations.Forward_E_InviteChan{
		RingNode_Invite_To_Forward_E_InviteChan: ringnode_invite_e_invitechan,
		RingNode_Invite_To_Forward_E: ringnode_invite_e,
	}

	roleChannels.S_Chan <- s_chan
	roleChannels.E_Chan <- e_chan

	inviteChannels.S_InviteChan <- s_inviteChan
	inviteChannels.E_InviteChan <- e_inviteChan

	wg.Add(1)

	ringnode_env := callbacks.New_Forward_RingNode_State()
	go Forward_RingNode(wg, ringnode_chan, ringnode_inviteChan, ringnode_env)
} 