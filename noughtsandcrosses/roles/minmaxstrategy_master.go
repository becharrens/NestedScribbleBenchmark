package roles

import "NestedScribbleBenchmark/noughtsandcrosses/messages"
import "NestedScribbleBenchmark/noughtsandcrosses/channels/minmaxstrategy"
import "NestedScribbleBenchmark/noughtsandcrosses/invitations"
import "NestedScribbleBenchmark/noughtsandcrosses/callbacks"
import minmaxstrategy_2 "NestedScribbleBenchmark/noughtsandcrosses/results/minmaxstrategy"
import "sync"

func MinMaxStrategy_Master(wg *sync.WaitGroup, roleChannels minmaxstrategy.Master_Chan, inviteChannels invitations.MinMaxStrategy_Master_InviteChan, env callbacks.MinMaxStrategy_Master_Env) minmaxstrategy_2.Master_Result {
	master_choice := env.Master_Choice()
	switch master_choice {
	case callbacks.MinMaxStrategy_Master_CurrState:
		board, currPlayer, toMove := env.CurrState_To_Worker()
		roleChannels.Label_To_Worker <- messages.CurrState
		roleChannels.Board_To_Worker <- board
		roleChannels.Player_To_Worker <- currPlayer
		roleChannels.Player_To_Worker <- toMove

		env.MinMaxStrategy_Setup()
		
		minmaxstrategy_rolechan := invitations.MinMaxStrategy_RoleSetupChan{
			Master_Chan: inviteChannels.Invite_Master_To_MinMaxStrategy_Master,
		}
		minmaxstrategy_invitechan := invitations.MinMaxStrategy_InviteSetupChan{
			Master_InviteChan: inviteChannels.Invite_Master_To_MinMaxStrategy_Master_InviteChan,
		}
		MinMaxStrategy_SendCommChannels(wg, minmaxstrategy_rolechan, minmaxstrategy_invitechan)

		minmaxstrategy_master_chan := <-inviteChannels.Invite_Master_To_MinMaxStrategy_Master
		minmaxstrategy_master_inviteChan := <-inviteChannels.Invite_Master_To_MinMaxStrategy_Master_InviteChan
		minmaxstrategy_master_env := env.To_MinMaxStrategy_Master_Env()
		minmaxstrategy_master_result := MinMaxStrategy_Master(wg, minmaxstrategy_master_chan, minmaxstrategy_master_inviteChan, minmaxstrategy_master_env)
		env.ResultFrom_MinMaxStrategy_Master(minmaxstrategy_master_result)

		<-roleChannels.Label_From_Worker
		score := <-roleChannels.Int_From_Worker
		env.Score_From_Worker(score)

		return env.Done()
	case callbacks.MinMaxStrategy_Master_FinalState:
		board_2, currPlayer_2, toMove_2 := env.FinalState_To_Worker()
		roleChannels.Label_To_Worker <- messages.FinalState
		roleChannels.Board_To_Worker <- board_2
		roleChannels.Player_To_Worker <- currPlayer_2
		roleChannels.Player_To_Worker <- toMove_2

		<-roleChannels.Label_From_Worker
		score_2 := <-roleChannels.Int_From_Worker
		env.Score_From_Worker_2(score_2)

		return env.Done()
	default:
		panic("Invalid choice was made")
	}
} 