package roles

import "NestedScribbleBenchmark/ring/messages"
import "NestedScribbleBenchmark/ring/channels/forward"
import "NestedScribbleBenchmark/ring/invitations"
import "NestedScribbleBenchmark/ring/callbacks"
import forward_2 "NestedScribbleBenchmark/ring/results/forward"
import "sync"

func Forward_E(wg *sync.WaitGroup, roleChannels forward.E_Chan, inviteChannels invitations.Forward_E_InviteChan, env callbacks.Forward_E_Env) forward_2.E_Result {
	ringnode_choice := <-roleChannels.Label_From_RingNode
	switch ringnode_choice {
	case messages.Forward_RingNode_E:
		forward_e_chan := <-inviteChannels.RingNode_Invite_To_Forward_E
		forward_e_inviteChan := <-inviteChannels.RingNode_Invite_To_Forward_E_InviteChan
		forward_e_env := env.To_Forward_E_Env()
		forward_e_result := Forward_E(wg, forward_e_chan, forward_e_inviteChan, forward_e_env)
		env.ResultFrom_Forward_E(forward_e_result)

		return env.Done()
	case messages.Msg:
		msg := <-roleChannels.String_From_RingNode
		hops := <-roleChannels.Int_From_RingNode
		env.Msg_From_RingNode(msg, hops)

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
} 