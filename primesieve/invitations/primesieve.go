package invitations

import "NestedScribbleBenchmark/primesieve/channels/sieve"
import "NestedScribbleBenchmark/primesieve/channels/primesieve"

type PrimeSieve_RoleSetupChan struct {
	Master_Chan chan primesieve.Master_Chan
	Worker_Chan chan primesieve.Worker_Chan
}

type PrimeSieve_InviteSetupChan struct {
	Master_InviteChan chan PrimeSieve_Master_InviteChan
	Worker_InviteChan chan PrimeSieve_Worker_InviteChan
}

type PrimeSieve_Master_InviteChan struct {
	Worker_Invite_To_Sieve_M            chan sieve.M_Chan
	Worker_Invite_To_Sieve_M_InviteChan chan Sieve_M_InviteChan
}

type PrimeSieve_Worker_InviteChan struct {
	Invite_Master_To_Sieve_M             chan sieve.M_Chan
	Invite_Master_To_Sieve_M_InviteChan  chan Sieve_M_InviteChan
	Invite_Worker_To_Sieve_W1            chan sieve.W1_Chan
	Invite_Worker_To_Sieve_W1_InviteChan chan Sieve_W1_InviteChan
}
