# :star: Difference between concurrent and parallelism in Golang
- Concurrency is when your program can handle multiple tasks at once while parallelism is when your program can execute multiple tasks at once using multiple processors.
- In other words, concurrency is a property of a program that allows you to have multiple tasks in progress at the same time, but not necessarily executing at the same time. (i.e. instead of keeping idle while waiting for the blocking call, we optimize and execute another code.)
- Parallelism is a runtime property where two or more tasks are executed at the same time.

# How can I control the number of CPUs?
- The number of CPUs available simultaneously to executing goroutines is controlled by the [GOMAXPROCS](https://pkg.go.dev/runtime) shell environment variable, whose default value is the number of CPU cores available.
- Programs with the potential for parallel execution should therefore achieve it by default on a multiple-CPU machine.