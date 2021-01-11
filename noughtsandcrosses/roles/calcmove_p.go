package roles

import "NestedScribbleBenchmark/noughtsandcrosses/channels/calcmove"
import "NestedScribbleBenchmark/noughtsandcrosses/invitations"
import "NestedScribbleBenchmark/noughtsandcrosses/callbacks"
import calcmove_2 "NestedScribbleBenchmark/noughtsandcrosses/results/calcmove"
import "sync"

func CalcMove_P(wg *sync.WaitGroup, roleChannels calcmove.P_Chan, inviteChannels invitations.CalcMove_P_InviteChan, env callbacks.CalcMove_P_Env) calcmove_2.P_Result {
	p_choice := env.P_Choice()
	switch p_choice {
	case callbacks.CalcMove_P_StandardStrategy:
		env.StandardStrategy_Setup()
		
		standardstrategy_rolechan := invitations.StandardStrategy_RoleSetupChan{
			P_Chan: inviteChannels.Invite_P_To_StandardStrategy_P,
		}
		standardstrategy_invitechan := invitations.StandardStrategy_InviteSetupChan{
			P_InviteChan: inviteChannels.Invite_P_To_StandardStrategy_P_InviteChan,
		}
		StandardStrategy_SendCommChannels(wg, standardstrategy_rolechan, standardstrategy_invitechan)

		standardstrategy_p_chan := <-inviteChannels.Invite_P_To_StandardStrategy_P
		standardstrategy_p_inviteChan := <-inviteChannels.Invite_P_To_StandardStrategy_P_InviteChan
		standardstrategy_p_env := env.To_StandardStrategy_P_Env()
		standardstrategy_p_result := StandardStrategy_P(wg, standardstrategy_p_chan, standardstrategy_p_inviteChan, standardstrategy_p_env)
		env.ResultFrom_StandardStrategy_P(standardstrategy_p_result)

		return env.Done()
	case callbacks.CalcMove_P_MinMaxStrategy:
		env.MinMaxStrategy_Setup()
		
		minmaxstrategy_rolechan := invitations.MinMaxStrategy_RoleSetupChan{
			Master_Chan: inviteChannels.Invite_P_To_MinMaxStrategy_Master,
		}
		minmaxstrategy_invitechan := invitations.MinMaxStrategy_InviteSetupChan{
			Master_InviteChan: inviteChannels.Invite_P_To_MinMaxStrategy_Master_InviteChan,
		}
		MinMaxStrategy_SendCommChannels(wg, minmaxstrategy_rolechan, minmaxstrategy_invitechan)

		minmaxstrategy_master_chan := <-inviteChannels.Invite_P_To_MinMaxStrategy_Master
		minmaxstrategy_master_inviteChan := <-inviteChannels.Invite_P_To_MinMaxStrategy_Master_InviteChan
		minmaxstrategy_master_env := env.To_MinMaxStrategy_Master_Env()
		minmaxstrategy_master_result := MinMaxStrategy_Master(wg, minmaxstrategy_master_chan, minmaxstrategy_master_inviteChan, minmaxstrategy_master_env)
		env.ResultFrom_MinMaxStrategy_Master(minmaxstrategy_master_result)

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
} 