package main

import (
	"errors"
	"log"
)

func CollatzConjecture(n int) (int, error) {
	steps := 0

	if n == 0 {
		return 0, errors.New("zero number")
	}

	if n < 0 {
		return 0, errors.New("negativ number")
	}

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		steps++
	}

	return steps, nil
}

func main() {
	log.Println(CollatzConjecture(12))
}
