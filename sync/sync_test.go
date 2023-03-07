package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("It should increase the count by 3", func(t *testing.T){
		counter := CreateCounter()
		counter.Increment()
		counter.Increment()
		counter.Increment()

		assertCount(t, counter, 3)
	})

	t.Run("It should work with concurrency", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}
		var wg sync.WaitGroup

		wg.Add(wantedCount)
		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Increment()
				wg.Done()
			}()
		}
		wg.Wait()
	})
}

func assertCount(t testing.TB, c *Counter, wantedCount int) {
	t.Helper()

	if (c.Value() != wantedCount) {
		t.Errorf("got %d, want %d", c.Value(), wantedCount)
	}
}