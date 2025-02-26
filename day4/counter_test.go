package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	// Test basic counter operations
	t.Run("basic counter operations", func(t *testing.T) {
		counter := getCounter("test1")
		if counter.GetValue() != 0 {
			t.Errorf("new counter should start at 0, got %d", counter.GetValue())
		}

		counter.Increment()
		if counter.GetValue() != 1 {
			t.Errorf("counter should be 1 after increment, got %d", counter.GetValue())
		}
	})

	// Test multiple counters
	t.Run("multiple counters", func(t *testing.T) {
		counter1 := getCounter("test2")
		counter2 := getCounter("test3")

		counter1.Increment()
		if counter2.GetValue() != 0 {
			t.Errorf("counter2 should be independent from counter1")
		}
	})

	// Test counter reuse
	t.Run("counter reuse", func(t *testing.T) {
		counter1 := getCounter("test4")
		counter1.Increment()

		counter2 := getCounter("test4")
		if counter1 != counter2 {
			t.Error("getCounter should return same instance for same ID")
		}
		if counter2.GetValue() != 1 {
			t.Error("reused counter should maintain its value")
		}
	})

	// Test concurrent access
	t.Run("concurrent access", func(t *testing.T) {
		counter := getCounter("test5")
		const numGoroutines = 10000
		const incrementsPerGoroutine = 100

		var syncGroup sync.WaitGroup
		var finishGroup sync.WaitGroup
		syncGroup.Add(numGoroutines)
		finishGroup.Add(numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func() {
				syncGroup.Done()
				syncGroup.Wait()
				defer finishGroup.Done()
				for j := 0; j < incrementsPerGoroutine; j++ {
					counter.Increment()
				}
			}()
		}

		finishGroup.Wait()

		expected := int64(numGoroutines * incrementsPerGoroutine)
		if counter.GetValue() != expected {
			t.Errorf("expected %d, got %d", expected, counter.GetValue())
		}
	})
}
