package invitations

import "NestedScribbleBenchmark/clientserver/channels/dyntaskgen"
import "NestedScribbleBenchmark/clientserver/channels/clientserver"

type ClientServer_RoleSetupChan struct {
	Client_Chan chan clientserver.Client_Chan
	Server_Chan chan clientserver.Server_Chan
}

type ClientServer_InviteSetupChan struct {
	Client_InviteChan chan ClientServer_Client_InviteChan
	Server_InviteChan chan ClientServer_Server_InviteChan
}

type ClientServer_Client_InviteChan struct {

}

type ClientServer_Server_InviteChan struct {
	Invite_Server_To_DynTaskGen_S chan dyntaskgen.S_Chan
	Invite_Server_To_DynTaskGen_S_InviteChan chan DynTaskGen_S_InviteChan
}