package internal

import "container/list"

type (
	// Channel interface, we show channel operations as an interface.
	Channel interface {
		Send(interface{})
		Recv() (interface{}, bool)
		Close()
		Next() bool
	}

	// Buffer interface, contains information of channel buffer.
	Buffer interface {
		Enqueue(interface{}) error
		Dequeue() interface{}
		IsFull() bool
		IsEmpty() bool
	}
)

// NewChannel builds a new channel for us
func NewChannel(size int, schedule bool) Channel {
	// create buffered channel
	if size > 0 {
		return &BufferedChannel{
			buf:   newListBuffer(size, schedule),
			sendQ: new(list.List).Init(),
			recvQ: new(list.List).Init(),
		}
	}

	// create unbuffered channel
	return &SyncChan{
		sendQ:    new(list.List).Init(),
		recvQ:    new(list.List).Init(),
		schedule: schedule,
	}
}
