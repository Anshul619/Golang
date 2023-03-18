package main

import "log"

func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}

	if year%100 == 0 {
		return false
	}

	if year%4 == 0 {
		return true
	}

	return false
}

func main() {
	log.Println(IsLeapYear(1997))
	log.Println(IsLeapYear(1996))
	log.Println(IsLeapYear(1900))
	log.Println(IsLeapYear(2000))
}
