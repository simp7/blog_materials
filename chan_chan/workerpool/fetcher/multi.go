//workerpool/fetcher/multi.go
package fetcher

import (
	"io"
)

type MultiFetcher struct {
	*fetcher
	workers int
}

func Multi(writer io.Writer, amount int) *MultiFetcher {
	return &MultiFetcher{fetcher: Simple(writer), workers: amount}
}

func (m *MultiFetcher) Run() {

	readCh := make(chan string)
	fetchCh := make(chan chan string, m.workers)
	distributeCh := make(chan string)

	go m.read(readCh)
	go m.fetch(readCh, fetchCh)
	go m.distribute(fetchCh, distributeCh)
	m.print(distributeCh)

}

func (m *MultiFetcher) fetch(inCh <-chan string, outChCh chan<- chan string) {

	defer close(outChCh)

	for line := range inCh {

		outCh := make(chan string)

		go m.fetchLine(line, outCh)
		outChCh <- outCh

	}

}

func (m *MultiFetcher) distribute(inCh <-chan chan string, outCh chan<- string) {
	defer close(outCh)
	for ch := range inCh {
		outCh <- <-ch
	}
}

func (m *MultiFetcher) fetchLine(line string, outCh chan<- string) {
	defer close(outCh)
	m.fetcher.fetchLine(line, outCh)
}
