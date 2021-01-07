package boundedfib_base

import (
	"sync"
)

// func startFstSender(wg *sync.WaitGroup, fib, idx, ubound int, sendChan chan int) {
// 	defer wg.Done()
// 	fstSender(fib, idx, ubound, sendChan)
// }
//
// func fstSender(fib, idx, ubound int, sendChan chan int) {
// 	sendChan <- ubound
// 	sendChan <- idx
// 	sendChan <- fib
//
// }
//
// func startSndSender(wg *sync.WaitGroup, fib, idx, ubound int, sendChan1, sendChan2 chan int, stopChan chan bool) {
// 	defer wg.Done()
// 	sndSender(fib, idx, ubound, sendChan1, sendChan2, stopChan)
// }
//
// func sndSender(fib, idx, ubound int, sendChan1, sendChan2 chan int, stopChan chan bool) {
// 	// sendChan1 <- idx
// 	sendChan1 <- fib
// 	stop := <-stopChan
// 	if !stop {
// 		fstSender(fib, idx, ubound, sendChan2)
// 	}
// }
//
// func fibCalc(wg *sync.WaitGroup, fib1Chan, fib2Chan, sendChan1, sendChan2, nextFib1Chan, resChan chan int,
// 	prevStopChan, stopChan chan bool) {
// 	defer wg.Done()
// 	ubound := <-fib1Chan
// 	idx1 := <-fib1Chan
// 	fib1 := <-fib1Chan
//
// 	// _ = <-fib2Chan
// 	fib2 := <-fib2Chan
//
// 	fib := fib1 + fib2
// 	idx := idx1 + 2
// 	returnFib := idx >= ubound
// 	prevStopChan <- returnFib // If worker is returning a result, prev worker should stop
// 	if returnFib {
// 		resChan <- fib
// 	} else {
// 		newFibChan1 := make(chan int, 1)
// 		newFibChan2 := make(chan int, 1)
// 		nextStopChan := make(chan bool, 1)
// 		wg.Add(1)
// 		go fibCalc(wg, nextFib1Chan, sendChan1, newFibChan1, newFibChan2, sendChan2, resChan, stopChan, nextStopChan)
// 		sndSender(fib, idx, ubound, sendChan1, sendChan2, stopChan)
// 	}
// }

// Channels and Invite Channels
type FibChannels struct {
	F2InviteChan    chan F2InviteChannels
	F2ChannelsChan  chan F2Channels
	F1ChannelsChan  chan F1Channels
	ResChannelsChan chan ResChannels
	ResInviteChan   chan ResInviteChannels
}

type StartF1Channels struct {
	IntFromStart chan int
}

type StartF1InviteChannels struct {
	ChannelsChan chan F1Channels
}

type StartF2Channels struct {
	IntFromStart chan int
}

type StartF2InviteChannels struct {
	ChannelsChan chan F2Channels
	InviteChan   chan F2InviteChannels
}

type StartChannels struct {
	IntToStartF1 chan int
	IntToStartF2 chan int
}

type StartInviteChannels struct {
	StartToResInviteChan   chan ResInviteChannels
	StartToResChannelsChan chan ResChannels
	F1ToF1ChannelsChan     chan F1Channels
	F2ToF2ChannelsChan     chan F2Channels
	F2ToF2InviteChan       chan F2InviteChannels
}

type F1Channels struct {
	IntToF3 chan int
}

type F2Channels struct {
	IntToF3    chan int
	BoolFromF3 chan bool
}

type F2InviteChannels struct {
	ChannelsChan chan F1Channels
}

type ResChannels struct {
	IntFromF3  chan int
	BoolFromF3 chan bool
}

type ResInviteChannels struct {
	InviteChan   chan ResInviteChannels
	ChannelsChan chan ResChannels
}

type F3Channels struct {
	IntFromF1 chan int
	IntFromF2 chan int
	BoolToF2  chan bool
	BoolToRes chan bool
	IntToRes  chan int
}

type F3InviteChannels struct {
	F3ToF2InviteChan     chan F2InviteChannels
	F3ToF2ChannelsChan   chan F2Channels
	F2ToF1ChannelsChan   chan F1Channels
	ResToResChannelsChan chan ResChannels
	ResToResInviteChan   chan ResInviteChannels
}

// Role functions

