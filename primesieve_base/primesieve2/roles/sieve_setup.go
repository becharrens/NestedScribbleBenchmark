package roles

import "NestedScribbleBenchmark/primesieve_base/primesieve2/messages"
import "NestedScribbleBenchmark/primesieve_base/primesieve2/channels/sieve_sendnums"
import "NestedScribbleBenchmark/primesieve_base/primesieve2/channels/sieve"
import "NestedScribbleBenchmark/primesieve_base/primesieve2/invitations"
import "NestedScribbleBenchmark/primesieve_base/primesieve2/callbacks"
import "sync"

func Sieve_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.Sieve_RoleSetupChan, inviteChannels invitations.Sieve_InviteSetupChan) {
	w2_invite_w2 := make(chan sieve.W1_Chan, 1)
	w2_invite_w2_invitechan := make(chan invitations.Sieve_W1_InviteChan, 1)
	w2_invite_m := make(chan sieve.M_Chan, 1)
	w2_invite_m_invitechan := make(chan invitations.Sieve_M_InviteChan, 1)
	w2_m_int := make(chan int, 1)
	w2_m_label := make(chan messages.PrimeSieve_Label, 1)
	w1_invite_w2 := make(chan sieve_sendnums.R_Chan, 1)
	w1_invite_w2_invitechan := make(chan invitations.Sieve_SendNums_R_InviteChan, 1)
	w1_invite_w1 := make(chan sieve_sendnums.S_Chan, 1)
	w1_invite_w1_invitechan := make(chan invitations.Sieve_SendNums_S_InviteChan, 1)
	w1_w2_int := make(chan int, 1)
	w1_w2_label := make(chan messages.PrimeSieve_Label, 1)

	w2_chan := sieve.W2_Chan{
		Label_To_M:    w2_m_label,
		Label_From_W1: w1_w2_label,
		Int_To_M:      w2_m_int,
		Int_From_W1:   w1_w2_int,
	}
	w1_chan := sieve.W1_Chan{
		Label_To_W2: w1_w2_label,
		Int_To_W2:   w1_w2_int,
	}
	m_chan := sieve.M_Chan{
		Label_From_W2: w2_m_label,
		Int_From_W2:   w2_m_int,
	}

	w2_inviteChan := invitations.Sieve_W2_InviteChan{
		W1_Invite_To_Sieve_SendNums_R_InviteChan: w1_invite_w2_invitechan,
		W1_Invite_To_Sieve_SendNums_R:            w1_invite_w2,
		Invite_W2_To_Sieve_W1_InviteChan:         w2_invite_w2_invitechan,
		Invite_W2_To_Sieve_W1:                    w2_invite_w2,
		Invite_M_To_Sieve_M_InviteChan:           w2_invite_m_invitechan,
		Invite_M_To_Sieve_M:                      w2_invite_m,
	}
	w1_inviteChan := invitations.Sieve_W1_InviteChan{
		Invite_W2_To_Sieve_SendNums_R_InviteChan: w1_invite_w2_invitechan,
		Invite_W2_To_Sieve_SendNums_R:            w1_invite_w2,
		Invite_W1_To_Sieve_SendNums_S_InviteChan: w1_invite_w1_invitechan,
		Invite_W1_To_Sieve_SendNums_S:            w1_invite_w1,
	}
	m_inviteChan := invitations.Sieve_M_InviteChan{
		W2_Invite_To_Sieve_M_InviteChan: w2_invite_m_invitechan,
		W2_Invite_To_Sieve_M:            w2_invite_m,
	}

	roleChannels.M_Chan <- m_chan
	roleChannels.W1_Chan <- w1_chan

	inviteChannels.M_InviteChan <- m_inviteChan
	inviteChannels.W1_InviteChan <- w1_inviteChan

	wg.Add(1)

	w2_env := callbacks.New_Sieve_W2_State()
	go Sieve_W2(wg, w2_chan, w2_inviteChan, w2_env)
}
