package invitations

import "NestedScribbleBenchmark/primesieve/channels/sieve_sendnums_send"
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
	R_Invite_To_Sieve_SendNums_SEND_S            chan sieve_sendnums_send.S_Chan
	R_Invite_To_Sieve_SendNums_SEND_S_InviteChan chan Sieve_SendNums_SEND_S_InviteChan
}

type Sieve_SendNums_R_InviteChan struct {
	Invite_R_To_Sieve_SendNums_SEND_R            chan sieve_sendnums_send.R_Chan
	Invite_R_To_Sieve_SendNums_SEND_R_InviteChan chan Sieve_SendNums_SEND_R_InviteChan
	Invite_S_To_Sieve_SendNums_SEND_S            chan sieve_sendnums_send.S_Chan
	Invite_S_To_Sieve_SendNums_SEND_S_InviteChan chan Sieve_SendNums_SEND_S_InviteChan
}
