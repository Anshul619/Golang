package main

import "log"

func ToRNA(dna string) string {

	out := ""

	for _, v := range dna {
		switch string(v) {
		case "G":
			out += "C"
		case "C":
			out += "G"
		case "T":
			out += "A"
		case "A":
			out += "U"
		}
	}

	return out
}

func main() {
	log.Println(ToRNA(""))
	log.Println(ToRNA("C"))
	log.Println(ToRNA("G"))
	log.Println(ToRNA("T"))
	log.Println(ToRNA("ACGTGGTCTTAA"))

}
