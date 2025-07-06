# Introduction
- [Panic & Recover](https://golangbot.com/panic-and-recover/) is like exception handling in GoLang.
- There is no try-&-catch in GoLang. (like java)

# Types of Error Logging

| Log                | Description                           |
|--------------------|---------------------------------------|
| log.Fatal("error") | Program will print error and exit     |
| log.Error("error") | Program will print error but NOT exit |

# Recovering from a Panic
- [recover](https://golangbot.com/panic-and-recover/#recoveringfromapanic) is a builtin function that is used to regain control of a panicking program.

# Error Group in GoRoutines
- [Error Group](https://pkg.go.dev/golang.org/x/sync/errgroup) is used to return errors in goRoutines.
