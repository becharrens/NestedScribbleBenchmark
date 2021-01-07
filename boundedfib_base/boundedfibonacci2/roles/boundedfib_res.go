package roles

import (
	"NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/messages"
	"sync"
)
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/channels/boundedfib"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/invitations"

func BoundedFib_Res(wg *sync.WaitGroup, roleChannels boundedfib.Res_Chan, inviteChannels invitations.BoundedFib_Res_InviteChan) int {
	f3_choice := <-roleChannels.Label_From_F3
	switch f3_choice {
	case messages.BoundedFib_Res_F2_F3:
		boundedfib_res_chan := <-inviteChannels.F3_Invite_To_BoundedFib_Res
		boundedfib_res_inviteChan := <-inviteChannels.F3_Invite_To_BoundedFib_Res_InviteChan
		fib := BoundedFib_Res(wg, boundedfib_res_chan, boundedfib_res_inviteChan)

		return fib
	case messages.Result:
		fib := <-roleChannels.Int_From_F3
		return fib
	default:
		panic("Invalid choice was made")
	}
}
