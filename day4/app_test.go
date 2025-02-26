package main

import (
	"fmt"
	"io"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestApi(t *testing.T) {
	if !APITests {
		return
	}

	app := setupApp(true)

	// Test GET endpoint
	t.Run("GET counter value", func(t *testing.T) {
		resp, err := app.Test(httptest.NewRequest("GET", "/test1", nil))
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", resp.StatusCode)
		}
		body, _ := io.ReadAll(resp.Body)
		if string(body) != "0" {
			t.Errorf("Expected body '0', got '%s'", string(body))
		}
	})

	// Test POST endpoint
	t.Run("POST increment counter", func(t *testing.T) {
		resp, err := app.Test(httptest.NewRequest("POST", "/test1", nil))
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != 200 {
			t.Errorf("Expected status 200, got %d", resp.StatusCode)
		}
		body, _ := io.ReadAll(resp.Body)
		if string(body) != "1" {
			t.Errorf("Expected body '1', got '%s'", string(body))
		}
	})

	// Test concurrent access
	t.Run("Concurrent access", func(t *testing.T) {
		const numGoroutines = 10000

		var syncGroup sync.WaitGroup
		var finishGroup sync.WaitGroup
		syncGroup.Add(numGoroutines)
		finishGroup.Add(numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func() {
				syncGroup.Done()
				syncGroup.Wait()
				defer finishGroup.Done()
				resp, err := app.Test(httptest.NewRequest("POST", "/concurrent", nil))
				if err != nil {
					t.Error(err)
					return
				}
				if resp.StatusCode != 200 {
					t.Errorf("Expected status 200, got %d", resp.StatusCode)
				}
			}()
		}

		finishGroup.Wait()

		// Verify final count
		resp, _ := app.Test(httptest.NewRequest("GET", "/concurrent", nil))
		body, _ := io.ReadAll(resp.Body)
		if string(body) != fmt.Sprintf("%d", numGoroutines) {
			t.Errorf("Expected count %d, got %s", numGoroutines, string(body))
		}
	})
}
