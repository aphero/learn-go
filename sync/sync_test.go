package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := newCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := newCounter()

		// Create a wait group
		var wg sync.WaitGroup
		// Give the wait group a counter of its own
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			// Run the following functions within a goroutine so they can be performed concurrently
			go func() {
				// Increment our counter and then
				counter.Inc()
				// Decrement the wait group counter
				wg.Done()
			}()
		}
		// Wait until the counter for the wait group hits zero
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}

func newCounter() *Counter {
	return &Counter{}
}

// Learned a new command in this lesson
// go vet - will check for any more subtle errors in a file
