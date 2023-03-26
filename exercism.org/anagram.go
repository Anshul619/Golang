package main

import (
	"log"
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {

	out := []string{}

	subject = strings.ToLower(subject)

	sBytes := []byte(subject)
	sort.Slice(sBytes, func(i, j int) bool { return sBytes[i] > sBytes[j] })

	for _, v := range candidates {

		log.Println(v, subject)

		if len(v) == len(subject) && strings.ToLower(v) != subject {

			cBytes := []byte(strings.ToLower(v))
			sort.Slice(cBytes, func(i, j int) bool { return cBytes[i] > cBytes[j] })

			if string(sBytes) == string(cBytes) {
				out = append(out, v)
			}
		}
	}

	return out
}

func main() {
	log.Println(Detect("allergy", []string{"gallery", "ballerina", "regally", "clergy", "largely", "leading"}))
	log.Println(Detect("BANANA", []string{"BANANA"}))
}
