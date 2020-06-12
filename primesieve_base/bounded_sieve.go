package primesieve_base

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func GeneratePrimes(n int, ch chan<- int, doneChan chan<- bool) {
	for i := 2; i <= n; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
	doneChan <- true
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func FilterPrimes(in <-chan int, out chan<- int, inDone <-chan bool,
	outDone chan<- bool, finished chan<- bool, prime int) {
	morePrimes := false
	for {
		select {
		case i := <-in:
			if i%prime != 0 {
				out <- i // Send 'i' to 'out'.
				morePrimes = true
			}
		case <-inDone:
			if !morePrimes {
				finished <- true
				return
			}
			outDone <- true
			return
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func PrimeSieve(n int) []int {
	if n < 2 {
		panic("n should be >= 2")
	}
	finishCh := make(chan bool)
	ch := make(chan int) // Create a new channel.
	doneChan := make(chan bool)
	go GeneratePrimes(n, ch, doneChan) // Launch Generate goroutine.
	var primes []int
LOOP:
	for {
		select {
		case prime := <-ch:
			primes = append(primes, prime)
			ch1 := make(chan int)
			doneChan1 := make(chan bool)
			go FilterPrimes(ch, ch1, doneChan, doneChan1, finishCh, prime)
			ch = ch1
			doneChan = doneChan1
		case <-finishCh:
			break LOOP
		}
	}
	return primes
}
