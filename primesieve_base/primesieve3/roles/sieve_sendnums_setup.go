package roles

import "NestedScribbleBenchmark/primesieve_base/primesieve3/messages"
import "NestedScribbleBenchmark/primesieve_base/primesieve3/channels/sieve_sendnums"
import "NestedScribbleBenchmark/primesieve_base/primesieve3/invitations"
import "sync"

func Sieve_SendNums_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.Sieve_SendNums_RoleSetupChan, inviteChannels invitations.Sieve_SendNums_InviteSetupChan) {
	s_r_int := make(chan int, 1)
	s_r_label := make(chan messages.PrimeSieve_Label, 1)

	s_chan := sieve_sendnums.S_Chan{
		Label_To_R: s_r_label,
		Int_To_R:   s_r_int,
	}
	r_chan := sieve_sendnums.R_Chan{
		Label_From_S: s_r_label,
		Int_From_S:   s_r_int,
	}

	s_inviteChan := invitations.Sieve_SendNums_S_InviteChan{}
	r_inviteChan := invitations.Sieve_SendNums_R_InviteChan{}

	roleChannels.S_Chan <- s_chan
	roleChannels.R_Chan <- r_chan

	inviteChannels.S_InviteChan <- s_inviteChan
	inviteChannels.R_InviteChan <- r_inviteChan
}
