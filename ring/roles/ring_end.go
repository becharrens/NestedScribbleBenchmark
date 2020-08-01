package roles

import "NestedScribbleBenchmark/ring/channels/ring"
import "NestedScribbleBenchmark/ring/invitations"
import "NestedScribbleBenchmark/ring/callbacks"
import ring_2 "NestedScribbleBenchmark/ring/results/ring"
import "sync"

func Ring_End(wg *sync.WaitGroup, roleChannels ring.End_Chan, inviteChannels invitations.Ring_End_InviteChan, env callbacks.Ring_End_Env) ring_2.End_Result {
	select {
	case forward_e_chan := <-inviteChannels.Start_Invite_To_Forward_E:
		forward_e_inviteChan := <-inviteChannels.Start_Invite_To_Forward_E_InviteChan
		forward_e_env := env.To_Forward_E_Env()
		forward_e_result := Forward_E(wg, forward_e_chan, forward_e_inviteChan, forward_e_env)
		env.ResultFrom_Forward_E(forward_e_result)

		msg_msg := env.Msg_To_Start()
		roleChannels.Start_Msg <- msg_msg

		return env.Done()
	case msg_msg_2 := <-roleChannels.Start_Msg_2:
		env.Msg_From_Start(msg_msg_2)

		msg_msg_3 := env.Msg_To_Start_2()
		roleChannels.Start_Msg_3 <- msg_msg_3

		return env.Done()
	}
}
