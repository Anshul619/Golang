package main

// [Errors](https://golangbot.com/custom-errors/) can be returned as multiple values from a function in Go

// error is inbuilt go interface
//type error interface {
//	Error() string
//}

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

func main() {
    wrappedErr := fmt.Errorf("operation failed: %w", err)
    errors.Is(wrappedErr, targetErr)        // checks if any error in the chain matches targetErr
    errors.As(wrappedErr, &errorString{}) // checks if any error in chain is of type NotFoundError
}