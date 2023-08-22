package internal

// Wrapper works as a scheduler on our channel
type Wrapper struct {
	InChannel  chan interface{}
	OutChannel chan interface{}

	channel chan interface{}
}

// New returns a new channel
func New(size int) *Wrapper {
	w := Wrapper{
		channel: make(chan interface{}),
	}

	if size > 0 {
		w.InChannel = make(chan interface{}, size)
		w.OutChannel = make(chan interface{}, size)
	} else {
		w.InChannel = make(chan interface{})
		w.OutChannel = make(chan interface{})
	}

	go w.onSend()
	go w.onRecv()

	return &w
}

// onSend handler
func (w Wrapper) onSend() {
	for {
		w.channel <- w.InChannel
	}
}

// onRecv handler
func (w Wrapper) onRecv() {
	for {
		w.OutChannel <- w.channel
	}
}
