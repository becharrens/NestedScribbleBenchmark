package roles

import "NestedScribbleBenchmark/ring/channels/forward"
import "NestedScribbleBenchmark/ring/invitations"
import "NestedScribbleBenchmark/ring/callbacks"
import forward_2 "NestedScribbleBenchmark/ring/results/forward"
import "sync"

func Forward_S(wg *sync.WaitGroup, roleChannels forward.S_Chan, inviteChannels invitations.Forward_S_InviteChan, env callbacks.Forward_S_Env) forward_2.S_Result {
	msg_msg := env.Msg_To_RingNode()
	roleChannels.RingNode_Msg <- msg_msg

	return env.Done()
}
