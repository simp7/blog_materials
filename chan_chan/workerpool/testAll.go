//workerpool/main.go
package main

import (
	"fmt"
	"github.com/simp7/blog_materials/chan_chan/workerpool/fetcher"
	"io"
	"os"
	"time"
)

type Fetcher interface {
	Run()
}

func Bench(f Fetcher) {
	start := time.Now()
	f.Run()
	fmt.Println("Fetching done:", time.Now().Sub(start))
}

func BenchAll(f ...Fetcher) {
	for i, target := range f {
		fmt.Printf("Bench fetcher %d\n", i+1)
		Bench(target)
	}
}

func main() {

	logAll := false
	var logger io.Writer
	logger, _ = os.Open("/dev/null")

	if logAll {
		logger = os.Stdout
	}

	f1 := fetcher.Simple(logger)
	f2 := fetcher.Parallel(logger, 1)
	f3 := fetcher.Parallel(logger, 30)
	f4 := fetcher.Multi(logger, 1)
	f5 := fetcher.Multi(logger, 30)
	BenchAll(f1, f2, f3, f4, f5)

}
