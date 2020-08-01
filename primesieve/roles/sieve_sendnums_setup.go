package roles

import "NestedScribbleBenchmark/primesieve/messages/sieve_sendnums"
import "NestedScribbleBenchmark/primesieve/channels/sieve_sendnums_send"
import sieve_sendnums_2 "NestedScribbleBenchmark/primesieve/channels/sieve_sendnums"
import "NestedScribbleBenchmark/primesieve/invitations"
import "sync"

func Sieve_SendNums_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.Sieve_SendNums_RoleSetupChan, inviteChannels invitations.Sieve_SendNums_InviteSetupChan) {
	s_r_end := make(chan sieve_sendnums.End, 1)
	r_invite_s := make(chan sieve_sendnums_send.S_Chan, 1)
	r_invite_s_invitechan := make(chan invitations.Sieve_SendNums_SEND_S_InviteChan, 1)
	r_invite_r := make(chan sieve_sendnums_send.R_Chan, 1)
	r_invite_r_invitechan := make(chan invitations.Sieve_SendNums_SEND_R_InviteChan, 1)
	s_r_num := make(chan sieve_sendnums.Num, 1)

	s_chan := sieve_sendnums_2.S_Chan{
		R_Num: s_r_num,
		R_End: s_r_end,
	}
	r_chan := sieve_sendnums_2.R_Chan{
		S_Num: s_r_num,
		S_End: s_r_end,
	}

	s_inviteChan := invitations.Sieve_SendNums_S_InviteChan{
		R_Invite_To_Sieve_SendNums_SEND_S_InviteChan: r_invite_s_invitechan,
		R_Invite_To_Sieve_SendNums_SEND_S:            r_invite_s,
	}
	r_inviteChan := invitations.Sieve_SendNums_R_InviteChan{
		Invite_S_To_Sieve_SendNums_SEND_S_InviteChan: r_invite_s_invitechan,
		Invite_S_To_Sieve_SendNums_SEND_S:            r_invite_s,
		Invite_R_To_Sieve_SendNums_SEND_R_InviteChan: r_invite_r_invitechan,
		Invite_R_To_Sieve_SendNums_SEND_R:            r_invite_r,
	}

	roleChannels.S_Chan <- s_chan
	roleChannels.R_Chan <- r_chan

	inviteChannels.S_InviteChan <- s_inviteChan
	inviteChannels.R_InviteChan <- r_inviteChan
}
