package protocol

import (
	"NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/channels/boundedfib"
	"NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/channels/boundedfibonacci"
	"NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/invitations"
	"NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/messages"
	"NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/roles"
	"sync"
)

// type BoundedFibonacci_Env interface {
// 	New_Start_Env() callbacks.BoundedFibonacci_Start_Env
// 	New_F1_Env() callbacks.BoundedFibonacci_F1_Env
// 	New_F2_Env() callbacks.BoundedFibonacci_F2_Env
// 	Start_Result(fib int)
// 	F1_Result()
// 	F2_Result()
// }

func Start_BoundedFibonacci_Start(wg *sync.WaitGroup, n int, res *int, roleChannels boundedfibonacci.Start_Chan, inviteChannels invitations.BoundedFibonacci_Start_InviteChan) {
	defer wg.Done()
	result := roles.BoundedFibonacci_Start(wg, n, roleChannels, inviteChannels)
	*res = result
}

func Start_BoundedFibonacci_F1(wg *sync.WaitGroup, roleChannels boundedfibonacci.F1_Chan, inviteChannels invitations.BoundedFibonacci_F1_InviteChan) {
	defer wg.Done()
	roles.BoundedFibonacci_F1(wg, roleChannels, inviteChannels)
}

func Start_BoundedFibonacci_F2(wg *sync.WaitGroup, roleChannels boundedfibonacci.F2_Chan, inviteChannels invitations.BoundedFibonacci_F2_InviteChan) {
	defer wg.Done()
	roles.BoundedFibonacci_F2(wg, roleChannels, inviteChannels)
}

func BoundedFibonacci(n int) int {
	start_invite_f2 := make(chan boundedfib.F2_Chan, 1)
	start_invite_f2_invitechan := make(chan invitations.BoundedFib_F2_InviteChan, 1)
	start_invite_f1 := make(chan boundedfib.F1_Chan, 1)
	// start_invite_f1_invitechan := make(chan invitations.BoundedFib_F1_InviteChan, 1)
	start_invite_start := make(chan boundedfib.Res_Chan, 1)
	start_invite_start_invitechan := make(chan invitations.BoundedFib_Res_InviteChan, 1)
	start_f2_int := make(chan int, 1)
	start_f2_label := make(chan messages.BoundedFibonacci_Label, 1)
	start_f1_int := make(chan int, 1)
	start_f1_label := make(chan messages.BoundedFibonacci_Label, 1)

	start_chan := boundedfibonacci.Start_Chan{
		Label_To_F2: start_f2_label,
		Label_To_F1: start_f1_label,
		Int_To_F2:   start_f2_int,
		Int_To_F1:   start_f1_int,
	}
	f2_chan := boundedfibonacci.F2_Chan{
		Label_From_Start: start_f2_label,
		Int_From_Start:   start_f2_int,
	}
	f1_chan := boundedfibonacci.F1_Chan{
		Label_From_Start: start_f1_label,
		Int_From_Start:   start_f1_int,
	}

	start_inviteChan := invitations.BoundedFibonacci_Start_InviteChan{
		Invite_Start_To_BoundedFib_Res_InviteChan: start_invite_start_invitechan,
		Invite_Start_To_BoundedFib_Res:            start_invite_start,
		Invite_F2_To_BoundedFib_F2_InviteChan:     start_invite_f2_invitechan,
		Invite_F2_To_BoundedFib_F2:                start_invite_f2,
		// Invite_F1_To_BoundedFib_F1_InviteChan:     start_invite_f1_invitechan,
		Invite_F1_To_BoundedFib_F1: start_invite_f1,
	}
	f2_inviteChan := invitations.BoundedFibonacci_F2_InviteChan{
		Start_Invite_To_BoundedFib_F2_InviteChan: start_invite_f2_invitechan,
		Start_Invite_To_BoundedFib_F2:            start_invite_f2,
	}
	f1_inviteChan := invitations.BoundedFibonacci_F1_InviteChan{
		// Start_Invite_To_BoundedFib_F1_InviteChan: start_invite_f1_invitechan,
		Start_Invite_To_BoundedFib_F1: start_invite_f1,
	}
	var wg sync.WaitGroup
	wg.Add(3)

	var res int
	go Start_BoundedFibonacci_Start(&wg, n, &res, start_chan, start_inviteChan)
	go Start_BoundedFibonacci_F1(&wg, f1_chan, f1_inviteChan)
	go Start_BoundedFibonacci_F2(&wg, f2_chan, f2_inviteChan)

	wg.Wait()
	return res
}
