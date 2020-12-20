package roles

import "NestedScribbleBenchmark/regexredux/messages"
import "NestedScribbleBenchmark/regexredux/channels/regexredux2"
import "NestedScribbleBenchmark/regexredux/invitations"
import "NestedScribbleBenchmark/regexredux/callbacks"
import "sync"

func RegexRedux2_W(wg *sync.WaitGroup, roleChannels regexredux2.W_Chan, inviteChannels invitations.RegexRedux2_W_InviteChan, env callbacks.RegexRedux2_W_Env) {
	defer wg.Done()
	m_choice := <-roleChannels.Label_From_M
	switch m_choice {
	case messages.Task:
		pattern := <-roleChannels.String_From_M
		b := <-roleChannels.ByteArr_From_M
		env.Task_From_M(pattern, b)

		nmatches := env.NumMatches_To_M()
		roleChannels.Label_To_M <- messages.NumMatches
		roleChannels.Int_To_M <- nmatches

		env.Done()
		return
	case messages.CalcLength:
		b_2 := <-roleChannels.ByteArr_From_M
		env.CalcLength_From_M(b_2)

		len := env.Length_To_M()
		roleChannels.Label_To_M <- messages.Length
		roleChannels.Int_To_M <- len

		env.Done()
		return
	default:
		panic("Invalid choice was made")
	}
}
