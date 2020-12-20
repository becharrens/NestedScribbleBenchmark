package invitations

import "NestedScribbleBenchmark/primesieve/channels/sieve_sendnums"

type Sieve_SendNums_RoleSetupChan struct {
	R_Chan chan sieve_sendnums.R_Chan
	S_Chan chan sieve_sendnums.S_Chan
}

type Sieve_SendNums_InviteSetupChan struct {
	R_InviteChan chan Sieve_SendNums_R_InviteChan
	S_InviteChan chan Sieve_SendNums_S_InviteChan
}

type Sieve_SendNums_S_InviteChan struct {
}

type Sieve_SendNums_R_InviteChan struct {
}
