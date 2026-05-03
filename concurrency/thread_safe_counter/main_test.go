package main

// How to run: go test -bench=Benchmark -run='^$' -count=1 -v

import (
	"sync"
	"testing"
)

func BenchmarkMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		counter := &MutexCounter{}
		var wg sync.WaitGroup

		for g := 0; g < goroutines; g++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < increments; j++ {
					counter.Inc()
				}
			}()
		}

		wg.Wait()
	}
}

func BenchmarkAtomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		counter := &AtomicCounter{}
		var wg sync.WaitGroup

		for g := 0; g < goroutines; g++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for j := 0; j < increments; j++ {
					counter.Inc()
				}
			}()
		}

		wg.Wait()
	}
}

func BenchmarkChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		channelCounter()
	}
}
