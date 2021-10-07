package main

import (
	"fmt"
)

type Request string

const (
	ODD Request = "ODD"
	EVN         = "EVEN"
	FIB         = "FIBONACCI"
	PRI         = "PRIME"
	EOS         = "EXIT" //End Of Stream
)

type server struct {
	requestCh  <-chan Request
	responseCh chan chan int
	data       map[Request][]int
}

func NewServer(inCh <-chan Request, outCh chan chan int) *server {

	s := new(server)

	s.requestCh = inCh
	s.responseCh = outCh

	s.data = map[Request][]int{
		ODD: {1, 3, 5, 7, 9},
		EVN: {2, 4, 6, 8},
		FIB: {1, 2, 3, 5, 8},
		PRI: {2, 3, 5, 7},
	}

	return s

}

func (s *server) Serve() {

	for {

		req := <-s.requestCh
		if req == EOS {
			close(s.responseCh)
			return
		}

		s.receive(req)

	}

}

func (s *server) receive(req Request) {

	ch := make(chan int)
	defer close(ch)

	s.responseCh <- ch

	for _, n := range s.data[req] {
		ch <- n
	}

}

func PrintAll(responses <-chan chan int) {
	for response := range responses {
		printEach(response)
		fmt.Println()
	}
}

func printEach(response <-chan int) {
	for data := range response {
		fmt.Printf("%d ", data)
	}
}

func Send(outCh chan<- Request, requests ...Request) {
	defer close(outCh)
	for _, request := range requests {
		outCh <- request
	}
}

func main() {

	requests := make(chan Request)
	responses := make(chan chan int)

	s := NewServer(requests, responses)
	go s.Serve()

	go Send(requests, ODD, FIB, PRI, EOS, EVN, PRI)
	PrintAll(responses)

}
