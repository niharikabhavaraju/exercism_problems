package raindrops

import (
	"strconv"
)

var mapping = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

// Convert number to string according to the number's factors
func Convert(num int) (result string) {
	for key, str := range mapping {
		if num%key == 0 {
			result += str
		}
	}

	if result == "" {
		result = strconv.Itoa(num)
	}
	return result
}
