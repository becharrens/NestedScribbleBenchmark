package roles

import "ScribbleBenchmark/ring/channels/forward"
import "ScribbleBenchmark/ring/invitations"
import "ScribbleBenchmark/ring/callbacks"
import forward_2 "ScribbleBenchmark/ring/results/forward"
import "sync"

func Forward_E(wg *sync.WaitGroup, roleChannels forward.E_Chan, inviteChannels invitations.Forward_E_InviteChan, env callbacks.Forward_E_Env) forward_2.E_Result {
	select {
		case forward_e_chan := <-inviteChannels.RingNode_Invite_To_Forward_E:
			forward_e_inviteChan := <-inviteChannels.RingNode_Invite_To_Forward_E_InviteChan
			forward_e_env := env.To_Forward_E_Env()
			forward_e_result := Forward_E(wg, forward_e_chan, forward_e_inviteChan, forward_e_env)
			env.ResultFrom_Forward_E(forward_e_result)

			return env.Done()
		case msg_msg := <-roleChannels.RingNode_Msg:
			env.Msg_From_RingNode(msg_msg)

			return env.Done()
	}
} 