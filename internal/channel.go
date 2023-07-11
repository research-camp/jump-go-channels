package internal

// Channel interface, we show channel operations as an interface
type Channel interface {
	Send(interface{})
	Recv() (interface{}, bool)
	Close()
	Next() bool
}
