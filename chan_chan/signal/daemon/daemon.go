package daemon

import (
	"fmt"
	"math/rand"
	"time"
)

type daemon struct {
	req chan chan struct{}
}

func New() *daemon {
	d := new(daemon)
	d.req = make(chan chan struct{})
	return d
}

func (d *daemon) Start() {
	go func() {
		for {
			select {
			case ch := <-d.req:
				rand.Seed(time.Now().UnixNano())
				time.Sleep(500 * time.Millisecond)
				fmt.Println(rand.Intn(100))
				close(ch)
			}
		}
	}()
}

func (d *daemon) Do() {
	ch := make(chan struct{})
	d.req <- ch
	<-ch
}
