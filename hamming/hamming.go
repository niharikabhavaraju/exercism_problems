package hamming

import "errors"

// Distance calculates the hamming distance between two DNA strands
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("Both the DNA strands should have same length")
	}
	hammingDistance := 0
	for i := range a {
		if a[i] != b[i] {
			hammingDistance++
		}
	}
	return hammingDistance, nil
}
