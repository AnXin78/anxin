package piscine

func LastRune(s string) rune {
	// Loop through the string starting from the end
	for i := len(s) - 1; i >= 0; i-- {
		return rune(s[i]) // Return the first rune (last character) found
	}
	return 0 // If the string is empty, return 0
}
