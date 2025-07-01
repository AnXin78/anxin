package piscine

func Index(s string, toFind string) int {
	// If toFind is empty, we return 0 as any string contains an empty string at index 0
	if toFind == "" {
		return 0
	}

	// Loop through string `s` and check for the first occurrence of `toFind`
	for i := 0; i <= len(s)-len(toFind); i++ {
		if s[i:i+len(toFind)] == toFind {
			return i
		}
	}

	// Return -1 if the substring was not found
	return -1
}
