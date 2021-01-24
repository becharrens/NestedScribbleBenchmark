package roles

import "NestedScribbleBenchmark/ring/messages"
import "NestedScribbleBenchmark/ring/channels/ring"
import "NestedScribbleBenchmark/ring/invitations"
import "NestedScribbleBenchmark/ring/callbacks"
import ring_2 "NestedScribbleBenchmark/ring/results/ring"
import "sync"

func Ring_End(wg *sync.WaitGroup, roleChannels ring.End_Chan, inviteChannels invitations.Ring_End_InviteChan, env callbacks.Ring_End_Env) ring_2.End_Result {
	start_choice := <-roleChannels.Label_From_Start
	switch start_choice {
	case messages.Forward_Start_End:
		forward_e_chan := <-inviteChannels.Start_Invite_To_Forward_E
		forward_e_inviteChan := <-inviteChannels.Start_Invite_To_Forward_E_InviteChan
		forward_e_env := env.To_Forward_E_Env()
		forward_e_result := Forward_E(wg, forward_e_chan, forward_e_inviteChan, forward_e_env)
		env.ResultFrom_Forward_E(forward_e_result)

		msg, hops := env.Msg_To_Start()
		roleChannels.Label_To_Start <- messages.Msg
		roleChannels.String_To_Start <- msg
		roleChannels.Int_To_Start <- hops

		return env.Done()
	case messages.Msg:
		msg_2 := <-roleChannels.String_From_Start
		hops_2 := <-roleChannels.Int_From_Start
		env.Msg_From_Start(msg_2, hops_2)

		msg_3, hops_3 := env.Msg_To_Start_2()
		roleChannels.Label_To_Start <- messages.Msg
		roleChannels.String_To_Start <- msg_3
		roleChannels.Int_To_Start <- hops_3

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
}
