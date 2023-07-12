package internal

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
