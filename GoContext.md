# Context in Go
- [Context](https://medium.com/@jamal.kaksouri/the-complete-guide-to-context-in-golang-efficient-concurrency-management-43d722f6eaea) provides a mechanism to control the lifecycle, cancellation, and propagation of requests across multiple goroutines.

# Best Practices

| Title                              | Description                                                                                                                                                                                                             |
|------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Pass Context Explicitly            | Always pass the context as an explicit argument to functions or goroutines instead of using global variables. This makes it easier to manage the context’s lifecycle and prevents potential data races.                 |
| Use `context.TODO()`               | If you are unsure which context to use in a particular scenario, consider using `context.TODO()`. However, make sure to replace it with the appropriate context later.                                                  |
| Avoid Using `context.Background()` | Instead of using `context.Background()` directly, create a specific context using `context.WithCancel()` or `context.WithTimeout()` to manage its lifecycle and avoid resource leaks.                                   |
| Prefer Cancel Over Timeout         | Use `context.WithCancel()` for cancellation when possible, as it allows you to explicitly trigger cancellation when needed. `context.WithTimeout()` is more suitable when you need an automatic cancellation mechanism. |
| Keep Context Size Small            | Avoid storing large or unnecessary data in the context. Only include the data required for the specific operation.                                                                                                      |
| Avoid Chaining Contexts            | Chaining contexts can lead to confusion and make it challenging to manage the context hierarchy. Instead, propagate a single context throughout the application.                                                        |
| Be Mindful of Goroutine Leaks      | Always ensure that goroutines associated with a context are properly closed or terminated to avoid goroutine leaks.                                                                                                     |

# Context in Real-World Scenarios

| Title                    | Description                                                                                                                                                                                                                                                                                                    |
|--------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Context in Microservices | In a microservices architecture, each service often relies on various external dependencies and communicates with other services. <br/>- Context can be used to propagate important information, such as authentication tokens, request metadata, or tracing identifiers, throughout the service interactions. |
| Context in Test Suites   | When writing test suites, context can be utilized to manage test timeouts, control test-specific configurations, and enable graceful termination of tests.                                                                                                                                                     |
| Context in Web Servers   | Web servers handle multiple concurrent requests, and context helps manage the lifecycle of each request.                                                                                                                                                                                                       |

# How to create context?
- To create a context, you can use the `context.Background()` function, which returns an empty, non-cancelable context as the root of the context tree.

````go
func main() {
    // create context with 5 mins timeout
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        results <- fmt.Sprintf("Error creating request for %s: %s", url, err.Error())
        return
    }
}
````

# Propagating context
- Once you have a context, you can propagate it to downstream functions or goroutines by passing it as an argument.
- Custom value can be assigned in the context.

````go
ctx := context.WithValue(context.Background(), "UserID", 123)
````

# Setting a Deadline for Context

````go
package main

import (
 "context"
 "fmt"
 "time"
)

func main() {
 ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
 defer cancel()

 go performTask(ctx)

 time.Sleep(3 * time.Second)
}

func performTask(ctx context.Context) {
 select {
 case <-ctx.Done():
  fmt.Println("Task completed or deadline exceeded:", ctx.Err())
  return
 }
}
````


# References
- [The Complete Guide to Context in Golang: Efficient Concurrency Management](https://medium.com/@jamal.kaksouri/the-complete-guide-to-context-in-golang-efficient-concurrency-management-43d722f6eaea)