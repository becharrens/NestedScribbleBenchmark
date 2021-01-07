package roles

import (
	"NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/channels/boundedfibonacci"
	"sync"
)
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/invitations"

func BoundedFibonacci_F1(wg *sync.WaitGroup, roleChannels boundedfibonacci.F1_Chan, inviteChannels invitations.BoundedFibonacci_F1_InviteChan) {
	<-roleChannels.Label_From_Start
	n := <-roleChannels.Int_From_Start
	val := <-roleChannels.Int_From_Start

	<-roleChannels.Label_From_Start
	boundedfib_f1_chan := <-inviteChannels.Start_Invite_To_BoundedFib_F1
	BoundedFib_F1(wg, n, 1, val, boundedfib_f1_chan)
}
