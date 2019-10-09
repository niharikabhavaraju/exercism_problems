package proverb

//Proverb builds the proverb according to the rhymes specified
func Proverb(rhyme []string) []string {
	var result []string
	if len(rhyme) == 1 {
		result = append(result, "And all for the want of a "+rhyme[0]+".")
	} else {
		for index, str := range rhyme {
			if index+1 != len(rhyme) {
				result = append(result, "For want of a "+str+" the "+rhyme[index+1]+" was lost"+".")
			} else {
				result = append(result, "And all for the want of a "+rhyme[0]+".")
			}
		}
	}
	return result
}
