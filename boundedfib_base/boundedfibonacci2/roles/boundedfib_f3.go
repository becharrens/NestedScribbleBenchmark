package roles

import (
	"NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/messages"
	"sync"
)
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/channels/boundedfib"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/invitations"

func BoundedFib_F3(wg *sync.WaitGroup, roleChannels boundedfib.F3_Chan, inviteChannels invitations.BoundedFib_F3_InviteChan) {
	defer wg.Done()
	<-roleChannels.Label_From_F1
	ubound := <-roleChannels.Int_From_F1
	idx := <-roleChannels.Int_From_F1
	val := <-roleChannels.Int_From_F1

	<-roleChannels.Label_From_F2
	// <-roleChannels.Int_From_F2
	val_2 := <-roleChannels.Int_From_F2

	next_fib := val + val_2
	curr_idx := idx + 2
	if curr_idx == ubound {
		roleChannels.Label_To_Res <- messages.Result
		roleChannels.Int_To_Res <- next_fib

		roleChannels.Label_To_F2 <- messages.End
	} else {
		roleChannels.Label_To_Res <- messages.BoundedFib_Res_F2_F3
		roleChannels.Label_To_F2 <- messages.BoundedFib_Res_F2_F3

		boundedfib_rolechan := invitations.BoundedFib_RoleSetupChan{
			Res_Chan: inviteChannels.Invite_Res_To_BoundedFib_Res,
			F1_Chan:  inviteChannels.Invite_F2_To_BoundedFib_F1,
			F2_Chan:  inviteChannels.Invite_F3_To_BoundedFib_F2,
		}
		boundedfib_invitechan := invitations.BoundedFib_InviteSetupChan{
			Res_InviteChan: inviteChannels.Invite_Res_To_BoundedFib_Res_InviteChan,
			// F1_InviteChan:  inviteChannels.Invite_F2_To_BoundedFib_F1_InviteChan,
			F2_InviteChan: inviteChannels.Invite_F3_To_BoundedFib_F2_InviteChan,
		}
		BoundedFib_SendCommChannels(wg, boundedfib_rolechan, boundedfib_invitechan)

		boundedfib_f2_chan := <-inviteChannels.Invite_F3_To_BoundedFib_F2
		boundedfib_f2_inviteChan := <-inviteChannels.Invite_F3_To_BoundedFib_F2_InviteChan
		BoundedFib_F2(wg, ubound, curr_idx, next_fib, boundedfib_f2_chan, boundedfib_f2_inviteChan)
	}
}
