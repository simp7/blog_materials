//user.go
package main

import "fmt"

type user struct {
	id       string
	notifier chan string
}

func User(id string) *user {
	u := new(user)
	u.id = id
	u.notifier = make(chan string)
	go u.listen()
	return u
}

func (u *user) Subscribe(b *blog) {
	b.SubscribedBy(u.notifier)
}

func (u *user) Unsubscribe(b *blog) {
	b.UnsubscribedBy(u.notifier)
}

func (u *user) listen() {
	for {
		s := <-u.notifier
		fmt.Printf("%s received new post - %s\n", u.id, s)
	}
}
