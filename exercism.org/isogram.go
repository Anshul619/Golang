package main

import (
	"log"
	"strings"
)

func IsIsogram(word string) bool {

	m := make(map[string]bool)

	for _, v := range word {
		if string(v) != " " && string(v) != "-" {

			str := strings.ToLower(string(v))

			if _, ok := m[str]; ok {
				return false
			}

			m[str] = true
		}
	}

	return true
}

func main() {
	log.Println(IsIsogram("lumberjacks"))
	log.Println(IsIsogram("background"))
	log.Println(IsIsogram("downstream"))
	log.Println(IsIsogram("six-year-old"))
	log.Println(IsIsogram("isograms"))
}
