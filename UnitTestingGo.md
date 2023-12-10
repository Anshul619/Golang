# Conventions

| Instruction                                                                    | Remarks                                                                                                                 |
|--------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------|
| Test files in Go end in `_test.go`                                             | Create a new file ending in `_test.go` in the same directory as your package sources.                                   |
| Individual tests are identified by function names matching `^Test[A-Z]`        |                                                                                                                         |
| Run subtests by `t.Run()`                                                      |                                                                                                                         |
| Log the error and mark the test failed by calling `t.Errorf()` or `t.Fatal()`. | Fatal is equivalent to Print() followed by a call to os.Exit(1).                                                        |
| go test -v                                                                     | Run `go test` command in that directory.<br/>- That script finds the Test functions, builds a test binary, and runs it. |

# Example Tests

````go
import testing

func TestFoo(t *testing.T) {
    ...
}
````

# References
- [How to write unit tests in GoLang?](https://blog.alexellis.io/golang-writing-unit-tests/)