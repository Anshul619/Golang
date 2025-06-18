package main

import (
	"errors"
	"log"
)

func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("disallow mismatching length")
	}

	out := 0

	for i, v := range a {

		if string(v) != string(b[i]) {
			out++
		}
	}

	return out, nil
}

func main() {
	log.Println(Distance("GGACTGAAATCTG", "GGACTGAAATCTG"))
	log.Println(Distance("GGACGGATTCTG", "AGGACGGATTCT"))
}
