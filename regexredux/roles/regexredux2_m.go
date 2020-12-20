package roles

import "NestedScribbleBenchmark/regexredux/messages"
import "NestedScribbleBenchmark/regexredux/channels/regexredux2"
import "NestedScribbleBenchmark/regexredux/invitations"
import "NestedScribbleBenchmark/regexredux/callbacks"
import regexredux2_2 "NestedScribbleBenchmark/regexredux/results/regexredux2"
import "sync"

func RegexRedux2_M(wg *sync.WaitGroup, roleChannels regexredux2.M_Chan, inviteChannels invitations.RegexRedux2_M_InviteChan, env callbacks.RegexRedux2_M_Env) regexredux2_2.M_Result {
	m_choice := env.M_Choice()
	switch m_choice {
	case callbacks.RegexRedux2_M_Task:
		pattern, b := env.Task_To_W()
		roleChannels.Label_To_W <- messages.Task
		roleChannels.String_To_W <- pattern
		roleChannels.ByteArr_To_W <- b

		env.RegexRedux2_Setup()

		regexredux2_rolechan := invitations.RegexRedux2_RoleSetupChan{
			M_Chan: inviteChannels.Invite_M_To_RegexRedux2_M,
		}
		regexredux2_invitechan := invitations.RegexRedux2_InviteSetupChan{
			M_InviteChan: inviteChannels.Invite_M_To_RegexRedux2_M_InviteChan,
		}
		RegexRedux2_SendCommChannels(wg, regexredux2_rolechan, regexredux2_invitechan)

		regexredux2_m_chan := <-inviteChannels.Invite_M_To_RegexRedux2_M
		regexredux2_m_inviteChan := <-inviteChannels.Invite_M_To_RegexRedux2_M_InviteChan
		regexredux2_m_env := env.To_RegexRedux2_M_Env()
		regexredux2_m_result := RegexRedux2_M(wg, regexredux2_m_chan, regexredux2_m_inviteChan, regexredux2_m_env)
		env.ResultFrom_RegexRedux2_M(regexredux2_m_result)

		<-roleChannels.Label_From_W
		nmatches := <-roleChannels.Int_From_W
		env.NumMatches_From_W(nmatches)

		return env.Done()
	case callbacks.RegexRedux2_M_CalcLength:
		b_2 := env.CalcLength_To_W()
		roleChannels.Label_To_W <- messages.CalcLength
		roleChannels.ByteArr_To_W <- b_2

		<-roleChannels.Label_From_W
		len := <-roleChannels.Int_From_W
		env.Length_From_W(len)

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
}
