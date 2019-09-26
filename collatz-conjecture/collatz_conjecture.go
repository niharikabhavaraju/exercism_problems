package collatzconjecture

import "errors"

// CollatzConjecture calculates the number of steps required to reach 1
func CollatzConjecture(input int) (int, error) {
	var steps int = 0
	if input < 1 {
		return 0, errors.New("Input must be positive integer")
	}
	for input != 1 {
		if input%2 == 0 {
			input = input / 2
			steps++

		} else {
			input = (input * 3) + 1
			steps++
		}
	}
	return steps, nil
}