func StartRes(wg *sync.WaitGroup, res *int, n int, channels StartChannels, inviteChannels StartInviteChannels) {
	defer wg.Done()
	channels.IntToStartF1 <- n
	channels.IntToStartF1 <- 1
	channels.IntToStartF2 <- n
	channels.IntToStartF2 <- 1
	fibInviteChannels := FibChannels{
		F2InviteChan:    inviteChannels.F2ToF2InviteChan,
		F2ChannelsChan:  inviteChannels.F2ToF2ChannelsChan,
		F1ChannelsChan:  inviteChannels.F1ToF1ChannelsChan,
		ResChannelsChan: inviteChannels.StartToResChannelsChan,
		ResInviteChan:   inviteChannels.StartToResInviteChan,
	}
	CreateAndSendChannels(wg, fibInviteChannels)
	resChan := <-inviteChannels.StartToResChannelsChan
	resInviteChan := <-inviteChannels.StartToResInviteChan
	*res = Res(resChan, resInviteChan)
}

func Res(channels ResChannels, inviteChannels ResInviteChannels) int {
	stop := <-channels.BoolFromF3
	if stop {
		fib := <-channels.IntFromF3
		return fib
	} else {
		resChan := <-inviteChannels.ChannelsChan
		resInviteChan := <-inviteChannels.InviteChan
		return Res(resChan, resInviteChan)
	}
}

func StartF1(wg *sync.WaitGroup, channels StartF1Channels, inviteChannels StartF1InviteChannels) {
	defer wg.Done()
	ubound := <-channels.IntFromStart
	fib := <-channels.IntFromStart
	f1Chan := <-inviteChannels.ChannelsChan
	F1(ubound, 1, fib, f1Chan)
}

func F1(ubound, idx, fib int, channels F1Channels) {
	channels.IntToF3 <- ubound
	channels.IntToF3 <- idx
	channels.IntToF3 <- fib
}

func StartF2(wg *sync.WaitGroup, channels StartF2Channels, inviteChannels StartF2InviteChannels) {
	defer wg.Done()
	ubound := <-channels.IntFromStart
	fib := <-channels.IntFromStart
	f2Chan := <-inviteChannels.ChannelsChan
	f2InviteChan := <-inviteChannels.InviteChan
	F2(ubound, 2, fib, f2Chan, f2InviteChan)
}

func F2(ubound, idx, fib int, channels F2Channels, inviteChannels F2InviteChannels) {
	channels.IntToF3 <- fib
	stop := <-channels.BoolFromF3
	if !stop {
		f1Chan := <-inviteChannels.ChannelsChan
		F1(ubound, idx, fib, f1Chan)
	}
}

func F3(wg *sync.WaitGroup, channels F3Channels, inviteChannels F3InviteChannels) {
	defer wg.Done()
	ubound := <-channels.IntFromF1
	idx := <-channels.IntFromF1
	fib1 := <-channels.IntFromF1

	fib2 := <-channels.IntFromF2
	fib3 := fib1 + fib2
	currIdx := idx + 2
	if currIdx == ubound {
		channels.BoolToRes <- true
		channels.IntToRes <- fib3
		channels.BoolToF2 <- true
	} else {
		channels.BoolToRes <- false
		channels.BoolToF2 <- false
		// 	Create channels
		fibInviteChans := FibChannels{
			F2InviteChan:    inviteChannels.F3ToF2InviteChan,
			F2ChannelsChan:  inviteChannels.F3ToF2ChannelsChan,
			F1ChannelsChan:  inviteChannels.F2ToF1ChannelsChan,
			ResChannelsChan: inviteChannels.ResToResChannelsChan,
			ResInviteChan:   inviteChannels.ResToResInviteChan,
		}
		CreateAndSendChannels(wg, fibInviteChans)
		// fmt.Println("f3")
		f2Channels := <-inviteChannels.F3ToF2ChannelsChan
		f2InviteChannels := <-inviteChannels.F3ToF2InviteChan
		F2(ubound, currIdx, fib3, f2Channels, f2InviteChannels)
	}
}

