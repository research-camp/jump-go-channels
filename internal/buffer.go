package internal

import "container/list"

type buffer struct {
	q        *list.List
	maxLen   int
	schedule bool
}

func newListBuffer(size int, schedule bool) *buffer {
	return &buffer{
		q:        new(list.List).Init(),
		maxLen:   size,
		schedule: schedule,
	}
}

func (c *buffer) IsFull() bool {
	return c.q.Len() >= c.maxLen
}

func (c *buffer) IsEmpty() bool {
	return c.q.Len() == 0
}

func (c *buffer) Enqueue(val interface{}) error {
	if c.IsFull() {
		return ErrQueueFull
	}

	c.q.PushBack(val)

	if c.schedule {
		c.q = schedule(c.q)
	}

	return nil
}

func (c *buffer) Dequeue() interface{} {
	if c.IsEmpty() {
		return nil
	}

	return c.q.Remove(c.q.Front())
}
