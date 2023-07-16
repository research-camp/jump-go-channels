package internal

import "errors"

var (
	// ErrQueueFull occurs when the queue is full
	ErrQueueFull = errors.New("queue is full")
)
