# Introduction
- Create a new file ending in `_test.go` in the same directory as your package sources. 
- Inside that file, import [testing](https://pkg.go.dev/testing) and write functions of the form.
- Run `go test` command in that directory. 
- That script finds the Test functions, builds a test binary, and runs it.

````go
func TestFoo(t *testing.T) {
    ...
}
````

# Conventions

| Instruction                                                                           | Remarks                                                          |
|---------------------------------------------------------------------------------------|------------------------------------------------------------------|
| Test files in Go  end in `_test.go`                                                   | -                                                                |
| Individual tests are identified by function names matching `^Test[A-Z]`               | -                                                                |
| You can run subtests by `t.Run()`                                                     | -                                                                |
| You log the error and mark the test failed by calling `t.Errorf()` or `t.Fatal()`.    | Fatal is equivalent to Print() followed by a call to os.Exit(1). |
| go test -v                                                                            | Print out the tests it is running                                |


# References
- [How to write unit tests in GoLang?](https://blog.alexellis.io/golang-writing-unit-tests/)