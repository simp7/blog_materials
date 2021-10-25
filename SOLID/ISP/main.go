package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"time"
)

func stream(r io.Reader, w io.Writer) error {

	b := make([]byte, 10)

	for {

		if _, err := r.Read(b); err != nil {
			return err
		}
		b = bytes.Trim(b, "\x00")

		time.Sleep(500 * time.Millisecond)

		if _, err := w.Write(b); err != nil {
			return err
		}

	}

}

func streaming(source io.ReadCloser, dest io.WriteCloser) {

	if err := stream(source, dest); err != io.EOF {
		log.Fatal("error while streaming")
	}

	log.Print("finish successfully")
	source.Close()
	dest.Close()

}

func main() {

	var source io.ReadCloser
	var dest io.WriteCloser

	source, _ = os.Open("test.txt")
	dest = os.Stdout
	streaming(source, dest)

	source = os.Stdin
	dest, _ = os.Create("input.txt")
	streaming(source, dest)

	source, _ = os.Open("input.txt")
	dest, _ = os.Create("log.txt")
	streaming(source, dest)

}
