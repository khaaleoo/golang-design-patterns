package singleton

import (
	"sync"
	"testing"
)

func TestGetInstanceIsSingleton(t *testing.T) {
	a := GetInstance()
	b := GetInstance()

	if a != b {
		t.Fatal("expected the same instance")
	}
}

func TestGetInstanceIsThreadSafe(t *testing.T) {
	const goroutines = 100
	const increasesPerGoroutine = 100

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < increasesPerGoroutine; j++ {
				GetInstance().Increase()
			}
		}()
	}

	wg.Wait()

	expected := goroutines * increasesPerGoroutine
	if got := GetInstance().Get(); got != expected {
		t.Fatalf("expected %d, got %d", expected, got)
	}
}
