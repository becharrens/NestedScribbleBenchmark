package roles

import (
	"NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/channels/boundedfibonacci"
	"sync"
)
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/invitations"

func BoundedFibonacci_F2(wg *sync.WaitGroup, roleChannels boundedfibonacci.F2_Chan, inviteChannels invitations.BoundedFibonacci_F2_InviteChan) {
	<-roleChannels.Label_From_Start
	n := <-roleChannels.Int_From_Start
	val := <-roleChannels.Int_From_Start

	<-roleChannels.Label_From_Start
	boundedfib_f2_chan := <-inviteChannels.Start_Invite_To_BoundedFib_F2
	boundedfib_f2_inviteChan := <-inviteChannels.Start_Invite_To_BoundedFib_F2_InviteChan
	BoundedFib_F2(wg, n, 2, val, boundedfib_f2_chan, boundedfib_f2_inviteChan)
}
