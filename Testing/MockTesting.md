# Mock Testing
- This is helpful to test a function by mocking internal function calls (external api, db calls)

# Mockgen command
- [mockgen](ttps://github.com/uber-go/mock/tree/main/mockgen) is a tool from the golang/mock package that automatically generates mock implementations from interfaces.

# Example usages

````shell
go install github.com/golang/mock/mockgen@latest
mockgen -source=cmd/payments/interfaces.go -destination=cmd/payments/interfaces_mock.go -package=payments
````
- It generates a mock struct with methods that allow you to specify expected inputs and outputs.