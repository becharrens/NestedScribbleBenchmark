package roles

import "ScribbleBenchmark/primesieve/messages/sieve"
import "ScribbleBenchmark/primesieve/channels/sieve_sendnums"
import sieve_2 "ScribbleBenchmark/primesieve/channels/sieve"
import "ScribbleBenchmark/primesieve/invitations"
import "ScribbleBenchmark/primesieve/callbacks"
import "sync"

func Sieve_SendCommChannels(wg *sync.WaitGroup, roleChannels invitations.Sieve_RoleSetupChan, inviteChannels invitations.Sieve_InviteSetupChan) {
	w2_m_finish := make(chan sieve.Finish, 1)
	w2_invite_w2 := make(chan sieve_2.W1_Chan, 1)
	w2_invite_w2_invitechan := make(chan invitations.Sieve_W1_InviteChan, 1)
	w2_invite_m := make(chan sieve_2.M_Chan, 1)
	w2_invite_m_invitechan := make(chan invitations.Sieve_M_InviteChan, 1)
	w2_m_prime := make(chan sieve.Prime, 1)
	w1_invite_w2 := make(chan sieve_sendnums.R_Chan, 1)
	w1_invite_w2_invitechan := make(chan invitations.Sieve_SendNums_R_InviteChan, 1)
	w1_invite_w1 := make(chan sieve_sendnums.S_Chan, 1)
	w1_invite_w1_invitechan := make(chan invitations.Sieve_SendNums_S_InviteChan, 1)
	w1_w2_filterprime := make(chan sieve.FilterPrime, 1)

	w2_chan := sieve_2.W2_Chan{
		W1_FilterPrime: w1_w2_filterprime,
		M_Prime:        w2_m_prime,
		M_Finish:       w2_m_finish,
	}
	w1_chan := sieve_2.W1_Chan{
		W2_FilterPrime: w1_w2_filterprime,
	}
	m_chan := sieve_2.M_Chan{
		W2_Prime:  w2_m_prime,
		W2_Finish: w2_m_finish,
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
