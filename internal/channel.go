package internal

import "container/list"

// Channel interface, we show channel operations as an interface
type Channel interface {
	Send(interface{})
	Recv() (interface{}, bool)
	Close()
	Next() bool
}

// Buffer interface, contains information of channel buffer
type Buffer interface {
	Enqueue(interface{}) error
	Dequeue() interface{}
	IsFull() bool
	IsEmpty() bool
}

type options struct {
	buffer Buffer
}

func WithBuffer(b Buffer) Option {
	return func(o *options) {
		o.buffer = b
	}
}

type Option func(o *options)

func NewChannel(size int, opts ...Option) Channel {
	o := &options{
		buffer: newListBuffer(size),
	}

	for _, f := range opts {
		f(o)
	}

	if size > 0 {
		return &BufferedChannel{
			buf:   o.buffer,
			sendQ: new(list.List).Init(),
			recvQ: new(list.List).Init(),
		}
	}

	return &SyncChan{
		sendQ: new(list.List).Init(),
		recvQ: new(list.List).Init(),
	}
}
