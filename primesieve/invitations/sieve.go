package invitations

import "NestedScribbleBenchmark/primesieve/channels/sieve_sendnums"
import "NestedScribbleBenchmark/primesieve/channels/sieve"

type Sieve_RoleSetupChan struct {
	M_Chan  chan sieve.M_Chan
	W1_Chan chan sieve.W1_Chan
}

type Sieve_InviteSetupChan struct {
	M_InviteChan  chan Sieve_M_InviteChan
	W1_InviteChan chan Sieve_W1_InviteChan
}

type Sieve_M_InviteChan struct {
	W2_Invite_To_Sieve_M            chan sieve.M_Chan
	W2_Invite_To_Sieve_M_InviteChan chan Sieve_M_InviteChan
}

type Sieve_W1_InviteChan struct {
	Invite_W1_To_Sieve_SendNums_S            chan sieve_sendnums.S_Chan
	Invite_W1_To_Sieve_SendNums_S_InviteChan chan Sieve_SendNums_S_InviteChan
	Invite_W2_To_Sieve_SendNums_R            chan sieve_sendnums.R_Chan
	Invite_W2_To_Sieve_SendNums_R_InviteChan chan Sieve_SendNums_R_InviteChan
}

type Sieve_W2_InviteChan struct {
	Invite_M_To_Sieve_M                      chan sieve.M_Chan
	Invite_M_To_Sieve_M_InviteChan           chan Sieve_M_InviteChan
	Invite_W2_To_Sieve_W1                    chan sieve.W1_Chan
	Invite_W2_To_Sieve_W1_InviteChan         chan Sieve_W1_InviteChan
	W1_Invite_To_Sieve_SendNums_R            chan sieve_sendnums.R_Chan
	W1_Invite_To_Sieve_SendNums_R_InviteChan chan Sieve_SendNums_R_InviteChan
}
