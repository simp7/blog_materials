package server

import (
	"github.com/simp7/blog_materials/chan_chan/stream/request"
	"math/rand"
	"time"
)

type server struct {
	data map[request.Request][]int
}

func New() *server {

	s := new(server)

	s.data = map[request.Request][]int{
		request.ODD: {1, 3, 5, 7, 9},
		request.EVN: {2, 4, 6, 8},
		request.FIB: {1, 2, 3, 5, 8},
		request.PRI: {2, 3, 5, 7},
	}

	return s

}

func (s *server) Serve(inCh <-chan request.Request, outChCh chan<- chan int) {

	var req request.Request

	for {
		if req = <-inCh; req == request.EOS {
			close(outChCh)
			return
		}
		s.send(req, outChCh)
	}

}

func (s *server) send(req request.Request, outChCh chan<- chan int) {
	ch := make(chan int)
	go s.sendEach(req, ch)
	outChCh <- ch
}

func (s *server) sendEach(req request.Request, outCh chan<- int) {
	defer close(outCh)
	for _, n := range s.data[req] {
		time.Sleep(time.Duration(rand.Intn(500))*time.Millisecond + 1)
		outCh <- n
	}
}
