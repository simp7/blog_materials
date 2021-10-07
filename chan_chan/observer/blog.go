package main

import "fmt"

type Post struct {
	Title   string
	Content string
}

type blog struct {
	name        string
	subscribers map[chan<- string]struct{}
	subscribe   chan chan<- string
	unsubscribe chan chan<- string
	publish     chan Post
}

func Blog(name string) *blog {
	b := new(blog)
	b.subscribers = make(map[chan<- string]struct{})
	b.subscribe = make(chan chan<- string)
	b.unsubscribe = make(chan chan<- string)
	b.publish = make(chan Post)
	b.name = name
	go b.run()
	return b
}

func (b *blog) SubscribedBy(subscriber chan<- string) {
	b.subscribe <- subscriber
}

func (b *blog) UnsubscribedBy(subscriber chan<- string) {
	b.unsubscribe <- subscriber
}

func (b *blog) Publish(p Post) {
	b.publish <- p
}

func (b *blog) run() {
	for {
		select {
		case sub := <-b.subscribe:
			b.subscribers[sub] = struct{}{}
		case sub := <-b.unsubscribe:
			delete(b.subscribers, sub)
		case p := <-b.publish:
			fmt.Printf("%s published %s\n", b.name, p.Title)
			for subscriber := range b.subscribers {
				subscriber <- p.Title
			}
		}
	}
}
