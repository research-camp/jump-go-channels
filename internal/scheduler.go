package internal

import (
	"container/list"

	"github.com/research-camp/go-channels-scheduling/pkg"

	"github.com/amirhnajafiz/pyramid"
)

func schedule(queue *list.List) *list.List {
	// creating a heap with scheduling interface
	heap := pyramid.NewHeap[pkg.Schedulable](func(a pkg.Schedulable, b pkg.Schedulable) bool {
		return a.Priority() > b.Priority()
	})

	// get queue size
	size := queue.Len()

	// scheduling logic in here
	for i := 0; i < size; i++ {
		heap.Push(queue.Front())
	}

	// creating a new list
	newQueue := new(list.List).Init()

	// put them in new queue
	for i := 0; i < size; i++ {
		newQueue.PushBack(heap.Pop())
	}

	return newQueue
}
