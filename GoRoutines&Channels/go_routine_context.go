package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func callAPI(ctx context.Context, name string, url string) error {
	// Create a new HTTP request with the context
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return fmt.Errorf("%s: failed to create request: %w", name, err)
	}

	// Use default HTTP client
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("%s: request failed: %w", name, err)
	}
	defer resp.Body.Close() // Always close body to avoid leaks

	fmt.Printf("%s: response code %d\n", name, resp.StatusCode)
	return nil
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	apis := []struct {
		name string
		url  string
	}{
		{"api1", "https://httpbin.org/delay/1"}, // responds in 1s
		{"api2", "https://httpbin.org/delay/3"}, // will timeout
		{"api3", "https://httpbin.org/status/200"},
	}

	var wg sync.WaitGroup
	errCh := make(chan error, len(apis))

	for _, api := range apis {
		wg.Add(1)
		go func(apiName, apiURL string) {
			defer wg.Done()
			if err := callAPI(ctx, apiName, apiURL); err != nil {
				fmt.Println("Error:", err)
				// Cancel other requests on first failure or timeout
				cancel()
				errCh <- err
			}
		}(api.name, api.url)
	}

	// Wait for all goroutines
	wg.Wait()
	close(errCh)

	if len(errCh) > 0 {
		fmt.Println("Some API calls failed or timed out.")
	} else {
		fmt.Println("All API calls succeeded.")
	}
}
