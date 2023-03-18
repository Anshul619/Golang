package main

import "log"

func Reverse(input string) string {
	out := []rune(input)

	l, r := 0, len(out)-1

	for l < r {
		t := out[l]
		out[l] = out[r]
		out[r] = t
		l++
		r--
	}

	return string(out)
}

func main() {
	log.Println(Reverse("cool"))
	log.Println(Reverse("Hello, 世界"))
}
