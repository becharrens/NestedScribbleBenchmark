package roles

import "NestedScribbleBenchmark/clientserver/messages"
import "NestedScribbleBenchmark/clientserver/channels/clientserver"
import "NestedScribbleBenchmark/clientserver/invitations"
import "NestedScribbleBenchmark/clientserver/callbacks"
import clientserver_2 "NestedScribbleBenchmark/clientserver/results/clientserver"
import "sync"

func ClientServer_Client(wg *sync.WaitGroup, roleChannels clientserver.Client_Chan, inviteChannels invitations.ClientServer_Client_InviteChan, env callbacks.ClientServer_Client_Env) clientserver_2.Client_Result {
	req := env.Req_To_Server()
	roleChannels.Label_To_Server <- messages.Req
	roleChannels.String_To_Server <- req

	<-roleChannels.Label_From_Server
	resp := <-roleChannels.String_From_Server
	env.Resp_From_Server(resp)

REPEAT:
	for {
		req_2 := env.Req_To_Server_2()
		roleChannels.Label_To_Server <- messages.Req
		roleChannels.String_To_Server <- req_2

		<-roleChannels.Label_From_Server
		resp_2 := <-roleChannels.String_From_Server
		env.Resp_From_Server_2(resp_2)

		continue REPEAT
	}
} 