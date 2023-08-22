package internal

import (
	"sync"
	"time"
)

// Wrapper works as a scheduler on our channel
type Wrapper struct {
	InChannel  chan interface{}
	OutChannel chan interface{}

	lock     sync.Mutex
	count    int
	limit    int
	schedule bool

	channel chan interface{}
}

// New returns a new channel
func New(schedule bool, size int) *Wrapper {
	w := Wrapper{
		channel:  make(chan interface{}),
		schedule: schedule,
	}

	if size > 0 {
		w.InChannel = make(chan interface{}, size)
		w.OutChannel = make(chan interface{}, size)
		w.limit = size
	} else {
		w.InChannel = make(chan interface{})
		w.OutChannel = make(chan interface{})
		w.limit = 1
	}

	go w.onSend()
	go w.onRecv()

	return &w
}

// onSend handler
func (w *Wrapper) onSend() {
	for {
		val := <-w.InChannel

		for {
			w.lock.Lock()

			if w.count < w.limit {
				w.count++

				break
			}

			w.lock.Unlock()

			time.Sleep(1 * time.Microsecond)
		}

		w.channel <- val
	}
}

// onRecv handler
func (w *Wrapper) onRecv() {
	for {
		val := <-w.channel

		w.lock.Lock()

		w.count--

		w.lock.Unlock()

		w.OutChannel <- val
	}
}
