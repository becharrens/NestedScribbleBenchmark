package roles

import (
	"NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/messages"
	"sync"
)
import "NestedScribbleBenchmark/boundedfib_base/boundedfibonacci2/channels/boundedfib"

func BoundedFib_F1(wg *sync.WaitGroup, ubound, idx, val int, roleChannels boundedfib.F1_Chan) {
	roleChannels.Label_To_F3 <- messages.Fib1
	roleChannels.Int_To_F3 <- ubound
	roleChannels.Int_To_F3 <- idx
	roleChannels.Int_To_F3 <- val
}
