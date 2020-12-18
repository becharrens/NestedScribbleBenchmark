package primesieve_base

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

var NumChannels = 0

func Main(primes []int, recvPrimes chan int, morePrimes chan bool) []int {
	if <-morePrimes {
		prime := <-recvPrimes
		primes := append(primes, prime)
		return Main(primes, recvPrimes, morePrimes)
	} else {
		return primes
	}
}

func PrimeSieve(n int) []int {
	if n < 2 {
		panic("n should be >= 2")
	}
	sendInt := make(chan int, 1)
	recvPrimes := make(chan int, 1)
	// filterPrime := make(chan int, 1)
	morePrimes := make(chan bool, 1)
	NumChannels += 3

	go FirstWorker(sendInt, recvPrimes, morePrimes)
	sendInt <- n
	sendInt <- 2
	primes := []int{2}
	return Main(primes, recvPrimes, morePrimes)
}

func FirstWorker(intChan <-chan int, resChan chan<- int, morePrimes chan<- bool) {
	ubound := <-intChan
	firstPrime := <-intChan
	primes := genPrimes(ubound, firstPrime)
	SieveChoice(primes, resChan, morePrimes)
}

func SieveWorker(primeChan, primesIn <-chan int, resChan chan<- int, morePrimes chan<- bool) {
	filterPrime := <-primeChan
	var primes []int
	for prime := range primesIn {
		if prime%filterPrime != 0 {
			primes = append(primes, prime)
		}
	}
	SieveChoice(primes, resChan, morePrimes)
}

func SieveChoice(primes []int, resChan chan<- int, morePrimes chan<- bool) {
	if len(primes) == 0 {
		morePrimes <- false
	} else {
		morePrimes <- true
		resChan <- primes[0]
		filterPrimeChan := make(chan int, 1)
		sendPrimes := make(chan int, 1)
		NumChannels += 2
		go SieveWorker(filterPrimeChan, sendPrimes, resChan, morePrimes)
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
