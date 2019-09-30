package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[byte]int

// DNA contains a string of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
func (d DNA) Counts() (Histogram, error) {
	var h = Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, n := range string(d) {
		nucleotide, ok := h[byte(n)]
		if !ok {
			return h, errors.New("Invalid nucleotide" + string(nucleotide))
		}
		h[byte(n)]++
	}
	return h, nil
}
