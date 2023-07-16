package internal

import "container/list"

type scheduler struct{}

func (s scheduler) schedule(queue *list.List) *list.List {
	// creating a new list
	newQueue := new(list.List).Init()

	// get queue size
	size := queue.Len()

	// scheduling logic in here
	for i := 0; i < size; i++ {
		newQueue.PushBack(queue.Back())
	}

	return newQueue
}
