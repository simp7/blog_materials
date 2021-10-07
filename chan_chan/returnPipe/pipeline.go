package main

import (
	"fmt"
	"math/rand"
	"time"
)

func readSomething() <-chan string {

	outCh := make(chan string)

	go func() {
		defer close(outCh)
		for i := 0; i < 1000; i++ {
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
			outCh <- fmt.Sprintf("line: %d", i)
		}
	}()

	return outCh

}

func fetchSomething(inCh <-chan string, intensity int) <-chan chan string {

	outChCh := make(chan chan string, intensity)

	go func() {

		defer close(outChCh)

		for line := range inCh {
			outChCh <- getLine(line)
		}

	}()

	return outChCh

}

func getLine(line string) chan string {

	outCh := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		outCh <- fmt.Sprintf("%s ... fetched!", line)
		close(outCh)
	}()

	return outCh

}

func distribute(inCh <-chan chan string) <-chan string {

	outCh := make(chan string)

	go func() {
		defer close(outCh)
		for ch := range inCh {
			outCh <- <-ch
		}
	}()

	return outCh

}

func printSomething(inCh <-chan string) {
	for line := range inCh {
		fmt.Println(line)
	}
}

func test(intensity int) {

	start := time.Now()

	reader := readSomething()
	fetcher := fetchSomething(reader, intensity)
	distributor := distribute(fetcher)
	printSomething(distributor)

	fmt.Println("done", time.Now().Sub(start))

}

func main() {
	test(20)
	test(40)
}
