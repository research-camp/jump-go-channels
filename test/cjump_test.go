package test

import (
	"github.com/research-camp/go-channels-scheduling/internal"
	"sync"
	"testing"
)

func BenchmarkCJump(b *testing.B) {
	wg := sync.WaitGroup{}
	ch := internal.NewChannel(0, true)

	go func() {
		for {
			ch.Recv()

			wg.Done()
		}
	}()

	for i := 0; i < b.N; i++ {
		wg.Add(1)

		ch.Send(dataMock{
			Value: i,
		})
	}

	wg.Wait()
}
