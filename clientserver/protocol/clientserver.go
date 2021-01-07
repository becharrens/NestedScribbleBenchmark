package protocol

import "NestedScribbleBenchmark/clientserver/messages"
import "NestedScribbleBenchmark/clientserver/channels/dyntaskgen"
import "NestedScribbleBenchmark/clientserver/channels/clientserver"
import "NestedScribbleBenchmark/clientserver/invitations"
import "NestedScribbleBenchmark/clientserver/callbacks"
import clientserver_2 "NestedScribbleBenchmark/clientserver/results/clientserver"
import "NestedScribbleBenchmark/clientserver/roles"
import "sync"

type ClientServer_Env interface {
	New_Client_Env() callbacks.ClientServer_Client_Env
	New_Server_Env() callbacks.ClientServer_Server_Env
	Client_Result(result clientserver_2.Client_Result) 
	Server_Result(result clientserver_2.Server_Result) 
}

func Start_ClientServer_Client(protocolEnv ClientServer_Env, wg *sync.WaitGroup, roleChannels clientserver.Client_Chan, inviteChannels invitations.ClientServer_Client_InviteChan, env callbacks.ClientServer_Client_Env)  {
	defer wg.Done()
	result := roles.ClientServer_Client(wg, roleChannels, inviteChannels, env)
	protocolEnv.Client_Result(result)
} 

func Start_ClientServer_Server(protocolEnv ClientServer_Env, wg *sync.WaitGroup, roleChannels clientserver.Server_Chan, inviteChannels invitations.ClientServer_Server_InviteChan, env callbacks.ClientServer_Server_Env)  {
	defer wg.Done()
	result := roles.ClientServer_Server(wg, roleChannels, inviteChannels, env)
	protocolEnv.Server_Result(result)
} 

func ClientServer(protocolEnv ClientServer_Env)  {
	server_client_string := make(chan string, 1)
	server_client_label := make(chan messages.ClientServer_Label, 1)
	server_invite_server := make(chan dyntaskgen.S_Chan, 1)
	server_invite_server_invitechan := make(chan invitations.DynTaskGen_S_InviteChan, 1)
	client_server_string := make(chan string, 1)
	client_server_label := make(chan messages.ClientServer_Label, 1)

	server_chan := clientserver.Server_Chan{
		String_To_Client: server_client_string,
		String_From_Client: client_server_string,
		Label_To_Client: server_client_label,
		Label_From_Client: client_server_label,
	}
	client_chan := clientserver.Client_Chan{
		String_To_Server: client_server_string,
		String_From_Server: server_client_string,
		Label_To_Server: client_server_label,
		Label_From_Server: server_client_label,
	}

	server_inviteChan := invitations.ClientServer_Server_InviteChan{
		Invite_Server_To_DynTaskGen_S_InviteChan: server_invite_server_invitechan,
		Invite_Server_To_DynTaskGen_S: server_invite_server,
	}
	client_inviteChan := invitations.ClientServer_Client_InviteChan{

	}

	var wg sync.WaitGroup

	wg.Add(2)

	client_env := protocolEnv.New_Client_Env()
	server_env := protocolEnv.New_Server_Env()

	go Start_ClientServer_Client(protocolEnv, &wg, client_chan, client_inviteChan, client_env)
	go Start_ClientServer_Server(protocolEnv, &wg, server_chan, server_inviteChan, server_env)

	wg.Wait()
} 