package roles

import "NestedScribbleBenchmark/clientserver/messages"
import "NestedScribbleBenchmark/clientserver/channels/dyntaskgen"
import "NestedScribbleBenchmark/clientserver/invitations"
import "NestedScribbleBenchmark/clientserver/callbacks"
import dyntaskgen_2 "NestedScribbleBenchmark/clientserver/results/dyntaskgen"
import "sync"

func DynTaskGen_S(wg *sync.WaitGroup, roleChannels dyntaskgen.S_Chan, inviteChannels invitations.DynTaskGen_S_InviteChan, env callbacks.DynTaskGen_S_Env) dyntaskgen_2.S_Result {
	s_choice := env.S_Choice()
	switch s_choice {
	case callbacks.DynTaskGen_S_Req:
		req := env.Req_To_W()
		roleChannels.Label_To_W <- messages.Req
		roleChannels.String_To_W <- req

		env.DynTaskGen_Setup()
		
		dyntaskgen_rolechan := invitations.DynTaskGen_RoleSetupChan{
			S_Chan: inviteChannels.Invite_S_To_DynTaskGen_S,
		}
		dyntaskgen_invitechan := invitations.DynTaskGen_InviteSetupChan{
			S_InviteChan: inviteChannels.Invite_S_To_DynTaskGen_S_InviteChan,
		}
		DynTaskGen_SendCommChannels(wg, dyntaskgen_rolechan, dyntaskgen_invitechan)

		dyntaskgen_s_chan := <-inviteChannels.Invite_S_To_DynTaskGen_S
		dyntaskgen_s_inviteChan := <-inviteChannels.Invite_S_To_DynTaskGen_S_InviteChan
		dyntaskgen_s_env := env.To_DynTaskGen_S_Env()
		dyntaskgen_s_result := DynTaskGen_S(wg, dyntaskgen_s_chan, dyntaskgen_s_inviteChan, dyntaskgen_s_env)
		env.ResultFrom_DynTaskGen_S(dyntaskgen_s_result)

		<-roleChannels.Label_From_W
		resp := <-roleChannels.String_From_W
		env.Resp_From_W(resp)

		return env.Done()
	case callbacks.DynTaskGen_S_LastReq:
		req_2 := env.LastReq_To_W()
		roleChannels.Label_To_W <- messages.LastReq
		roleChannels.String_To_W <- req_2

		<-roleChannels.Label_From_W
		resp_2 := <-roleChannels.String_From_W
		env.Resp_From_W_2(resp_2)

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
} 