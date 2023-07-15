package main

import (
	"log"
	"sync"

	"github.com/research-camp/go-channels-scheduling/internal"
)

func main() {
	// unbuffered channel
	ch := internal.NewChannel(0)
	wg := sync.WaitGroup{}

	wg.Add(1)

	// create a go routine
	go func() {
		msg, _ := ch.Recv()

		log.Println(msg)
		wg.Done()
	}()

	ch.Send("Hello World")

	wg.Wait()
}
