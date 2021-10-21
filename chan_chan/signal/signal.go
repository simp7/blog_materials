package main

import (
	"fmt"
	"github.com/simp7/blog_materials/chan_chan/signal/daemon"
)

func main() {
	d := daemon.New()
	d.Start()
	d.Do()
	d.Do()
	d.Do()
	fmt.Println("Disconnected.")
}
