package main

import (
    "fmt"
    "log"
    "strings"
)
// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
    border := strings.Repeat("*", numStarsPerLine)
    return fmt.Sprintf("%v\n%s\n%v", border, welcomeMsg, border)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
    str1 := strings.ReplaceAll(oldMsg, "*", "")
    str2 := strings.TrimSpace(str1)
    return str2
}

func main() {
    log.Println(WelcomeMessage("Judy"))
    log.Println(AddBorder("Welcome!", 10))
    log.Println(CleanupMessage("B **************************A  \n A"))
}