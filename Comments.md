# Comments
- In Go, [comments](https://exercism.org/tracks/go/concepts/comments) play an important role in documenting code. 
- They are used by the [godoc](https://pkg.go.dev/golang.org/x/tools/cmd/godoc) command, which extracts these comments to create documentation about Go packages.

# Package Comments
- Package comments should be written directly before a package clause (`package x`) and begin with `Package x ...` like this.

````go
// Package kelvin provides tools to convert temperatures to and from Kelvin.
package kelvin
````

# Function comments
- A function comment should be written directly before the function declaration.
- It should also explain what arguments the function takes, what it does with them, and what its return values mean, ending in a period):

````go
// CelsiusFreezingTemp returns an integer value equal to the temperature at which water freezes in degrees Celsius.
func CelsiusFreezingTemp() int {
	return 0
}
````

# Golint
- [Golint](https://github.com/golang/lint) is a great tool to check for missing comments and other common stylistic issues, which you can read more about at [Effective Go](https://go.dev/doc/effective_go).

# References
- [Comments in Go](https://exercism.org/tracks/go/concepts/comments)