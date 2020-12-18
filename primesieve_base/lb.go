package primesieve_base

import (
	"fmt"
)

// ======================================================================
// # Load Balancing #

func C(req_chan, resp_chan chan string) {
	req_chan <- "request"
	fmt.Println(<-resp_chan)
}

func LB(req_chan, resp_chan chan string) {
	fmt.Println(<-req_chan)
	resp_chan <- "resp"
}

func S(req_chan_chan chan chan string, resp_chan_chan chan chan string) {
	for {
		req_chan := <-req_chan_chan
		resp_chan := <-resp_chan_chan
		req_chan2 := make(chan string, 1)
		go LB(req_chan2, resp_chan)
		req := <-req_chan
		req_chan2 <- req
	}
}

func LoadBalancer() {
	req_chan_chan := make(chan chan string, 1)
	resp_chan_chan := make(chan chan string, 1)
	go S(req_chan_chan, resp_chan_chan)
	for {
		req_chan := make(chan string, 1)
		resp_chan := make(chan string, 1)
		go C(req_chan, resp_chan)
		req_chan_chan <- req_chan
		resp_chan_chan <- resp_chan
	}
}

// Possible bugs: mixing up channels

// ======================================================================

// # Dynamic Task Generation #

func Worker(idx int, req string, resp_chan chan string) {
	resp_chan <- req[idx : idx+1]
}

func Master(req_chan, resp_chan chan string) {
	for {
		req := <-req_chan
		worker_chan := make(chan string, 1)
		n_workers := len(req)
		for i := 0; i < n_workers; i++ {
			go Worker(i, req, worker_chan)
		}
		res := ""
		for i := 0; i < n_workers; i++ {
			res += <-worker_chan
		}
		resp_chan <- res
	}
}

func Client(req_chan, resp_chan chan string) {
	for i := 0; ; i++ {
		if i%2 == 0 {
			req_chan <- "request"
		} else {
			req_chan <- "req"
		}
		fmt.Println(<-resp_chan)
	}
}

func DynTaskGen() {
	req_chan := make(chan string, 1)
	resp_chan := make(chan string, 1)
	go Master(req_chan, resp_chan)
	go Client(req_chan, resp_chan)
	ch := make(chan bool)
	<-ch
}

// Posible bugs:
//  - Utilizar el mismo res chan para los workers y el client
//  - Orden de respuestas de workers no es deterministico (obvio)
