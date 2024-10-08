# :star: Why is Golang fast compared to other languages?
- Golang is faster than other programming languages because of its simple and efficient memory management and [concurrency model](GoRoutines&Channels).
- The compilation process to machine code is very fast and efficient.
- Additionally, the dependencies are linked to a single binary file thereby putting off dependencies on servers.
- It also uses a compile-link model for generating executable binaries from the source code.

> To put it simply, if Docker had not been written in Go, it would not have been as successful. ~Solomon Hykes

# :+1: What are the advantages of Golang over other languages?

| Advantage                      | Description                                                                                                                                                                       |
|--------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| :star: Support for concurrency | Go provides very good support for [concurrency using  Go Routines or channels](GoRoutines&Channels).                                                                              |
| Simple and Understandable      | Go was developed by keeping simplicity, maintainability and readability in mind. Hence easy to learn.                                                                             |
| Standard Powerful Library      | Go supports all standard libraries and packages that help in writing code easily and efficiently.                                                                                 |
| Static Type Checking           | Go is a very strong and [statically typed programming language.](OOPs/TypesGo.md)<br/>- This ensures that the code is type-safe and all type conversions are handled efficiently. |
| Easy to install binaries       | Go provides support for generating binaries for the applications with all required dependencies.                                                                                  |
| Good Testing Support           | Go has good support for [writing unit test cases](UnitTesting.md) along with our code.                                                                                            |
| Fast Compiled Time             | Go is designed for quick compilation without the need for dependency checking, so it addresses the build pain.                                                                    |
| Compiled Language              | Go is a compiled language, and thus runs several orders of magnitude faster than interpreted languages, like Python or Ruby.                                                      |
| Transparent Type System        | Go is easy to figure out how much everything your program’s trying to do will cost.                                                                                               |
| Strings Processing             | Google built a rich library of string functions into Go, Garbage collecting makes strings in Go simple to think about, and efficient in ways some other string libraries are not. |
| Efficient Garbage Collection   |                                                                                                                                                                                   |
| Memory Efficient               | Go is designed to be memory efficient with lessor memory footprint, which saves hardware costs as well.                                                                           |

# References
- [Is there a favoured language for microservices development at Google (say, Go)? Why?](https://www.quora.com/Is-there-a-favoured-language-for-microservices-development-at-Google-say-Go-Why)