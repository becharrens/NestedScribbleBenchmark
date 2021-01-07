package roles

import "NestedScribbleBenchmark/clientserver/messages"
import "NestedScribbleBenchmark/clientserver/channels/dyntaskgen"
import "NestedScribbleBenchmark/clientserver/channels/clientserver"
import "NestedScribbleBenchmark/clientserver/invitations"
import "sync"

func ClientServer_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.ClientServer_RoleSetupChan, inviteChannels invitations.ClientServer_InviteSetupChan)  {
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

	roleChannels.Client_Chan <- client_chan
	roleChannels.Server_Chan <- server_chan

	inviteChannels.Client_InviteChan <- client_inviteChan
	inviteChannels.Server_InviteChan <- server_inviteChan
} 