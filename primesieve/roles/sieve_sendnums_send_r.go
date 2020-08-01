package roles

import "NestedScribbleBenchmark/primesieve/channels/sieve_sendnums_send"
import "NestedScribbleBenchmark/primesieve/invitations"
import "NestedScribbleBenchmark/primesieve/callbacks"
import sieve_sendnums_send_2 "NestedScribbleBenchmark/primesieve/results/sieve_sendnums_send"
import "sync"

func Sieve_SendNums_SEND_R(wg *sync.WaitGroup, roleChannels sieve_sendnums_send.R_Chan, inviteChannels invitations.Sieve_SendNums_SEND_R_InviteChan, env callbacks.Sieve_SendNums_SEND_R_Env) sieve_sendnums_send_2.R_Result {
	select {
	case num_msg := <-roleChannels.S_Num:
		env.Num_From_S(num_msg)

		env.Sieve_SendNums_SEND_Setup()
		sieve_sendnums_send_rolechan := invitations.Sieve_SendNums_SEND_RoleSetupChan{
			R_Chan: inviteChannels.Invite_R_To_Sieve_SendNums_SEND_R,
			S_Chan: inviteChannels.Invite_S_To_Sieve_SendNums_SEND_S,
		}
		sieve_sendnums_send_invitechan := invitations.Sieve_SendNums_SEND_InviteSetupChan{
			R_InviteChan: inviteChannels.Invite_R_To_Sieve_SendNums_SEND_R_InviteChan,
			S_InviteChan: inviteChannels.Invite_S_To_Sieve_SendNums_SEND_S_InviteChan,
		}
		Sieve_SendNums_SEND_SendCommChannels(wg, sieve_sendnums_send_rolechan, sieve_sendnums_send_invitechan)

		sieve_sendnums_send_r_chan := <-inviteChannels.Invite_R_To_Sieve_SendNums_SEND_R
		sieve_sendnums_send_r_inviteChan := <-inviteChannels.Invite_R_To_Sieve_SendNums_SEND_R_InviteChan
		sieve_sendnums_send_r_env := env.To_Sieve_SendNums_SEND_R_Env()
		sieve_sendnums_send_r_result := Sieve_SendNums_SEND_R(wg, sieve_sendnums_send_r_chan, sieve_sendnums_send_r_inviteChan, sieve_sendnums_send_r_env)
		env.ResultFrom_Sieve_SendNums_SEND_R(sieve_sendnums_send_r_result)

		return env.Done()
	case end_msg := <-roleChannels.S_End:
		env.End_From_S(end_msg)

		return env.Done()
	}
}
