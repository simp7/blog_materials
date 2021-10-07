//workerpool/fetcher/simple.go
package fetcher

import (
	"fmt"
	"io"
	"math/rand"
	"time"
)

type fetcher struct {
	io.Writer
}

func Simple(writer io.Writer) *fetcher {
	return &fetcher{Writer: writer}
}

func (f *fetcher) Run() {

	readCh := make(chan string)
	fetchCh := make(chan string)

	go f.read(readCh)
	go f.fetch(readCh, fetchCh)
	f.print(fetchCh)

}

func (f *fetcher) read(outCh chan<- string) {
	defer close(outCh)
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		outCh <- fmt.Sprintf("line: %d", i)
	}
}

func (f *fetcher) fetch(inCh <-chan string, outCh chan<- string) {
	defer close(outCh)
	for i := range inCh {
		f.fetchLine(i, outCh)
	}
}

func (f *fetcher) print(inCh <-chan string) {
	for line := range inCh {
		_, _ = f.Write([]byte(line))
	}
}

func (f *fetcher) fetchLine(line string, outCh chan<- string) {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	outCh <- fmt.Sprintf("%s ... fetched!\n", line)
}
