package main

import (
	"log"
	"strconv"
	"unicode"
)

func IsValidISBN(isbn string) bool {

	sum := 0
	counter := 10
	isCheckDigit := false

	for i, v := range isbn {

		if string(v) != "-" {
			switch {
			case unicode.IsDigit(v):
				{
					num, _ := strconv.Atoi(string(v))
					sum += num * counter
				}
			case string(v) == "X" && i == len(isbn)-1:
				{
					isCheckDigit = true
				}
			default:
				return false
			}
			counter--
		}
	}
	switch {
	case counter != 0:
		return false
	case isCheckDigit:
		{
			checkDigit := (11 - sum%11) % 11
			if checkDigit > 0 && checkDigit < 11 {
				return true
			}
		}
	case sum%11 == 0:
		return true
	}

	return false
}

func main() {
	// log.Println(IsValidISBN("3-598-21508-8"))
	// log.Println(IsValidISBN("3-598-21508-9"))
	// log.Println(IsValidISBN("3-598-21507-X"))
	// log.Println(IsValidISBN("3-598-21507-A"))
	// log.Println(IsValidISBN("4-598-21507-B"))
	// log.Println(IsValidISBN("3-598-2X507-9"))
	log.Println(IsValidISBN("3598215078X"))
	log.Println(IsValidISBN("00"))
}
