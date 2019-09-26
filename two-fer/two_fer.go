package twofer

// ShareWith returns a string with name specified
//if name is not specified it returns a default string
func ShareWith(name string) string {
	if name != "" {
		return "One for " + name + ", one for me."
	}
	return "One for you, one for me."
}
