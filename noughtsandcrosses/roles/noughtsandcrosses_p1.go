package roles

import "NestedScribbleBenchmark/noughtsandcrosses/messages"
import "NestedScribbleBenchmark/noughtsandcrosses/channels/noughtsandcrosses"
import "NestedScribbleBenchmark/noughtsandcrosses/invitations"
import "NestedScribbleBenchmark/noughtsandcrosses/callbacks"
import noughtsandcrosses_2 "NestedScribbleBenchmark/noughtsandcrosses/results/noughtsandcrosses"
import "sync"

func NoughtsAndCrosses_P1(wg *sync.WaitGroup, roleChannels noughtsandcrosses.P1_Chan, inviteChannels invitations.NoughtsAndCrosses_P1_InviteChan, env callbacks.NoughtsAndCrosses_P1_Env) noughtsandcrosses_2.P1_Result {
	env.CalcMove_Setup()
	
	calcmove_rolechan := invitations.CalcMove_RoleSetupChan{
		P_Chan: inviteChannels.Invite_P1_To_CalcMove_P,
	}
	calcmove_invitechan := invitations.CalcMove_InviteSetupChan{
		P_InviteChan: inviteChannels.Invite_P1_To_CalcMove_P_InviteChan,
	}
	CalcMove_SendCommChannels(wg, calcmove_rolechan, calcmove_invitechan)

	calcmove_p_chan := <-inviteChannels.Invite_P1_To_CalcMove_P
	calcmove_p_inviteChan := <-inviteChannels.Invite_P1_To_CalcMove_P_InviteChan
	calcmove_p_env := env.To_CalcMove_P_Env()
	calcmove_p_result := CalcMove_P(wg, calcmove_p_chan, calcmove_p_inviteChan, calcmove_p_env)
	env.ResultFrom_CalcMove_P(calcmove_p_result)

	p1_choice_2 := env.P1_Choice()
	switch p1_choice_2 {
	case callbacks.NoughtsAndCrosses_P1_Win:
		move := env.Win_To_P2()
		roleChannels.Label_To_P2 <- messages.Win
		roleChannels.Int_To_P2 <- move

		return env.Done()
	case callbacks.NoughtsAndCrosses_P1_Draw:
		move_2 := env.Draw_To_P2()
		roleChannels.Label_To_P2 <- messages.Draw
		roleChannels.Int_To_P2 <- move_2

		return env.Done()
	case callbacks.NoughtsAndCrosses_P1_Move:
		move_3 := env.Move_To_P2()
		roleChannels.Label_To_P2 <- messages.Move
		roleChannels.Int_To_P2 <- move_3

		p2_choice_2 := <-roleChannels.Label_From_P2
		switch p2_choice_2 {
		case messages.Win:
			move_4 := <-roleChannels.Int_From_P2
			env.Win_From_P2(move_4)

			return env.Done()
		case messages.Draw:
			move_5 := <-roleChannels.Int_From_P2
			env.Draw_From_P2(move_5)

			return env.Done()
		case messages.Move:
			move_6 := <-roleChannels.Int_From_P2
			env.Move_From_P2(move_6)

P1MOVE:
			for {
				env.CalcMove_Setup_2()
				
				calcmove_rolechan_2 := invitations.CalcMove_RoleSetupChan{
					P_Chan: inviteChannels.Invite_P1_To_CalcMove_P,
				}
				calcmove_invitechan_2 := invitations.CalcMove_InviteSetupChan{
					P_InviteChan: inviteChannels.Invite_P1_To_CalcMove_P_InviteChan,
				}
				CalcMove_SendCommChannels(wg, calcmove_rolechan_2, calcmove_invitechan_2)

				calcmove_p_chan_2 := <-inviteChannels.Invite_P1_To_CalcMove_P
				calcmove_p_inviteChan_2 := <-inviteChannels.Invite_P1_To_CalcMove_P_InviteChan
				calcmove_p_env_2 := env.To_CalcMove_P_Env_2()
				calcmove_p_result_2 := CalcMove_P(wg, calcmove_p_chan_2, calcmove_p_inviteChan_2, calcmove_p_env_2)
				env.ResultFrom_CalcMove_P_2(calcmove_p_result_2)

				p1_choice := env.P1_Choice_2()
				switch p1_choice {
				case callbacks.NoughtsAndCrosses_P1_Win_2:
					move_7 := env.Win_To_P2_2()
					roleChannels.Label_To_P2 <- messages.Win
					roleChannels.Int_To_P2 <- move_7

					return env.Done()
				case callbacks.NoughtsAndCrosses_P1_Draw_2:
					move_8 := env.Draw_To_P2_2()
					roleChannels.Label_To_P2 <- messages.Draw
					roleChannels.Int_To_P2 <- move_8

					return env.Done()
				case callbacks.NoughtsAndCrosses_P1_Move_2:
					move_9 := env.Move_To_P2_2()
					roleChannels.Label_To_P2 <- messages.Move
					roleChannels.Int_To_P2 <- move_9

					p2_choice := <-roleChannels.Label_From_P2
					switch p2_choice {
					case messages.Win:
						move_10 := <-roleChannels.Int_From_P2
						env.Win_From_P2_2(move_10)

						return env.Done()
					case messages.Draw:
						move_11 := <-roleChannels.Int_From_P2
						env.Draw_From_P2_2(move_11)

						return env.Done()
					case messages.Move:
						move_12 := <-roleChannels.Int_From_P2
						env.Move_From_P2_2(move_12)

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