package strand

var dnaToRnaMapping = map[string]string{
	"C": "G",
	"G": "C",
	"T": "A",
	"A": "U",
}

//ToRNA converts dna to  rna
func ToRNA(dna string) string {
	var rna string
	if dna == "" {
		return ""
	}
	for _, elem := range dna {
		rna += dnaToRnaMapping[string(elem)]
	}
	return rna
}
