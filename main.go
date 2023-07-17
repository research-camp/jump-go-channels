package main

import (
	"log"
	"sync"

	"github.com/research-camp/go-channels-scheduling/internal"
)

type Data struct {
	Value string
	p     int
}

func (d Data) Priority() int {
	return d.p
}

func main() {
	// unbuffered channel
	ch := internal.NewChannel(0, true)
	wg := sync.WaitGroup{}

	wg.Add(2)

	// create a go routine
	go func() {
		for {
			log.Println("waiting ...")
			msg, _ := ch.Recv()

			log.Println(msg.(Data).Value)

			wg.Done()
		}
	}()

	ch.Send(Data{
		Value: "low value",
		p:     1,
	})
	ch.Send(Data{
		Value: "high value",
		p:     2,
	})

	wg.Wait()
	ch.Close()
}
