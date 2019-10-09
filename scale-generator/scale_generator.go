package scale

import "strings"

var (
	intervalMapping = map[string]int{
		"A": 3,
		"M": 2,
		"m": 1,
	}
	flatTonics = "Fdcgf"
	flatScale  = []string{"C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab", "A", "Bb", "B"}
	sharpScale = []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
)

//getScale returns sharps scale or flats scale accorrding to the tonic specified
func getScale(tonic string) []string {
	switch {
	case strings.Contains(flatTonics, tonic):
		return flatScale
	case len(tonic) == 1:
		return sharpScale
	case tonic[1] == 'b':
		return flatScale
	default:
		return sharpScale
	}
}

//Scale generates scale with specified tonic and interval
func Scale(tonic string, interval string) []string {
	var result []string
	scale := getScale(tonic)
	tonicindex := findindex(scale, tonic)
	if interval == "" {
		result = append(scale[tonicindex:], scale[:tonicindex]...)
	} else {
		for _, step := range interval[:len(interval)-1] {
			tonicindex = tonicindex + intervalMapping[string(step)]
			if tonicindex >= len(scale) {
				tonicindex = tonicindex - len(scale)
			}
			result = append(result, scale[tonicindex])
		}
		result = append([]string{strings.Title(tonic)}, result...)
	}
	return result
}

//findIndex returns the index of the tonic
func findindex(scale []string, item string) int {
	for index := range scale {
		if strings.EqualFold(scale[index], item) {
			return index
		}
	}
	return -1
}
