package roles

import (
	"NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/messages"
	"sync"
)
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/channels/boundedfibonacci"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/invitations"

func BoundedFibonacci_Start(wg *sync.WaitGroup, n int, roleChannels boundedfibonacci.Start_Chan, inviteChannels invitations.BoundedFibonacci_Start_InviteChan) int {
	roleChannels.Label_To_F1 <- messages.StartFib1
	roleChannels.Int_To_F1 <- n
	roleChannels.Int_To_F1 <- 1

	roleChannels.Label_To_F2 <- messages.StartFib2
	roleChannels.Int_To_F2 <- n
	roleChannels.Int_To_F2 <- 1

	roleChannels.Label_To_F1 <- messages.BoundedFib_Start_F1_F2
	roleChannels.Label_To_F2 <- messages.BoundedFib_Start_F1_F2
	boundedfib_rolechan := invitations.BoundedFib_RoleSetupChan{
		Res_Chan: inviteChannels.Invite_Start_To_BoundedFib_Res,
		F1_Chan:  inviteChannels.Invite_F1_To_BoundedFib_F1,
		F2_Chan:  inviteChannels.Invite_F2_To_BoundedFib_F2,
	}
	boundedfib_invitechan := invitations.BoundedFib_InviteSetupChan{
		Res_InviteChan: inviteChannels.Invite_Start_To_BoundedFib_Res_InviteChan,
		// F1_InviteChan:  inviteChannels.Invite_F1_To_BoundedFib_F1_InviteChan,
		F2_InviteChan: inviteChannels.Invite_F2_To_BoundedFib_F2_InviteChan,
	}
	BoundedFib_SendCommChannels(wg, boundedfib_rolechan, boundedfib_invitechan)

	boundedfib_res_chan := <-inviteChannels.Invite_Start_To_BoundedFib_Res
	boundedfib_res_inviteChan := <-inviteChannels.Invite_Start_To_BoundedFib_Res_InviteChan
	fib := BoundedFib_Res(wg, boundedfib_res_chan, boundedfib_res_inviteChan)

	return fib
}
