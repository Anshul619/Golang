package main

import (
	"log"
	"strings"
	"unicode"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {

	freq := Frequency{}
	word := ""

	for i, v := range phrase {

		switch {
		case unicode.IsLetter(v),
			unicode.IsDigit(v),
			string(v) == "'" && len(word) != 0 && i+1 < len(phrase)-1 && unicode.IsLetter(rune(phrase[i+1])):
			word += string(v)
		case len(word) != 0:
			{
				freq[strings.ToLower(word)]++
				word = ""
			}
		}
	}

	if len(word) != 0 {
		freq[strings.ToLower(word)]++
	}

	return freq
}

func main() {
	// log.Println(WordCount("one fish two fish red fish blue fish"))
	// log.Println(WordCount("one,\ntwo,\nthree"))
	// log.Println(WordCount("car: carpet as java: javascript!!&@$%^&"))
	// log.Println(WordCount("'First: don't laugh. Then: don't cry. You're getting it.'"))
	// log.Println(WordCount("go Go GO Stop stop"))
	log.Println(WordCount("testing, 1, 2 testing"))
	// log.Println(WordCount("Joe can't tell between 'large' and large."))
	log.Println(WordCount("testing, 1234, 243, 1 testing"))
}
