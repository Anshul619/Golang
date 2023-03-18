package main

import (
	"log"
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	bucket := make([]int, 26)

	input = strings.ToLower(input)

	for _, v := range input {
		if unicode.IsLetter(v) {
			bucket[v-rune('a')]++
		}

	}

	for _, v := range bucket {
		if v == 0 {
			return false
		}
	}

	return true
}

func main() {
	log.Println(IsPangram("The quick brown fox jumps over the lazy dog."))
}
