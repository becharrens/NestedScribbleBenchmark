package roles

import "NestedScribbleBenchmark/clientserver/messages"
import "NestedScribbleBenchmark/clientserver/channels/clientserver"
import "NestedScribbleBenchmark/clientserver/invitations"
import "NestedScribbleBenchmark/clientserver/callbacks"
import clientserver_2 "NestedScribbleBenchmark/clientserver/results/clientserver"
import "sync"

func ClientServer_Server(wg *sync.WaitGroup, roleChannels clientserver.Server_Chan, inviteChannels invitations.ClientServer_Server_InviteChan, env callbacks.ClientServer_Server_Env) clientserver_2.Server_Result {
	<-roleChannels.Label_From_Client
	req := <-roleChannels.String_From_Client
	env.Req_From_Client(req)

	env.DynTaskGen_Setup()
	
	dyntaskgen_rolechan := invitations.DynTaskGen_RoleSetupChan{
		S_Chan: inviteChannels.Invite_Server_To_DynTaskGen_S,
	}
	dyntaskgen_invitechan := invitations.DynTaskGen_InviteSetupChan{
		S_InviteChan: inviteChannels.Invite_Server_To_DynTaskGen_S_InviteChan,
	}
	DynTaskGen_SendCommChannels(wg, dyntaskgen_rolechan, dyntaskgen_invitechan)

	dyntaskgen_s_chan := <-inviteChannels.Invite_Server_To_DynTaskGen_S
	dyntaskgen_s_inviteChan := <-inviteChannels.Invite_Server_To_DynTaskGen_S_InviteChan
	dyntaskgen_s_env := env.To_DynTaskGen_S_Env()
	dyntaskgen_s_result := DynTaskGen_S(wg, dyntaskgen_s_chan, dyntaskgen_s_inviteChan, dyntaskgen_s_env)
	env.ResultFrom_DynTaskGen_S(dyntaskgen_s_result)

	resp := env.Resp_To_Client()
	roleChannels.Label_To_Client <- messages.Resp
	roleChannels.String_To_Client <- resp

REPEAT:
	for {
		<-roleChannels.Label_From_Client
		req_2 := <-roleChannels.String_From_Client
		env.Req_From_Client_2(req_2)

		env.DynTaskGen_Setup_2()
		
		dyntaskgen_rolechan_2 := invitations.DynTaskGen_RoleSetupChan{
			S_Chan: inviteChannels.Invite_Server_To_DynTaskGen_S,
		}
		dyntaskgen_invitechan_2 := invitations.DynTaskGen_InviteSetupChan{
			S_InviteChan: inviteChannels.Invite_Server_To_DynTaskGen_S_InviteChan,
		}
		DynTaskGen_SendCommChannels(wg, dyntaskgen_rolechan_2, dyntaskgen_invitechan_2)

		dyntaskgen_s_chan_2 := <-inviteChannels.Invite_Server_To_DynTaskGen_S
		dyntaskgen_s_inviteChan_2 := <-inviteChannels.Invite_Server_To_DynTaskGen_S_InviteChan
		dyntaskgen_s_env_2 := env.To_DynTaskGen_S_Env_2()
		dyntaskgen_s_result_2 := DynTaskGen_S(wg, dyntaskgen_s_chan_2, dyntaskgen_s_inviteChan_2, dyntaskgen_s_env_2)
		env.ResultFrom_DynTaskGen_S_2(dyntaskgen_s_result_2)

		resp_2 := env.Resp_To_Client_2()
		roleChannels.Label_To_Client <- messages.Resp
		roleChannels.String_To_Client <- resp_2

		continue REPEAT
	}
} 