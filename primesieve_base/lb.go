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
	resp_chan <- req[idx:idx+1] + req[idx:idx+1] /* Compute worker result */
}

func Master(req_chan, resp_chan chan string) {
	for {
		req := <-req_chan /* Service request from Client */
		worker_chan := make(chan string, 1)
		n_workers := len(req)            /* Spawn workers to calculate result */
		for i := 0; i < n_workers; i++ { /* n_workers depends on runtime value */
			go Worker(i, req, worker_chan)
		}
		res := ""
		for i := 0; i < n_workers; i++ { /* Aggregate worker results */
			res += <-worker_chan
		}
		resp_chan <- res /* Send response to Client */
	}
}

func Client(req_chan, resp_chan chan string) {
	requests := []string{"req", "short", "longreq"}
	for i := 0; ; i++ { /* Send requests to Master */
		req_chan <- requests[i%len(requests)]
		fmt.Println(<-resp_chan) /* Process response */
	}
}

func DynTaskGen() {
	req_chan := make(chan string, 1)
	resp_chan := make(chan string, 1)
	go Master(req_chan, resp_chan) /* Spawn Client and Master */
	go Client(req_chan, resp_chan)
	ch := make(chan bool)
	<-ch /* Block so main thread does not exit prematurely */
}

// Posible bugs:
//  - Utilizar el mismo res chan para los workers y el client
//  - Orden de respuestas de workers no es deterministico (obvio)
