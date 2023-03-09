# Introduction
- [Panic & Recover](https://golangbot.com/panic-and-recover/) is like exception in GoLang.
- There is no try-&-catch in GoLang.

# Return errors 
- [Errors](https://golangbot.com/custom-errors/) can be returned as multiple values from a function in Go

```go

package main
import (
	"fmt"
)

func reverseValues(a,b string)(string, string){
    return b,a    //notice how multiple values are returned
}

func main(){
    val1,val2:= reverseValues("interview","bit")    // notice how multiple values are assigned
    fmt.Println(val1, val2)
}
```

# Custom Errors in GoLang

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
