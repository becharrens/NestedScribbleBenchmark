package roles

import "NestedScribbleBenchmark/noughtsandcrosses/messages"
import "NestedScribbleBenchmark/noughtsandcrosses/channels/noughtsandcrosses"
import "NestedScribbleBenchmark/noughtsandcrosses/invitations"
import "NestedScribbleBenchmark/noughtsandcrosses/callbacks"
import noughtsandcrosses_2 "NestedScribbleBenchmark/noughtsandcrosses/results/noughtsandcrosses"
import "sync"

func NoughtsAndCrosses_P2(wg *sync.WaitGroup, roleChannels noughtsandcrosses.P2_Chan, inviteChannels invitations.NoughtsAndCrosses_P2_InviteChan, env callbacks.NoughtsAndCrosses_P2_Env) noughtsandcrosses_2.P2_Result {
	p1_choice_2 := <-roleChannels.Label_From_P1
	switch p1_choice_2 {
	case messages.Win:
		move := <-roleChannels.Int_From_P1
		env.Win_From_P1(move)

		return env.Done()
	case messages.Draw:
		move_2 := <-roleChannels.Int_From_P1
		env.Draw_From_P1(move_2)

		return env.Done()
	case messages.Move:
		move_3 := <-roleChannels.Int_From_P1
		env.Move_From_P1(move_3)

		env.CalcMove_Setup()
		
		calcmove_rolechan := invitations.CalcMove_RoleSetupChan{
			P_Chan: inviteChannels.Invite_P2_To_CalcMove_P,
		}
		calcmove_invitechan := invitations.CalcMove_InviteSetupChan{
			P_InviteChan: inviteChannels.Invite_P2_To_CalcMove_P_InviteChan,
		}
		CalcMove_SendCommChannels(wg, calcmove_rolechan, calcmove_invitechan)

		calcmove_p_chan := <-inviteChannels.Invite_P2_To_CalcMove_P
		calcmove_p_inviteChan := <-inviteChannels.Invite_P2_To_CalcMove_P_InviteChan
		calcmove_p_env := env.To_CalcMove_P_Env()
		calcmove_p_result := CalcMove_P(wg, calcmove_p_chan, calcmove_p_inviteChan, calcmove_p_env)
		env.ResultFrom_CalcMove_P(calcmove_p_result)

		p2_choice_2 := env.P2_Choice()
		switch p2_choice_2 {
		case callbacks.NoughtsAndCrosses_P2_Win:
			move_4 := env.Win_To_P1()
			roleChannels.Label_To_P1 <- messages.Win
			roleChannels.Int_To_P1 <- move_4

			return env.Done()
		case callbacks.NoughtsAndCrosses_P2_Draw:
			move_5 := env.Draw_To_P1()
			roleChannels.Label_To_P1 <- messages.Draw
			roleChannels.Int_To_P1 <- move_5

			return env.Done()
		case callbacks.NoughtsAndCrosses_P2_Move:
			move_6 := env.Move_To_P1()
			roleChannels.Label_To_P1 <- messages.Move
			roleChannels.Int_To_P1 <- move_6

P1MOVE:
			for {
				p1_choice := <-roleChannels.Label_From_P1
				switch p1_choice {
				case messages.Win:
					move_7 := <-roleChannels.Int_From_P1
					env.Win_From_P1_2(move_7)

					return env.Done()
				case messages.Draw:
					move_8 := <-roleChannels.Int_From_P1
					env.Draw_From_P1_2(move_8)

					return env.Done()
				case messages.Move:
					move_9 := <-roleChannels.Int_From_P1
					env.Move_From_P1_2(move_9)

					env.CalcMove_Setup_2()
					
					calcmove_rolechan_2 := invitations.CalcMove_RoleSetupChan{
						P_Chan: inviteChannels.Invite_P2_To_CalcMove_P,
					}
					calcmove_invitechan_2 := invitations.CalcMove_InviteSetupChan{
						P_InviteChan: inviteChannels.Invite_P2_To_CalcMove_P_InviteChan,
					}
					CalcMove_SendCommChannels(wg, calcmove_rolechan_2, calcmove_invitechan_2)

					calcmove_p_chan_2 := <-inviteChannels.Invite_P2_To_CalcMove_P
					calcmove_p_inviteChan_2 := <-inviteChannels.Invite_P2_To_CalcMove_P_InviteChan
					calcmove_p_env_2 := env.To_CalcMove_P_Env_2()
					calcmove_p_result_2 := CalcMove_P(wg, calcmove_p_chan_2, calcmove_p_inviteChan_2, calcmove_p_env_2)
					env.ResultFrom_CalcMove_P_2(calcmove_p_result_2)

					p2_choice := env.P2_Choice_2()
					switch p2_choice {
					case callbacks.NoughtsAndCrosses_P2_Win_2:
						move_10 := env.Win_To_P1_2()
						roleChannels.Label_To_P1 <- messages.Win
						roleChannels.Int_To_P1 <- move_10

						return env.Done()
					case callbacks.NoughtsAndCrosses_P2_Draw_2:
						move_11 := env.Draw_To_P1_2()
						roleChannels.Label_To_P1 <- messages.Draw
						roleChannels.Int_To_P1 <- move_11

						return env.Done()
					case callbacks.NoughtsAndCrosses_P2_Move_2:
						move_12 := env.Move_To_P1_2()
						roleChannels.Label_To_P1 <- messages.Move
						roleChannels.Int_To_P1 <- move_12

						continue P1MOVE
					default:
						panic("Invalid choice was made")
					}
				default:
					panic("Invalid choice was made")
				}
			}
		default:
			panic("Invalid choice was made")
		}
	default:
		panic("Invalid choice was made")
	}
} 