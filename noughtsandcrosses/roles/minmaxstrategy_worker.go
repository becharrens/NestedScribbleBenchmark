package roles

import "NestedScribbleBenchmark/noughtsandcrosses/messages"
import "NestedScribbleBenchmark/noughtsandcrosses/channels/minmaxstrategy"
import "NestedScribbleBenchmark/noughtsandcrosses/invitations"
import "NestedScribbleBenchmark/noughtsandcrosses/callbacks"
import "sync"

func MinMaxStrategy_Worker(wg *sync.WaitGroup, roleChannels minmaxstrategy.Worker_Chan, inviteChannels invitations.MinMaxStrategy_Worker_InviteChan, env callbacks.MinMaxStrategy_Worker_Env)  {
	defer wg.Done()
	master_choice := <-roleChannels.Label_From_Master
	switch master_choice {
	case messages.CurrState:
		board := <-roleChannels.Board_From_Master
		currPlayer := <-roleChannels.Player_From_Master
		toMove := <-roleChannels.Player_From_Master
		env.CurrState_From_Master(board, currPlayer, toMove)

		worker_choice := env.Worker_Choice()
		switch worker_choice {
		case callbacks.MinMaxStrategy_Worker_MinMaxStrategy:
			env.MinMaxStrategy_Setup()
			
			minmaxstrategy_rolechan := invitations.MinMaxStrategy_RoleSetupChan{
				Master_Chan: inviteChannels.Invite_Worker_To_MinMaxStrategy_Master,
			}
			minmaxstrategy_invitechan := invitations.MinMaxStrategy_InviteSetupChan{
				Master_InviteChan: inviteChannels.Invite_Worker_To_MinMaxStrategy_Master_InviteChan,
			}
			MinMaxStrategy_SendCommChannels(wg, minmaxstrategy_rolechan, minmaxstrategy_invitechan)

			minmaxstrategy_master_chan := <-inviteChannels.Invite_Worker_To_MinMaxStrategy_Master
			minmaxstrategy_master_inviteChan := <-inviteChannels.Invite_Worker_To_MinMaxStrategy_Master_InviteChan
			minmaxstrategy_master_env := env.To_MinMaxStrategy_Master_Env()
			minmaxstrategy_master_result := MinMaxStrategy_Master(wg, minmaxstrategy_master_chan, minmaxstrategy_master_inviteChan, minmaxstrategy_master_env)
			env.ResultFrom_MinMaxStrategy_Master(minmaxstrategy_master_result)

			score := env.Score_To_Master()
			roleChannels.Label_To_Master <- messages.Score
			roleChannels.Int_To_Master <- score

			env.Done()
			return 
		case callbacks.MinMaxStrategy_Worker_MinMaxStrategy_EvalBoard:
			env.MinMaxStrategy_EvalBoard_Setup()
			
			minmaxstrategy_evalboard_rolechan := invitations.MinMaxStrategy_EvalBoard_RoleSetupChan{
				W_Chan: inviteChannels.Invite_Worker_To_MinMaxStrategy_EvalBoard_W,
			}
			minmaxstrategy_evalboard_invitechan := invitations.MinMaxStrategy_EvalBoard_InviteSetupChan{
				W_InviteChan: inviteChannels.Invite_Worker_To_MinMaxStrategy_EvalBoard_W_InviteChan,
			}
			MinMaxStrategy_EvalBoard_SendCommChannels(wg, minmaxstrategy_evalboard_rolechan, minmaxstrategy_evalboard_invitechan)

			minmaxstrategy_evalboard_w_chan := <-inviteChannels.Invite_Worker_To_MinMaxStrategy_EvalBoard_W
			minmaxstrategy_evalboard_w_inviteChan := <-inviteChannels.Invite_Worker_To_MinMaxStrategy_EvalBoard_W_InviteChan
			minmaxstrategy_evalboard_w_env := env.To_MinMaxStrategy_EvalBoard_W_Env()
			minmaxstrategy_evalboard_w_result := MinMaxStrategy_EvalBoard_W(wg, minmaxstrategy_evalboard_w_chan, minmaxstrategy_evalboard_w_inviteChan, minmaxstrategy_evalboard_w_env)
			env.ResultFrom_MinMaxStrategy_EvalBoard_W(minmaxstrategy_evalboard_w_result)

			score_2 := env.Score_To_Master_2()
			roleChannels.Label_To_Master <- messages.Score
			roleChannels.Int_To_Master <- score_2

			env.Done()
			return 
		default:
			panic("Invalid choice was made")
		}
	case messages.FinalState:
		board_2 := <-roleChannels.Board_From_Master
		currPlayer_2 := <-roleChannels.Player_From_Master
		toMove_2 := <-roleChannels.Player_From_Master
		env.FinalState_From_Master(board_2, currPlayer_2, toMove_2)

		worker_choice_2 := env.Worker_Choice_2()
		switch worker_choice_2 {
		case callbacks.MinMaxStrategy_Worker_MinMaxStrategy_2:
			env.MinMaxStrategy_Setup_2()
			
			minmaxstrategy_rolechan_2 := invitations.MinMaxStrategy_RoleSetupChan{
				Master_Chan: inviteChannels.Invite_Worker_To_MinMaxStrategy_Master,
			}
			minmaxstrategy_invitechan_2 := invitations.MinMaxStrategy_InviteSetupChan{
				Master_InviteChan: inviteChannels.Invite_Worker_To_MinMaxStrategy_Master_InviteChan,
			}
			MinMaxStrategy_SendCommChannels(wg, minmaxstrategy_rolechan_2, minmaxstrategy_invitechan_2)

			minmaxstrategy_master_chan_2 := <-inviteChannels.Invite_Worker_To_MinMaxStrategy_Master
			minmaxstrategy_master_inviteChan_2 := <-inviteChannels.Invite_Worker_To_MinMaxStrategy_Master_InviteChan
			minmaxstrategy_master_env_2 := env.To_MinMaxStrategy_Master_Env_2()
			minmaxstrategy_master_result_2 := MinMaxStrategy_Master(wg, minmaxstrategy_master_chan_2, minmaxstrategy_master_inviteChan_2, minmaxstrategy_master_env_2)
			env.ResultFrom_MinMaxStrategy_Master_2(minmaxstrategy_master_result_2)

			score_3 := env.Score_To_Master_3()
			roleChannels.Label_To_Master <- messages.Score
			roleChannels.Int_To_Master <- score_3

			env.Done()
			return 
		case callbacks.MinMaxStrategy_Worker_MinMaxStrategy_EvalBoard_2:
			env.MinMaxStrategy_EvalBoard_Setup_2()
			
			minmaxstrategy_evalboard_rolechan_2 := invitations.MinMaxStrategy_EvalBoard_RoleSetupChan{
				W_Chan: inviteChannels.Invite_Worker_To_MinMaxStrategy_EvalBoard_W,
			}
			minmaxstrategy_evalboard_invitechan_2 := invitations.MinMaxStrategy_EvalBoard_InviteSetupChan{
				W_InviteChan: inviteChannels.Invite_Worker_To_MinMaxStrategy_EvalBoard_W_InviteChan,
			}
			MinMaxStrategy_EvalBoard_SendCommChannels(wg, minmaxstrategy_evalboard_rolechan_2, minmaxstrategy_evalboard_invitechan_2)

			minmaxstrategy_evalboard_w_chan_2 := <-inviteChannels.Invite_Worker_To_MinMaxStrategy_EvalBoard_W
			minmaxstrategy_evalboard_w_inviteChan_2 := <-inviteChannels.Invite_Worker_To_MinMaxStrategy_EvalBoard_W_InviteChan
			minmaxstrategy_evalboard_w_env_2 := env.To_MinMaxStrategy_EvalBoard_W_Env_2()
			minmaxstrategy_evalboard_w_result_2 := MinMaxStrategy_EvalBoard_W(wg, minmaxstrategy_evalboard_w_chan_2, minmaxstrategy_evalboard_w_inviteChan_2, minmaxstrategy_evalboard_w_env_2)
			env.ResultFrom_MinMaxStrategy_EvalBoard_W_2(minmaxstrategy_evalboard_w_result_2)

			score_4 := env.Score_To_Master_4()
			roleChannels.Label_To_Master <- messages.Score
			roleChannels.Int_To_Master <- score_4

			env.Done()
			return 
		default:
			panic("Invalid choice was made")
		}
	default:
		panic("Invalid choice was made")
	}
} 