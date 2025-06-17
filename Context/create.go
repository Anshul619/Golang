package main

import "fmt"

// To create a context, you can use the `context.Background()` function, which returns an empty, non-cancelable context as the root of the context tree.
func main() {
    // create context with 5 mins timeout
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        results <- fmt.Sprintf("Error creating request for %s: %s", url, err.Error())
        return
    }

    // context with custom value
    ctx := context.WithValue(context.Background(), "UserID", 123)
}