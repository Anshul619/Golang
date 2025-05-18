# What is Golang?
- [Go](https://go.dev/) is a high level, general-purpose programming language that is [very strongly and statically typed](OOPs/TypesGo.md) by providing support for garbage collection and [concurrent programming](GoRoutines&Channels).
- Go technically is [pass by value](https://stackoverflow.com/questions/47296325/passing-by-reference-and-value-in-go-to-functions).
- Go is a [case-sensitive language](https://en.wikipedia.org/wiki/Case_sensitivity).

# Key Features

| Title                                                                                                    | Remarks                                                                                                                                                                                                                                                                                         |
|----------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| :star: [Types in GoLang](OOPs/TypesGo.md)                                                                | Type parameters permit what is known as generic programming, in which functions and data structures are defined in terms of types that are specified later, when those functions and data structures are used.                                                                                  |
| :star: [Coding Helpers & Guidelines in GoLang](CodingHelpers&Guidelines.md)                              | Coding Helpers and guidelines for coding in GoLang.                                                                                                                                                                                                                                             |
| :star: [Pointers in GoLang](Pointers.md)                                                                 | GoLang supports pointers using (*, & operators)                                                                                                                                                                                                                                                 |
| :star: [Concurrency in GoLang](GoRoutines&Channels)                                                      | Go provides very good support for concurrency using [Go Routines or channels](https://go.dev/tour/concurrency/1)                                                                                                                                                                                |
| [Slices in GoLang](Slices.md)                                                                            | Slice in Go is a lightweight data structure of variable length sequence for storing homogeneous data.                                                                                                                                                                                           |
| [OOPs in GoLang](OOPs/Readme.md)                                                                         | Although Go has types and methods and allows an object-oriented style of programming, there is no type hierarchy.                                                                                                                                                                               |
| [Panic & Recover in GoLang](Panics&Recover.md)                                                           | [Panic & Recover](https://golangbot.com/panic-and-recover/) is like exception in GoLang.                                                                                                                                                                                                        |
| [Unit Testing in GoLang](UnitTesting.md)                                                                 | GoLang supports unit testing using [Testing](https://pkg.go.dev/testing) package                                                                                                                                                                                                                |
| [Labels in GoLang](https://medium.com/golangspec/labels-in-go-4ffd81932339)                              | Label is used in break and continue statement where it’s optional but it’s required in goto statement.                                                                                                                                                                                          |
| [How to Work With SQL in GoLang?](https://betterprogramming.pub/how-to-work-with-sql-in-go-ca8bc0b30722) | [database/sql package](https://pkg.go.dev/database/sql) helps to query SQL databases.                                                                                                                                                                                                           |
| [DB Transaction in GoLang](DBTransaction.md)                                                             | [Using Begin, Commit code](https://dev.to/techschoolguru/a-clean-way-to-implement-database-transaction-in-golang-2ba) block, [atomicity](https://github.com/Anshul619/System-Designs/blob/main/src/1_HLDDesignComponents/0_SystemGlossaries/Database/Atomicity.md) can be implemented in GoLang. |
| [Packages](Packages.md)                                                                                  | In go, code is organized using packages.                                                                                                                                                                                                                                                        |
| [Modules](Modules.md)                                                                                    | Modules represent an app/service in Go.                                                                                                                                                                                                                                                         |
| [Comments & Documentation in Go](Comments.md)                                                            |                                                                                                                                                                                                                                                                                                 |

# What do you understand by Golang string literals?
- String literals are those variables storing string constants that can be a single character or that can be obtained as a result of the concatenation of a sequence of characters. 
- Go provides two types of string literals - Raw & Interpreted string literals.

## Raw string literals

```go
`interviewbit`
```

## Interpreted string literals

```go
"Interviewbit
Website"
```

# What is the syntax used for the for loop in Golang?

```go
for [condition |  ( init; condition; increment ) | Range]  
{  
statement(s);  
//more statements
}  
```

# What do you understand by the scope of variables in Go?

## Local variables
- These are declared inside a function or a block and is accessible only within these entities.

## Global variables
- These are declared outside function or block and is accessible by the whole source file.

# How can we check if the Go map contains a key?

![img.png](assests/go_map_img.png)

```go
if val, isExists := map_obj["foo"]; isExists {
    //do steps needed here
}
```

# Can you format a string without printing?

```go
return fmt.Sprintf ("Size: %d MB.", 50)
```

# Closures in Golang
- Go language provides a special feature known as an anonymous function. 
- [An anonymous function can form a closure](https://www.geeksforgeeks.org/closures-in-golang/).

````go
// Golang program to illustrate how
// to create a Closure
package main

import "fmt"

func main() {
	
	// Declaring the variable
	GFG := 0

	// Assigning an anonymous
	// function to a variable
	counter := func() int {
	GFG += 1
	return GFG
	}

	fmt.Println(counter())
	fmt.Println(counter())	
}
````

# Scan Function in GoLang
- The [Scan function](https://www.educative.io/answers/what-is-the-scan-function-in-golang) in the Go programming language is used to read data from the standard input, format the string, and store the resultant strings into the destinations specified by the additional arguments.

````go
package main
 
import (
    "fmt"
)
 
func main() {
 
  var name string
  var unit string
  var amount int
  var temp string

  // taking input and storing in variable using a sample input string would be:
  // "Faraz owns 500 acres of land"
  fmt.Scan(&name, &temp, &amount, &unit)
 
  // print out new string using the extracted values 
  fmt.Printf ("% d %s of land is owned by %s\n",amount, unit, name);
}
````

# Format the string
- In Go, the %s verb is used to format a string. 
- When used with a custom type that has a String() method defined, the String() method will be automatically called and its return value will be used in the formatted string.

# Variadic Functions in Go
- The function that is called with the varying number of arguments is known as variadic function.

`````go
function function_name(para1, para2...type)type {// code...}
`````

# Switch
- One interesting thing about [switch statements](https://exercism.org/tracks/go/exercises/blackjack/edit), is that the value after the switch keyword can be omitted, and we can have boolean conditions for each case.

````go
age := 21

switch {
case age > 20 && age < 30:
    // do something if age is between 20 and 30
case age == 10:
    // do something if age is equal to 10
default:
    // do something else for every other case
}
````

# References
- [Frequently Asked Questions (FAQ)](https://go.dev/doc/faq)
- [Go at Google: Language Design in the Service of Software Engineering](https://go.dev/talks/2012/splash.article#TOC_1)
- [Golang Interview Questions](https://www.interviewbit.com/golang-interview-questions/)
- [Using Modules and Packages in Go](https://levelup.gitconnected.com/using-modules-and-packages-in-go-36a418960556)
- [Why Golang Is Taking Over the Software Industry?](https://betterprogramming.pub/why-golang-is-about-to-take-over-the-software-industry-fb48174a4cf)
- [Companies Using Golang: Top 7 Famous Apps You Use Daily](https://brainhub.eu/library/companies-using-golang)
- [Golang In Ecommerce](https://reemishirsath.medium.com/golang-in-ecommerce-ac87a8512e75)
- [Optimizing Go Microservices for Low Latency & High Throughput](https://muratdemirci.com.tr/en/optimizing-go-microservices/)
- [Data_Structures_and_Algorithms_Using_GoLang](https://github.com/sraynitjsr/Data_Structures_and_Algorithms_Using_GoLang/tree/main)
