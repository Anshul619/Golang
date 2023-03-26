package main

import (
	"log"
	"strings"
	"unicode"
)

func Hey(remark string) string {

	remark = strings.TrimSpace(remark)
	hasQuestion := false

	spaces, letters, upper := 0, 0, 0

	for i, v := range remark {
		switch {
		case unicode.IsLetter(v):
			letters++

			if unicode.IsUpper(v) {
				upper++
			}
		case unicode.IsSpace(v):
			spaces++
		case i == len(remark)-1 && string(v) == "?":
			hasQuestion = true
		}
	}

	switch {
	case len(remark) == spaces:
		return "Fine. Be that way!"
	case hasQuestion:
		if letters > 0 && letters == upper {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Sure."
		}
	case letters > 0 && letters == upper:
		return "Whoa, chill out!"
	}

	return "Whatever."
}

func main() {
	log.Println(Hey("Wait! Hang on. Are you going to be OK?"))
	log.Println(Hey("I HATE THE DENTIST"))
	log.Println(Hey("1, 2, 3 GO!"))
	log.Println(Hey("WHAT'S GOING ON?"))
	log.Println(Hey("fffbbcbeab?"))
	log.Println(Hey("1, 2, 3"))
	log.Println(Hey("4?"))
	log.Println(Hey(":) ?"))
	log.Println(Hey("Okay if like my  spacebar  quite a bit?   "))
}
