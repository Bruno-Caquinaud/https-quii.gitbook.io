package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Test 1 : Check Increment method", func(t *testing.T) {
		counter := Counter{}
		expected := 3000
		var sync sync.WaitGroup

		sync.Add(expected)

		for range expected {
			go func() {
				counter.Inc()
				sync.Done()
			}()
		}

		sync.Wait()

		if int(counter.count.Load()) != expected {
			t.Errorf("counter.count : %d != expected count : %d", counter.count.Load(), expected)
		}
	})
}
