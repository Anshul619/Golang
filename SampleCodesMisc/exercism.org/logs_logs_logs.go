package main

import (
	"fmt"
	"log"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {

	for _, v := range log {

		u := fmt.Sprintf("%U", v)

		switch u {
		case "U+2757":
			return "recommendation"
		case "U+1F50D":
			return "search"
		case "U+2600":
			return "weather"
		}
	}

	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	out := []rune{}

	for _, v := range log {

		if v == oldRune {
			out = append(out, newRune)
		} else {
			out = append(out, v)
		}
	}

	return string(out)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	if utf8.RuneCountInString(log) <= limit {
		return true
	}
	return false
}

func main() {
	log.Println(Application("❗ recommended search product 🔍"))
	log.Println(Replace("please replace '👎' with '👍'", '👎', '👍'))
	log.Println(WithinLimit("hello❗", 6))
}
