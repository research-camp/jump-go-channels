package test

import (
	"sync"
	"testing"
)

type dataMock struct {
	Value int
}

func (d dataMock) Priority() int {
	return d.Value
}

func BenchmarkChannel(b *testing.B) {
	wg := sync.WaitGroup{}
	ch := make(chan dataMock)

	go func() {
		for {
			<-ch

			wg.Done()
		}
	}()

	for i := 0; i < b.N; i++ {
		wg.Add(1)

		ch <- dataMock{
			Value: b.N - i,
		}
	}

	wg.Wait()
}
