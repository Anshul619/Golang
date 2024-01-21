# Introduction
- [Panic & Recover](https://golangbot.com/panic-and-recover/) is like exception in GoLang.
- There is no try-&-catch in GoLang.

# Return errors 
- [Errors](https://golangbot.com/custom-errors/) can be returned as multiple values from a function in Go

````go
package errors

// New returns an error that formats as the given text.
// Each call to New returns a distinct error value even if the text is identical.
func New(text string) error {  
        return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {  
        s string
}

func (e *errorString) Error() string {  
        return e.s
}
````

# Types of Error Logging

| Log                | Description                           |
|--------------------|---------------------------------------|
| log.Fatal("error") | Program will print error and exit     |
| log.Error("error") | Program will print error but NOT exit |

# Defer
- A [defer](https://go.dev/tour/flowcontrol/12) statement defers the execution of a function until the surrounding function returns.
- The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

## Not use defer in for-loop
- Don’t [defer](http://go-database-sql.org/retrieving.html) within a loop.
- A deferred statement doesn't get executed until the function exits, so a long-running function shouldn’t use it.
- If you do, you will slowly accumulate memory.
- If you are repeatedly querying and consuming result sets within a loop, you should explicitly call rows.Close() when you’re done with each result, and not use defer.

# Recovering from a Panic
- [recover](https://golangbot.com/panic-and-recover/#recoveringfromapanic) is a builtin function that is used to regain control of a panicking program.

````GO
package main

import (  
    "fmt"
)

func recoverFullName() {  
    if r := recover(); r!= nil {
        fmt.Println("recovered from ", r)
    }
}

func fullName(firstName *string, lastName *string) {  
    defer recoverFullName()
    if firstName == nil {
        panic("runtime error: first name cannot be nil")
    }
    if lastName == nil {
        panic("runtime error: last name cannot be nil")
    }
    fmt.Printf("%s %s\n", *firstName, *lastName)
    fmt.Println("returned normally from fullName")
}

func main() {  
    defer fmt.Println("deferred call in main")
    firstName := "Elon"
    fullName(&firstName, nil)
    fmt.Println("returned normally from main")
}
````

[Read more](https://golangbot.com/panic-and-recover/#recoveringfromapanic)

# Error Group in GoRoutines
- [Error Group](https://pkg.go.dev/golang.org/x/sync/errgroup) is used to return errors in goRoutines.

````go
package main

import (
        "context"
        "fmt"
        "math/rand"
        "time"

        "golang.org/x/sync/errgroup"
)

func fetchAll(ctx context.Context) error {
        errs, ctx := errgroup.WithContext(ctx)

        // run all the http requests in parallel
        for i := 0; i < 4; i++ {
                errs.Go(func() error {
                        // pretend this does an http request and returns an error                                                  
                        time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)                                               
                        return fmt.Errorf("error in go routine, bailing")                                                      
                })
        }

        // Wait for completion and return the first error (if any)                                                                 
        return errs.Wait()
}

func main() {
        fmt.Println(fetchAll(context.Background()))
}
````
