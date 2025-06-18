package main

import (
	"errors"
	"log"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {

	h := Histogram{rune('A'): 0, rune('C'): 0, rune('G'): 0, rune('T'): 0}

	for _, v := range d {

		switch string(v) {
		case "A", "C", "G", "T":
			h[v]++
		default:
			return Histogram{}, errors.New("strand with invalid nucleotides")
		}
	}

	return h, nil
}

func main() {
	var dna DNA
	dna = "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC"
	log.Println(dna.Counts())

	dna = "AGXXACT"
	log.Println(dna.Counts())
}
