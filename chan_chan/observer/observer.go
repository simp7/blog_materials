//main.go
package main

import (
	"time"
)

func main() {

	b1 := Blog("Just_foo_bar")
	b2 := Blog("Interesting_Blog")

	u1 := User("Kim")
	u2 := User("Lee")
	u3 := User("Park")

	u1.Subscribe(b1)
	u1.Subscribe(b2)
	u2.Subscribe(b1)
	u3.Subscribe(b2)

	b1.Publish(Post{"Foo", "bar"})
	b2.Publish(Post{"Hello", "world!"})

	u1.Unsubscribe(b1)
	b1.Publish(Post{"Another foo", "another bar"})

	u2.Subscribe(b2)
	u1.Subscribe(b2)
	b2.Publish(Post{"The answer", "42"})

	time.Sleep(1 * time.Millisecond)

}
