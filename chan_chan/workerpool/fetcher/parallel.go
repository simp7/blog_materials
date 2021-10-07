//workerpool/fetcher/parallel.go
package fetcher

import (
	"io"
	"sync"
)

type ParallelFetcher struct {
	*fetcher
	workers int
}

func Parallel(writer io.Writer, amount int) *ParallelFetcher {
	return &ParallelFetcher{fetcher: Simple(writer), workers: amount}
}

func (p *ParallelFetcher) Run() {

	readCh := make(chan string)
	printCh := make(chan string)

	go p.read(readCh)
	go p.fetch(readCh, printCh)
	p.print(printCh)

}

func (p *ParallelFetcher) fetch(inCh <-chan string, outCh chan<- string) {

	defer close(outCh)

	var wg sync.WaitGroup

	for i := 0; i < p.workers; i++ {
		wg.Add(1)
		go p.work(&wg, inCh, outCh)
	}

	wg.Wait()

}

func (p *ParallelFetcher) work(wg *sync.WaitGroup, inCh <-chan string, outCh chan<- string) {
	for line := range inCh {
		p.fetchLine(line, outCh)
	}
	wg.Done()
}
