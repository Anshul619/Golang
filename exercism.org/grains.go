package main

import (
	"errors"
	"log"
	"math"
)

func Square(number int) (uint64, error) {

	if number < 1 || number > 64 {
		return 0, errors.New("out of range")
	}

	return uint64(math.Pow(2, float64(number)-1)), nil
}

func Total() uint64 {
	return uint64(math.Pow(2, 64) - 1)
	//return 1<<64 - 1
}

func main() {
	// log.Println(Square(4))
	// log.Println(Square(64))
	// log.Println(Square(-1))
	// log.Println(Square(0))
	// log.Println(Square(65))
	log.Println(Total())
}
