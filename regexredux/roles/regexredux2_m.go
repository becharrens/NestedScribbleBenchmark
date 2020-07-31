package roles

import "ScribbleBenchmark/regexredux/channels/regexredux2"
import "ScribbleBenchmark/regexredux/invitations"
import "ScribbleBenchmark/regexredux/callbacks"
import regexredux2_2 "ScribbleBenchmark/regexredux/results/regexredux2"
import "sync"

func RegexRedux2_M(wg *sync.WaitGroup, roleChannels regexredux2.M_Chan, inviteChannels invitations.RegexRedux2_M_InviteChan, env callbacks.RegexRedux2_M_Env) regexredux2_2.M_Result {
	m_choice := env.M_Choice()
	switch m_choice {
	case callbacks.RegexRedux2_M_Task:
		task_msg := env.Task_To_W()
		roleChannels.W_Task <- task_msg

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

		nummatches_msg := <-roleChannels.W_NumMatches
		env.NumMatches_From_W(nummatches_msg)

		return env.Done()
	case callbacks.RegexRedux2_M_CalcLength:
		calclength_msg := env.CalcLength_To_W()
		roleChannels.W_CalcLength <- calclength_msg

		length_msg := <-roleChannels.W_Length
		env.Length_From_W(length_msg)

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
} 