# What is Golang?
- [Go](https://go.dev/) is a high level, general-purpose programming language that is [very strongly and statically typed](TypesGo.md) by providing support for garbage collection and [concurrent programming](GoRoutines&Channels.md).
- Go technically is [pass by value](https://stackoverflow.com/questions/47296325/passing-by-reference-and-value-in-go-to-functions).
- Go is a [case-sensitive language](https://en.wikipedia.org/wiki/Case_sensitivity).

# :star: Why is Golang fast compared to other languages?
- Golang is faster than other programming languages because of its simple and efficient memory management and [concurrency model](GoRoutines&Channels.md).
- The compilation process to machine code is very fast and efficient.
- Additionally, the dependencies are linked to a single binary file thereby putting off dependencies on servers.
- It also uses a compile-link model for generating executable binaries from the source code.

# Key Features

| Title                                                                                                    | Remarks                                                                                                                                                                                                                                                                                          |
|----------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| :star: [Coding Helpers & Guidelines in GoLang](CodingHelpers&GuidelinesGo.md)                            | Coding Helpers and guidelines for coding in GoLang.                                                                                                                                                                                                                                              |
| :star: [Pointers in GoLang](PointersGo.md)                                                               | GoLang supports pointers using (*, & operators)                                                                                                                                                                                                                                                  |
| :star: [Concurrency in GoLang](GoRoutines&Channels.md)                                                            | Go provides very good support for concurrency using [Go Routines or channels](https://go.dev/tour/concurrency/1)                                                                                                                                                                                 |
| [Types in GoLang](TypesGo.md)                                                                            | Type parameters permit what is known as generic programming, in which functions and data structures are defined in terms of types that are specified later, when those functions and data structures are used.                                                                                   |
| [Slices in GoLang](SlicesGo.md)                                                                          | Slice in Go is a lightweight data structure of variable length sequence for storing homogeneous data.                                                                                                                                                                                            |
| [OOPs in GoLang](OOPsGo/Readme.md)                                                                       | Although Go has types and methods and allows an object-oriented style of programming, there is no type hierarchy.                                                                                                                                                                                |
| [Panic & Recover in GoLang](Panics&Recover.md)                                                           | [Panic & Recover](https://golangbot.com/panic-and-recover/) is like exception in GoLang.                                                                                                                                                                                                         |
| [Unit Testing in GoLang](UnitTestingGo.md)                                                               | GoLang supports unit testing using [Testing](https://pkg.go.dev/testing) package                                                                                                                                                                                                                 |
| [Labels in GoLang](https://medium.com/golangspec/labels-in-go-4ffd81932339)                              | Label is used in break and continue statement where it’s optional but it’s required in goto statement.                                                                                                                                                                                           |
| [How to Work With SQL in GoLang?](https://betterprogramming.pub/how-to-work-with-sql-in-go-ca8bc0b30722) | [database/sql package](https://pkg.go.dev/database/sql) helps to query SQL databases.                                                                                                                                                                                                            |
| [DB Transaction in GoLang](DBTransactionGo.md)                                                           | [Using Begin, Commit code](https://dev.to/techschoolguru/a-clean-way-to-implement-database-transaction-in-golang-2ba) block, [atomicity](https://github.com/Anshul619/System-Designs/blob/main/src/1_HLDDesignComponents/0_SystemGlossaries/Database/Atomicity.md) can be implemented in GoLang. |
| [Packages](Packages.md)                                                                                  | In go, code is organized using packages.                                                                                                                                                                                                                                                         |
| [Modules](Modules.md)                                                                                    | Modules represent an app/service in Go.                                                                                                                                                                                                                                                          |
| [Comments & Documentation in Go](CommentsGo.md)                                                          | -                                                                                                                                                                                                                                                                                                |

# :+1: What are the advantages of Golang over other languages?

| Advantage                      | Description                                                                                                                                                    |
|--------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Simple and Understandable      | Go was developed by keeping simplicity, maintainability and readability in mind.                                                                               |
| Standard Powerful Library      | Go supports all standard libraries and packages that help in writing code easily and efficiently.                                                              |
| :star: Support for concurrency | Go provides very good support for [concurrency using  Go Routines or channels](GoRoutines&Channels.md).                                                              |
| Static Type Checking           | Go is a very strong and [statically typed programming language.](TypesGo.md)<br/>- This ensures that the code is type-safe and all type conversions are handled efficiently. |
| Easy to install binaries       | Go provides support for generating binaries for the applications with all required dependencies.                                                                                                                                                             |
| Good Testing Support           | Go has good support for [writing unit test cases](UnitTestingGo.md) along with our code.                                                                                                                                                               |

# Who uses GoLang?

| Company         | Title                                                                                                                                                                   |
|-----------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Open Source     | [Docker](https://github.com/Anshul619/System-Designs/blob/main/src/1_HLDDesignComponents/6_ContainerOrchestrationServices/Docker/Readme.md) is developed in [GoLang]()  |
| Open Source     | [Kubernates](https://github.com/Anshul619/System-Designs/blob/main/src/1_HLDDesignComponents/6_ContainerOrchestrationServices/Kubernates.md) is developed in [GoLang]() |
| Google          | Google's download server (dl.google.com), golang.org.                                                                                                                   |
| Google          | Google Chrome and Google Earth were created in this way. It is also used in YouTube and Google App Engine.                                                              |
| Uber            | Geofence service, which serves the user’s location and product availability.                                                                                            |
| Twitch          | Go enabled Twitch to improve 20 times the GC (garbage collection) factor responsible for automatically managing dynamically allocated memory.                           |
| Dailymotion     | Automation of APIs, tests has been improved.                                                                                                                            |
| SendGrid        | Simultaneous asynchronous programming<br/>- reduced maintenance costs and resolved concurrency problems<br/>- efficiently process over 500 million messages a day       |
| Dropbox         | -                                                                                                                                                                       |
| SoundCloud      | Static typing and fast compilation enabled by Go.                                                                                                                       |


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
# References
- [Frequently Asked Questions (FAQ)](https://go.dev/doc/faq)
- [Golang Interview Questions](https://www.interviewbit.com/golang-interview-questions/)
- [Using Modules and Packages in Go](https://levelup.gitconnected.com/using-modules-and-packages-in-go-36a418960556)
- [Why Golang Is Taking Over the Software Industry?](https://betterprogramming.pub/why-golang-is-about-to-take-over-the-software-industry-fb48174a4cf)
- [Companies Using Golang: Top 7 Famous Apps You Use Daily](https://brainhub.eu/library/companies-using-golang)