package main

import (
	"log"
	"strconv"
	"strings"
)

func Valid(id string) bool {

	id = strings.ReplaceAll(id, " ", "")

	if len(id) < 2 {
		return false
	}

	sum := 0
	for i := len(id) - 1; i >= 0; i-- {
		n, err := strconv.Atoi(string(id[i]))

		if err != nil {
			return false
		}

		if (len(id)-i)%2 == 0 {
			double := n * 2
			if double > 9 {
				sum += double - 9
			} else {
				sum += double
			}
		} else {
			sum += n
		}
	}

	return sum%10 == 0
}

func main() {
	log.Println(Valid("4539 3195 0343 6467"))
	// // log.Println(Valid("8273 1232 7352 0569"))
	// log.Println(Valid("0"))
	log.Println(Valid("055 444 285"))
	// log.Println(Valid("055 444 286"))
	// log.Println(Valid("095 245 88"))
	log.Println(Valid("059a"))
	log.Println(Valid("055# 444$ 285"))
}
