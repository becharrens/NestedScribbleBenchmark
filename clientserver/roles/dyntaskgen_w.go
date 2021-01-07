package roles

import "NestedScribbleBenchmark/clientserver/messages"
import "NestedScribbleBenchmark/clientserver/channels/dyntaskgen"
import "NestedScribbleBenchmark/clientserver/invitations"
import "NestedScribbleBenchmark/clientserver/callbacks"
import "sync"

func DynTaskGen_W(wg *sync.WaitGroup, roleChannels dyntaskgen.W_Chan, inviteChannels invitations.DynTaskGen_W_InviteChan, env callbacks.DynTaskGen_W_Env)  {
	defer wg.Done()
	s_choice := <-roleChannels.Label_From_S
	switch s_choice {
	case messages.Req:
		req := <-roleChannels.String_From_S
		env.Req_From_S(req)

		resp := env.Resp_To_S()
		roleChannels.Label_To_S <- messages.Resp
		roleChannels.String_To_S <- resp

		env.Done()
		return 
	case messages.LastReq:
		req_2 := <-roleChannels.String_From_S
		env.LastReq_From_S(req_2)

		resp_2 := env.Resp_To_S_2()
		roleChannels.Label_To_S <- messages.Resp
		roleChannels.String_To_S <- resp_2

		env.Done()
		return 
	default:
		panic("Invalid choice was made")
	}
} 