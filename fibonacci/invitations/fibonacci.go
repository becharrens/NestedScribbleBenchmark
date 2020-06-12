package invitations

import "ScribbleBenchmark/fibonacci/channels/fibonacci"
import "ScribbleBenchmark/fibonacci/channels/fib"

type Fibonacci_RoleSetupChan struct {
	F1_Chan chan fibonacci.F1_Chan
	F2_Chan chan fibonacci.F2_Chan
	Start_Chan chan fibonacci.Start_Chan
}

type Fibonacci_InviteSetupChan struct {
	F1_InviteChan chan Fibonacci_F1_InviteChan
	F2_InviteChan chan Fibonacci_F2_InviteChan
	Start_InviteChan chan Fibonacci_Start_InviteChan
}

type Fibonacci_Start_InviteChan struct {
	Invite_F1_To_Fib_F1 chan fib.F1_Chan
	Invite_F1_To_Fib_F1_InviteChan chan Fib_F1_InviteChan
	Invite_F2_To_Fib_F2 chan fib.F2_Chan
	Invite_F2_To_Fib_F2_InviteChan chan Fib_F2_InviteChan
	Invite_Start_To_Fib_Res chan fib.Res_Chan
	Invite_Start_To_Fib_Res_InviteChan chan Fib_Res_InviteChan
}

type Fibonacci_F1_InviteChan struct {
	Start_Invite_To_Fib_F1 chan fib.F1_Chan
	Start_Invite_To_Fib_F1_InviteChan chan Fib_F1_InviteChan
}

type Fibonacci_F2_InviteChan struct {
	Start_Invite_To_Fib_F2 chan fib.F2_Chan
	Start_Invite_To_Fib_F2_InviteChan chan Fib_F2_InviteChan
}