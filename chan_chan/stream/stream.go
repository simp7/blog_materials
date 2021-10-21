package main

import (
	"fmt"
	"github.com/simp7/blog_materials/chan_chan/stream/request"
	"github.com/simp7/blog_materials/chan_chan/stream/server"
	"math/rand"
	"time"
)

func SendRequests(outCh chan<- request.Request, requests ...request.Request) {
	defer close(outCh)
	for _, req := range requests {
		time.Sleep(time.Duration(rand.Intn(2500))*time.Millisecond + 1)
		outCh <- req
	}
}

func PrintAll(responses <-chan chan int) {
	for response := range responses {
		printEach(response)
		fmt.Println()
	}
	fmt.Println("Disconnected")
}

func printEach(response <-chan int) {
	for data := range response {
		fmt.Printf("%d ", data)
	}
}

func main() {

	requests := make(chan request.Request)
	responses := make(chan chan int)

	s := server.New()
	go s.Serve(requests, responses)
	go SendRequests(requests, request.ODD, request.FIB, request.PRI, request.EOS, request.EVN, request.PRI)
	PrintAll(responses)

}
