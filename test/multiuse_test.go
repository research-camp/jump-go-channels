package test

import (
	"sync"
	"testing"

	"github.com/research-camp/go-channels-scheduling/internal"
)

type dataMock2 struct {
	Value int
}

func (d dataMock2) Priority() int {
	return d.Value
}

func BenchmarkMultiUse(b *testing.B) {
	wg := sync.WaitGroup{}
	ch := internal.NewChannel(0, true)

	for i := 0; i < b.N; i++ {
		go func() {
			for {
				ch.Recv()
			}
		}()
	}

	for i := 0; i < b.N; i++ {
		go func(index int) {
			for j := 0; j < b.N; j++ {
				ch.Send(dataMock2{
					Value: index*j + 1,
				})
			}

			wg.Done()
		}(i)

		wg.Add(1)
	}

	wg.Wait()
}
