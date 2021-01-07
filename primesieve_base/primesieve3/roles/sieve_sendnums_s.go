package roles

import (
	"NestedScribbleBenchmark/primesieve_base/primesieve3/channels/sieve_sendnums"
	"NestedScribbleBenchmark/primesieve_base/primesieve3/invitations"
	"NestedScribbleBenchmark/primesieve_base/primesieve3/messages"
	"sync"
)

func Sieve_SendNums_S(wg *sync.WaitGroup, primeCandidates []int, roleChannels sieve_sendnums.S_Chan, inviteChannels invitations.Sieve_SendNums_S_InviteChan) {
	if len(primeCandidates) > 0 {
		idx := 0
		roleChannels.Label_To_R <- messages.Num
		roleChannels.Int_To_R <- primeCandidates[idx]
		idx++
	SEND:
		for {
			if idx < len(primeCandidates) {
				prime := primeCandidates[idx]
				idx++
				roleChannels.Label_To_R <- messages.Num
				roleChannels.Int_To_R <- prime
				continue SEND
			} else {
				roleChannels.Label_To_R <- messages.End
				return
			}
		}
	} else {
		roleChannels.Label_To_R <- messages.End
		return
	}
}
