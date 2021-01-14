package roles

import "NestedScribbleBenchmark/ring/messages"
import "NestedScribbleBenchmark/ring/channels/ring"
import "NestedScribbleBenchmark/ring/invitations"
import "NestedScribbleBenchmark/ring/callbacks"
import ring_2 "NestedScribbleBenchmark/ring/results/ring"
import "sync"

func Ring_Start(wg *sync.WaitGroup, roleChannels ring.Start_Chan, inviteChannels invitations.Ring_Start_InviteChan, env callbacks.Ring_Start_Env) ring_2.Start_Result {
	start_choice := env.Start_Choice()
	switch start_choice {
	case callbacks.Ring_Start_Forward:
		env.Forward_Setup()
		
		roleChannels.Label_To_End <- messages.Forward_Start_End
		forward_rolechan := invitations.Forward_RoleSetupChan{
			S_Chan: inviteChannels.Invite_Start_To_Forward_S,
			E_Chan: inviteChannels.Invite_End_To_Forward_E,
		}
		forward_invitechan := invitations.Forward_InviteSetupChan{
			S_InviteChan: inviteChannels.Invite_Start_To_Forward_S_InviteChan,
			E_InviteChan: inviteChannels.Invite_End_To_Forward_E_InviteChan,
		}
		Forward_SendCommChannels(wg, forward_rolechan, forward_invitechan)

		forward_s_chan := <-inviteChannels.Invite_Start_To_Forward_S
		forward_s_inviteChan := <-inviteChannels.Invite_Start_To_Forward_S_InviteChan
		forward_s_env := env.To_Forward_S_Env()
		forward_s_result := Forward_S(wg, forward_s_chan, forward_s_inviteChan, forward_s_env)
		env.ResultFrom_Forward_S(forward_s_result)

		<-roleChannels.Label_From_End
		msg := <-roleChannels.String_From_End
		hops := <-roleChannels.Int_From_End
		env.Msg_From_End(msg, hops)

		return env.Done()
	case callbacks.Ring_Start_Msg:
		msg_2, hops_2 := env.Msg_To_End()
		roleChannels.Label_To_End <- messages.Msg
		roleChannels.String_To_End <- msg_2
		roleChannels.Int_To_End <- hops_2

		<-roleChannels.Label_From_End
		msg_3 := <-roleChannels.String_From_End
		hops_3 := <-roleChannels.Int_From_End
		env.Msg_From_End_2(msg_3, hops_3)

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
} 