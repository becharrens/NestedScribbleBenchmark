package primesieve_base

import "sync"

// Send the sequence 2, 3, 4, ... to channel 'ch'.
// func GeneratePrimes(n int, ch chan<- int, doneChan chan<- bool) {
// 	for i := 2; i <= n; i++ {
// 		ch <- i // Send 'i' to channel 'ch'.
// 	}
// 	doneChan <- true
// }
//
// // Copy the values from channel 'in' to channel 'out',
// // removing those divisible by 'prime'.
// func FilterPrimes(in <-chan int, out chan<- int, inDone <-chan bool,
// 	outDone chan<- bool, finished chan<- bool, prime int) {
// 	morePrimes := false
// 	for {
// 		select {
// 		case i := <-in:
// 			if i%prime != 0 {
// 				out <- i // Send 'i' to 'out'.
// 				morePrimes = true
// 			}
// 		case <-inDone:
// 			if !morePrimes {
// 				finished <- true
// 				return
// 			}
// 			outDone <- true
// 			return
// 		}
// 	}
// }

// The prime sieve: Daisy-chain Filter processes.
// func PrimeSieve(n int) []int {
// 	if n < 2 {
// 		panic("n should be >= 2")
// 	}
// 	finishCh := make(chan bool)
// 	ch := make(chan int) // Create a new channel.
// 	doneChan := make(chan bool)
// 	go GeneratePrimes(n, ch, doneChan) // Launch Generate goroutine.
// 	var primes []int
// LOOP:
// 	for {
// 		select {
// 		case prime := <-ch:
// 			primes = append(primes, prime)
// 			ch1 := make(chan int)
// 			doneChan1 := make(chan bool)
// 			go FilterPrimes(ch, ch1, doneChan, doneChan1, finishCh, prime)
// 			ch = ch1
// 			doneChan = doneChan1
// 		case <-finishCh:
// 			break LOOP
// 		}
// 	}
// 	return primes
// }

func StartMain(wg *sync.WaitGroup, primes *[]int, n int, sendInt chan int, recvPrimes chan int, morePrimes chan bool) {
	defer wg.Done()
	sendInt <- n
	sendInt <- 2
	Main(primes, recvPrimes, morePrimes)
}

func Main(primes *[]int, recvPrimes chan int, morePrimes chan bool) {
	if <-morePrimes {
		prime := <-recvPrimes
		*primes = append(*primes, prime)
		Main(primes, recvPrimes, morePrimes)
	}
}

func PrimeSieve(n int) []int {
	if n < 2 {
		panic("n should be >= 2")
	}
	sendInt := make(chan int, 1)
	recvPrimes := make(chan int, 1)
	morePrimes := make(chan bool, 1)

	primes := []int{2}

	var wg sync.WaitGroup
	wg.Add(2)

	go FirstWorker(&wg, sendInt, recvPrimes, morePrimes)
	go StartMain(&wg, &primes, n, sendInt, recvPrimes, morePrimes)

	wg.Wait()
	return primes
}

func FirstWorker(wg *sync.WaitGroup, intChan <-chan int, resChan chan<- int, morePrimes chan<- bool) {
	defer wg.Done()
	ubound := <-intChan
	firstPrime := <-intChan
	primes := genPrimes(ubound, firstPrime)
	SieveChoice(wg, primes, resChan, morePrimes)
}

func SieveWorker(wg *sync.WaitGroup, primeChan, primesIn <-chan int, resChan chan<- int, morePrimes chan<- bool) {
	defer wg.Done()
	filterPrime := <-primeChan
	var primes []int
	for prime := range primesIn {
		if prime%filterPrime != 0 {
			primes = append(primes, prime)
		}
	}
	SieveChoice(wg, primes, resChan, morePrimes)
}

func SieveChoice(wg *sync.WaitGroup, primes []int, resChan chan<- int, morePrimes chan<- bool) {
	if len(primes) == 0 {
		morePrimes <- false
	} else {
		morePrimes <- true
		resChan <- primes[0]
		filterPrimeChan := make(chan int, 1)
		sendPrimes := make(chan int, 1)
		wg.Add(1)
		go SieveWorker(wg, filterPrimeChan, sendPrimes, resChan, morePrimes)
		forwardPrimes(primes, filterPrimeChan, sendPrimes)
	}
}

func genPrimes(ubound int, filterPrime int) []int {
	var primes []int
	for prime := 3; prime <= ubound; prime++ {
		if prime%filterPrime != 0 {
			primes = append(primes, prime)
		}
	}
	return primes
}

func forwardPrimes(primes []int, primeChan, sendPrimes chan<- int) {
	primeChan <- primes[0]
	for i := 1; i < len(primes); i++ {
		sendPrimes <- primes[i]
	}
	close(sendPrimes)
}

// 100 - Scribble: 305, 178.750591, Base: 113, 31.324126, Ratio: 0.370492
// 1100 - Scribble: 7668, 452.370821, Base: 4169, 288.037503, Ratio: 0.543688
// 2100 - Scribble: 21332, 1071.216199, Base: 11990, 517.326917, Ratio: 0.562066
// 3100 - Scribble: 40331, 1056.366749, Base: 22780, 632.720595, Ratio: 0.564826
// 4100 - Scribble: 65228, 1956.849161, Base: 36788, 779.933996, Ratio: 0.563991
// 5100 - Scribble: 93846, 2259.329353, Base: 54142, 3312.552018, Ratio: 0.576924
// 6100 - Scribble: 126499, 2662.512696, Base: 72355, 1267.370471, Ratio: 0.571981
// 7100 - Scribble: 164252, 3041.754332, Base: 94528, 2959.454203, Ratio: 0.575506
// 8100 - Scribble: 203316, 1618.411284, Base: 117686, 2019.867681, Ratio: 0.578833
// 9100 - Scribble: 249860, 2836.200837, Base: 144499, 4472.291520, Ratio: 0.578320
