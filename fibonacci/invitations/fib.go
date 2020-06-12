package invitations

import "ScribbleBenchmark/fibonacci/channels/fib"

type Fib_RoleSetupChan struct {
	F1_Chan chan fib.F1_Chan
	F2_Chan chan fib.F2_Chan
	Res_Chan chan fib.Res_Chan
}

type Fib_InviteSetupChan struct {
	F1_InviteChan chan Fib_F1_InviteChan
	F2_InviteChan chan Fib_F2_InviteChan
	Res_InviteChan chan Fib_Res_InviteChan
}

type Fib_Res_InviteChan struct {
	F3_Invite_To_Fib_Res chan fib.Res_Chan
	F3_Invite_To_Fib_Res_InviteChan chan Fib_Res_InviteChan
}

type Fib_F1_InviteChan struct {

}

type Fib_F2_InviteChan struct {
	F3_Invite_To_Fib_F1 chan fib.F1_Chan
	F3_Invite_To_Fib_F1_InviteChan chan Fib_F1_InviteChan
}

type Fib_F3_InviteChan struct {
	Invite_F2_To_Fib_F1 chan fib.F1_Chan
	Invite_F2_To_Fib_F1_InviteChan chan Fib_F1_InviteChan
	Invite_F3_To_Fib_F2 chan fib.F2_Chan
	Invite_F3_To_Fib_F2_InviteChan chan Fib_F2_InviteChan
	Invite_Res_To_Fib_Res chan fib.Res_Chan
	Invite_Res_To_Fib_Res_InviteChan chan Fib_Res_InviteChan
}