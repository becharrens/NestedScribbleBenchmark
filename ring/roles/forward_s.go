package roles

import "ScribbleBenchmark/ring/channels/forward"
import "ScribbleBenchmark/ring/invitations"
import "ScribbleBenchmark/ring/callbacks"
import forward_2 "ScribbleBenchmark/ring/results/forward"
import "sync"

func Forward_S(wg *sync.WaitGroup, roleChannels forward.S_Chan, inviteChannels invitations.Forward_S_InviteChan, env callbacks.Forward_S_Env) forward_2.S_Result {
	msg_msg := env.Msg_To_RingNode()
	roleChannels.RingNode_Msg <- msg_msg

	return env.Done()
} 