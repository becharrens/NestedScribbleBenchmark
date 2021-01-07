package roles

import (
	"NestedScribbleBenchmark/primesieve_base/primesieve4/callbacks"
	"NestedScribbleBenchmark/primesieve_base/primesieve4/channels/sieve"
	"NestedScribbleBenchmark/primesieve_base/primesieve4/channels/sieve_sendnums"
	"NestedScribbleBenchmark/primesieve_base/primesieve4/invitations"
	"NestedScribbleBenchmark/primesieve_base/primesieve4/messages"
	"sync"
)

func Sieve_W2(wg *sync.WaitGroup, roleChannels sieve.W2_Chan, inviteChannels invitations.Sieve_W2_InviteChan, env callbacks.Sieve_W2_Env) {
	defer wg.Done()
	<-roleChannels.Label_From_W1
	int := <-roleChannels.Int_From_W1
	env.FilterPrime_From_W1(int)

	<-roleChannels.Label_From_W1
	// sieve_sendnums_r_chan := <-inviteChannels.W1_Invite_To_Sieve_SendNums_R
	// sieve_sendnums_r_inviteChan := <-inviteChannels.W1_Invite_To_Sieve_SendNums_R_InviteChan
	sieve_sendnums_r_chan := sieve_sendnums.R_Chan{
		Int_From_S:   roleChannels.Int_From_W1,
		Label_From_S: roleChannels.Label_From_W1,
	}
	sieve_sendnums_r_env := env.To_Sieve_SendNums_R_Env()
	sieve_sendnums_r_result := Sieve_SendNums_R(wg, sieve_sendnums_r_chan, sieve_sendnums_r_env)
	env.ResultFrom_Sieve_SendNums_R(sieve_sendnums_r_result)

	w2_choice := env.W2_Choice()
	switch w2_choice {
	case callbacks.Sieve_W2_Prime:
		n := env.Prime_To_M()
		roleChannels.Label_To_M <- messages.Prime
		roleChannels.Int_To_M <- n

		env.Sieve_Setup()
		roleChannels.Label_To_M <- messages.Sieve_M_W2

		sieve_rolechan := invitations.Sieve_RoleSetupChan{
			M_Chan:       inviteChannels.Invite_M_To_Sieve_M,
			M_InviteChan: inviteChannels.Invite_M_To_Sieve_M_InviteChan,
		}

		sieve_w1_chan := Sieve_SendCommChannels(wg, sieve_rolechan)

		// sieve_w1_inviteChan := <-inviteChannels.Invite_W2_To_Sieve_W1_InviteChan
		sieve_w1_env := env.To_Sieve_W1_Env()
		sieve_w1_result := Sieve_W1(wg, sieve_w1_chan, sieve_w1_env)
		env.ResultFrom_Sieve_W1(sieve_w1_result)

		env.Done()
		return
	case callbacks.Sieve_W2_Finish:
		env.Finish_To_M()
		roleChannels.Label_To_M <- messages.Finish
		env.Done()
		return
	default:
		panic("Invalid choice was made")
	}
}
