# Go runtime
- Go does have an extensive library, called the [runtime](https://pkg.go.dev/runtime), that is part of every Go program.
- The runtime library implements [garbage collection](GarbageCollector/Readme.md), [concurrency](https://github.com/Anshul619/Concurrency-Go), stack management, and other critical features of the Go language.
- It is important to understand, however, that Go's runtime **does not include** a virtual machine, such as is provided by the Java runtime.
- Go programs are compiled ahead of time to native machine code.