func CreateAndSendChannels(wg *sync.WaitGroup, inviteChannels FibChannels) {
	f1ToF3Int := make(chan int, 1)
	f2ToF3Int := make(chan int, 1)
	f3ToResInt := make(chan int, 1)
	f3ToF2Stop := make(chan bool, 1)
	f3ToResStop := make(chan bool, 1)

	f2ToF1Chan := make(chan F1Channels, 1)
	f3ToF2Chan := make(chan F2Channels, 1)
	resToResChan := make(chan ResChannels, 1)

	f3ToF2Invite := make(chan F2InviteChannels, 1)
	resToResInvite := make(chan ResInviteChannels, 1)

	resChan := ResChannels{
		IntFromF3:  f3ToResInt,
		BoolFromF3: f3ToResStop,
	}
	f1Chan := F1Channels{IntToF3: f1ToF3Int}
	f2Chan := F2Channels{
		IntToF3:    f2ToF3Int,
		BoolFromF3: f3ToF2Stop,
	}
	f3Chan := F3Channels{
		IntFromF1: f1ToF3Int,
		IntFromF2: f2ToF3Int,
		BoolToF2:  f3ToF2Stop,
		BoolToRes: f3ToResStop,
		IntToRes:  f3ToResInt,
	}

	f2InviteChan := F2InviteChannels{ChannelsChan: f2ToF1Chan}
	f3InviteChan := F3InviteChannels{
		F3ToF2InviteChan:     f3ToF2Invite,
		F3ToF2ChannelsChan:   f3ToF2Chan,
		F2ToF1ChannelsChan:   f2ToF1Chan,
		ResToResChannelsChan: resToResChan,
		ResToResInviteChan:   resToResInvite,
	}
	resInviteChan := ResInviteChannels{
		InviteChan:   resToResInvite,
		ChannelsChan: resToResChan,
	}

	inviteChannels.F1ChannelsChan <- f1Chan

	inviteChannels.ResChannelsChan <- resChan
	inviteChannels.ResInviteChan <- resInviteChan

	inviteChannels.F2ChannelsChan <- f2Chan
	inviteChannels.F2InviteChan <- f2InviteChan

	//
	wg.Add(1)
	//
	go F3(wg, f3Chan, f3InviteChan)
}

func Fib(n int) int {
	if n < 3 {
		panic("n should always be > 2")
	}

	startF1Int := make(chan int, 1)
	startF2Int := make(chan int, 1)

	f1ToF1Chan := make(chan F1Channels, 1)
	f2ToF2Chan := make(chan F2Channels, 1)
	startToResChan := make(chan ResChannels, 1)

	f2ToF2InviteChan := make(chan F2InviteChannels, 1)
	startToResInviteChan := make(chan ResInviteChannels, 1)

	f1Chan := StartF1Channels{IntFromStart: startF1Int}
	f2Chan := StartF2Channels{IntFromStart: startF2Int}
	startChan := StartChannels{
		IntToStartF1: startF1Int,
		IntToStartF2: startF2Int,
	}

	f1InviteChan := StartF1InviteChannels{ChannelsChan: f1ToF1Chan}
	f2InviteChan := StartF2InviteChannels{
		ChannelsChan: f2ToF2Chan,
		InviteChan:   f2ToF2InviteChan,
	}
	startInviteChan := StartInviteChannels{
		StartToResInviteChan:   startToResInviteChan,
		StartToResChannelsChan: startToResChan,
		F1ToF1ChannelsChan:     f1ToF1Chan,
		F2ToF2ChannelsChan:     f2ToF2Chan,
		F2ToF2InviteChan:       f2ToF2InviteChan,
	}

	var wg sync.WaitGroup
	wg.Add(3)

	var res int

	go StartF1(&wg, f1Chan, f1InviteChan)
	go StartF2(&wg, f2Chan, f2InviteChan)
	go StartRes(&wg, &res, n, startChan, startInviteChan)

	wg.Wait()
	return res
}

// No callbacks
// No results
// 4 initial messages less
// 2 messages between roles less per protocol call
// 3 fewer channels per protocol call
// Remove empty invite channel struct

// 10 - Scribble: 65, 23.931811, Base: 37, 14.812260, Ratio: 0.569231
// 25 - Scribble: 174, 50.056384, Base: 100, 27.872956, Ratio: 0.574713
// 40 - Scribble: 278, 62.872710, Base: 153, 35.201322, Ratio: 0.550360
// 55 - Scribble: 384, 77.821024, Base: 211, 45.559231, Ratio: 0.549479
// 70 - Scribble: 492, 92.235661, Base: 297, 88.990919, Ratio: 0.603659
// 80 - Scribble: 561, 100.195011, Base: 306, 57.868833, Ratio: 0.545455
// 90 - Scribble: 627, 103.661244, Base: 342, 59.608328, Ratio: 0.545455
