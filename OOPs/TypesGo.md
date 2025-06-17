# Static Typing in Go
- Go is a very strong and statically typed programming language.
- This ensures that the code is type-safe and all type conversions are handled efficiently.

# Why does Go have type parameters?
- Type parameters permit what is known as [generic programming]((https://go.dev/doc/faq#overloading)), in which functions and data structures are defined in terms of types that are specified later, when those functions and data structures are used.
- For example, they make it possible to write a function that returns the minimum of two values of any ordered type, without having to write a separate version for each possible type.

# Features

| Feature                                                                    | Description                                                                                                                                                                                                 | Example Code                                                                 |
|----------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------|
| [Type Assertion](https://www.geeksforgeeks.org/type-assertions-in-golang/) | Type casting of the generic interface to the specific type, is done like `interfaceVariable.(concreteType)`.                                                                                                | var input interface{} = 123<br/>str := input.(string)                        |
| Type Conversion                                                            | Convert a type to another type                                                                                                                                                                              | t := []byte("String")                                                        |
| [Type Switches](https://exercism.org/tracks/go/concepts/type-assertion)    | A type switch can perform several type assertions in a row. It has the same syntax as a type assertion (interfaceVariable.(concreteType)), but instead of a specific concreteType it uses the keyword type. | var i interface{} = 12<br/>switch v := i.(type) {case int:<br/>case string:} |

# Is it possible to declare variables of different types in a single line of code in Golang?

```go
var a,b,c= 9, 7.1, "interviewbit"
```

# Type & Receiver Function

````go
package main
 
import "fmt"
 
type book string

// Receiver function
func (b book) printTitle() {
    fmt.Println(b)
}
 
func main() {
    var b book = "Harry Potter"
    b.printTitle()
}
````