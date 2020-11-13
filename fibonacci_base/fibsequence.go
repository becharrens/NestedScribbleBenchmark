package fibonacci_base

import (
	"fmt"
	"strconv"
)

// func Fib1(fib int, next chan int) {
// 	next <- fib
// }
func Fib2(fib int, next chan int) {
	// fmt.Println("sending", fib)
	next <- fib
	// fmt.Println("sending", fib)
	next <- fib
}
func Fib3(resChan, fib1Chan, fib2Chan chan int) {
	fib1, fib2 := <-fib1Chan, <-fib2Chan
	fib := fib1 + fib2
	go Fib3(resChan, fib2Chan, fib1Chan)
	Fib2(fib, fib1Chan)
	resChan <- fib
}
func FibSequence() {
	// Create async (buffered) channels
	fib1Chan := make(chan int, 10)
	fib2Chan := make(chan int, 10)
	resChan := make(chan int, 10)
	// Create fst two workers
	go Fib2(1, fib2Chan)
	go Fib3(resChan, fib1Chan, fib2Chan)
	// Send first old_fibonacci number
	fib1Chan <- 1
	// Print Fibonacci sequence
	for i := 3; ; i++ {
		fmt.Println("Fib "+strconv.Itoa(i)+":", <-resChan)
	}
}
