package piscine

func SplitWhiteSpaces(s string) []string {
	word := ""
	var result []string

	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' {
			if word != "" {
				result = append(result, word)
				word = ""
			}
		} else {
			word += string(r)
		}
	}
	if word != "" {
		result = append(result, word)
	}
	return result
}
