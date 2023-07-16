package internal

import (
	"container/list"
	"sync"
	"sync/atomic"
	"time"
)

type SyncChan struct {
	val         *interface{}
	lock        sync.Mutex
	closed      bool
	sendQ       *list.List
	recvQ       *list.List
	sendCounter int32
	recvCounter int32
	schedule    bool
}

func (c *SyncChan) Close() {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.closed = true
}

func (c *SyncChan) Next() bool {
	defer c.lock.Unlock()

	for {
		c.lock.Lock()

		if c.closed && c.val == nil {
			return false
		}

		if c.val != nil {
			return true
		}

		c.lock.Unlock()
		time.Sleep(10 * time.Millisecond)
	}
}

func (c *SyncChan) Send(val interface{}) {
	if c.closed {
		panic("channel is closed")
	}

	c.lock.Lock()
	defer c.lock.Unlock()

	if c.val == nil {
		c.val = &val

		return
	}

	ticket := atomic.AddInt32(&c.sendCounter, 1)

	c.sendQ.PushBack(ticket)

	if c.schedule {
		c.sendQ = schedule(c.sendQ)
	}

	c.lock.Unlock()

	for {
		c.lock.Lock()

		if ticket == c.sendQ.Front().Value.(int32) {
			break
		}

		c.lock.Unlock()
		time.Sleep(10 * time.Millisecond)
	}

	c.sendQ.Remove(c.sendQ.Front())

	c.val = &val
}

func (c *SyncChan) Recv() (interface{}, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.val == nil && c.closed {
		return nil, false
	}

	if c.val != nil {
		val := *c.val
		c.val = nil

		return val, true
	}

	ticket := atomic.AddInt32(&c.recvCounter, 1)

	c.recvQ.PushBack(ticket)

	c.lock.Unlock()

	for {
		c.lock.Lock()

		if c.val == nil && c.closed {
			return nil, false
		}

		if c.val != nil && ticket == c.recvQ.Front().Value.(int32) {
			break
		}

		c.lock.Unlock()
		time.Sleep(10 * time.Millisecond)
	}

	val := *c.val
	c.val = nil

	c.recvQ.Remove(c.recvQ.Front())

	return val, true
}
