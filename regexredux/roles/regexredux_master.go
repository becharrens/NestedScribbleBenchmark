package roles

import "NestedScribbleBenchmark/regexredux/channels/regexredux"
import "NestedScribbleBenchmark/regexredux/invitations"
import "NestedScribbleBenchmark/regexredux/callbacks"
import regexredux_2 "NestedScribbleBenchmark/regexredux/results/regexredux"
import "sync"

func RegexRedux_Master(wg *sync.WaitGroup, roleChannels regexredux.Master_Chan, inviteChannels invitations.RegexRedux_Master_InviteChan, env callbacks.RegexRedux_Master_Env) regexredux_2.Master_Result {
	env.RegexRedux2_Setup()

	regexredux2_rolechan := invitations.RegexRedux2_RoleSetupChan{
		M_Chan: inviteChannels.Invite_Master_To_RegexRedux2_M,
	}
	regexredux2_invitechan := invitations.RegexRedux2_InviteSetupChan{
		M_InviteChan: inviteChannels.Invite_Master_To_RegexRedux2_M_InviteChan,
	}
	RegexRedux2_SendCommChannels(wg, regexredux2_rolechan, regexredux2_invitechan)

	regexredux2_m_chan := <-inviteChannels.Invite_Master_To_RegexRedux2_M
	regexredux2_m_inviteChan := <-inviteChannels.Invite_Master_To_RegexRedux2_M_InviteChan
	regexredux2_m_env := env.To_RegexRedux2_M_Env()
	regexredux2_m_result := RegexRedux2_M(wg, regexredux2_m_chan, regexredux2_m_inviteChan, regexredux2_m_env)
	env.ResultFrom_RegexRedux2_M(regexredux2_m_result)

	return env.Done()
}
