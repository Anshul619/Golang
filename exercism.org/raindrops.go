package main

import (
	"log"
	"strconv"
)

func Convert(number int) string {

	var out string

	if number%3 == 0 {
		out += "Pling"
	}

	if number%5 == 0 {
		out += "Plang"
	}

	if number%7 == 0 {
		out += "Plong"
	}

	if len(out) == 0 {
		out = strconv.Itoa(number)
	}

	return out
}

func main() {
	log.Println(Convert(28))
	log.Println(Convert(30))
	log.Println(Convert(34))
}
