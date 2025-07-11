# Conventions

|                                                                               | Remarks                                                                                                                 |
|-------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------|
| Test files in Go end in `_test.go`                                            | Create a new file ending in `_test.go` in the same directory as your package sources.                                   |
| Individual tests are identified by function names matching `^Test[A-Z]`       |                                                                                                                         |
| Run subtests by `t.Run()`                                                     |                                                                                                                         |
| Log the error and mark the test failed by calling `t.Errorf()` or `t.Fatal()` | Fatal is equivalent to Print() followed by a call to os.Exit(1).                                                        |
| go test -v                                                                    | Run `go test` command in that directory.<br/>- That script finds the Test functions, builds a test binary, and runs it. |
| go test -v -run TestSquare                                                    | Run specific `TestSquare` test                                                                                          |
| Use `t.Parallel()` for independent tests                                      |                                                                                                                         |
| Use `-race` to catch race conditions in both test types                       |                                                                                                                         |
| Use `go test -cover` for coverage reports                                     |                                                                                                                         |
| Use test containers or dockertest for DBs in integration tests                |                                                                                                                         |
| Use gomock, testify, or fake structs for mocking in unit tests                |                                                                                                                         |

# Table Tests
- Table tests in Go are a popular approach to organizing multiple use-cases for a single test.
- They are a method of validating a function or method against multiple parameters and results. 
- When a table contains multiple test cases, the test simply iterates through all table entries and runs the necessary tests.

# References
- [How to write unit tests in GoLang?](https://blog.alexellis.io/golang-writing-unit-tests/)
- [Comprehensive Guide to Testing in Go](https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/)