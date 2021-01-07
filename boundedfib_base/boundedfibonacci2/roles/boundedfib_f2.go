package roles

import (
	"NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/messages"
	"sync"
)
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/channels/boundedfib"
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/invitations"

func BoundedFib_F2(wg *sync.WaitGroup, ubound, idx, val int, roleChannels boundedfib.F2_Chan, inviteChannels invitations.BoundedFib_F2_InviteChan) {
	roleChannels.Label_To_F3 <- messages.Fib2
	// roleChannels.Int_To_F3 <- idx
	roleChannels.Int_To_F3 <- val

	f3_choice := <-roleChannels.Label_From_F3
	switch f3_choice {
	case messages.End:
		// break
	case messages.BoundedFib_Res_F2_F3:
		boundedfib_f1_chan := <-inviteChannels.F3_Invite_To_BoundedFib_F1
		BoundedFib_F1(wg, ubound, idx, val, boundedfib_f1_chan)
	default:
		panic("Invalid choice was made")
	}
}
