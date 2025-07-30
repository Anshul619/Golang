# Unit testing in Go

| Command/Code               | Title                                                                         | Remarks                                                                               |
|----------------------------|-------------------------------------------------------------------------------|---------------------------------------------------------------------------------------|
| user_test.go               | Test files in Go end in `_test.go`                                            | Create a new file ending in `_test.go` in the same directory as your package sources. |
| ^Test[A-Z]                 | Individual tests are identified by function names matching `^Test[A-Z]`       |                                                                                       |
| t.Run()                    | Run subtests by `t.Run()`                                                     |                                                                                       |
| t.Errorf() or t.Fatal()    | Log the error and mark the test failed by calling `t.Errorf()` or `t.Fatal()` | Fatal is equivalent to Print() followed by a call to os.Exit(1).                      |
| go test -v                 | Run `go test` command in that directory                                       | That script finds the Test functions, builds a test binary, and runs it.              |
| go test -v -run TestSquare | Run specific `TestSquare` test                                                |                                                                                       |
| t.Parallel()               | To run independent tests                                                      |                                                                                       |
| go test -race              | To catch race conditions in both test types                                   |                                                                                       |
| go test -cover             | For coverage reports                                                          |                                                                                       |
| go test -bench=.           | To run all benchmarks in the package                                          |                                                                                       |

# Table Tests
- Table tests in Go are a popular approach to organizing multiple use-cases for a single test.
- They are a method of validating a function or method against multiple parameters and results. 
- When a table contains multiple test cases, the test simply iterates through all table entries and runs the necessary tests.

# References
- [How to write unit tests in GoLang?](https://blog.alexellis.io/golang-writing-unit-tests/)
- [Comprehensive Guide to Testing in Go](https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/)