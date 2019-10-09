package darts

import "math"

//Score returns the dart score according to the dart position
func Score(x float64, y float64) int {
	dartposition := math.Sqrt(x*x + y*y)
	switch {
	case dartposition <= 1.0:
		return 10
	case dartposition <= 5.0:
		return 5
	case dartposition <= 10.0:
		return 1
	default:
		return 0
	}

}
