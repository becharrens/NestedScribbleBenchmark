package roles

import "NestedScribbleBenchmark/ring/messages"
import "NestedScribbleBenchmark/ring/channels/forward"
import "NestedScribbleBenchmark/ring/invitations"
import "NestedScribbleBenchmark/ring/callbacks"
import "sync"

func Forward_RingNode(wg *sync.WaitGroup, roleChannels forward.RingNode_Chan, inviteChannels invitations.Forward_RingNode_InviteChan, env callbacks.Forward_RingNode_Env)  {
	defer wg.Done()
	<-roleChannels.Label_From_S
	msg := <-roleChannels.String_From_S
	hops := <-roleChannels.Int_From_S
	env.Msg_From_S(msg, hops)

	ringnode_choice := env.RingNode_Choice()
	switch ringnode_choice {
	case callbacks.Forward_RingNode_Forward:
		env.Forward_Setup()
		
		roleChannels.Label_To_E <- messages.Forward_RingNode_E
		forward_rolechan := invitations.Forward_RoleSetupChan{
			S_Chan: inviteChannels.Invite_RingNode_To_Forward_S,
			E_Chan: inviteChannels.Invite_E_To_Forward_E,
		}
		forward_invitechan := invitations.Forward_InviteSetupChan{
			S_InviteChan: inviteChannels.Invite_RingNode_To_Forward_S_InviteChan,
			E_InviteChan: inviteChannels.Invite_E_To_Forward_E_InviteChan,
		}
		Forward_SendCommChannels(wg, forward_rolechan, forward_invitechan)

		forward_s_chan := <-inviteChannels.Invite_RingNode_To_Forward_S
		forward_s_inviteChan := <-inviteChannels.Invite_RingNode_To_Forward_S_InviteChan
		forward_s_env := env.To_Forward_S_Env()
		forward_s_result := Forward_S(wg, forward_s_chan, forward_s_inviteChan, forward_s_env)
		env.ResultFrom_Forward_S(forward_s_result)

		env.Done()
		return 
	case callbacks.Forward_RingNode_Msg:
		msg_2, hops_2 := env.Msg_To_E()
		roleChannels.Label_To_E <- messages.Msg
		roleChannels.String_To_E <- msg_2
		roleChannels.Int_To_E <- hops_2

		env.Done()
		return 
	default:
		panic("Invalid choice was made")
	}
} 