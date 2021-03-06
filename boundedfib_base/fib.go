package boundedfib_base

import "sync"

func startFstSender(wg *sync.WaitGroup, fib, idx, ubound int, sendChan chan int) {
	defer wg.Done()
	fstSender(fib, idx, ubound, sendChan)
}

func fstSender(fib, idx, ubound int, sendChan chan int) {
	sendChan <- ubound
	sendChan <- idx
	sendChan <- fib
}

func startSndSender(wg *sync.WaitGroup, fib, idx, ubound int, sendChan1, sendChan2 chan int, stopChan chan bool) {
	defer wg.Done()
	sndSender(fib, idx, ubound, sendChan1, sendChan2, stopChan)
}

func sndSender(fib, idx, ubound int, sendChan1, sendChan2 chan int, stopChan chan bool) {
	// sendChan1 <- idx
	sendChan1 <- fib
	stop := <-stopChan
	if !stop {
		fstSender(fib, idx, ubound, sendChan2)
	}
}

func fibCalc(wg *sync.WaitGroup, fib1Chan, fib2Chan, sendChan1, sendChan2, nextFib1Chan, resChan chan int,
	prevStopChan, stopChan chan bool) {
	defer wg.Done()
	ubound := <-fib1Chan
	idx1 := <-fib1Chan
	fib1 := <-fib1Chan

	fib2 := <-fib2Chan

	fib := fib1 + fib2
	idx := idx1 + 2
	returnFib := idx >= ubound
	prevStopChan <- returnFib // If worker is returning a result, prev worker should stop
	if returnFib {
		resChan <- fib
	} else {
		newFibChan1 := make(chan int, 1)
		newFibChan2 := make(chan int, 1)
		nextStopChan := make(chan bool, 1)
		wg.Add(1)
		go fibCalc(wg, nextFib1Chan, sendChan1, newFibChan1, newFibChan2, sendChan2, resChan, stopChan, nextStopChan)
		sndSender(fib, idx, ubound, sendChan1, sendChan2, stopChan)
	}
}

func Fibonacci(n int) int {
	if n < 3 {
		panic("n should always be > 2")
	}

	resChan := make(chan int, 1)
	fstChan := make(chan int, 1)
	sndChan1 := make(chan int, 1)
	sndChan2 := make(chan int, 1)
	stopChan := make(chan bool, 1)
	nextFibChan1 := make(chan int, 1)
	nextFibChan2 := make(chan int, 1)
	nextStopChan := make(chan bool, 1)
	var wg sync.WaitGroup
	wg.Add(3)

	go startFstSender(&wg, 1, 1, n, fstChan)
	go startSndSender(&wg, 1, 2, n, sndChan1, sndChan2, stopChan)
	go fibCalc(&wg, fstChan, sndChan1, nextFibChan1, nextFibChan2, sndChan2, resChan, stopChan, nextStopChan)
	// for i := 3; i <= n; i++ {
	// 	nextFibChan1 := make(chan int, 1)
	// 	nextFibChan2 := make(chan int, 1)
	// 	nextStopChan := make(chan bool, 1)
	// 	go fibCalc(fstChan, sndChan1, nextFibChan1, nextFibChan2, sndChan2,resChan, stopChan, nextStopChan, i >= n, n)
	// 	stopChan = nextStopChan
	// 	fstChan = sndChan2
	// 	sndChan1 = nextFibChan1
	// 	sndChan2 = nextFibChan2
	// }
	res := <-resChan
	wg.Wait()
	return res
}
